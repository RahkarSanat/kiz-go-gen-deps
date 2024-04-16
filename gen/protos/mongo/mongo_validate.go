package mongo

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var _ primitive.A
var _ fmt.Stringer

type FieldOptionsMongo struct {
	Ignore    bool   `bson:"ignore" json:"ignore"`
	Required  bool   `bson:"required" json:"required"`
	MinLength int32  `bson:"min_length" json:"min_length"`
	MaxLength int32  `bson:"max_length" json:"max_length"`
	Pattern   string `bson:"pattern" json:"pattern"`
}
type FileOptionsMongo struct {
	Ignore    bool   `bson:"ignore" json:"ignore"`
	Extension string `bson:"extension" json:"extension"`
}
type MessageOptionsMongo struct {
	NotIgnore                    bool `bson:"not_ignore" json:"not_ignore"`
	AllFieldsRequired            bool `bson:"all_fields_required" json:"all_fields_required"`
	AllowNullValues              bool `bson:"allow_null_values" json:"allow_null_values"`
	DisallowAdditionalProperties bool `bson:"disallow_additional_properties" json:"disallow_additional_properties"`
	EnumsAsConstants             bool `bson:"enums_as_constants" json:"enums_as_constants"`
}
type EnumOptionsMongo struct {
	EnumsAsConstants   bool `bson:"enums_as_constants" json:"enums_as_constants"`
	EnumsAsStringsOnly bool `bson:"enums_as_strings_only" json:"enums_as_strings_only"`
	EnumsTrimPrefix    bool `bson:"enums_trim_prefix" json:"enums_trim_prefix"`
	Ignore             bool `bson:"ignore" json:"ignore"`
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
