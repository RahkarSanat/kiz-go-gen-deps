package mongo

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

type FieldOptionsMongo struct {
	Ignore    *bool   `bson:"ignore,omitempty" json:"ignore,omitempty"`
	Required  *bool   `bson:"required,omitempty" json:"required,omitempty"`
	MinLength *int32  `bson:"min_length,omitempty" json:"min_length,omitempty"`
	MaxLength *int32  `bson:"max_length,omitempty" json:"max_length,omitempty"`
	Pattern   *string `bson:"pattern,omitempty" json:"pattern,omitempty"`
}
type FileOptionsMongo struct {
	Ignore    *bool   `bson:"ignore,omitempty" json:"ignore,omitempty"`
	Extension *string `bson:"extension,omitempty" json:"extension,omitempty"`
}
type MessageOptionsMongo struct {
	NotIgnore                    *bool `bson:"not_ignore,omitempty" json:"not_ignore,omitempty"`
	AllFieldsRequired            *bool `bson:"all_fields_required,omitempty" json:"all_fields_required,omitempty"`
	AllowNullValues              *bool `bson:"allow_null_values,omitempty" json:"allow_null_values,omitempty"`
	DisallowAdditionalProperties *bool `bson:"disallow_additional_properties,omitempty" json:"disallow_additional_properties,omitempty"`
	EnumsAsConstants             *bool `bson:"enums_as_constants,omitempty" json:"enums_as_constants,omitempty"`
}
type EnumOptionsMongo struct {
	EnumsAsConstants   *bool `bson:"enums_as_constants,omitempty" json:"enums_as_constants,omitempty"`
	EnumsAsStringsOnly *bool `bson:"enums_as_strings_only,omitempty" json:"enums_as_strings_only,omitempty"`
	EnumsTrimPrefix    *bool `bson:"enums_trim_prefix,omitempty" json:"enums_trim_prefix,omitempty"`
	Ignore             *bool `bson:"ignore,omitempty" json:"ignore,omitempty"`
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

func (m *FieldOptions) Validate() error {

	return nil
}

func (m *FileOptions) Validate() error {

	return nil
}

func (m *MessageOptions) Validate() error {

	return nil
}

func (m *EnumOptions) Validate() error {

	return nil
}
