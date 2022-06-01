package datastore

import (
	"bytes"
	"database/sql/driver"
	"encoding/json"
	"errors"
	"strings"

	"github.com/frain-dev/convoy/auth"
	"github.com/frain-dev/convoy/config"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

type Pageable struct {
	Page    int `json:"page" bson:"page"`
	PerPage int `json:"per_page" bson:"per_page"`
	Sort    int `json:"sort" bson:"sort"`
}

type PaginationData struct {
	Total     int64 `json:"total"`
	Page      int64 `json:"page"`
	PerPage   int64 `json:"perPage"`
	Prev      int64 `json:"prev"`
	Next      int64 `json:"next"`
	TotalPage int64 `json:"totalPage"`
}

type Period int

var PeriodValues = map[string]Period{
	"daily":   Daily,
	"weekly":  Weekly,
	"monthly": Monthly,
	"yearly":  Yearly,
}

const (
	Daily Period = iota
	Weekly
	Monthly
	Yearly
)

func IsValidPeriod(period string) bool {
	_, ok := PeriodValues[period]
	return ok
}

type DocumentStatus string

type SearchParams struct {
	CreatedAtStart int64 `json:"created_at_start" bson:"created_at_start"`
	CreatedAtEnd   int64 `json:"created_at_end" bson:"created_at_end"`
}

const (
	ActiveDocumentStatus   DocumentStatus = "Active"
	InactiveDocumentStatus DocumentStatus = "Inactive"
	DeletedDocumentStatus  DocumentStatus = "Deleted"
)

type StrategyProvider string

type GroupType string

type SourceType string

type VerifierType string

type EncodingType string

const (
	HTTPSource     SourceType = "http"
	RestApiSource  SourceType = "rest_api"
	PubSubSource   SourceType = "pub_sub"
	DBChangeStream SourceType = "db_change_stream"
)

const (
	HMacVerifier      VerifierType = "hmac"
	BasicAuthVerifier VerifierType = "basic_auth"
	APIKeyVerifier    VerifierType = "api_key"
)

const (
	Base64Encoding EncodingType = "base64"
	HexEncoding    EncodingType = "hex"
)

const (
	OutgoingGroup GroupType = "outgoing"
	IncomingGroup GroupType = "incoming"
)

const (
	DefaultStrategyProvider     = LinearStrategyProvider
	LinearStrategyProvider      = "linear"
	ExponentialStrategyProvider = "exponential"
)

var (
	DefaultStrategyConfig = StrategyConfiguration{
		Type:       "linear",
		Duration:   100,
		RetryCount: 10,
	}

	DefaultSignatureConfig = SignatureConfiguration{
		Header: "X-Convoy-Signature",
		Hash:   "SHA256",
	}

	DefaultRateLimitConfig = RateLimitConfiguration{
		Count:    1000,
		Duration: "1m",
	}
)

var (
	ErrApplicationNotFound = errors.New("application not found")
	ErrEndpointNotFound    = errors.New("endpoint not found")
	ErrSourceNotFound      = errors.New("source not found")
	ErrUserNotFound        = errors.New("user not found")
)

const (
	ActiveEndpointStatus   EndpointStatus = "active"
	InactiveEndpointStatus EndpointStatus = "inactive"
	PendingEndpointStatus  EndpointStatus = "pending"
)

type Application struct {
	ID              primitive.ObjectID `json:"-" bson:"_id"`
	UID             string             `json:"uid" bson:"uid"`
	GroupID         string             `json:"group_id" bson:"group_id"`
	Title           string             `json:"name" bson:"title"`
	SupportEmail    string             `json:"support_email" bson:"support_email"`
	SlackWebhookURL string             `json:"slack_webhook_url,omitempty" bson:"slack_webhook_url"`
	IsDisabled      bool               `json:"is_disabled" bson:"is_disabled"`

	Endpoints []Endpoint         `json:"endpoints" bson:"endpoints"`
	CreatedAt primitive.DateTime `json:"created_at,omitempty" bson:"created_at,omitempty" swaggertype:"string"`
	UpdatedAt primitive.DateTime `json:"updated_at,omitempty" bson:"updated_at,omitempty" swaggertype:"string"`
	DeletedAt primitive.DateTime `json:"deleted_at,omitempty" bson:"deleted_at,omitempty" swaggertype:"string"`

	Events int64 `json:"events" bson:"-"`

	DocumentStatus DocumentStatus `json:"-" bson:"document_status"`
}

type EndpointStatus string

type Endpoint struct {
	UID         string         `json:"uid" bson:"uid"`
	TargetURL   string         `json:"target_url" bson:"target_url"`
	Description string         `json:"description" bson:"description"`
	Status      EndpointStatus `json:"status" bson:"status"`
	Secret      string         `json:"secret" bson:"secret"`

	HttpTimeout       string `json:"http_timeout" bson:"http_timeout"`
	RateLimit         int    `json:"rate_limit" bson:"rate_limit"`
	RateLimitDuration string `json:"rate_limit_duration" bson:"rate_limit_duration"`

	Events []string `json:"events" bson:"events"`

	CreatedAt primitive.DateTime `json:"created_at,omitempty" bson:"created_at,omitempty" swaggertype:"string"`
	UpdatedAt primitive.DateTime `json:"updated_at,omitempty" bson:"updated_at,omitempty" swaggertype:"string"`
	DeletedAt primitive.DateTime `json:"deleted_at,omitempty" bson:"deleted_at,omitempty" swaggertype:"string"`

	DocumentStatus DocumentStatus `json:"-" bson:"document_status"`
}

var ErrGroupNotFound = errors.New("group not found")
var ErrOrgNotFound = errors.New("organisation not found")
var ErrOrgMemberNotFound = errors.New("organisation member not found")

type Group struct {
	ID         primitive.ObjectID `json:"-" bson:"_id"`
	UID        string             `json:"uid" bson:"uid"`
	Name       string             `json:"name" bson:"name"`
	LogoURL    string             `json:"logo_url" bson:"logo_url"`
	Type       GroupType          `json:"type" bson:"type"`
	Config     *GroupConfig       `json:"config" bson:"config"`
	Statistics *GroupStatistics   `json:"statistics" bson:"-"`

	// TODO(subomi): refactor this into the Instance API.
	RateLimit         int    `json:"rate_limit" bson:"rate_limit"`
	RateLimitDuration string `json:"rate_limit_duration" bson:"rate_limit_duration"`

	CreatedAt primitive.DateTime `json:"created_at,omitempty" bson:"created_at,omitempty" swaggertype:"string"`
	UpdatedAt primitive.DateTime `json:"updated_at,omitempty" bson:"updated_at,omitempty" swaggertype:"string"`
	DeletedAt primitive.DateTime `json:"deleted_at,omitempty" bson:"deleted_at,omitempty" swaggertype:"string"`

	DocumentStatus DocumentStatus `json:"-" bson:"document_status"`
}

type GroupConfig struct {
	RateLimit       RateLimitConfiguration `json:"ratelimit"`
	Strategy        StrategyConfiguration  `json:"strategy"`
	Signature       SignatureConfiguration `json:"signature"`
	DisableEndpoint bool                   `json:"disable_endpoint"`
	ReplayAttacks   bool                   `json:"replay_attacks"`
}

type RateLimitConfiguration struct {
	Count    int    `json:"count" bson:"count"`
	Duration string `json:"duration" bson:"duration"`
}

type StrategyConfiguration struct {
	Type       StrategyProvider `json:"type" valid:"required~please provide a valid strategy type, in(linear|exponential)~unsupported strategy type"`
	Duration   uint64           `json:"duration" valid:"required~please provide a valid duration in seconds,int"`
	RetryCount uint64           `json:"retry_count" valid:"required~please provide a valid retry count,int"`
}

type SignatureConfiguration struct {
	Header config.SignatureHeaderProvider `json:"header" valid:"required~please provide a valid signature header"`
	Hash   string                         `json:"hash" valid:"required~please provide a valid hash,supported_hash~unsupported hash type"`
}

type SignatureValues struct {
	Header config.SignatureHeaderProvider `json:"header" valid:"required~please provide a valid signature header"`
	Hash   string                         `json:"hash" valid:"required~please provide a valid hash,supported_hash~unsupported hash type"`
}

type GroupStatistics struct {
	GroupID      string `json:"-" bson:"group_id"`
	MessagesSent int64  `json:"messages_sent" bson:"messages_sent"`
	TotalApps    int64  `json:"total_apps" bson:"total_apps"`
}

type GroupFilter struct {
	Names []string `json:"name" bson:"name"`
}

func (g *GroupFilter) WithNamesTrimmed() *GroupFilter {
	f := GroupFilter{[]string{}}

	for _, s := range g.Names {
		f.Names = append(f.Names, strings.TrimSpace(s))
	}

	return &f
}

func (o *Group) IsDeleted() bool { return o.DeletedAt > 0 }

func (o *Group) IsOwner(a *Application) bool { return o.UID == a.GroupID }

var (
	ErrEventNotFound = errors.New("event not found")
)

type AppMetadata struct {
	UID          string `json:"uid" bson:"uid"`
	Title        string `json:"title" bson:"title"`
	GroupID      string `json:"group_id" bson:"group_id"`
	SupportEmail string `json:"support_email" bson:"support_email"`
}

// EventType is used to identify an specific event.
// This could be "user.new"
// This will be used for data indexing
// Makes it easy to filter by a list of events
type EventType string

//Event defines a payload to be sent to an application
type Event struct {
	ID               primitive.ObjectID `json:"-" bson:"_id"`
	UID              string             `json:"uid" bson:"uid"`
	EventType        EventType          `json:"event_type" bson:"event_type"`
	MatchedEndpoints int                `json:"matched_endpoints" bson:"matched_enpoints"`

	// ProviderID is a custom ID that can be used to reconcile this Event
	// with your internal systems.
	// This is optional
	// If not provided, we will generate one for you
	ProviderID string `json:"provider_id" bson:"provider_id"`

	// Data is an arbitrary JSON value that gets sent as the body of the
	// webhook to the endpoints
	Data json.RawMessage `json:"data" bson:"data"`

	AppMetadata *AppMetadata `json:"app_metadata,omitempty" bson:"app_metadata"`

	CreatedAt primitive.DateTime `json:"created_at,omitempty" bson:"created_at,omitempty" swaggertype:"string"`
	UpdatedAt primitive.DateTime `json:"updated_at,omitempty" bson:"updated_at,omitempty" swaggertype:"string"`
	DeletedAt primitive.DateTime `json:"deleted_at,omitempty" bson:"deleted_at,omitempty" swaggertype:"string"`

	DocumentStatus DocumentStatus `json:"-" bson:"document_status"`
}

type EventDeliveryStatus string
type HttpHeader map[string]string

var (
	ErrEventDeliveryNotFound        = errors.New("event not found")
	ErrEventDeliveryAttemptNotFound = errors.New("delivery attempt not found")
)

const (
	// ScheduledEventStatus : when  a Event has been scheduled for delivery
	ScheduledEventStatus  EventDeliveryStatus = "Scheduled"
	ProcessingEventStatus EventDeliveryStatus = "Processing"
	DiscardedEventStatus  EventDeliveryStatus = "Discarded"
	FailureEventStatus    EventDeliveryStatus = "Failure"
	SuccessEventStatus    EventDeliveryStatus = "Success"
	RetryEventStatus      EventDeliveryStatus = "Retry"
)

func (e EventDeliveryStatus) IsValid() bool {
	switch e {
	case ScheduledEventStatus,
		ProcessingEventStatus,
		DiscardedEventStatus,
		FailureEventStatus,
		SuccessEventStatus,
		RetryEventStatus:
		return true
	default:
		return false
	}
}

type Metadata struct {
	// Data to be sent to endpoint.
	Data     json.RawMessage  `json:"data" bson:"data"`
	Strategy StrategyProvider `json:"strategy" bson:"strategy"`
	// NextSendTime denotes the next time a Event will be published in
	// case it failed the first time
	NextSendTime primitive.DateTime `json:"next_send_time" bson:"next_send_time"`

	// NumTrials: number of times we have tried to deliver this Event to
	// an application
	NumTrials uint64 `json:"num_trials" bson:"num_trials"`

	IntervalSeconds uint64 `json:"interval_seconds" bson:"interval_seconds"`

	RetryLimit uint64 `json:"retry_limit" bson:"retry_limit"`
}

func (em Metadata) Value() (driver.Value, error) {
	b := new(bytes.Buffer)

	if err := json.NewEncoder(b).Encode(em); err != nil {
		return driver.Value(""), err
	}

	return driver.Value(b.String()), nil
}

type EndpointMetadata struct {
	UID               string         `json:"uid" bson:"uid"`
	TargetURL         string         `json:"target_url" bson:"target_url"`
	Status            EndpointStatus `json:"status" bson:"status"`
	Secret            string         `json:"secret" bson:"secret"`
	HttpTimeout       string         `json:"http_timeout" bson:"http_timeout"`
	RateLimit         int            `json:"rate_limit" bson:"rate_limit"`
	RateLimitDuration string         `json:"rate_limit_duration" bson:"rate_limit_duration"`

	Sent bool `json:"sent" bson:"sent"`
}

type EventIntervalData struct {
	Interval int64  `json:"index" bson:"index"`
	Time     string `json:"date" bson:"total_time"`
}

type EventInterval struct {
	Data  EventIntervalData `json:"data" bson:"_id"`
	Count uint64            `json:"count" bson:"count"`
}

type EventMetadata struct {
	UID       string    `json:"uid" bson:"uid"`
	EventType EventType `json:"name" bson:"name"`
}

type DeliveryAttempt struct {
	ID         primitive.ObjectID `json:"-" bson:"_id"`
	UID        string             `json:"uid" bson:"uid"`
	MsgID      string             `json:"msg_id" bson:"msg_id"`
	URL        string             `json:"url" bson:"url"`
	Method     string             `json:"method" bson:"method"`
	EndpointID string             `json:"endpoint_id" bson:"endpoint_id"`
	APIVersion string             `json:"api_version" bson:"api_version"`

	IPAddress        string     `json:"ip_address,omitempty" bson:"ip_address,omitempty"`
	RequestHeader    HttpHeader `json:"request_http_header,omitempty" bson:"request_http_header,omitempty"`
	ResponseHeader   HttpHeader `json:"response_http_header,omitempty" bson:"response_http_header,omitempty"`
	HttpResponseCode string     `json:"http_status,omitempty" bson:"http_status,omitempty"`
	ResponseData     string     `json:"response_data,omitempty" bson:"response_data,omitempty"`
	Error            string     `json:"error,omitempty" bson:"error,omitempty"`
	Status           bool       `json:"status,omitempty" bson:"status,omitempty"`

	CreatedAt primitive.DateTime `json:"created_at,omitempty" bson:"created_at,omitempty" swaggertype:"string"`
	UpdatedAt primitive.DateTime `json:"updated_at,omitempty" bson:"updated_at,omitempty" swaggertype:"string"`
	DeletedAt primitive.DateTime `json:"deleted_at,omitempty" bson:"deleted_at,omitempty" swaggertype:"string"`
}

//Event defines a payload to be sent to an application
type EventDelivery struct {
	ID            primitive.ObjectID `json:"-" bson:"_id"`
	UID           string             `json:"uid" bson:"uid"`
	EventMetadata *EventMetadata     `json:"event_metadata" bson:"event_metadata"`

	// Endpoint contains the destination of the event.
	EndpointMetadata *EndpointMetadata `json:"endpoint" bson:"endpoint"`

	AppMetadata      *AppMetadata        `json:"app_metadata,omitempty" bson:"app_metadata"`
	Metadata         *Metadata           `json:"metadata" bson:"metadata"`
	Description      string              `json:"description,omitempty" bson:"description"`
	Status           EventDeliveryStatus `json:"status" bson:"status"`
	DeliveryAttempts []DeliveryAttempt   `json:"-" bson:"attempts"`

	CreatedAt primitive.DateTime `json:"created_at,omitempty" bson:"created_at,omitempty" swaggertype:"string"`
	UpdatedAt primitive.DateTime `json:"updated_at,omitempty" bson:"updated_at,omitempty" swaggertype:"string"`
	DeletedAt primitive.DateTime `json:"deleted_at,omitempty" bson:"deleted_at,omitempty" swaggertype:"string"`

	DocumentStatus DocumentStatus `json:"-" bson:"document_status"`
}

var (
	ErrAPIKeyNotFound   = errors.New("api key not found")
	ErrDuplicateAppName = errors.New("an application with this name exists")
)

type KeyType string

type APIKey struct {
	ID        primitive.ObjectID `json:"-" bson:"_id"`
	UID       string             `json:"uid" bson:"uid"`
	MaskID    string             `json:"mask_id,omitempty" bson:"mask_id"`
	Name      string             `json:"name" bson:"name"`
	Role      auth.Role          `json:"role" bson:"role"`
	Hash      string             `json:"hash,omitempty" bson:"hash"`
	Salt      string             `json:"salt,omitempty" bson:"salt"`
	Type      KeyType            `json:"key_type" bson:"key_type"`
	ExpiresAt primitive.DateTime `json:"expires_at,omitempty" bson:"expires_at,omitempty" swaggertype:"string"`
	CreatedAt primitive.DateTime `json:"created_at,omitempty" bson:"created_at" swaggertype:"string"`
	UpdatedAt primitive.DateTime `json:"updated_at,omitempty" bson:"updated_at" swaggertype:"string"`
	DeletedAt primitive.DateTime `json:"delted_at,omitempty" bson:"deleted_at" swaggertype:"string"`

	DocumentStatus DocumentStatus `json:"-" bson:"document_status"`
}

type Source struct {
	ID         primitive.ObjectID `json:"-" bson:"_id"`
	UID        string             `json:"uid" bson:"uid"`
	GroupID    string             `json:"group_id" bson:"group_id"`
	MaskID     string             `json:"mask_id" bson:"mask_id"`
	Name       string             `json:"name" bson:"name"`
	Type       SourceType         `json:"type" bson:"type"`
	IsDisabled bool               `json:"is_disabled" bson:"is_disabled"`
	Verifier   *VerifierConfig    `json:"verifier" bson:"verifier"`

	CreatedAt primitive.DateTime `json:"created_at,omitempty" bson:"created_at" swaggertype:"string"`
	UpdatedAt primitive.DateTime `json:"updated_at,omitempty" bson:"updated_at" swaggertype:"string"`
	DeletedAt primitive.DateTime `json:"deleted_at,omitempty" bson:"deleted_at" swaggertype:"string"`

	DocumentStatus DocumentStatus `json:"-" bson:"document_status"`
}

type User struct {
	ID        primitive.ObjectID `json:"-" bson:"_id"`
	UID       string             `json:"uid" bson:"uid"`
	FirstName string             `json:"first_name" bson:"first_name"`
	LastName  string             `json:"last_name" bson:"last_name"`
	Email     string             `json:"email" bson:"email"`
	Password  string             `json:"-" bson:"password"`
	Role      auth.Role          `json:"role" bson:"role"`

	CreatedAt primitive.DateTime `json:"created_at,omitempty" bson:"created_at,omitempty" swaggertype:"string"`
	UpdatedAt primitive.DateTime `json:"updated_at,omitempty" bson:"updated_at,omitempty" swaggertype:"string"`
	DeletedAt primitive.DateTime `json:"deleted_at,omitempty" bson:"deleted_at,omitempty" swaggertype:"string"`

	DocumentStatus DocumentStatus `json:"-" bson:"document_status"`
}
type VerifierConfig struct {
	Type      VerifierType `json:"type,omitempty" bson:"type" valid:"supported_verifier~please provide a valid verifier type,optional"`
	HMac      HMac         `json:"hmac" bson:"hmac"`
	BasicAuth BasicAuth    `json:"basic_auth" bson:"basic_auth"`
	ApiKey    ApiKey       `json:"api_key" bson:"api_key"`
}

type HMac struct {
	Header   string       `json:"header,omitempty" bson:"header"`
	Hash     string       `json:"hash,omitempty" bson:"hash" valid:"supported_hash,optional"`
	Secret   string       `json:"secret,omitempty" bson:"secret"`
	Encoding EncodingType `json:"encoding,omitempty" bson:"encoding" valid:"supported_encoding~please provide a valid encoding type,optional"`
}

type BasicAuth struct {
	UserName string `json:"username,omitempty" bson:"username"`
	Password string `json:"password,omitempty" bson:"password"`
}

type ApiKey struct {
	APIKey       string `json:"key,omitempty" bson:"key"`
	APIKeyHeader string `json:"header,omitempty" bson:"header"`
}

type Organisation struct {
	ID             primitive.ObjectID `json:"-" bson:"_id"`
	UID            string             `json:"uid" bson:"uid"`
	OwnerID        string             `json:"owner_id" bson:"owner_id"`
	Name           string             `json:"name" bson:"name"`
	DocumentStatus DocumentStatus     `json:"-" bson:"document_status"`
	CreatedAt      primitive.DateTime `json:"created_at,omitempty" bson:"created_at,omitempty" swaggertype:"string"`
	UpdatedAt      primitive.DateTime `json:"updated_at,omitempty" bson:"updated_at,omitempty" swaggertype:"string"`
	DeletedAt      primitive.DateTime `json:"deleted_at,omitempty" bson:"deleted_at,omitempty" swaggertype:"string"`
}

type OrganisationMember struct {
	ID             primitive.ObjectID `json:"-" bson:"_id"`
	UID            string             `json:"uid" bson:"uid"`
	OrganisationID string             `json:"organisation_id" bson:"organisation_id"`
	UserID         string             `json:"user_id" bson:"user_id"`
	Role           auth.Role          `json:"role" bson:"role"`
	DocumentStatus DocumentStatus     `json:"-" bson:"document_status"`
	CreatedAt      primitive.DateTime `json:"created_at,omitempty" bson:"created_at,omitempty" swaggertype:"string"`
	UpdatedAt      primitive.DateTime `json:"updated_at,omitempty" bson:"updated_at,omitempty" swaggertype:"string"`
	DeletedAt      primitive.DateTime `json:"deleted_at,omitempty" bson:"deleted_at,omitempty" swaggertype:"string"`
}

type Password struct {
	Plaintext string
	Hash      []byte
}

func (p *Password) GenerateHash() error {
	hash, err := bcrypt.GenerateFromPassword([]byte(p.Plaintext), 12)
	if err != nil {
		return err
	}

	p.Hash = hash
	return nil
}

func (p *Password) Matches() (bool, error) {
	err := bcrypt.CompareHashAndPassword(p.Hash, []byte(p.Plaintext))
	if err != nil {
		switch {
		case errors.Is(err, bcrypt.ErrMismatchedHashAndPassword):
			return false, nil
		default:
			return false, err
		}
	}

	return true, err
}
