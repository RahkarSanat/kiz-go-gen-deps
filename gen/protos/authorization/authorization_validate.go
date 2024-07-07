package authorization

import (
	"fmt"
	"encoding/json"
	reflect "reflect"
	"strings"

	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"

	base "github.com/RahkarSanat/kiz-go-gen-deps/gen/protos/base"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

var _ codes.Code
var _ status.Status
var _ reflect.Kind
var _ primitive.A
var _ fmt.Stringer

var _ base.BaseAccess

var _ json.Marshaler
var _ strings.Builder

type MqttFieldFilterMongo struct {
	Message    *string `bson:"message,omitempty" json:"message,omitempty"`
	ClientInfo *string `bson:"client_info,omitempty" json:"client_info,omitempty"`
}
type ResponseMqttFieldFilterMongo struct {
	FieldMessage  *string `bson:"field_message,omitempty" json:"field_message,omitempty"`
	FilterMessage *string `bson:"filter_message,omitempty" json:"filter_message,omitempty"`
}
type AuthorizeRequestMongo struct {
	Action *string     `bson:"action,omitempty" json:"action,omitempty"`
	Object *string     `bson:"object,omitempty" json:"object,omitempty"`
	Token  *TokenMongo `bson:"token,omitempty" json:"token,omitempty"`
}
type AccessRequestMongo struct {
	AccessToken *string  `bson:"access_token,omitempty" json:"access_token,omitempty"`
	Action      *string  `bson:"action,omitempty" json:"action,omitempty"`
	Object      *string  `bson:"object,omitempty" json:"object,omitempty"`
	Resources   []string `bson:"resources,omitempty" json:"resources,omitempty"`
	ApiKey      *string  `bson:"api_key,omitempty" json:"api_key,omitempty"`
}
type TokenMongo struct {
	T        *string `bson:"t,omitempty" json:"t,omitempty"`
	Ip       *string `bson:"ip,omitempty" json:"ip,omitempty"`
	Cid      *string `bson:"cid,omitempty" json:"cid,omitempty"`
	Aid      *string `bson:"aid,omitempty" json:"aid,omitempty"`
	Uid      *string `bson:"uid,omitempty" json:"uid,omitempty"`
	Subjects *string `bson:"subjects,omitempty" json:"subjects,omitempty"`
	Scopes   *string `bson:"scopes,omitempty" json:"scopes,omitempty"`
	Session  *string `bson:"session,omitempty" json:"session,omitempty"`
}
type AuthorizeResponseMongo struct {
	Granted *bool        `bson:"granted,omitempty" json:"granted,omitempty"`
	Grants  []GrantMongo `bson:"grants,omitempty" json:"grants,omitempty"`
}
type GrantMongo struct {
	Subject  *string          `bson:"subject,omitempty" json:"subject,omitempty"`
	Action   *string          `bson:"action,omitempty" json:"action,omitempty"`
	Object   *string          `bson:"object,omitempty" json:"object,omitempty"`
	Field    []string         `bson:"field,omitempty" json:"field,omitempty"`
	Filter   []string         `bson:"filter,omitempty" json:"filter,omitempty"`
	Location []string         `bson:"location,omitempty" json:"location,omitempty"`
	Time     []GrantTimeMongo `bson:"time,omitempty" json:"time,omitempty"`
}
type GrantTimeMongo struct {
	Duration *int64  `bson:"duration,omitempty" json:"duration,omitempty"`
	CronExp  *string `bson:"cron_exp,omitempty" json:"cron_exp,omitempty"`
}
type RequestMongo struct {
	Query *string `bson:"query,omitempty" json:"query,omitempty"`
	Body  *string `bson:"body,omitempty" json:"body,omitempty"`
}
type SanitizeRequestMongo struct {
	AccessToken     *string       `bson:"access_token,omitempty" json:"access_token,omitempty"`
	Action          *string       `bson:"action,omitempty" json:"action,omitempty"`
	Object          *string       `bson:"object,omitempty" json:"object,omitempty"`
	Scope           []string      `bson:"scope,omitempty" json:"scope,omitempty"`
	Request         *RequestMongo `bson:"request,omitempty" json:"request,omitempty"`
	IsInflux        *bool         `bson:"is_influx,omitempty" json:"is_influx,omitempty"`
	Db              *string       `bson:"db,omitempty" json:"db,omitempty"`
	ApiKey          *string       `bson:"api_key,omitempty" json:"api_key,omitempty"`
	ExcludeOwn      *bool         `bson:"exclude_own,omitempty" json:"exclude_own,omitempty"`
	ExcludeZone     *bool         `bson:"exclude_zone,omitempty" json:"exclude_zone,omitempty"`
	ExcludeShare    *bool         `bson:"exclude_share,omitempty" json:"exclude_share,omitempty"`
	ExcludeRelation *bool         `bson:"exclude_relation,omitempty" json:"exclude_relation,omitempty"`
}
type SanitizeDecryptedRequestMongo struct {
	Token           *TokenMongo   `bson:"token,omitempty" json:"token,omitempty"`
	Action          *string       `bson:"action,omitempty" json:"action,omitempty"`
	Object          *string       `bson:"object,omitempty" json:"object,omitempty"`
	Scope           []string      `bson:"scope,omitempty" json:"scope,omitempty"`
	Request         *RequestMongo `bson:"request,omitempty" json:"request,omitempty"`
	IsInflux        *bool         `bson:"is_influx,omitempty" json:"is_influx,omitempty"`
	Db              *string       `bson:"db,omitempty" json:"db,omitempty"`
	ApiKey          *string       `bson:"api_key,omitempty" json:"api_key,omitempty"`
	ExcludeOwn      *bool         `bson:"exclude_own,omitempty" json:"exclude_own,omitempty"`
	ExcludeZone     *bool         `bson:"exclude_zone,omitempty" json:"exclude_zone,omitempty"`
	ExcludeShare    *bool         `bson:"exclude_share,omitempty" json:"exclude_share,omitempty"`
	ExcludeRelation *bool         `bson:"exclude_relation,omitempty" json:"exclude_relation,omitempty"`
}
type SanitizeResponseMongo struct {
	Grants  []GrantMongo  `bson:"grants,omitempty" json:"grants,omitempty"`
	Request *RequestMongo `bson:"request,omitempty" json:"request,omitempty"`
}
type GrantsListMongo struct {
	Items []GrantMongo `bson:"items,omitempty" json:"items,omitempty"`
}
type MqttAuthorizationMongo struct {
	Topic         *string `bson:"topic,omitempty" json:"topic,omitempty"`
	ActiveSession *string `bson:"active_session,omitempty" json:"active_session,omitempty"`
	MqttUsername  *string `bson:"mqtt_username,omitempty" json:"mqtt_username,omitempty"`
	Action        *string `bson:"action,omitempty" json:"action,omitempty"`
}
type TopicIsAvailableMongo struct {
	Available *string `bson:"available,omitempty" json:"available,omitempty"`
	Allow     *bool   `bson:"allow,omitempty" json:"allow,omitempty"`
}

func isZero(x interface{}) bool {
	return reflect.DeepEqual(x, reflect.Zero(reflect.TypeOf(x)).Interface())
}

func FlattenHelper(prefix string, obj interface{}, flatMap map[string]interface{}) {
	val := reflect.ValueOf(obj)

	if val.Kind() == reflect.Ptr {
		nval := val.Elem()
		if nval.Kind() != reflect.Invalid {
			val = nval
		} else {
		}
	}
	switch val.Kind() {
	case reflect.Struct:
		for i := 0; i < val.NumField(); i++ {
			field := val.Type().Field(i)
			name := strings.Split(field.Tag.Get("json"), ",")[0]
			FlattenHelper(prefix+name+".", val.Field(i).Interface(), flatMap)
		}

	case reflect.Map:
		switch t := obj.(type) {
		case *map[string]any:
			for k, v := range *t {
				FlattenHelper(prefix+k+".", v, flatMap)
			}
		case map[string]any:
			for k, v := range t {
				FlattenHelper(prefix+k+".", v, flatMap)
			}
		case *map[string]string:
			for k, v := range *t {
				FlattenHelper(prefix+k+".", v, flatMap)
			}
		case map[string]string:
			for k, v := range t {
				FlattenHelper(prefix+k+".", v, flatMap)
			}
		}

	default:
		if obj != nil && !isZero(obj) {
			flatMap[prefix[:len(prefix)-1]] = obj
		}
	}
}

func (m *MqttFieldFilter) Validate() error {

	return nil
}

func (m *ResponseMqttFieldFilter) Validate() error {

	return nil
}

func (m *AuthorizeRequest) Validate() error {

	return nil
}

func (m *AccessRequest) Validate() error {

	return nil
}

func (m *Token) Validate() error {

	return nil
}

func (m *AuthorizeResponse) Validate() error {

	return nil
}

func (m *Grant) Validate() error {

	return nil
}

func (m *GrantTime) Validate() error {

	return nil
}

func (m *Request) Validate() error {

	return nil
}

func (m *SanitizeRequest) Validate() error {

	return nil
}

func (m *SanitizeDecryptedRequest) Validate() error {

	return nil
}

func (m *SanitizeResponse) Validate() error {

	return nil
}

func (m *GrantsList) Validate() error {

	return nil
}

func (m *MqttAuthorization) Validate() error {

	return nil
}

func (m *TopicIsAvailable) Validate() error {

	return nil
}
