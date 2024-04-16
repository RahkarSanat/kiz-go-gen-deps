package storage

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var _ primitive.A
var _ fmt.Stringer

type ReadObjectRequestMongo struct {
	Prefix string `bson:"prefix" json:"prefix"`
	Object string `bson:"object" json:"object"`
	Title  string `bson:"title" json:"title"`
}
type WriteObjectRequestMongo struct {
	FirstMessage FirstMessageMongo    `bson:"first_message" json:"first_message"`
	Data         ChecksummedDataMongo `bson:"data" json:"data"`
	FinishWrite  bool                 `bson:"finish_write" json:"finish_write"`
}
type FirstMessageMongo struct {
	Prefix string `bson:"prefix" json:"prefix"`
	Name   string `bson:"name" json:"name"`
}
type ReadObjectResponseMongo struct {
	FirstMessage    FirstMessageMongo    `bson:"first_message" json:"first_message"`
	ChecksummedData ChecksummedDataMongo `bson:"checksummed_data" json:"checksummed_data"`
}
type WriteObjectResponseMongo struct {
	PersistedSize int64 `bson:"persisted_size" json:"persisted_size"`
}
type ChecksummedDataMongo struct {
	Content bytes   `bson:"content" json:"content"`
	Crc32C  fixed32 `bson:"crc32c" json:"crc32c"`
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
