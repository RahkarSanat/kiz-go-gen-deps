package storage

import (
	"fmt"
)

var _ fmt.Stringer

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
