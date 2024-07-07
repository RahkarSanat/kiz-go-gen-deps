package authentication

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

type VoidMongo struct {
	Result *string `bson:"result,omitempty" json:"result,omitempty"`
}
type AccessTokenMongo struct {
	AccessToken *string `bson:"access_token,omitempty" json:"access_token,omitempty"`
	ApiKey      *string `bson:"api_key,omitempty" json:"api_key,omitempty"`
}
type AuthRequestMongo struct {
	GrantType        *string  `bson:"grant_type,omitempty" json:"grant_type,omitempty"`
	AppId            *string  `bson:"app_id,omitempty" json:"app_id,omitempty"`
	Subjects         []string `bson:"subjects,omitempty" json:"subjects,omitempty"`
	ClientId         *string  `bson:"client_id,omitempty" json:"client_id,omitempty"`
	ClientSecret     *string  `bson:"client_secret,omitempty" json:"client_secret,omitempty"`
	Phone            *string  `bson:"phone,omitempty" json:"phone,omitempty"`
	Email            *string  `bson:"email,omitempty" json:"email,omitempty"`
	Username         *string  `bson:"username,omitempty" json:"username,omitempty"`
	Password         *string  `bson:"password,omitempty" json:"password,omitempty"`
	ConfirmationCode *string  `bson:"confirmation_code,omitempty" json:"confirmation_code,omitempty"`
	RefreshToken     *string  `bson:"refresh_token,omitempty" json:"refresh_token,omitempty"`
	Scopes           []string `bson:"scopes,omitempty" json:"scopes,omitempty"`
	Code             *string  `bson:"code,omitempty" json:"code,omitempty"`
}
type AuthResponseMongo struct {
	AccessToken  *string  `bson:"access_token,omitempty" json:"access_token,omitempty"`
	RefreshToken *string  `bson:"refresh_token,omitempty" json:"refresh_token,omitempty"`
	TokenType    *string  `bson:"token_type,omitempty" json:"token_type,omitempty"`
	Code         *string  `bson:"code,omitempty" json:"code,omitempty"`
	ExpiresIn    *string  `bson:"expires_in,omitempty" json:"expires_in,omitempty"`
	Session      *string  `bson:"session,omitempty" json:"session,omitempty"`
	Scopes       []string `bson:"scopes,omitempty" json:"scopes,omitempty"`
	Subjects     []string `bson:"subjects,omitempty" json:"subjects,omitempty"`
	Uid          *string  `bson:"uid,omitempty" json:"uid,omitempty"`
	Registered   *bool    `bson:"registered,omitempty" json:"registered,omitempty"`
}
type AuthTokenMongo struct {
	Cid      *string `bson:"cid,omitempty" json:"cid,omitempty"`
	Aid      *string `bson:"aid,omitempty" json:"aid,omitempty"`
	Uid      *string `bson:"uid,omitempty" json:"uid,omitempty"`
	Subjects *string `bson:"subjects,omitempty" json:"subjects,omitempty"`
	Scopes   *string `bson:"scopes,omitempty" json:"scopes,omitempty"`
	Session  *string `bson:"session,omitempty" json:"session,omitempty"`
	ClientId *string `bson:"client_id,omitempty" json:"client_id,omitempty"`
	Iat      *int64  `bson:"iat,omitempty" json:"iat,omitempty"`
	Exp      *int64  `bson:"exp,omitempty" json:"exp,omitempty"`
}
type CodeMongo struct {
	Code *string `bson:"code,omitempty" json:"code,omitempty"`
}
type NewPasswordRequestMongo struct {
	ConfirmationCode *string `bson:"confirmation_code,omitempty" json:"confirmation_code,omitempty"`
	Code             *string `bson:"code,omitempty" json:"code,omitempty"`
	NewPassword      *string `bson:"new_password,omitempty" json:"new_password,omitempty"`
}
type MqttAuthenticationMongo struct {
	AccessToken   *string `bson:"access_token,omitempty" json:"access_token,omitempty"`
	ActiveSession *string `bson:"active_session,omitempty" json:"active_session,omitempty"`
	MqttUsername  *string `bson:"mqtt_username,omitempty" json:"mqtt_username,omitempty"`
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

func (m *Void) Validate() error {

	return nil
}

func (m *AccessToken) Validate() error {

	return nil
}

func (m *AuthRequest) Validate() error {

	return nil
}

func (m *AuthResponse) Validate() error {

	return nil
}

func (m *AuthToken) Validate() error {

	return nil
}

func (m *Code) Validate() error {

	return nil
}

func (m *NewPasswordRequest) Validate() error {

	return nil
}

func (m *MqttAuthentication) Validate() error {

	return nil
}
