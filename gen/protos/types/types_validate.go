package types

import (
	"fmt"
)

var _ fmt.Stringer

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
