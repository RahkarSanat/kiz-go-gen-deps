package gorm

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var _ primitive.A
var _ fmt.Stringer

type GormFileOptionsMongo struct {
}
type GormMessageOptionsMongo struct {
	Ormable      bool              `bson:"ormable" json:"ormable"`
	Include      []ExtraFieldMongo `bson:"include" json:"include"`
	Table        string            `bson:"table" json:"table"`
	MultiAccount bool              `bson:"multi_account" json:"multi_account"`
}
type ExtraFieldMongo struct {
	Type    string       `bson:"type" json:"type"`
	Name    string       `bson:"name" json:"name"`
	Tag     GormTagMongo `bson:"tag" json:"tag"`
	Package string       `bson:"package" json:"package"`
}
type GormFieldOptionsMongo struct {
	Tag         GormTagMongo           `bson:"tag" json:"tag"`
	Drop        bool                   `bson:"drop" json:"drop"`
	HasOne      HasOneOptionsMongo     `bson:"has_one" json:"has_one"`
	BelongsTo   BelongsToOptionsMongo  `bson:"belongs_to" json:"belongs_to"`
	HasMany     HasManyOptionsMongo    `bson:"has_many" json:"has_many"`
	ManyToMany  ManyToManyOptionsMongo `bson:"many_to_many" json:"many_to_many"`
	ReferenceOf string                 `bson:"reference_of" json:"reference_of"`
}
type GormTagMongo struct {
	Column                         string `bson:"column" json:"column"`
	Type                           string `bson:"type" json:"type"`
	Size                           int32  `bson:"size" json:"size"`
	Precision                      int32  `bson:"precision" json:"precision"`
	PrimaryKey                     bool   `bson:"primary_key" json:"primary_key"`
	Unique                         bool   `bson:"unique" json:"unique"`
	Default                        string `bson:"default" json:"default"`
	NotNull                        bool   `bson:"not_null" json:"not_null"`
	AutoIncrement                  bool   `bson:"auto_increment" json:"auto_increment"`
	Index                          string `bson:"index" json:"index"`
	UniqueIndex                    string `bson:"unique_index" json:"unique_index"`
	Embedded                       bool   `bson:"embedded" json:"embedded"`
	EmbeddedPrefix                 string `bson:"embedded_prefix" json:"embedded_prefix"`
	Ignore                         bool   `bson:"ignore" json:"ignore"`
	Foreignkey                     string `bson:"foreignkey" json:"foreignkey"`
	AssociationForeignkey          string `bson:"association_foreignkey" json:"association_foreignkey"`
	ManyToMany                     string `bson:"many_to_many" json:"many_to_many"`
	JointableForeignkey            string `bson:"jointable_foreignkey" json:"jointable_foreignkey"`
	AssociationJointableForeignkey string `bson:"association_jointable_foreignkey" json:"association_jointable_foreignkey"`
	DisableAssociationAutoupdate   bool   `bson:"disable_association_autoupdate" json:"disable_association_autoupdate"`
	DisableAssociationAutocreate   bool   `bson:"disable_association_autocreate" json:"disable_association_autocreate"`
	AssociationSaveReference       bool   `bson:"association_save_reference" json:"association_save_reference"`
	Preload                        bool   `bson:"preload" json:"preload"`
	Serializer                     string `bson:"serializer" json:"serializer"`
}
type HasOneOptionsMongo struct {
	Foreignkey                   string       `bson:"foreignkey" json:"foreignkey"`
	ForeignkeyTag                GormTagMongo `bson:"foreignkey_tag" json:"foreignkey_tag"`
	AssociationForeignkey        string       `bson:"association_foreignkey" json:"association_foreignkey"`
	DisableAssociationAutoupdate bool         `bson:"disable_association_autoupdate" json:"disable_association_autoupdate"`
	DisableAssociationAutocreate bool         `bson:"disable_association_autocreate" json:"disable_association_autocreate"`
	AssociationSaveReference     bool         `bson:"association_save_reference" json:"association_save_reference"`
	Preload                      bool         `bson:"preload" json:"preload"`
	Replace                      bool         `bson:"replace" json:"replace"`
	Append                       bool         `bson:"append" json:"append"`
	Clear                        bool         `bson:"clear" json:"clear"`
}
type BelongsToOptionsMongo struct {
	Foreignkey                   string       `bson:"foreignkey" json:"foreignkey"`
	ForeignkeyTag                GormTagMongo `bson:"foreignkey_tag" json:"foreignkey_tag"`
	AssociationForeignkey        string       `bson:"association_foreignkey" json:"association_foreignkey"`
	DisableAssociationAutoupdate bool         `bson:"disable_association_autoupdate" json:"disable_association_autoupdate"`
	DisableAssociationAutocreate bool         `bson:"disable_association_autocreate" json:"disable_association_autocreate"`
	AssociationSaveReference     bool         `bson:"association_save_reference" json:"association_save_reference"`
	Preload                      bool         `bson:"preload" json:"preload"`
}
type HasManyOptionsMongo struct {
	Foreignkey                   string       `bson:"foreignkey" json:"foreignkey"`
	ForeignkeyTag                GormTagMongo `bson:"foreignkey_tag" json:"foreignkey_tag"`
	AssociationForeignkey        string       `bson:"association_foreignkey" json:"association_foreignkey"`
	PositionField                string       `bson:"position_field" json:"position_field"`
	PositionFieldTag             GormTagMongo `bson:"position_field_tag" json:"position_field_tag"`
	DisableAssociationAutoupdate bool         `bson:"disable_association_autoupdate" json:"disable_association_autoupdate"`
	DisableAssociationAutocreate bool         `bson:"disable_association_autocreate" json:"disable_association_autocreate"`
	AssociationSaveReference     bool         `bson:"association_save_reference" json:"association_save_reference"`
	Preload                      bool         `bson:"preload" json:"preload"`
	Replace                      bool         `bson:"replace" json:"replace"`
	Append                       bool         `bson:"append" json:"append"`
	Clear                        bool         `bson:"clear" json:"clear"`
}
type ManyToManyOptionsMongo struct {
	Jointable                      string `bson:"jointable" json:"jointable"`
	Foreignkey                     string `bson:"foreignkey" json:"foreignkey"`
	JointableForeignkey            string `bson:"jointable_foreignkey" json:"jointable_foreignkey"`
	AssociationForeignkey          string `bson:"association_foreignkey" json:"association_foreignkey"`
	AssociationJointableForeignkey string `bson:"association_jointable_foreignkey" json:"association_jointable_foreignkey"`
	DisableAssociationAutoupdate   bool   `bson:"disable_association_autoupdate" json:"disable_association_autoupdate"`
	DisableAssociationAutocreate   bool   `bson:"disable_association_autocreate" json:"disable_association_autocreate"`
	AssociationSaveReference       bool   `bson:"association_save_reference" json:"association_save_reference"`
	Preload                        bool   `bson:"preload" json:"preload"`
	Replace                        bool   `bson:"replace" json:"replace"`
	Append                         bool   `bson:"append" json:"append"`
	Clear                          bool   `bson:"clear" json:"clear"`
}
type AutoServerOptionsMongo struct {
	Autogen       bool `bson:"autogen" json:"autogen"`
	TxnMiddleware bool `bson:"txn_middleware" json:"txn_middleware"`
	WithTracing   bool `bson:"with_tracing" json:"with_tracing"`
}
type MethodOptionsMongo struct {
	ObjectType string `bson:"object_type" json:"object_type"`
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
