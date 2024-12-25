package base

import (
	"fmt"
	"encoding/json"
	reflect "reflect"
	"strings"

	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

var _ codes.Code
var _ status.Status
var _ reflect.Kind
var _ primitive.A
var _ fmt.Stringer

var _ json.Marshaler
var _ strings.Builder

type CountRequestMongo struct {
	Filter *FilterMongo `bson:"filter,omitempty" json:"filter,omitempty"`
}
type FindRequestMongo struct {
	Filter *FilterMongo `bson:"filter,omitempty" json:"filter,omitempty"`
}
type FindByIdArchiveRequestMongo struct {
	Id      string       `bson:"id" json:"id"`
	Version *int32       `bson:"version,omitempty" json:"version,omitempty"`
	Filter  *FilterMongo `bson:"filter,omitempty" json:"filter,omitempty"`
	Queries []QueryMongo `bson:"queries,omitempty" json:"queries,omitempty"`
}
type FindByIdArchiveResponseMongo struct {
	Id        *string  `bson:"id,omitempty" json:"id,omitempty"`
	Owner     *string  `bson:"owner,omitempty" json:"owner,omitempty"`
	Clients   []string `bson:"clients,omitempty" json:"clients,omitempty"`
	Zones     []string `bson:"zones,omitempty" json:"zones,omitempty"`
	Relations []string `bson:"relations,omitempty" json:"relations,omitempty"`
	Shares    []string `bson:"shares,omitempty" json:"shares,omitempty"`
	CreatedBy *string  `bson:"created_by,omitempty" json:"created_by,omitempty"`
	CreatedAt *uint64  `bson:"created_at,omitempty" json:"created_at,omitempty"`
	CreatedIn *string  `bson:"created_in,omitempty" json:"created_in,omitempty"`
	Version   *int32   `bson:"version,omitempty" json:"version,omitempty"`
	Data      *[]byte  `bson:"data,omitempty" json:"data,omitempty"`
}
type FindByIDRequestMongo struct {
	Id     string       `bson:"id" json:"id"`
	Filter *FilterMongo `bson:"filter,omitempty" json:"filter,omitempty"`
}
type CustomFileOptionsMongo struct {
	Ignore *bool   `bson:"ignore,omitempty" json:"ignore,omitempty"`
	DbName *string `bson:"db_name,omitempty" json:"db_name,omitempty"`
}
type MongoOptionMongo struct {
	MongoType *string `bson:"mongo_type,omitempty" json:"mongo_type,omitempty"`
	BsonTag   *string `bson:"bson_tag,omitempty" json:"bson_tag,omitempty"`
}
type AggregateOptionMongo struct {
	IsAggregate *bool   `bson:"is_aggregate,omitempty" json:"is_aggregate,omitempty"`
	Aggregate   *string `bson:"aggregate,omitempty" json:"aggregate,omitempty"`
}
type ViewOptionMongo struct {
	IsView *bool    `bson:"is_view,omitempty" json:"is_view,omitempty"`
	Names  []string `bson:"names,omitempty" json:"names,omitempty"`
}
type MongoMessageOptionsMongo struct {
	ShouldValidate *bool `bson:"should_validate,omitempty" json:"should_validate,omitempty"`
}
type MinioOptionMongo struct {
	Ignore *bool `bson:"ignore,omitempty" json:"ignore,omitempty"`
}
type AclOptionMongo struct {
	Sub *string `bson:"sub,omitempty" json:"sub,omitempty"`
	Obj *string `bson:"obj,omitempty" json:"obj,omitempty"`
	Act *string `bson:"act,omitempty" json:"act,omitempty"`
}
type SortMongo struct {
	Key   *string `bson:"key,omitempty" json:"key,omitempty"`
	Value *int32  `bson:"value,omitempty" json:"value,omitempty"`
}
type PaginationMongo struct {
	Limit *int32      `bson:"limit,omitempty" json:"limit,omitempty"`
	Skip  *int64      `bson:"skip,omitempty" json:"skip,omitempty"`
	Sort  []SortMongo `bson:"sort,omitempty" json:"sort,omitempty"`
}
type QueryMongo struct {
	Key   *string `bson:"key,omitempty" json:"key,omitempty"`
	Value *string `bson:"value,omitempty" json:"value,omitempty"`
	Op    *string `bson:"op,omitempty" json:"op,omitempty"`
	Func  *string `bson:"func,omitempty" json:"func,omitempty"`
}
type FilterMongo struct {
	Pagination      *PaginationMongo `bson:"pagination,omitempty" json:"pagination,omitempty"`
	Projection      *string          `bson:"projection,omitempty" json:"projection,omitempty"`
	Query           *string          `bson:"query,omitempty" json:"query,omitempty"`
	ExcludeOwn      *bool            `bson:"exclude_own,omitempty" json:"exclude_own,omitempty"`
	ExcludeZone     *bool            `bson:"exclude_zone,omitempty" json:"exclude_zone,omitempty"`
	ExcludeShare    *bool            `bson:"exclude_share,omitempty" json:"exclude_share,omitempty"`
	ExcludeRelation *bool            `bson:"exclude_relation,omitempty" json:"exclude_relation,omitempty"`
	IsTrash         *bool            `bson:"is_trash,omitempty" json:"is_trash,omitempty"`
}
type OperationsMongo struct {
	AddToSet    *string `bson:"add_to_set,omitempty" json:"add_to_set,omitempty"`
	Bit         *string `bson:"bit,omitempty" json:"bit,omitempty"`
	CurrentDate *string `bson:"current_date,omitempty" json:"current_date,omitempty"`
	Inc         *string `bson:"inc,omitempty" json:"inc,omitempty"`
	Max         *string `bson:"max,omitempty" json:"max,omitempty"`
	Min         *string `bson:"min,omitempty" json:"min,omitempty"`
	Mul         *string `bson:"mul,omitempty" json:"mul,omitempty"`
	Pop         *string `bson:"pop,omitempty" json:"pop,omitempty"`
	Pull        *string `bson:"pull,omitempty" json:"pull,omitempty"`
	PullAll     *string `bson:"pull_all,omitempty" json:"pull_all,omitempty"`
	Push        *string `bson:"push,omitempty" json:"push,omitempty"`
	Rename      *string `bson:"rename,omitempty" json:"rename,omitempty"`
	Set         *string `bson:"set,omitempty" json:"set,omitempty"`
	SetOnInsert *string `bson:"set_on_insert,omitempty" json:"set_on_insert,omitempty"`
	Unset       *string `bson:"unset,omitempty" json:"unset,omitempty"`
}
type TotalMongo struct {
	Count *int64 `bson:"count,omitempty" json:"count,omitempty"`
}
type BaseAccessMongo struct {
	Id         *string             `bson:"id,omitempty" json:"id,omitempty"`
	Owner      string              `bson:"owner" json:"owner"`
	Zones      []string            `bson:"zones,omitempty" json:"zones,omitempty"`
	Clients    []string            `bson:"clients" json:"clients"`
	Shares     []string            `bson:"shares,omitempty" json:"shares,omitempty"`
	Relations  []string            `bson:"relations,omitempty" json:"relations,omitempty"`
	CreatedBy  *primitive.ObjectID `bson:"created_by,omitempty" json:"created_by,omitempty"`
	UpdatedBy  *primitive.ObjectID `bson:"updated_by,omitempty" json:"updated_by,omitempty"`
	DeletedBy  *primitive.ObjectID `bson:"deleted_by,omitempty" json:"deleted_by,omitempty"`
	RestoredBy *primitive.ObjectID `bson:"restored_by,omitempty" json:"restored_by,omitempty"`
}
type BasePropertiesMongo struct {
	Id         *string             `bson:"id,omitempty" json:"id,omitempty"`
	CreatedIn  *primitive.ObjectID `bson:"created_in,omitempty" json:"created_in,omitempty"`
	UpdatedIn  *primitive.ObjectID `bson:"updated_in,omitempty" json:"updated_in,omitempty"`
	DeletedIn  *primitive.ObjectID `bson:"deleted_in,omitempty" json:"deleted_in,omitempty"`
	RestoredIn *primitive.ObjectID `bson:"restored_in,omitempty" json:"restored_in,omitempty"`
	Tags       []string            `bson:"tags,omitempty" json:"tags,omitempty"`
	Version    *int32              `bson:"version,omitempty" json:"version,omitempty"`
}
type BaseDatesMongo struct {
	Id         *string             `bson:"id,omitempty" json:"id,omitempty"`
	CreatedAt  *primitive.DateTime `bson:"created_at,omitempty" json:"created_at,omitempty"`
	UpdatedAt  *primitive.DateTime `bson:"updated_at,omitempty" json:"updated_at,omitempty"`
	DeletedAt  *primitive.DateTime `bson:"deleted_at,omitempty" json:"deleted_at,omitempty"`
	RestoredAt *primitive.DateTime `bson:"restored_at,omitempty" json:"restored_at,omitempty"`
}
type BaseAccessUpdateMongo struct {
	Owner     *string  `bson:"owner,omitempty" json:"owner,omitempty"`
	Zones     []string `bson:"zones,omitempty" json:"zones,omitempty"`
	Shares    []string `bson:"shares,omitempty" json:"shares,omitempty"`
	Relations []string `bson:"relations,omitempty" json:"relations,omitempty"`
}
type BasePropertiesUpdateMongo struct {
	Tags []string `bson:"tags,omitempty" json:"tags,omitempty"`
}
type BaseDatesUpdateMongo struct {
}
type BaseAccessReadMongo struct {
	Owner      *string             `bson:"owner,omitempty" json:"owner,omitempty"`
	Zones      []string            `bson:"zones,omitempty" json:"zones,omitempty"`
	Clients    []string            `bson:"clients,omitempty" json:"clients,omitempty"`
	Shares     []string            `bson:"shares,omitempty" json:"shares,omitempty"`
	Relations  []string            `bson:"relations,omitempty" json:"relations,omitempty"`
	CreatedBy  *primitive.ObjectID `bson:"created_by,omitempty" json:"created_by,omitempty"`
	UpdatedBy  *primitive.ObjectID `bson:"updated_by,omitempty" json:"updated_by,omitempty"`
	DeletedBy  *primitive.ObjectID `bson:"deleted_by,omitempty" json:"deleted_by,omitempty"`
	RestoredBy *primitive.ObjectID `bson:"restored_by,omitempty" json:"restored_by,omitempty"`
}
type BasePropertiesReadMongo struct {
	CreatedIn  *primitive.ObjectID `bson:"created_in,omitempty" json:"created_in,omitempty"`
	UpdatedIn  *primitive.ObjectID `bson:"updated_in,omitempty" json:"updated_in,omitempty"`
	DeletedIn  *primitive.ObjectID `bson:"deleted_in,omitempty" json:"deleted_in,omitempty"`
	RestoredIn *primitive.ObjectID `bson:"restored_in,omitempty" json:"restored_in,omitempty"`
	Tags       []string            `bson:"tags,omitempty" json:"tags,omitempty"`
	Version    *int32              `bson:"version,omitempty" json:"version,omitempty"`
}
type BaseDatesReadMongo struct {
	CreatedAt  *primitive.DateTime `bson:"created_at,omitempty" json:"created_at,omitempty"`
	UpdatedAt  *primitive.DateTime `bson:"updated_at,omitempty" json:"updated_at,omitempty"`
	DeletedAt  *primitive.DateTime `bson:"deleted_at,omitempty" json:"deleted_at,omitempty"`
	RestoredAt *primitive.DateTime `bson:"restored_at,omitempty" json:"restored_at,omitempty"`
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

func (m *CountRequest) Validate() error {

	return nil
}

func (m *FindRequest) Validate() error {

	return nil
}

func (m *FindByIdArchiveRequest) Validate() error {

	if m.Id == "" {
		err := status.Error(codes.InvalidArgument, "Id is required")
		return err
	}

	return nil
}

func (m *FindByIdArchiveResponse) Validate() error {

	return nil
}

func (m *FindByIDRequest) Validate() error {

	if m.Id == "" {
		err := status.Error(codes.InvalidArgument, "Id is required")
		return err
	}

	return nil
}

func (m *CustomFileOptions) Validate() error {

	return nil
}

func (m *MongoOption) Validate() error {

	return nil
}

func (m *AggregateOption) Validate() error {

	return nil
}

func (m *ViewOption) Validate() error {

	return nil
}

func (m *MongoMessageOptions) Validate() error {

	return nil
}

func (m *MinioOption) Validate() error {

	return nil
}

func (m *AclOption) Validate() error {

	return nil
}

func (m *Sort) Validate() error {

	return nil
}

func (m *Pagination) Validate() error {

	return nil
}

func (m *Query) Validate() error {

	return nil
}

func (m *Filter) Validate() error {

	return nil
}

func (m *Operations) Validate() error {

	return nil
}

func (m *Total) Validate() error {

	return nil
}

func (m *BaseAccess) Validate() error {

	if m.Owner == "" {
		err := status.Error(codes.InvalidArgument, "Owner is required")
		return err
	}

	if m.Clients == nil {
		err := status.Error(codes.InvalidArgument, "Clients is required")
		return err
	}

	return nil
}

func (m *BaseProperties) Validate() error {

	return nil
}

func (m *BaseDates) Validate() error {

	return nil
}

func (m *BaseAccessUpdate) Validate() error {

	return nil
}

func (m *BasePropertiesUpdate) Validate() error {

	return nil
}

func (m *BaseDatesUpdate) Validate() error {

	return nil
}

func (m *BaseAccessRead) Validate() error {

	return nil
}

func (m *BasePropertiesRead) Validate() error {

	return nil
}

func (m *BaseDatesRead) Validate() error {

	return nil
}
