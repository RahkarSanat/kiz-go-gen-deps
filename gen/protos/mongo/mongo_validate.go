package mongo

import (
	"fmt"
)

var _ fmt.Stringer

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
