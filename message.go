package hookcamp

import (
	"bytes"
	"context"
	"database/sql/driver"
	"encoding/json"
	"errors"
	pager "github.com/gobeam/mongo-go-pagination"
	"github.com/hookcamp/hookcamp/server/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

type MessageStatus string

var (
	ErrMessageNotFound = errors.New("message not found")
)

const (
	// UnknownMessageStatus when we don't know the state of a message
	UnknownMessageStatus MessageStatus = "Unknown"
	// ScheduledMessageStatus : when  a message has been scheduled for
	// delivery
	ScheduledMessageStatus  MessageStatus = "Scheduled"
	ProcessingMessageStatus MessageStatus = "Processing"
	FailureMessageStatus    MessageStatus = "Failure"
	SuccessMessageStatus    MessageStatus = "Success"
	RetryMessageStatus      MessageStatus = "Retry"
)

type MessageMetadata struct {
	// NextSendTime denotes the next time a message will be published in
	// case it failed the first time
	NextSendTime primitive.DateTime `json:"next_send_time" bson:"next_send_time"`

	// NumTrials: number of times we have tried to deliver this message to
	// an application
	NumTrials int64 `json:"num_trials" bson:"num_trials"`

	RetryLimit int64 `json:"retry_limit" bson:"retry_limit"`
}

type AppMetadata struct {
	OrgID string `json:"org_id" bson:"org_id"`

	Endpoints []EndpointMetadata `json:"endpoints" bson:"endpoints"`
}

type EndpointMetadata struct {
	UID       string `json:"uid" bson:"uid"`
	TargetURL string `json:"target_url" bson:"target_url"`

	Merged bool `json:"merged" bson:"merged"`
}

func (m MessageMetadata) Value() (driver.Value, error) {
	b := new(bytes.Buffer)

	if err := json.NewEncoder(b).Encode(m); err != nil {
		return driver.Value(""), err
	}

	return driver.Value(b.String()), nil
}

// EventType is used to identify an specific event.
// This could be "user.new"
// This will be used for data indexing
// Makes it easy to filter by a list of events
type EventType string

// Message defines a payload to be sent to an application
type Message struct {
	ID        primitive.ObjectID `json:"-" bson:"_id"`
	UID       string             `json:"uid" bson:"uid"`
	AppID     string             `json:"app_id" bson:"app_id"`
	EventType EventType          `json:"event_type" bson:"event_type"`

	// ProviderID is a custom ID that can be used to reconcile this message
	// with your internal systems.
	// This is optional
	// If not provided, we will generate one for you
	ProviderID string `json:"provider_id" bson:"provider_id"`

	// Data is an arbitrary JSON value that gets sent as the body of the
	// webhook to the endpoints
	Data json.RawMessage `json:"data" bson:"data"`

	Metadata *MessageMetadata `json:"metadata" bson:"metadata"`

	Description string `json:"description,omitempty" bson:"description"`

	Status MessageStatus `json:"status" bson:"status"`

	AppMetadata *AppMetadata `json:"app_metadata,omitempty" bson:"app_metadata"`

	MessageAttempts []MessageAttempt `json:"attempts" bson:"attempts"`

	CreatedAt primitive.DateTime `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt primitive.DateTime `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
	DeletedAt primitive.DateTime `json:"deleted_at,omitempty" bson:"deleted_at,omitempty"`
}

type MessageAttempt struct {
	ID         primitive.ObjectID `json:"-" bson:"_id"`
	UID        string             `json:"uid" bson:"uid"`
	MsgID      string             `json:"msg_id" bson:"msg_id"`
	EndpointID string             `json:"endpoint_id" bson:"endpoint_id"`
	APIVersion string             `json:"api_version" bson:"api_version"`

	IPAddress        string        `json:"ip_address,omitempty" bson:"ip_address,omitempty"`
	ContentType      string        `json:"content_type,omitempty" bson:"content_type,omitempty"`
	Header           http.Header   `json:"http_header,omitempty" bson:"http_header,omitempty"`
	HttpResponseCode string        `json:"http_status,omitempty" bson:"http_status,omitempty"`
	ResponseData     string        `json:"response_data,omitempty" bson:"response_data,omitempty"`
	Error            string        `json:"error,omitempty" bson:"error,omitempty"`
	Status           MessageStatus `json:"status,omitempty" bson:"status,omitempty"`

	CreatedAt primitive.DateTime `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt primitive.DateTime `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
	DeletedAt primitive.DateTime `json:"deleted_at,omitempty" bson:"deleted_at,omitempty"`
}

type MessageRepository interface {
	CreateMessage(context.Context, *Message) error
	LoadMessageIntervals(context.Context, string, models.SearchParams, Period, int) ([]models.MessageInterval, error)
	LoadMessagesByAppId(context.Context, string) ([]Message, error)
	FindMessageByID(ctx context.Context, id string) (*Message, error)
	LoadMessagesScheduledForPosting(context.Context) ([]Message, error)
	LoadMessagesForPostingRetry(context.Context) ([]Message, error)
	LoadAbandonedMessagesForPostingRetry(context.Context) ([]Message, error)
	UpdateStatusOfMessages(context.Context, []Message, MessageStatus) error
	UpdateMessage(ctx context.Context, m Message) error
	LoadMessagesPaged(context.Context, string, models.Pageable) ([]Message, pager.PaginationData, error)
}
