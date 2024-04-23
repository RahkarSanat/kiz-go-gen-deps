package gorm

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

type GormFileOptionsMongo struct {
}
type GormMessageOptionsMongo struct {
	Ormable      *bool             `bson:"ormable,omitempty" json:"ormable,omitempty"`
	Include      []ExtraFieldMongo `bson:"include,omitempty" json:"include,omitempty"`
	Table        *string           `bson:"table,omitempty" json:"table,omitempty"`
	MultiAccount *bool             `bson:"multi_account,omitempty" json:"multi_account,omitempty"`
}
type ExtraFieldMongo struct {
	Type    *string       `bson:"type,omitempty" json:"type,omitempty"`
	Name    *string       `bson:"name,omitempty" json:"name,omitempty"`
	Tag     *GormTagMongo `bson:"tag,omitempty" json:"tag,omitempty"`
	Package *string       `bson:"package,omitempty" json:"package,omitempty"`
}
type GormFieldOptionsMongo struct {
	Tag         *GormTagMongo           `bson:"tag,omitempty" json:"tag,omitempty"`
	Drop        *bool                   `bson:"drop,omitempty" json:"drop,omitempty"`
	HasOne      *HasOneOptionsMongo     `bson:"has_one,omitempty" json:"has_one,omitempty"`
	BelongsTo   *BelongsToOptionsMongo  `bson:"belongs_to,omitempty" json:"belongs_to,omitempty"`
	HasMany     *HasManyOptionsMongo    `bson:"has_many,omitempty" json:"has_many,omitempty"`
	ManyToMany  *ManyToManyOptionsMongo `bson:"many_to_many,omitempty" json:"many_to_many,omitempty"`
	ReferenceOf *string                 `bson:"reference_of,omitempty" json:"reference_of,omitempty"`
}
type GormTagMongo struct {
	Column                         *string `bson:"column,omitempty" json:"column,omitempty"`
	Type                           *string `bson:"type,omitempty" json:"type,omitempty"`
	Size                           *int32  `bson:"size,omitempty" json:"size,omitempty"`
	Precision                      *int32  `bson:"precision,omitempty" json:"precision,omitempty"`
	PrimaryKey                     *bool   `bson:"primary_key,omitempty" json:"primary_key,omitempty"`
	Unique                         *bool   `bson:"unique,omitempty" json:"unique,omitempty"`
	Default                        *string `bson:"default,omitempty" json:"default,omitempty"`
	NotNull                        *bool   `bson:"not_null,omitempty" json:"not_null,omitempty"`
	AutoIncrement                  *bool   `bson:"auto_increment,omitempty" json:"auto_increment,omitempty"`
	Index                          *string `bson:"index,omitempty" json:"index,omitempty"`
	UniqueIndex                    *string `bson:"unique_index,omitempty" json:"unique_index,omitempty"`
	Embedded                       *bool   `bson:"embedded,omitempty" json:"embedded,omitempty"`
	EmbeddedPrefix                 *string `bson:"embedded_prefix,omitempty" json:"embedded_prefix,omitempty"`
	Ignore                         *bool   `bson:"ignore,omitempty" json:"ignore,omitempty"`
	Foreignkey                     *string `bson:"foreignkey,omitempty" json:"foreignkey,omitempty"`
	AssociationForeignkey          *string `bson:"association_foreignkey,omitempty" json:"association_foreignkey,omitempty"`
	ManyToMany                     *string `bson:"many_to_many,omitempty" json:"many_to_many,omitempty"`
	JointableForeignkey            *string `bson:"jointable_foreignkey,omitempty" json:"jointable_foreignkey,omitempty"`
	AssociationJointableForeignkey *string `bson:"association_jointable_foreignkey,omitempty" json:"association_jointable_foreignkey,omitempty"`
	DisableAssociationAutoupdate   *bool   `bson:"disable_association_autoupdate,omitempty" json:"disable_association_autoupdate,omitempty"`
	DisableAssociationAutocreate   *bool   `bson:"disable_association_autocreate,omitempty" json:"disable_association_autocreate,omitempty"`
	AssociationSaveReference       *bool   `bson:"association_save_reference,omitempty" json:"association_save_reference,omitempty"`
	Preload                        *bool   `bson:"preload,omitempty" json:"preload,omitempty"`
	Serializer                     *string `bson:"serializer,omitempty" json:"serializer,omitempty"`
}
type HasOneOptionsMongo struct {
	Foreignkey                   *string       `bson:"foreignkey,omitempty" json:"foreignkey,omitempty"`
	ForeignkeyTag                *GormTagMongo `bson:"foreignkey_tag,omitempty" json:"foreignkey_tag,omitempty"`
	AssociationForeignkey        *string       `bson:"association_foreignkey,omitempty" json:"association_foreignkey,omitempty"`
	DisableAssociationAutoupdate *bool         `bson:"disable_association_autoupdate,omitempty" json:"disable_association_autoupdate,omitempty"`
	DisableAssociationAutocreate *bool         `bson:"disable_association_autocreate,omitempty" json:"disable_association_autocreate,omitempty"`
	AssociationSaveReference     *bool         `bson:"association_save_reference,omitempty" json:"association_save_reference,omitempty"`
	Preload                      *bool         `bson:"preload,omitempty" json:"preload,omitempty"`
	Replace                      *bool         `bson:"replace,omitempty" json:"replace,omitempty"`
	Append                       *bool         `bson:"append,omitempty" json:"append,omitempty"`
	Clear                        *bool         `bson:"clear,omitempty" json:"clear,omitempty"`
}
type BelongsToOptionsMongo struct {
	Foreignkey                   *string       `bson:"foreignkey,omitempty" json:"foreignkey,omitempty"`
	ForeignkeyTag                *GormTagMongo `bson:"foreignkey_tag,omitempty" json:"foreignkey_tag,omitempty"`
	AssociationForeignkey        *string       `bson:"association_foreignkey,omitempty" json:"association_foreignkey,omitempty"`
	DisableAssociationAutoupdate *bool         `bson:"disable_association_autoupdate,omitempty" json:"disable_association_autoupdate,omitempty"`
	DisableAssociationAutocreate *bool         `bson:"disable_association_autocreate,omitempty" json:"disable_association_autocreate,omitempty"`
	AssociationSaveReference     *bool         `bson:"association_save_reference,omitempty" json:"association_save_reference,omitempty"`
	Preload                      *bool         `bson:"preload,omitempty" json:"preload,omitempty"`
}
type HasManyOptionsMongo struct {
	Foreignkey                   *string       `bson:"foreignkey,omitempty" json:"foreignkey,omitempty"`
	ForeignkeyTag                *GormTagMongo `bson:"foreignkey_tag,omitempty" json:"foreignkey_tag,omitempty"`
	AssociationForeignkey        *string       `bson:"association_foreignkey,omitempty" json:"association_foreignkey,omitempty"`
	PositionField                *string       `bson:"position_field,omitempty" json:"position_field,omitempty"`
	PositionFieldTag             *GormTagMongo `bson:"position_field_tag,omitempty" json:"position_field_tag,omitempty"`
	DisableAssociationAutoupdate *bool         `bson:"disable_association_autoupdate,omitempty" json:"disable_association_autoupdate,omitempty"`
	DisableAssociationAutocreate *bool         `bson:"disable_association_autocreate,omitempty" json:"disable_association_autocreate,omitempty"`
	AssociationSaveReference     *bool         `bson:"association_save_reference,omitempty" json:"association_save_reference,omitempty"`
	Preload                      *bool         `bson:"preload,omitempty" json:"preload,omitempty"`
	Replace                      *bool         `bson:"replace,omitempty" json:"replace,omitempty"`
	Append                       *bool         `bson:"append,omitempty" json:"append,omitempty"`
	Clear                        *bool         `bson:"clear,omitempty" json:"clear,omitempty"`
}
type ManyToManyOptionsMongo struct {
	Jointable                      *string `bson:"jointable,omitempty" json:"jointable,omitempty"`
	Foreignkey                     *string `bson:"foreignkey,omitempty" json:"foreignkey,omitempty"`
	JointableForeignkey            *string `bson:"jointable_foreignkey,omitempty" json:"jointable_foreignkey,omitempty"`
	AssociationForeignkey          *string `bson:"association_foreignkey,omitempty" json:"association_foreignkey,omitempty"`
	AssociationJointableForeignkey *string `bson:"association_jointable_foreignkey,omitempty" json:"association_jointable_foreignkey,omitempty"`
	DisableAssociationAutoupdate   *bool   `bson:"disable_association_autoupdate,omitempty" json:"disable_association_autoupdate,omitempty"`
	DisableAssociationAutocreate   *bool   `bson:"disable_association_autocreate,omitempty" json:"disable_association_autocreate,omitempty"`
	AssociationSaveReference       *bool   `bson:"association_save_reference,omitempty" json:"association_save_reference,omitempty"`
	Preload                        *bool   `bson:"preload,omitempty" json:"preload,omitempty"`
	Replace                        *bool   `bson:"replace,omitempty" json:"replace,omitempty"`
	Append                         *bool   `bson:"append,omitempty" json:"append,omitempty"`
	Clear                          *bool   `bson:"clear,omitempty" json:"clear,omitempty"`
}
type AutoServerOptionsMongo struct {
	Autogen       *bool `bson:"autogen,omitempty" json:"autogen,omitempty"`
	TxnMiddleware *bool `bson:"txn_middleware,omitempty" json:"txn_middleware,omitempty"`
	WithTracing   *bool `bson:"with_tracing,omitempty" json:"with_tracing,omitempty"`
}
type MethodOptionsMongo struct {
	ObjectType *string `bson:"object_type,omitempty" json:"object_type,omitempty"`
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

func (m *GormFileOptions) Validate() error {

	return nil
}

func (m *GormMessageOptions) Validate() error {

	return nil
}

func (m *ExtraField) Validate() error {

	return nil
}

func (m *GormFieldOptions) Validate() error {

	// mamad

	// mamad

	// mamad

	// mamad

	return nil
}

func (m *GormTag) Validate() error {

	return nil
}

func (m *HasOneOptions) Validate() error {

	return nil
}

func (m *BelongsToOptions) Validate() error {

	return nil
}

func (m *HasManyOptions) Validate() error {

	return nil
}

func (m *ManyToManyOptions) Validate() error {

	return nil
}

func (m *AutoServerOptions) Validate() error {

	return nil
}

func (m *MethodOptions) Validate() error {

	return nil
}
