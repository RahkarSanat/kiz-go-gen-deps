package types

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var _ primitive.A
var _ fmt.Stringer

type UUIDValueMongo struct {
	Value string `bson:"value" json:"value"`
}
type JSONValueMongo struct {
	Value string `bson:"value" json:"value"`
}
type UUIDMongo struct {
	Value string `bson:"value" json:"value"`
}
type InetValueMongo struct {
	Value string `bson:"value" json:"value"`
}
type TimeOnlyMongo struct {
	Value uint32 `bson:"value" json:"value"`
}
type BigIntMongo struct {
	Value string `bson:"value" json:"value"`
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
