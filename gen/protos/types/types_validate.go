package types

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

type UUIDValueMongo struct {
	Value *string `bson:"value,omitempty" json:"value,omitempty"`
}
type JSONValueMongo struct {
	Value *string `bson:"value,omitempty" json:"value,omitempty"`
}
type UUIDMongo struct {
	Value *string `bson:"value,omitempty" json:"value,omitempty"`
}
type InetValueMongo struct {
	Value *string `bson:"value,omitempty" json:"value,omitempty"`
}
type TimeOnlyMongo struct {
	Value *uint32 `bson:"value,omitempty" json:"value,omitempty"`
}
type BigIntMongo struct {
	Value *string `bson:"value,omitempty" json:"value,omitempty"`
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

func (m *UUIDValue) Validate() error {

	return nil
}

func (m *JSONValue) Validate() error {

	return nil
}

func (m *UUID) Validate() error {

	return nil
}

func (m *InetValue) Validate() error {

	return nil
}

func (m *TimeOnly) Validate() error {

	return nil
}

func (m *BigInt) Validate() error {

	return nil
}
