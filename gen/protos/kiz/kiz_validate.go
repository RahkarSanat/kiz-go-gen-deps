package kiz

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var _ primitive.A
var _ fmt.Stringer

type InterceptorOptionsMongo struct {
	ForceAll bool `bson:"force_all" json:"force_all"`
}

func (m *InterceptorOptions) Validate() error {

	return nil
}
