package base

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var _ primitive.A
var _ fmt.Stringer

type CountRequestMongo struct {
}
type FindRequestMongo struct {
}
type FindByIDRequestMongo struct {
	Id string `bson:"id" json:"id"`
}
type CustomFileOptionsMongo struct {
	Ignore bool   `bson:"ignore" json:"ignore"`
	DbName string `bson:"db_name" json:"db_name"`
}
type MongoOptionMongo struct {
	MongoType string `bson:"mongo_type" json:"mongo_type"`
	BsonTag   string `bson:"bson_tag" json:"bson_tag"`
}
type AggregateOptionMongo struct {
	IsAggregate bool   `bson:"is_aggregate" json:"is_aggregate"`
	Aggregate   string `bson:"aggregate" json:"aggregate"`
}
type ViewOptionMongo struct {
	IsView bool     `bson:"is_view" json:"is_view"`
	Names  []string `bson:"names" json:"names"`
}
type MongoMessageOptionsMongo struct {
	ShouldValidate bool `bson:"should_validate" json:"should_validate"`
}
type MinioOptionMongo struct {
	Ignore bool `bson:"ignore" json:"ignore"`
}
type AclOptionMongo struct {
	Sub string `bson:"sub" json:"sub"`
	Obj string `bson:"obj" json:"obj"`
	Act string `bson:"act" json:"act"`
}
type SortMongo struct {
	Key   string `bson:"key" json:"key"`
	Value int32  `bson:"value" json:"value"`
}
type PaginationMongo struct {
	Limit int32       `bson:"limit" json:"limit"`
	Skip  int64       `bson:"skip" json:"skip"`
	Sort  []SortMongo `bson:"sort" json:"sort"`
}
type FilterMongo struct {
	Pagination      PaginationMongo `bson:"pagination" json:"pagination"`
	Projection      string          `bson:"projection" json:"projection"`
	Query           string          `bson:"query" json:"query"`
	ExcludeOwn      bool            `bson:"exclude_own" json:"exclude_own"`
	ExcludeZone     bool            `bson:"exclude_zone" json:"exclude_zone"`
	ExcludeShare    bool            `bson:"exclude_share" json:"exclude_share"`
	ExcludeRelatoin bool            `bson:"exclude_relatoin" json:"exclude_relatoin"`
	IsTrash         bool            `bson:"is_trash" json:"is_trash"`
}
type OperationsMongo struct {
	AddToSet    string `bson:"add_to_set" json:"add_to_set"`
	Bit         string `bson:"bit" json:"bit"`
	CurrentDate string `bson:"current_date" json:"current_date"`
	Inc         string `bson:"inc" json:"inc"`
	Max         string `bson:"max" json:"max"`
	Min         string `bson:"min" json:"min"`
	Mul         string `bson:"mul" json:"mul"`
	Pop         string `bson:"pop" json:"pop"`
	Pull        string `bson:"pull" json:"pull"`
	PullAll     string `bson:"pull_all" json:"pull_all"`
	Push        string `bson:"push" json:"push"`
	Rename      string `bson:"rename" json:"rename"`
	Set         string `bson:"set" json:"set"`
	SetOnInsert string `bson:"set_on_insert" json:"set_on_insert"`
	Unset       string `bson:"unset" json:"unset"`
}
type TotalMongo struct {
	Count int64 `bson:"count" json:"count"`
}
type BaseAccessMongo struct {
	Id         string              `bson:"id" json:"id"`
	Owner      string              `bson:"owner" json:"owner"`
	Zones      []string            `bson:"zones" json:"zones"`
	Clients    []string            `bson:"clients" json:"clients"`
	Shares     []string            `bson:"shares" json:"shares"`
	Relations  []string            `bson:"relations" json:"relations"`
	CreatedBy  *primitive.ObjectID `bson:"created_by" json:"created_by"`
	UpdatedBy  *primitive.ObjectID `bson:"updated_by" json:"updated_by"`
	DeletedBy  *primitive.ObjectID `bson:"deleted_by" json:"deleted_by"`
	RestoredBy *primitive.ObjectID `bson:"restored_by" json:"restored_by"`
}
type BasePropertiesMongo struct {
	Id         string              `bson:"id" json:"id"`
	CreatedIn  *primitive.ObjectID `bson:"created_in" json:"created_in"`
	UpdatedIn  *primitive.ObjectID `bson:"updated_in" json:"updated_in"`
	DeletedIn  *primitive.ObjectID `bson:"deleted_in" json:"deleted_in"`
	RestoredIn *primitive.ObjectID `bson:"restored_in" json:"restored_in"`
	Tags       []string            `bson:"tags" json:"tags"`
	Version    string              `bson:"version" json:"version"`
}
type BaseDatesMongo struct {
	Id         string              `bson:"id" json:"id"`
	CreatedAt  *primitive.DateTime `bson:"created_at" json:"created_at"`
	UpdatedAt  *primitive.DateTime `bson:"updated_at" json:"updated_at"`
	DeletedAt  *primitive.DateTime `bson:"deleted_at" json:"deleted_at"`
	RestoredAt *primitive.DateTime `bson:"restored_at" json:"restored_at"`
}
type BaseAccessUpdateMongo struct {
	Owner     string   `bson:"owner" json:"owner"`
	Zones     []string `bson:"zones" json:"zones"`
	Shares    []string `bson:"shares" json:"shares"`
	Relations []string `bson:"relations" json:"relations"`
}
type BasePropertiesUpdateMongo struct {
	Tags []string `bson:"tags" json:"tags"`
}
type BaseDatesUpdateMongo struct {
}
type BaseAccessReadMongo struct {
	Owner      string              `bson:"owner" json:"owner"`
	Zones      []string            `bson:"zones" json:"zones"`
	Clients    []string            `bson:"clients" json:"clients"`
	Shares     []string            `bson:"shares" json:"shares"`
	Relations  []string            `bson:"relations" json:"relations"`
	CreatedBy  *primitive.ObjectID `bson:"created_by" json:"created_by"`
	UpdatedBy  *primitive.ObjectID `bson:"updated_by" json:"updated_by"`
	DeletedBy  *primitive.ObjectID `bson:"deleted_by" json:"deleted_by"`
	RestoredBy *primitive.ObjectID `bson:"restored_by" json:"restored_by"`
}
type BasePropertiesReadMongo struct {
	CreatedIn  *primitive.ObjectID `bson:"created_in" json:"created_in"`
	UpdatedIn  *primitive.ObjectID `bson:"updated_in" json:"updated_in"`
	DeletedIn  *primitive.ObjectID `bson:"deleted_in" json:"deleted_in"`
	RestoredIn *primitive.ObjectID `bson:"restored_in" json:"restored_in"`
	Tags       []string            `bson:"tags" json:"tags"`
	Version    string              `bson:"version" json:"version"`
}
type BaseDatesReadMongo struct {
	CreatedAt  *primitive.DateTime `bson:"created_at" json:"created_at"`
	UpdatedAt  *primitive.DateTime `bson:"updated_at" json:"updated_at"`
	DeletedAt  *primitive.DateTime `bson:"deleted_at" json:"deleted_at"`
	RestoredAt *primitive.DateTime `bson:"restored_at" json:"restored_at"`
}

func (m *CountRequest) Validate() error {

	return nil
}

func (m *FindRequest) Validate() error {

	return nil
}

func (m *FindByIDRequest) Validate() error {

	if m.Id == "" {
		return fmt.Errorf("Id is required")
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
		return fmt.Errorf("Owner is required")
	}

	if m.Clients == nil {
		return fmt.Errorf("Clients is required")
	}

	// ali

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
