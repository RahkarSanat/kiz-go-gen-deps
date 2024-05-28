package storage

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

type ReadObjectRequestMongo struct {
	Prefix *string `bson:"prefix,omitempty" json:"prefix,omitempty"`
	Object *string `bson:"object,omitempty" json:"object,omitempty"`
	Title  *string `bson:"title,omitempty" json:"title,omitempty"`
}
type WriteObjectRequestMongo struct {
	FirstMessage *FirstMessageMongo    `bson:"first_message,omitempty" json:"first_message,omitempty"`
	Data         *ChecksummedDataMongo `bson:"data,omitempty" json:"data,omitempty"`
	FinishWrite  *bool                 `bson:"finish_write,omitempty" json:"finish_write,omitempty"`
}
type FirstMessageMongo struct {
	Prefix *string `bson:"prefix,omitempty" json:"prefix,omitempty"`
	Name   *string `bson:"name,omitempty" json:"name,omitempty"`
}
type ReadObjectResponseMongo struct {
	FirstMessage    *FirstMessageMongo    `bson:"first_message,omitempty" json:"first_message,omitempty"`
	ChecksummedData *ChecksummedDataMongo `bson:"checksummed_data,omitempty" json:"checksummed_data,omitempty"`
}
type WriteObjectResponseMongo struct {
	PersistedSize *int64 `bson:"persisted_size,omitempty" json:"persisted_size,omitempty"`
}
type ChecksummedDataMongo struct {
	Content *[]byte `bson:"content,omitempty" json:"content,omitempty"`
	Crc32C  *uint32 `bson:"crc32c,omitempty" json:"crc32c,omitempty"`
}
type DeleteObjectRequestMongo struct {
	Bucket *string `bson:"bucket,omitempty" json:"bucket,omitempty"`
	Object *string `bson:"object,omitempty" json:"object,omitempty"`
}
type DeleteObjectResponseMongo struct {
	Status *bool `bson:"status,omitempty" json:"status,omitempty"`
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
		if obj != nil && !isZero(obj) {
			flatMap[prefix[:len(prefix)-1]] = obj
		}
	}
}

func (m *ReadObjectRequest) Validate() error {

	return nil
}

func (m *WriteObjectRequest) Validate() error {

	return nil
}

func (m *FirstMessage) Validate() error {

	return nil
}

func (m *ReadObjectResponse) Validate() error {

	return nil
}

func (m *WriteObjectResponse) Validate() error {

	return nil
}

func (m *ChecksummedData) Validate() error {

	return nil
}

func (m *DeleteObjectRequest) Validate() error {

	return nil
}

func (m *DeleteObjectResponse) Validate() error {

	return nil
}
