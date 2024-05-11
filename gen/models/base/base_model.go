package models

type CountRequest struct {
	Filter Filter ``
}

type FindRequest struct {
	Filter Filter ``

	Filter Filter ``
}

type FindByIDRequest struct {
	Filter Filter ``

	Filter Filter ``

	// NOT MESSAGE
	Id string ``

	Filter Filter ``
}

type CustomFileOptions struct {
	Filter Filter ``

	Filter Filter ``

	// NOT MESSAGE
	Id string ``

	Filter Filter ``

	// NOT MESSAGE
	Ignore bool ``

	// NOT MESSAGE
	DbName string ``
}

type MongoOption struct {
	Filter Filter ``

	Filter Filter ``

	// NOT MESSAGE
	Id string ``

	Filter Filter ``

	// NOT MESSAGE
	Ignore bool ``

	// NOT MESSAGE
	DbName string ``

	// NOT MESSAGE
	MongoType string ``

	// NOT MESSAGE
	BsonTag string ``
}

type AggregateOption struct {
	Filter Filter ``

	Filter Filter ``

	// NOT MESSAGE
	Id string ``

	Filter Filter ``

	// NOT MESSAGE
	Ignore bool ``

	// NOT MESSAGE
	DbName string ``

	// NOT MESSAGE
	MongoType string ``

	// NOT MESSAGE
	BsonTag string ``

	// NOT MESSAGE
	IsAggregate bool ``

	// NOT MESSAGE
	Aggregate string ``
}

type ViewOption struct {
	Filter Filter ``

	Filter Filter ``

	// NOT MESSAGE
	Id string ``

	Filter Filter ``

	// NOT MESSAGE
	Ignore bool ``

	// NOT MESSAGE
	DbName string ``

	// NOT MESSAGE
	MongoType string ``

	// NOT MESSAGE
	BsonTag string ``

	// NOT MESSAGE
	IsAggregate bool ``

	// NOT MESSAGE
	Aggregate string ``

	// NOT MESSAGE
	IsView bool ``

	// NOT MESSAGE
	Names []string ``
}

type MongoMessageOptions struct {
	Filter Filter ``

	Filter Filter ``

	// NOT MESSAGE
	Id string ``

	Filter Filter ``

	// NOT MESSAGE
	Ignore bool ``

	// NOT MESSAGE
	DbName string ``

	// NOT MESSAGE
	MongoType string ``

	// NOT MESSAGE
	BsonTag string ``

	// NOT MESSAGE
	IsAggregate bool ``

	// NOT MESSAGE
	Aggregate string ``

	// NOT MESSAGE
	IsView bool ``

	// NOT MESSAGE
	Names []string ``

	// NOT MESSAGE
	ShouldValidate bool ``
}

type MinioOption struct {
	Filter Filter ``

	Filter Filter ``

	// NOT MESSAGE
	Id string ``

	Filter Filter ``

	// NOT MESSAGE
	Ignore bool ``

	// NOT MESSAGE
	DbName string ``

	// NOT MESSAGE
	MongoType string ``

	// NOT MESSAGE
	BsonTag string ``

	// NOT MESSAGE
	IsAggregate bool ``

	// NOT MESSAGE
	Aggregate string ``

	// NOT MESSAGE
	IsView bool ``

	// NOT MESSAGE
	Names []string ``

	// NOT MESSAGE
	ShouldValidate bool ``

	// NOT MESSAGE
	Ignore bool ``
}

type AclOption struct {
	Filter Filter ``

	Filter Filter ``

	// NOT MESSAGE
	Id string ``

	Filter Filter ``

	// NOT MESSAGE
	Ignore bool ``

	// NOT MESSAGE
	DbName string ``

	// NOT MESSAGE
	MongoType string ``

	// NOT MESSAGE
	BsonTag string ``

	// NOT MESSAGE
	IsAggregate bool ``

	// NOT MESSAGE
	Aggregate string ``

	// NOT MESSAGE
	IsView bool ``

	// NOT MESSAGE
	Names []string ``

	// NOT MESSAGE
	ShouldValidate bool ``

	// NOT MESSAGE
	Ignore bool ``

	// NOT MESSAGE
	Sub string ``

	// NOT MESSAGE
	Obj string ``

	// NOT MESSAGE
	Act string ``
}

type Sort struct {
	Filter Filter ``

	Filter Filter ``

	// NOT MESSAGE
	Id string ``

	Filter Filter ``

	// NOT MESSAGE
	Ignore bool ``

	// NOT MESSAGE
	DbName string ``

	// NOT MESSAGE
	MongoType string ``

	// NOT MESSAGE
	BsonTag string ``

	// NOT MESSAGE
	IsAggregate bool ``

	// NOT MESSAGE
	Aggregate string ``

	// NOT MESSAGE
	IsView bool ``

	// NOT MESSAGE
	Names []string ``

	// NOT MESSAGE
	ShouldValidate bool ``

	// NOT MESSAGE
	Ignore bool ``

	// NOT MESSAGE
	Sub string ``

	// NOT MESSAGE
	Obj string ``

	// NOT MESSAGE
	Act string ``

	// NOT MESSAGE
	Key string ``

	// NOT MESSAGE
	Value int32 ``
}

type Pagination struct {
	Filter Filter ``

	Filter Filter ``

	// NOT MESSAGE
	Id string ``

	Filter Filter ``

	// NOT MESSAGE
	Ignore bool ``

	// NOT MESSAGE
	DbName string ``

	// NOT MESSAGE
	MongoType string ``

	// NOT MESSAGE
	BsonTag string ``

	// NOT MESSAGE
	IsAggregate bool ``

	// NOT MESSAGE
	Aggregate string ``

	// NOT MESSAGE
	IsView bool ``

	// NOT MESSAGE
	Names []string ``

	// NOT MESSAGE
	ShouldValidate bool ``

	// NOT MESSAGE
	Ignore bool ``

	// NOT MESSAGE
	Sub string ``

	// NOT MESSAGE
	Obj string ``

	// NOT MESSAGE
	Act string ``

	// NOT MESSAGE
	Key string ``

	// NOT MESSAGE
	Value int32 ``

	// NOT MESSAGE
	Limit int32 ``

	// NOT MESSAGE
	Skip int64 ``

	Sort []Sort ``
}

type Filter struct {
	Filter Filter ``

	Filter Filter ``

	// NOT MESSAGE
	Id string ``

	Filter Filter ``

	// NOT MESSAGE
	Ignore bool ``

	// NOT MESSAGE
	DbName string ``

	// NOT MESSAGE
	MongoType string ``

	// NOT MESSAGE
	BsonTag string ``

	// NOT MESSAGE
	IsAggregate bool ``

	// NOT MESSAGE
	Aggregate string ``

	// NOT MESSAGE
	IsView bool ``

	// NOT MESSAGE
	Names []string ``

	// NOT MESSAGE
	ShouldValidate bool ``

	// NOT MESSAGE
	Ignore bool ``

	// NOT MESSAGE
	Sub string ``

	// NOT MESSAGE
	Obj string ``

	// NOT MESSAGE
	Act string ``

	// NOT MESSAGE
	Key string ``

	// NOT MESSAGE
	Value int32 ``

	// NOT MESSAGE
	Limit int32 ``

	// NOT MESSAGE
	Skip int64 ``

	Sort []Sort ``

	Pagination Pagination ``

	// NOT MESSAGE
	Projection string ``

	// NOT MESSAGE
	Query string ``

	// NOT MESSAGE
	ExcludeOwn bool ``

	// NOT MESSAGE
	ExcludeZone bool ``

	// NOT MESSAGE
	ExcludeShare bool ``

	// NOT MESSAGE
	ExcludeRelatoin bool ``

	// NOT MESSAGE
	IsTrash bool ``
}

type Operations struct {
	Filter Filter ``

	Filter Filter ``

	// NOT MESSAGE
	Id string ``

	Filter Filter ``

	// NOT MESSAGE
	Ignore bool ``

	// NOT MESSAGE
	DbName string ``

	// NOT MESSAGE
	MongoType string ``

	// NOT MESSAGE
	BsonTag string ``

	// NOT MESSAGE
	IsAggregate bool ``

	// NOT MESSAGE
	Aggregate string ``

	// NOT MESSAGE
	IsView bool ``

	// NOT MESSAGE
	Names []string ``

	// NOT MESSAGE
	ShouldValidate bool ``

	// NOT MESSAGE
	Ignore bool ``

	// NOT MESSAGE
	Sub string ``

	// NOT MESSAGE
	Obj string ``

	// NOT MESSAGE
	Act string ``

	// NOT MESSAGE
	Key string ``

	// NOT MESSAGE
	Value int32 ``

	// NOT MESSAGE
	Limit int32 ``

	// NOT MESSAGE
	Skip int64 ``

	Sort []Sort ``

	Pagination Pagination ``

	// NOT MESSAGE
	Projection string ``

	// NOT MESSAGE
	Query string ``

	// NOT MESSAGE
	ExcludeOwn bool ``

	// NOT MESSAGE
	ExcludeZone bool ``

	// NOT MESSAGE
	ExcludeShare bool ``

	// NOT MESSAGE
	ExcludeRelatoin bool ``

	// NOT MESSAGE
	IsTrash bool ``

	// NOT MESSAGE
	AddToSet string ``

	// NOT MESSAGE
	Bit string ``

	// NOT MESSAGE
	CurrentDate string ``

	// NOT MESSAGE
	Inc string ``

	// NOT MESSAGE
	Max string ``

	// NOT MESSAGE
	Min string ``

	// NOT MESSAGE
	Mul string ``

	// NOT MESSAGE
	Pop string ``

	// NOT MESSAGE
	Pull string ``

	// NOT MESSAGE
	PullAll string ``

	// NOT MESSAGE
	Push string ``

	// NOT MESSAGE
	Rename string ``

	// NOT MESSAGE
	Set string ``

	// NOT MESSAGE
	SetOnInsert string ``

	// NOT MESSAGE
	Unset string ``
}

type Total struct {
	Filter Filter ``

	Filter Filter ``

	// NOT MESSAGE
	Id string ``

	Filter Filter ``

	// NOT MESSAGE
	Ignore bool ``

	// NOT MESSAGE
	DbName string ``

	// NOT MESSAGE
	MongoType string ``

	// NOT MESSAGE
	BsonTag string ``

	// NOT MESSAGE
	IsAggregate bool ``

	// NOT MESSAGE
	Aggregate string ``

	// NOT MESSAGE
	IsView bool ``

	// NOT MESSAGE
	Names []string ``

	// NOT MESSAGE
	ShouldValidate bool ``

	// NOT MESSAGE
	Ignore bool ``

	// NOT MESSAGE
	Sub string ``

	// NOT MESSAGE
	Obj string ``

	// NOT MESSAGE
	Act string ``

	// NOT MESSAGE
	Key string ``

	// NOT MESSAGE
	Value int32 ``

	// NOT MESSAGE
	Limit int32 ``

	// NOT MESSAGE
	Skip int64 ``

	Sort []Sort ``

	Pagination Pagination ``

	// NOT MESSAGE
	Projection string ``

	// NOT MESSAGE
	Query string ``

	// NOT MESSAGE
	ExcludeOwn bool ``

	// NOT MESSAGE
	ExcludeZone bool ``

	// NOT MESSAGE
	ExcludeShare bool ``

	// NOT MESSAGE
	ExcludeRelatoin bool ``

	// NOT MESSAGE
	IsTrash bool ``

	// NOT MESSAGE
	AddToSet string ``

	// NOT MESSAGE
	Bit string ``

	// NOT MESSAGE
	CurrentDate string ``

	// NOT MESSAGE
	Inc string ``

	// NOT MESSAGE
	Max string ``

	// NOT MESSAGE
	Min string ``

	// NOT MESSAGE
	Mul string ``

	// NOT MESSAGE
	Pop string ``

	// NOT MESSAGE
	Pull string ``

	// NOT MESSAGE
	PullAll string ``

	// NOT MESSAGE
	Push string ``

	// NOT MESSAGE
	Rename string ``

	// NOT MESSAGE
	Set string ``

	// NOT MESSAGE
	SetOnInsert string ``

	// NOT MESSAGE
	Unset string ``

	// NOT MESSAGE
	Count int64 ``
}

type BaseAccess struct {
	Filter Filter ``

	Filter Filter ``

	// NOT MESSAGE
	Id string ``

	Filter Filter ``

	// NOT MESSAGE
	Ignore bool ``

	// NOT MESSAGE
	DbName string ``

	// NOT MESSAGE
	MongoType string ``

	// NOT MESSAGE
	BsonTag string ``

	// NOT MESSAGE
	IsAggregate bool ``

	// NOT MESSAGE
	Aggregate string ``

	// NOT MESSAGE
	IsView bool ``

	// NOT MESSAGE
	Names []string ``

	// NOT MESSAGE
	ShouldValidate bool ``

	// NOT MESSAGE
	Ignore bool ``

	// NOT MESSAGE
	Sub string ``

	// NOT MESSAGE
	Obj string ``

	// NOT MESSAGE
	Act string ``

	// NOT MESSAGE
	Key string ``

	// NOT MESSAGE
	Value int32 ``

	// NOT MESSAGE
	Limit int32 ``

	// NOT MESSAGE
	Skip int64 ``

	Sort []Sort ``

	Pagination Pagination ``

	// NOT MESSAGE
	Projection string ``

	// NOT MESSAGE
	Query string ``

	// NOT MESSAGE
	ExcludeOwn bool ``

	// NOT MESSAGE
	ExcludeZone bool ``

	// NOT MESSAGE
	ExcludeShare bool ``

	// NOT MESSAGE
	ExcludeRelatoin bool ``

	// NOT MESSAGE
	IsTrash bool ``

	// NOT MESSAGE
	AddToSet string ``

	// NOT MESSAGE
	Bit string ``

	// NOT MESSAGE
	CurrentDate string ``

	// NOT MESSAGE
	Inc string ``

	// NOT MESSAGE
	Max string ``

	// NOT MESSAGE
	Min string ``

	// NOT MESSAGE
	Mul string ``

	// NOT MESSAGE
	Pop string ``

	// NOT MESSAGE
	Pull string ``

	// NOT MESSAGE
	PullAll string ``

	// NOT MESSAGE
	Push string ``

	// NOT MESSAGE
	Rename string ``

	// NOT MESSAGE
	Set string ``

	// NOT MESSAGE
	SetOnInsert string ``

	// NOT MESSAGE
	Unset string ``

	// NOT MESSAGE
	Count int64 ``

	// NOT MESSAGE
	Id string ``

	// NOT MESSAGE
	Owner string ``

	// NOT MESSAGE
	Zones []string ``

	// NOT MESSAGE
	Clients []string ``

	// NOT MESSAGE
	Shares []string ``

	// NOT MESSAGE
	Relations []string ``

	// NOT MESSAGE
	CreatedBy string ``

	// NOT MESSAGE
	UpdatedBy string ``

	// NOT MESSAGE
	DeletedBy string ``

	// NOT MESSAGE
	RestoredBy string ``
}

type BaseProperties struct {
	Filter Filter ``

	Filter Filter ``

	// NOT MESSAGE
	Id string ``

	Filter Filter ``

	// NOT MESSAGE
	Ignore bool ``

	// NOT MESSAGE
	DbName string ``

	// NOT MESSAGE
	MongoType string ``

	// NOT MESSAGE
	BsonTag string ``

	// NOT MESSAGE
	IsAggregate bool ``

	// NOT MESSAGE
	Aggregate string ``

	// NOT MESSAGE
	IsView bool ``

	// NOT MESSAGE
	Names []string ``

	// NOT MESSAGE
	ShouldValidate bool ``

	// NOT MESSAGE
	Ignore bool ``

	// NOT MESSAGE
	Sub string ``

	// NOT MESSAGE
	Obj string ``

	// NOT MESSAGE
	Act string ``

	// NOT MESSAGE
	Key string ``

	// NOT MESSAGE
	Value int32 ``

	// NOT MESSAGE
	Limit int32 ``

	// NOT MESSAGE
	Skip int64 ``

	Sort []Sort ``

	Pagination Pagination ``

	// NOT MESSAGE
	Projection string ``

	// NOT MESSAGE
	Query string ``

	// NOT MESSAGE
	ExcludeOwn bool ``

	// NOT MESSAGE
	ExcludeZone bool ``

	// NOT MESSAGE
	ExcludeShare bool ``

	// NOT MESSAGE
	ExcludeRelatoin bool ``

	// NOT MESSAGE
	IsTrash bool ``

	// NOT MESSAGE
	AddToSet string ``

	// NOT MESSAGE
	Bit string ``

	// NOT MESSAGE
	CurrentDate string ``

	// NOT MESSAGE
	Inc string ``

	// NOT MESSAGE
	Max string ``

	// NOT MESSAGE
	Min string ``

	// NOT MESSAGE
	Mul string ``

	// NOT MESSAGE
	Pop string ``

	// NOT MESSAGE
	Pull string ``

	// NOT MESSAGE
	PullAll string ``

	// NOT MESSAGE
	Push string ``

	// NOT MESSAGE
	Rename string ``

	// NOT MESSAGE
	Set string ``

	// NOT MESSAGE
	SetOnInsert string ``

	// NOT MESSAGE
	Unset string ``

	// NOT MESSAGE
	Count int64 ``

	// NOT MESSAGE
	Id string ``

	// NOT MESSAGE
	Owner string ``

	// NOT MESSAGE
	Zones []string ``

	// NOT MESSAGE
	Clients []string ``

	// NOT MESSAGE
	Shares []string ``

	// NOT MESSAGE
	Relations []string ``

	// NOT MESSAGE
	CreatedBy string ``

	// NOT MESSAGE
	UpdatedBy string ``

	// NOT MESSAGE
	DeletedBy string ``

	// NOT MESSAGE
	RestoredBy string ``

	// NOT MESSAGE
	Id string ``

	// NOT MESSAGE
	CreatedIn string ``

	// NOT MESSAGE
	UpdatedIn string ``

	// NOT MESSAGE
	DeletedIn string ``

	// NOT MESSAGE
	RestoredIn string ``

	// NOT MESSAGE
	Tags []string ``

	// NOT MESSAGE
	Version string ``
}

type BaseDates struct {
	Filter Filter ``

	Filter Filter ``

	// NOT MESSAGE
	Id string ``

	Filter Filter ``

	// NOT MESSAGE
	Ignore bool ``

	// NOT MESSAGE
	DbName string ``

	// NOT MESSAGE
	MongoType string ``

	// NOT MESSAGE
	BsonTag string ``

	// NOT MESSAGE
	IsAggregate bool ``

	// NOT MESSAGE
	Aggregate string ``

	// NOT MESSAGE
	IsView bool ``

	// NOT MESSAGE
	Names []string ``

	// NOT MESSAGE
	ShouldValidate bool ``

	// NOT MESSAGE
	Ignore bool ``

	// NOT MESSAGE
	Sub string ``

	// NOT MESSAGE
	Obj string ``

	// NOT MESSAGE
	Act string ``

	// NOT MESSAGE
	Key string ``

	// NOT MESSAGE
	Value int32 ``

	// NOT MESSAGE
	Limit int32 ``

	// NOT MESSAGE
	Skip int64 ``

	Sort []Sort ``

	Pagination Pagination ``

	// NOT MESSAGE
	Projection string ``

	// NOT MESSAGE
	Query string ``

	// NOT MESSAGE
	ExcludeOwn bool ``

	// NOT MESSAGE
	ExcludeZone bool ``

	// NOT MESSAGE
	ExcludeShare bool ``

	// NOT MESSAGE
	ExcludeRelatoin bool ``

	// NOT MESSAGE
	IsTrash bool ``

	// NOT MESSAGE
	AddToSet string ``

	// NOT MESSAGE
	Bit string ``

	// NOT MESSAGE
	CurrentDate string ``

	// NOT MESSAGE
	Inc string ``

	// NOT MESSAGE
	Max string ``

	// NOT MESSAGE
	Min string ``

	// NOT MESSAGE
	Mul string ``

	// NOT MESSAGE
	Pop string ``

	// NOT MESSAGE
	Pull string ``

	// NOT MESSAGE
	PullAll string ``

	// NOT MESSAGE
	Push string ``

	// NOT MESSAGE
	Rename string ``

	// NOT MESSAGE
	Set string ``

	// NOT MESSAGE
	SetOnInsert string ``

	// NOT MESSAGE
	Unset string ``

	// NOT MESSAGE
	Count int64 ``

	// NOT MESSAGE
	Id string ``

	// NOT MESSAGE
	Owner string ``

	// NOT MESSAGE
	Zones []string ``

	// NOT MESSAGE
	Clients []string ``

	// NOT MESSAGE
	Shares []string ``

	// NOT MESSAGE
	Relations []string ``

	// NOT MESSAGE
	CreatedBy string ``

	// NOT MESSAGE
	UpdatedBy string ``

	// NOT MESSAGE
	DeletedBy string ``

	// NOT MESSAGE
	RestoredBy string ``

	// NOT MESSAGE
	Id string ``

	// NOT MESSAGE
	CreatedIn string ``

	// NOT MESSAGE
	UpdatedIn string ``

	// NOT MESSAGE
	DeletedIn string ``

	// NOT MESSAGE
	RestoredIn string ``

	// NOT MESSAGE
	Tags []string ``

	// NOT MESSAGE
	Version string ``

	// NOT MESSAGE
	Id string ``

	// NOT MESSAGE
	CreatedAt string ``

	// NOT MESSAGE
	UpdatedAt string ``

	// NOT MESSAGE
	DeletedAt string ``

	// NOT MESSAGE
	RestoredAt string ``
}

type BaseAccessUpdate struct {
	Filter Filter ``

	Filter Filter ``

	// NOT MESSAGE
	Id string ``

	Filter Filter ``

	// NOT MESSAGE
	Ignore bool ``

	// NOT MESSAGE
	DbName string ``

	// NOT MESSAGE
	MongoType string ``

	// NOT MESSAGE
	BsonTag string ``

	// NOT MESSAGE
	IsAggregate bool ``

	// NOT MESSAGE
	Aggregate string ``

	// NOT MESSAGE
	IsView bool ``

	// NOT MESSAGE
	Names []string ``

	// NOT MESSAGE
	ShouldValidate bool ``

	// NOT MESSAGE
	Ignore bool ``

	// NOT MESSAGE
	Sub string ``

	// NOT MESSAGE
	Obj string ``

	// NOT MESSAGE
	Act string ``

	// NOT MESSAGE
	Key string ``

	// NOT MESSAGE
	Value int32 ``

	// NOT MESSAGE
	Limit int32 ``

	// NOT MESSAGE
	Skip int64 ``

	Sort []Sort ``

	Pagination Pagination ``

	// NOT MESSAGE
	Projection string ``

	// NOT MESSAGE
	Query string ``

	// NOT MESSAGE
	ExcludeOwn bool ``

	// NOT MESSAGE
	ExcludeZone bool ``

	// NOT MESSAGE
	ExcludeShare bool ``

	// NOT MESSAGE
	ExcludeRelatoin bool ``

	// NOT MESSAGE
	IsTrash bool ``

	// NOT MESSAGE
	AddToSet string ``

	// NOT MESSAGE
	Bit string ``

	// NOT MESSAGE
	CurrentDate string ``

	// NOT MESSAGE
	Inc string ``

	// NOT MESSAGE
	Max string ``

	// NOT MESSAGE
	Min string ``

	// NOT MESSAGE
	Mul string ``

	// NOT MESSAGE
	Pop string ``

	// NOT MESSAGE
	Pull string ``

	// NOT MESSAGE
	PullAll string ``

	// NOT MESSAGE
	Push string ``

	// NOT MESSAGE
	Rename string ``

	// NOT MESSAGE
	Set string ``

	// NOT MESSAGE
	SetOnInsert string ``

	// NOT MESSAGE
	Unset string ``

	// NOT MESSAGE
	Count int64 ``

	// NOT MESSAGE
	Id string ``

	// NOT MESSAGE
	Owner string ``

	// NOT MESSAGE
	Zones []string ``

	// NOT MESSAGE
	Clients []string ``

	// NOT MESSAGE
	Shares []string ``

	// NOT MESSAGE
	Relations []string ``

	// NOT MESSAGE
	CreatedBy string ``

	// NOT MESSAGE
	UpdatedBy string ``

	// NOT MESSAGE
	DeletedBy string ``

	// NOT MESSAGE
	RestoredBy string ``

	// NOT MESSAGE
	Id string ``

	// NOT MESSAGE
	CreatedIn string ``

	// NOT MESSAGE
	UpdatedIn string ``

	// NOT MESSAGE
	DeletedIn string ``

	// NOT MESSAGE
	RestoredIn string ``

	// NOT MESSAGE
	Tags []string ``

	// NOT MESSAGE
	Version string ``

	// NOT MESSAGE
	Id string ``

	// NOT MESSAGE
	CreatedAt string ``

	// NOT MESSAGE
	UpdatedAt string ``

	// NOT MESSAGE
	DeletedAt string ``

	// NOT MESSAGE
	RestoredAt string ``

	// NOT MESSAGE
	Owner string ``

	// NOT MESSAGE
	Zones []string ``

	// NOT MESSAGE
	Shares []string ``

	// NOT MESSAGE
	Relations []string ``
}

type BasePropertiesUpdate struct {
	Filter Filter ``

	Filter Filter ``

	// NOT MESSAGE
	Id string ``

	Filter Filter ``

	// NOT MESSAGE
	Ignore bool ``

	// NOT MESSAGE
	DbName string ``

	// NOT MESSAGE
	MongoType string ``

	// NOT MESSAGE
	BsonTag string ``

	// NOT MESSAGE
	IsAggregate bool ``

	// NOT MESSAGE
	Aggregate string ``

	// NOT MESSAGE
	IsView bool ``

	// NOT MESSAGE
	Names []string ``

	// NOT MESSAGE
	ShouldValidate bool ``

	// NOT MESSAGE
	Ignore bool ``

	// NOT MESSAGE
	Sub string ``

	// NOT MESSAGE
	Obj string ``

	// NOT MESSAGE
	Act string ``

	// NOT MESSAGE
	Key string ``

	// NOT MESSAGE
	Value int32 ``

	// NOT MESSAGE
	Limit int32 ``

	// NOT MESSAGE
	Skip int64 ``

	Sort []Sort ``

	Pagination Pagination ``

	// NOT MESSAGE
	Projection string ``

	// NOT MESSAGE
	Query string ``

	// NOT MESSAGE
	ExcludeOwn bool ``

	// NOT MESSAGE
	ExcludeZone bool ``

	// NOT MESSAGE
	ExcludeShare bool ``

	// NOT MESSAGE
	ExcludeRelatoin bool ``

	// NOT MESSAGE
	IsTrash bool ``

	// NOT MESSAGE
	AddToSet string ``

	// NOT MESSAGE
	Bit string ``

	// NOT MESSAGE
	CurrentDate string ``

	// NOT MESSAGE
	Inc string ``

	// NOT MESSAGE
	Max string ``

	// NOT MESSAGE
	Min string ``

	// NOT MESSAGE
	Mul string ``

	// NOT MESSAGE
	Pop string ``

	// NOT MESSAGE
	Pull string ``

	// NOT MESSAGE
	PullAll string ``

	// NOT MESSAGE
	Push string ``

	// NOT MESSAGE
	Rename string ``

	// NOT MESSAGE
	Set string ``

	// NOT MESSAGE
	SetOnInsert string ``

	// NOT MESSAGE
	Unset string ``

	// NOT MESSAGE
	Count int64 ``

	// NOT MESSAGE
	Id string ``

	// NOT MESSAGE
	Owner string ``

	// NOT MESSAGE
	Zones []string ``

	// NOT MESSAGE
	Clients []string ``

	// NOT MESSAGE
	Shares []string ``

	// NOT MESSAGE
	Relations []string ``

	// NOT MESSAGE
	CreatedBy string ``

	// NOT MESSAGE
	UpdatedBy string ``

	// NOT MESSAGE
	DeletedBy string ``

	// NOT MESSAGE
	RestoredBy string ``

	// NOT MESSAGE
	Id string ``

	// NOT MESSAGE
	CreatedIn string ``

	// NOT MESSAGE
	UpdatedIn string ``

	// NOT MESSAGE
	DeletedIn string ``

	// NOT MESSAGE
	RestoredIn string ``

	// NOT MESSAGE
	Tags []string ``

	// NOT MESSAGE
	Version string ``

	// NOT MESSAGE
	Id string ``

	// NOT MESSAGE
	CreatedAt string ``

	// NOT MESSAGE
	UpdatedAt string ``

	// NOT MESSAGE
	DeletedAt string ``

	// NOT MESSAGE
	RestoredAt string ``

	// NOT MESSAGE
	Owner string ``

	// NOT MESSAGE
	Zones []string ``

	// NOT MESSAGE
	Shares []string ``

	// NOT MESSAGE
	Relations []string ``

	// NOT MESSAGE
	Tags []string ``
}

type BaseDatesUpdate struct {
	Filter Filter ``

	Filter Filter ``

	// NOT MESSAGE
	Id string ``

	Filter Filter ``

	// NOT MESSAGE
	Ignore bool ``

	// NOT MESSAGE
	DbName string ``

	// NOT MESSAGE
	MongoType string ``

	// NOT MESSAGE
	BsonTag string ``

	// NOT MESSAGE
	IsAggregate bool ``

	// NOT MESSAGE
	Aggregate string ``

	// NOT MESSAGE
	IsView bool ``

	// NOT MESSAGE
	Names []string ``

	// NOT MESSAGE
	ShouldValidate bool ``

	// NOT MESSAGE
	Ignore bool ``

	// NOT MESSAGE
	Sub string ``

	// NOT MESSAGE
	Obj string ``

	// NOT MESSAGE
	Act string ``

	// NOT MESSAGE
	Key string ``

	// NOT MESSAGE
	Value int32 ``

	// NOT MESSAGE
	Limit int32 ``

	// NOT MESSAGE
	Skip int64 ``

	Sort []Sort ``

	Pagination Pagination ``

	// NOT MESSAGE
	Projection string ``

	// NOT MESSAGE
	Query string ``

	// NOT MESSAGE
	ExcludeOwn bool ``

	// NOT MESSAGE
	ExcludeZone bool ``

	// NOT MESSAGE
	ExcludeShare bool ``

	// NOT MESSAGE
	ExcludeRelatoin bool ``

	// NOT MESSAGE
	IsTrash bool ``

	// NOT MESSAGE
	AddToSet string ``

	// NOT MESSAGE
	Bit string ``

	// NOT MESSAGE
	CurrentDate string ``

	// NOT MESSAGE
	Inc string ``

	// NOT MESSAGE
	Max string ``

	// NOT MESSAGE
	Min string ``

	// NOT MESSAGE
	Mul string ``

	// NOT MESSAGE
	Pop string ``

	// NOT MESSAGE
	Pull string ``

	// NOT MESSAGE
	PullAll string ``

	// NOT MESSAGE
	Push string ``

	// NOT MESSAGE
	Rename string ``

	// NOT MESSAGE
	Set string ``

	// NOT MESSAGE
	SetOnInsert string ``

	// NOT MESSAGE
	Unset string ``

	// NOT MESSAGE
	Count int64 ``

	// NOT MESSAGE
	Id string ``

	// NOT MESSAGE
	Owner string ``

	// NOT MESSAGE
	Zones []string ``

	// NOT MESSAGE
	Clients []string ``

	// NOT MESSAGE
	Shares []string ``

	// NOT MESSAGE
	Relations []string ``

	// NOT MESSAGE
	CreatedBy string ``

	// NOT MESSAGE
	UpdatedBy string ``

	// NOT MESSAGE
	DeletedBy string ``

	// NOT MESSAGE
	RestoredBy string ``

	// NOT MESSAGE
	Id string ``

	// NOT MESSAGE
	CreatedIn string ``

	// NOT MESSAGE
	UpdatedIn string ``

	// NOT MESSAGE
	DeletedIn string ``

	// NOT MESSAGE
	RestoredIn string ``

	// NOT MESSAGE
	Tags []string ``

	// NOT MESSAGE
	Version string ``

	// NOT MESSAGE
	Id string ``

	// NOT MESSAGE
	CreatedAt string ``

	// NOT MESSAGE
	UpdatedAt string ``

	// NOT MESSAGE
	DeletedAt string ``

	// NOT MESSAGE
	RestoredAt string ``

	// NOT MESSAGE
	Owner string ``

	// NOT MESSAGE
	Zones []string ``

	// NOT MESSAGE
	Shares []string ``

	// NOT MESSAGE
	Relations []string ``

	// NOT MESSAGE
	Tags []string ``
}

type BaseAccessRead struct {
	Filter Filter ``

	Filter Filter ``

	// NOT MESSAGE
	Id string ``

	Filter Filter ``

	// NOT MESSAGE
	Ignore bool ``

	// NOT MESSAGE
	DbName string ``

	// NOT MESSAGE
	MongoType string ``

	// NOT MESSAGE
	BsonTag string ``

	// NOT MESSAGE
	IsAggregate bool ``

	// NOT MESSAGE
	Aggregate string ``

	// NOT MESSAGE
	IsView bool ``

	// NOT MESSAGE
	Names []string ``

	// NOT MESSAGE
	ShouldValidate bool ``

	// NOT MESSAGE
	Ignore bool ``

	// NOT MESSAGE
	Sub string ``

	// NOT MESSAGE
	Obj string ``

	// NOT MESSAGE
	Act string ``

	// NOT MESSAGE
	Key string ``

	// NOT MESSAGE
	Value int32 ``

	// NOT MESSAGE
	Limit int32 ``

	// NOT MESSAGE
	Skip int64 ``

	Sort []Sort ``

	Pagination Pagination ``

	// NOT MESSAGE
	Projection string ``

	// NOT MESSAGE
	Query string ``

	// NOT MESSAGE
	ExcludeOwn bool ``

	// NOT MESSAGE
	ExcludeZone bool ``

	// NOT MESSAGE
	ExcludeShare bool ``

	// NOT MESSAGE
	ExcludeRelatoin bool ``

	// NOT MESSAGE
	IsTrash bool ``

	// NOT MESSAGE
	AddToSet string ``

	// NOT MESSAGE
	Bit string ``

	// NOT MESSAGE
	CurrentDate string ``

	// NOT MESSAGE
	Inc string ``

	// NOT MESSAGE
	Max string ``

	// NOT MESSAGE
	Min string ``

	// NOT MESSAGE
	Mul string ``

	// NOT MESSAGE
	Pop string ``

	// NOT MESSAGE
	Pull string ``

	// NOT MESSAGE
	PullAll string ``

	// NOT MESSAGE
	Push string ``

	// NOT MESSAGE
	Rename string ``

	// NOT MESSAGE
	Set string ``

	// NOT MESSAGE
	SetOnInsert string ``

	// NOT MESSAGE
	Unset string ``

	// NOT MESSAGE
	Count int64 ``

	// NOT MESSAGE
	Id string ``

	// NOT MESSAGE
	Owner string ``

	// NOT MESSAGE
	Zones []string ``

	// NOT MESSAGE
	Clients []string ``

	// NOT MESSAGE
	Shares []string ``

	// NOT MESSAGE
	Relations []string ``

	// NOT MESSAGE
	CreatedBy string ``

	// NOT MESSAGE
	UpdatedBy string ``

	// NOT MESSAGE
	DeletedBy string ``

	// NOT MESSAGE
	RestoredBy string ``

	// NOT MESSAGE
	Id string ``

	// NOT MESSAGE
	CreatedIn string ``

	// NOT MESSAGE
	UpdatedIn string ``

	// NOT MESSAGE
	DeletedIn string ``

	// NOT MESSAGE
	RestoredIn string ``

	// NOT MESSAGE
	Tags []string ``

	// NOT MESSAGE
	Version string ``

	// NOT MESSAGE
	Id string ``

	// NOT MESSAGE
	CreatedAt string ``

	// NOT MESSAGE
	UpdatedAt string ``

	// NOT MESSAGE
	DeletedAt string ``

	// NOT MESSAGE
	RestoredAt string ``

	// NOT MESSAGE
	Owner string ``

	// NOT MESSAGE
	Zones []string ``

	// NOT MESSAGE
	Shares []string ``

	// NOT MESSAGE
	Relations []string ``

	// NOT MESSAGE
	Tags []string ``

	// NOT MESSAGE
	Owner string ``

	// NOT MESSAGE
	Zones []string ``

	// NOT MESSAGE
	Clients []string ``

	// NOT MESSAGE
	Shares []string ``

	// NOT MESSAGE
	Relations []string ``

	// NOT MESSAGE
	CreatedBy string ``

	// NOT MESSAGE
	UpdatedBy string ``

	// NOT MESSAGE
	DeletedBy string ``

	// NOT MESSAGE
	RestoredBy string ``
}

type BasePropertiesRead struct {
	Filter Filter ``

	Filter Filter ``

	// NOT MESSAGE
	Id string ``

	Filter Filter ``

	// NOT MESSAGE
	Ignore bool ``

	// NOT MESSAGE
	DbName string ``

	// NOT MESSAGE
	MongoType string ``

	// NOT MESSAGE
	BsonTag string ``

	// NOT MESSAGE
	IsAggregate bool ``

	// NOT MESSAGE
	Aggregate string ``

	// NOT MESSAGE
	IsView bool ``

	// NOT MESSAGE
	Names []string ``

	// NOT MESSAGE
	ShouldValidate bool ``

	// NOT MESSAGE
	Ignore bool ``

	// NOT MESSAGE
	Sub string ``

	// NOT MESSAGE
	Obj string ``

	// NOT MESSAGE
	Act string ``

	// NOT MESSAGE
	Key string ``

	// NOT MESSAGE
	Value int32 ``

	// NOT MESSAGE
	Limit int32 ``

	// NOT MESSAGE
	Skip int64 ``

	Sort []Sort ``

	Pagination Pagination ``

	// NOT MESSAGE
	Projection string ``

	// NOT MESSAGE
	Query string ``

	// NOT MESSAGE
	ExcludeOwn bool ``

	// NOT MESSAGE
	ExcludeZone bool ``

	// NOT MESSAGE
	ExcludeShare bool ``

	// NOT MESSAGE
	ExcludeRelatoin bool ``

	// NOT MESSAGE
	IsTrash bool ``

	// NOT MESSAGE
	AddToSet string ``

	// NOT MESSAGE
	Bit string ``

	// NOT MESSAGE
	CurrentDate string ``

	// NOT MESSAGE
	Inc string ``

	// NOT MESSAGE
	Max string ``

	// NOT MESSAGE
	Min string ``

	// NOT MESSAGE
	Mul string ``

	// NOT MESSAGE
	Pop string ``

	// NOT MESSAGE
	Pull string ``

	// NOT MESSAGE
	PullAll string ``

	// NOT MESSAGE
	Push string ``

	// NOT MESSAGE
	Rename string ``

	// NOT MESSAGE
	Set string ``

	// NOT MESSAGE
	SetOnInsert string ``

	// NOT MESSAGE
	Unset string ``

	// NOT MESSAGE
	Count int64 ``

	// NOT MESSAGE
	Id string ``

	// NOT MESSAGE
	Owner string ``

	// NOT MESSAGE
	Zones []string ``

	// NOT MESSAGE
	Clients []string ``

	// NOT MESSAGE
	Shares []string ``

	// NOT MESSAGE
	Relations []string ``

	// NOT MESSAGE
	CreatedBy string ``

	// NOT MESSAGE
	UpdatedBy string ``

	// NOT MESSAGE
	DeletedBy string ``

	// NOT MESSAGE
	RestoredBy string ``

	// NOT MESSAGE
	Id string ``

	// NOT MESSAGE
	CreatedIn string ``

	// NOT MESSAGE
	UpdatedIn string ``

	// NOT MESSAGE
	DeletedIn string ``

	// NOT MESSAGE
	RestoredIn string ``

	// NOT MESSAGE
	Tags []string ``

	// NOT MESSAGE
	Version string ``

	// NOT MESSAGE
	Id string ``

	// NOT MESSAGE
	CreatedAt string ``

	// NOT MESSAGE
	UpdatedAt string ``

	// NOT MESSAGE
	DeletedAt string ``

	// NOT MESSAGE
	RestoredAt string ``

	// NOT MESSAGE
	Owner string ``

	// NOT MESSAGE
	Zones []string ``

	// NOT MESSAGE
	Shares []string ``

	// NOT MESSAGE
	Relations []string ``

	// NOT MESSAGE
	Tags []string ``

	// NOT MESSAGE
	Owner string ``

	// NOT MESSAGE
	Zones []string ``

	// NOT MESSAGE
	Clients []string ``

	// NOT MESSAGE
	Shares []string ``

	// NOT MESSAGE
	Relations []string ``

	// NOT MESSAGE
	CreatedBy string ``

	// NOT MESSAGE
	UpdatedBy string ``

	// NOT MESSAGE
	DeletedBy string ``

	// NOT MESSAGE
	RestoredBy string ``

	// NOT MESSAGE
	CreatedIn string ``

	// NOT MESSAGE
	UpdatedIn string ``

	// NOT MESSAGE
	DeletedIn string ``

	// NOT MESSAGE
	RestoredIn string ``

	// NOT MESSAGE
	Tags []string ``

	// NOT MESSAGE
	Version string ``
}

type BaseDatesRead struct {
	Filter Filter ``

	Filter Filter ``

	// NOT MESSAGE
	Id string ``

	Filter Filter ``

	// NOT MESSAGE
	Ignore bool ``

	// NOT MESSAGE
	DbName string ``

	// NOT MESSAGE
	MongoType string ``

	// NOT MESSAGE
	BsonTag string ``

	// NOT MESSAGE
	IsAggregate bool ``

	// NOT MESSAGE
	Aggregate string ``

	// NOT MESSAGE
	IsView bool ``

	// NOT MESSAGE
	Names []string ``

	// NOT MESSAGE
	ShouldValidate bool ``

	// NOT MESSAGE
	Ignore bool ``

	// NOT MESSAGE
	Sub string ``

	// NOT MESSAGE
	Obj string ``

	// NOT MESSAGE
	Act string ``

	// NOT MESSAGE
	Key string ``

	// NOT MESSAGE
	Value int32 ``

	// NOT MESSAGE
	Limit int32 ``

	// NOT MESSAGE
	Skip int64 ``

	Sort []Sort ``

	Pagination Pagination ``

	// NOT MESSAGE
	Projection string ``

	// NOT MESSAGE
	Query string ``

	// NOT MESSAGE
	ExcludeOwn bool ``

	// NOT MESSAGE
	ExcludeZone bool ``

	// NOT MESSAGE
	ExcludeShare bool ``

	// NOT MESSAGE
	ExcludeRelatoin bool ``

	// NOT MESSAGE
	IsTrash bool ``

	// NOT MESSAGE
	AddToSet string ``

	// NOT MESSAGE
	Bit string ``

	// NOT MESSAGE
	CurrentDate string ``

	// NOT MESSAGE
	Inc string ``

	// NOT MESSAGE
	Max string ``

	// NOT MESSAGE
	Min string ``

	// NOT MESSAGE
	Mul string ``

	// NOT MESSAGE
	Pop string ``

	// NOT MESSAGE
	Pull string ``

	// NOT MESSAGE
	PullAll string ``

	// NOT MESSAGE
	Push string ``

	// NOT MESSAGE
	Rename string ``

	// NOT MESSAGE
	Set string ``

	// NOT MESSAGE
	SetOnInsert string ``

	// NOT MESSAGE
	Unset string ``

	// NOT MESSAGE
	Count int64 ``

	// NOT MESSAGE
	Id string ``

	// NOT MESSAGE
	Owner string ``

	// NOT MESSAGE
	Zones []string ``

	// NOT MESSAGE
	Clients []string ``

	// NOT MESSAGE
	Shares []string ``

	// NOT MESSAGE
	Relations []string ``

	// NOT MESSAGE
	CreatedBy string ``

	// NOT MESSAGE
	UpdatedBy string ``

	// NOT MESSAGE
	DeletedBy string ``

	// NOT MESSAGE
	RestoredBy string ``

	// NOT MESSAGE
	Id string ``

	// NOT MESSAGE
	CreatedIn string ``

	// NOT MESSAGE
	UpdatedIn string ``

	// NOT MESSAGE
	DeletedIn string ``

	// NOT MESSAGE
	RestoredIn string ``

	// NOT MESSAGE
	Tags []string ``

	// NOT MESSAGE
	Version string ``

	// NOT MESSAGE
	Id string ``

	// NOT MESSAGE
	CreatedAt string ``

	// NOT MESSAGE
	UpdatedAt string ``

	// NOT MESSAGE
	DeletedAt string ``

	// NOT MESSAGE
	RestoredAt string ``

	// NOT MESSAGE
	Owner string ``

	// NOT MESSAGE
	Zones []string ``

	// NOT MESSAGE
	Shares []string ``

	// NOT MESSAGE
	Relations []string ``

	// NOT MESSAGE
	Tags []string ``

	// NOT MESSAGE
	Owner string ``

	// NOT MESSAGE
	Zones []string ``

	// NOT MESSAGE
	Clients []string ``

	// NOT MESSAGE
	Shares []string ``

	// NOT MESSAGE
	Relations []string ``

	// NOT MESSAGE
	CreatedBy string ``

	// NOT MESSAGE
	UpdatedBy string ``

	// NOT MESSAGE
	DeletedBy string ``

	// NOT MESSAGE
	RestoredBy string ``

	// NOT MESSAGE
	CreatedIn string ``

	// NOT MESSAGE
	UpdatedIn string ``

	// NOT MESSAGE
	DeletedIn string ``

	// NOT MESSAGE
	RestoredIn string ``

	// NOT MESSAGE
	Tags []string ``

	// NOT MESSAGE
	Version string ``

	// NOT MESSAGE
	CreatedAt string ``

	// NOT MESSAGE
	UpdatedAt string ``

	// NOT MESSAGE
	DeletedAt string ``

	// NOT MESSAGE
	RestoredAt string ``
}
