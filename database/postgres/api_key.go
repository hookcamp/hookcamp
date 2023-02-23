package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/frain-dev/convoy/datastore"
	"github.com/frain-dev/convoy/util"
	"github.com/jmoiron/sqlx"
	"gopkg.in/guregu/null.v4"
)

const (
	createAPIKey = `
    INSERT INTO convoy.api_keys (id,name,key_type,mask_id,role_type,role_project,role_endpoint,hash,salt,user_id,expires_at)
    VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11);
    `

	updateAPIKeyById = `
	UPDATE convoy.api_keys SET
	    name = $2,
		role_type= $3,
		role_project=$4,
		role_endpoint=$5,
		updated_at = now()
	WHERE id = $1 AND deleted_at IS NULL ;
	`

	fetchAPIKey = `
	SELECT
	    id,
		name,
	    key_type,
	    mask_id,
	    role_type as "role.type",
	    role_project as "role.project",
	    role_endpoint as "role.endpoint",
	    hash,
	    salt,
	    COALESCE(user_id, '') AS user_id,
	    created_at,
	    updated_at,
	    expires_at
	FROM convoy.api_keys
	WHERE %s = $1 AND deleted_at IS NULL
	`

	deleteAPIKeys = `
	UPDATE convoy.api_keys SET
	deleted_at = now()
	WHERE id IN (?);
	`
	baseAPIKeysCount = `
	WITH table_count AS (
		SELECT count(distinct(id)) as count
		FROM convoy.api_keys WHERE deleted_at IS NULL
		%s
	)
	`

	fetchAPIKeysPaginated = `
	SELECT
	   table_count.count,
	    id,
		name,
	    key_type,
	    mask_id,
	    role_type as "role.type",
	    role_project as "role.project",
	    role_endpoint as "role.endpoint",
	    hash,
	    salt,
	    COALESCE(user_id, '') AS user_id,
	    created_at,
	    updated_at,
	    expires_at
	FROM table_count, convoy.api_keys
	WHERE deleted_at IS NULL
	%s
	ORDER BY id LIMIT ? OFFSET ?;
	`

	baseFilter = `AND (role_project = ? OR ? = '') AND (role_endpoint = ? OR ? = '') AND (user_id = ? OR ? = '') AND (key_type = ? OR ? = '')`
)

var (
	ErrAPIKeyNotCreated = errors.New("api key could not be created")
	ErrAPIKeyNotUpdated = errors.New("api key could not be updated")
	ErrAPIKeyNotRevoked = errors.New("api key could not be revoked")
)

type apiKeyRepo struct {
	db *sqlx.DB
}

func NewAPIKeyRepo(db *sqlx.DB) datastore.APIKeyRepository {
	return &apiKeyRepo{db: db}
}

func (a *apiKeyRepo) CreateAPIKey(ctx context.Context, key *datastore.APIKey) error {
	var userID null.String

	if !util.IsStringEmpty(key.UserID) {
		userID = null.StringFrom(key.UserID)
	}

	result, err := a.db.ExecContext(
		ctx, createAPIKey, key.UID, key.Name, key.Type, key.MaskID,
		key.Role.Type, key.Role.Project, key.Role.Endpoint, key.Hash,
		key.Salt, userID, key.ExpiresAt,
	)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected < 1 {
		return ErrAPIKeyNotCreated
	}

	return nil
}

func (a *apiKeyRepo) UpdateAPIKey(ctx context.Context, key *datastore.APIKey) error {
	result, err := a.db.ExecContext(
		ctx, updateAPIKeyById, key.UID, key.Name, key.Role.Type, key.Role.Project, key.Role.Endpoint,
	)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected < 1 {
		return ErrAPIKeyNotUpdated
	}

	return nil
}

func (a *apiKeyRepo) FindAPIKeyByID(ctx context.Context, id string) (*datastore.APIKey, error) {
	apiKey := &datastore.APIKey{}
	err := a.db.QueryRowxContext(ctx, fmt.Sprintf(fetchAPIKey, "id"), id).StructScan(apiKey)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, datastore.ErrAPIKeyNotFound
		}
		return nil, err
	}

	return apiKey, nil
}

func (a *apiKeyRepo) FindAPIKeyByMaskID(ctx context.Context, maskID string) (*datastore.APIKey, error) {
	apiKey := &datastore.APIKey{}
	err := a.db.QueryRowxContext(ctx, fmt.Sprintf(fetchAPIKey, "mask_id"), maskID).StructScan(apiKey)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, datastore.ErrAPIKeyNotFound
		}
		return nil, err
	}

	return apiKey, nil
}

func (a *apiKeyRepo) FindAPIKeyByHash(ctx context.Context, hash string) (*datastore.APIKey, error) {
	apiKey := &datastore.APIKey{}
	err := a.db.QueryRowxContext(ctx, fmt.Sprintf(fetchAPIKey, "hash"), hash).StructScan(apiKey)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, datastore.ErrAPIKeyNotFound
		}
		return nil, err
	}

	return apiKey, nil
}

func (a *apiKeyRepo) RevokeAPIKeys(ctx context.Context, ids []string) error {
	query, args, err := sqlx.In(deleteAPIKeys, ids)
	if err != nil {
		return err
	}

	result, err := a.db.ExecContext(ctx, a.db.Rebind(query), args...)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected < 1 {
		return ErrAPIKeyNotRevoked
	}

	return nil
}

func (a *apiKeyRepo) LoadAPIKeysPaged(ctx context.Context, filter *datastore.ApiKeyFilter, pageable *datastore.Pageable) ([]datastore.APIKey, datastore.PaginationData, error) {
	var query string
	var err error
	var args []interface{}

	if len(filter.EndpointIDs) > 0 {
		filterQuery := `AND role_endpoint IN (?) ` + baseFilter
		q := fmt.Sprintf(baseAPIKeysCount, filterQuery) + fmt.Sprintf(fetchAPIKeysPaginated, filterQuery)
		args = []interface{}{filter.EndpointIDs, filter.ProjectID, filter.ProjectID, filter.EndpointID, filter.EndpointID, filter.UserID, filter.UserID, filter.KeyType, filter.KeyType}
		args = append(args, args...)
		args = append(args, pageable.Limit(), pageable.Offset())
		query, args, err = sqlx.In(q, args...)
		if err != nil {
			return nil, datastore.PaginationData{}, err
		}
		query = a.db.Rebind(query)
	} else {
		q := fmt.Sprintf(baseAPIKeysCount, baseFilter) + fmt.Sprintf(fetchAPIKeysPaginated, baseFilter)
		query = a.db.Rebind(q)
		args = []interface{}{filter.ProjectID, filter.ProjectID, filter.EndpointID, filter.EndpointID, filter.UserID, filter.UserID, filter.KeyType, filter.KeyType}
		args = append(args, args...)
		args = append(args, pageable.Limit(), pageable.Offset())
	}

	rows, err := a.db.QueryxContext(ctx, query, args...)
	if err != nil {
		return nil, datastore.PaginationData{}, err
	}

	var apiKeys []datastore.APIKey

	var count int
	ak := ApiKeyPaginated{}
	for rows.Next() {
		err = rows.StructScan(&ak)
		if err != nil {
			return nil, datastore.PaginationData{}, err
		}

		apiKeys = append(apiKeys, ak.APIKey)
		count = ak.Count
	}

	pagination := calculatePaginationData(count, pageable.Page, pageable.PerPage)
	return apiKeys, pagination, nil
}

type ApiKeyPaginated struct {
	Count int `db:"count"`
	datastore.APIKey
}