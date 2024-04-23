package kiz

import (
	"fmt"
	"encoding/json"
	reflect "reflect"
	"strings"

	base "github.com/RahkarSanat/kiz-go-gen-deps/gen/protos/base"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

var _ reflect.Kind
var _ primitive.A
var _ fmt.Stringer

var _ base.BaseAccess

var _ json.Marshaler
var _ strings.Builder

type InterceptorOptionsMongo struct {
	ForceAll *bool `bson:"force_all,omitempty" json:"force_all,omitempty"`
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
	default:
		fmt.Println(obj)
		if obj != nil && !isZero(obj) {
			flatMap[prefix[:len(prefix)-1]] = obj
		}
	}
}

func (m *InterceptorOptions) Validate() error {

	return nil
}
