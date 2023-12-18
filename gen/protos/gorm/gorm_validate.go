package gorm

import (
	"fmt"
)

var _ fmt.Stringer

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
