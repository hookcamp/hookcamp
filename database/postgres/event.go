package postgres

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"time"

	"github.com/frain-dev/convoy/datastore"
	"github.com/frain-dev/convoy/pkg/httpheader"
	"github.com/frain-dev/convoy/util"
	"github.com/jmoiron/sqlx"
	"gopkg.in/guregu/null.v4"
)

var (
	ErrEventNotCreated = errors.New("event could not be created")
	ErrEventNotFound   = errors.New("event not found")
)

const (
	createEvent = `
	INSERT INTO convoy.events (id, event_type, project_id, source_id, headers, raw, data) 
	VALUES ($1, $2, $3, $4, $5, $6, $7)
	`

	createEventEndpoints = `
	INSERT INTO convoy.events_endpoints (endpoint_id, event_id) VALUES (:endpoint_id, :event_id)
	`

	fetchEventById = `
	SELECT * from convoy.events WHERE id = $1 AND deleted_at is NULL;
	`

	fetchEventsByIds = ` 
	SELECT * from convoy.events WHERE id IN (?) AND deleted_at IS NULL;
	`

	countProjectMessages = `
	SELECT count(*) from convoy.events WHERE project_id = $1 AND deleted_at IS NULL;
	`
	countEvents = `
	SELECT count(distinct(ev.id)) from convoy.events ev 
	LEFT JOIN convoy.events_endpoints ee on ee.event_id = ev.id 
	LEFT JOIN convoy.endpoints e on ee.endpoint_id = e.id
	WHERE (ev.project_id = $1 OR $1 = '') AND (e.id = $2 OR $2 = '' ) 
	AND (ev.source_id = $3 OR $3 = '') AND ev.created_at >= $4 AND ev.created_at <= $5 AND ev.deleted_at IS NULL;
	`

	baseEventsPaged = `
	SELECT ev.id,
	ev.project_id, ev.source_id, ev.headers, ev.raw,
	ev.data, ev.created_at, ev.updated_at, ev.deleted_at,
	e.id AS "endpoint.id", e.title AS "endpoint.title", 
	e.project_id AS "endpoint.project_id", e.support_email AS "endpoint.support_email", 
	e.target_url AS "endpoint.target_url", s.id AS "source.id", 
	s.name AS "source.name" FROM convoy.events AS ev
	LEFT JOIN convoy.events_endpoints ee ON ee.event_id = ev.id
	LEFT JOIN convoy.endpoints e ON e.id = ee.endpoint_id
	LEFT JOIN convoy.sources s ON s.id = ev.source_id
	WHERE ev.deleted_at IS NULL
	`
	baseWhere = `AND (ev.project_id = ? OR ? = '') AND (ev.source_id = ? OR ? = '') AND ev.created_at >= ? AND ev.created_at <= ? GROUP BY ev.id, s.id, e.id LIMIT ? OFFSET ?`

	fetchEventsPaginatedFilterByEndpoints = baseEventsPaged + `AND e.id IN (?)` + baseWhere

	fetchEventsPaginated = baseEventsPaged + baseWhere

	softDeleteProjectEvents = `
	UPDATE convoy.events SET deleted_at = now()
	WHERE project_id = $1 AND created_at >= $2 AND created_at <= $3 
	AND deleted_at IS NULL
	`
	hardDeleteProjectEvents = `
	DELETE from convoy.events WHERE project_id = $1 AND created_at
	>= $2 AND created_at <= $3 AND deleted_at IS NULL
	`
)

type eventRepo struct {
	db *sqlx.DB
}

func NewEventRepo(db *sqlx.DB) datastore.EventRepository {
	return &eventRepo{db: db}
}

func (e *eventRepo) CreateEvent(ctx context.Context, event *datastore.Event) error {
	var sourceID *string

	if !util.IsStringEmpty(event.SourceID) {
		sourceID = &event.SourceID
	}

	tx, err := e.db.BeginTxx(ctx, &sql.TxOptions{})
	if err != nil {
		return err
	}

	headers, err := json.Marshal(event.Headers)
	if err != nil {
		return err
	}

	_, err = tx.ExecContext(ctx, createEvent,
		event.UID,
		event.EventType,
		event.ProjectID,
		sourceID,
		headers,
		event.Raw,
		event.Data,
	)
	if err != nil {
		return err
	}

	var ids []interface{}
	if len(event.Endpoints) > 0 {
		for _, endpointID := range event.Endpoints {
			ids = append(ids, &EventEndpoint{EventID: event.UID, EndpointID: endpointID})
		}

		_, err = tx.NamedExecContext(ctx, createEventEndpoints, ids)
		if err != nil {
			return err
		}

	}

	return tx.Commit()
}

func (e *eventRepo) FindEventByID(ctx context.Context, id string) (*datastore.Event, error) {
	event := &datastore.Event{}
	err := e.db.QueryRowxContext(ctx, fetchEventById, id).StructScan(event)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrEventNotFound
		}

		return nil, err
	}
	return event, nil
}

func (e *eventRepo) FindEventsByIDs(ctx context.Context, ids []string) ([]datastore.Event, error) {
	query, args, err := sqlx.In(fetchEventsByIds, ids)
	if err != nil {
		return nil, err
	}

	query = e.db.Rebind(query)
	rows, err := e.db.QueryxContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	events := make([]datastore.Event, 0)
	for rows.Next() {
		var event datastore.Event

		err := rows.StructScan(&event)
		if err != nil {
			return nil, err
		}

		events = append(events, event)
	}

	return events, nil
}

func (e *eventRepo) CountProjectMessages(ctx context.Context, projectID string) (int64, error) {
	var count int64

	err := e.db.QueryRowxContext(ctx, countProjectMessages, projectID).Scan(&count)
	if err != nil {
		return count, err
	}

	return count, nil
}

func (e *eventRepo) CountEvents(ctx context.Context, filter *datastore.Filter) (int64, error) {
	var count int64
	var projectID string
	startDate, endDate := getCreatedDateFilter(filter.SearchParams.CreatedAtStart, filter.SearchParams.CreatedAtEnd)

	if filter.Project != nil {
		projectID = filter.Project.UID
	}

	err := e.db.QueryRowxContext(ctx, countEvents, projectID, filter.EndpointID, filter.SourceID, startDate, endDate).Scan(&count)
	if err != nil {
		return count, err
	}

	return count, nil
}

func (e *eventRepo) LoadEventsPaged(ctx context.Context, filter *datastore.Filter) ([]datastore.Event, datastore.PaginationData, error) {
	var query string
	var args []interface{}
	var err error
	var projectID string
	startDate, endDate := getCreatedDateFilter(filter.SearchParams.CreatedAtStart, filter.SearchParams.CreatedAtEnd)

	if filter.Project != nil {
		projectID = filter.Project.UID
	}

	if !util.IsStringEmpty(filter.EndpointID) {
		filter.EndpointIDs = append(filter.EndpointIDs, filter.EndpointID)
	}

	if len(filter.EndpointIDs) > 0 {
		query, args, err = sqlx.In(fetchEventsPaginatedFilterByEndpoints, filter.EndpointIDs, projectID, projectID, filter.SourceID, filter.SourceID, startDate, endDate, filter.Pageable.Limit(), filter.Pageable.Offset())
		if err != nil {
			return nil, datastore.PaginationData{}, err
		}

		query = e.db.Rebind(query)
	} else {
		query = e.db.Rebind(fetchEventsPaginated)
		args = []interface{}{projectID, projectID, filter.SourceID, filter.SourceID, startDate, endDate, filter.Pageable.Limit(), filter.Pageable.Offset()}
	}

	rows, err := e.db.QueryxContext(ctx, query, args...)
	if err != nil {
		return nil, datastore.PaginationData{}, err
	}

	eventMap := make(map[string]*datastore.Event)
	var events []datastore.Event
	for rows.Next() {
		var data EventPaginated

		err = rows.StructScan(&data)
		if err != nil {
			return nil, datastore.PaginationData{}, err
		}

		record, exists := eventMap[data.UID]
		if exists {
			endpoint := data.Endpoint
			if !util.IsStringEmpty(endpoint.UID.String) {
				record.EndpointMetadata = append(record.EndpointMetadata, getEventEndpoint(endpoint))
			}

		} else {
			event := getEvent(data)
			endpoint := data.Endpoint
			if !util.IsStringEmpty(endpoint.UID.String) {
				event.EndpointMetadata = append(event.EndpointMetadata, getEventEndpoint(endpoint))
			}

			source := data.Source
			if !util.IsStringEmpty(source.UID.String) {
				event.Source = &datastore.Source{
					UID:  source.UID.String,
					Name: source.Name.String,
				}
			}

			eventMap[event.UID] = event
		}
	}

	var count int
	err = e.db.Get(&count, countProjectMessages, projectID)
	if err != nil {
		return nil, datastore.PaginationData{}, err
	}

	for _, event := range eventMap {
		events = append(events, *event)
	}

	pagination := calculatePaginationData(count, filter.Pageable.Page, filter.Pageable.PerPage)
	return events, pagination, nil
}

func (e *eventRepo) DeleteProjectEvents(ctx context.Context, filter *datastore.EventFilter, hardDelete bool) error {
	query := softDeleteProjectEvents
	startDate, endDate := getCreatedDateFilter(filter.CreatedAtStart, filter.CreatedAtEnd)

	if hardDelete {
		query = hardDeleteProjectEvents
	}

	_, err := e.db.ExecContext(ctx, query, filter.ProjectID, startDate, endDate)
	if err != nil {
		return err
	}

	return nil
}

func getCreatedDateFilter(startDate, endDate int64) (time.Time, time.Time) {
	return time.Unix(startDate, 0), time.Unix(endDate, 0)
}

func getEvent(data EventPaginated) *datastore.Event {
	return &datastore.Event{
		UID:              data.UID,
		EventType:        datastore.EventType(data.EventType),
		SourceID:         data.SourceID.String,
		AppID:            data.AppID,
		ProjectID:        data.ProjectID,
		Headers:          data.Headers,
		Data:             data.Data,
		Raw:              data.Raw,
		EndpointMetadata: []*datastore.Endpoint{},
		CreatedAt:        data.CreatedAt,
		UpdatedAt:        data.UpdatedAt,
	}
}

func getEventEndpoint(endpoint *Endpoint) *datastore.Endpoint {
	return &datastore.Endpoint{
		UID:          endpoint.UID.String,
		Title:        endpoint.Title.String,
		ProjectID:    endpoint.ProjectID.String,
		SupportEmail: endpoint.SupportEmail.String,
		TargetURL:    endpoint.TargetUrl.String,
	}
}

type EventEndpoint struct {
	EventID    string `db:"event_id"`
	EndpointID string `db:"endpoint_id"`
}

type EventPaginated struct {
	Count     int
	UID       string                `db:"id"`
	EventType string                `db:"event_type"`
	SourceID  null.String           `db:"source_id"`
	AppID     string                `db:"app_id"`
	ProjectID string                `db:"project_id"`
	Headers   httpheader.HTTPHeader `db:"headers"`
	Data      json.RawMessage       `db:"data"`
	Raw       string                `db:"raw"`
	CreatedAt time.Time             `db:"created_at"`
	UpdatedAt time.Time             `db:"updated_at"`
	DeletedAt null.Time             `db:"deleted_at"`
	Endpoint  *Endpoint             `db:"endpoint"`
	Source    *Source               `db:"source"`
}

type Endpoint struct {
	UID          null.String `db:"id"`
	Title        null.String `db:"title"`
	ProjectID    null.String `db:"project_id"`
	SupportEmail null.String `db:"support_email"`
	TargetUrl    null.String `db:"target_url"`
}

type Source struct {
	UID  null.String `db:"id"`
	Name null.String `db:"name"`
}
