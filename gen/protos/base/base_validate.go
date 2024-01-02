package base

import (
	"fmt"
)

var _ fmt.Stringer

func (m *CountRequest) Validate() error {

	return nil
}

func (m *FindRequest) Validate() error {

	return nil
}

func (m *FindByIDRequest) Validate() error {

	if m.Id == "" {
		return fmt.Errorf("Id is required")
	}

	return nil
}

func (m *CustomFileOptions) Validate() error {

	return nil
}

func (m *MongoOption) Validate() error {

	return nil
}

func (m *AggregateOption) Validate() error {

	return nil
}

func (m *ViewOption) Validate() error {

	return nil
}

func (m *AclOption) Validate() error {

	return nil
}

func (m *Sort) Validate() error {

	return nil
}

func (m *Pagination) Validate() error {

	return nil
}

func (m *Filter) Validate() error {

	return nil
}

func (m *Operations) Validate() error {

	return nil
}

func (m *Total) Validate() error {

	return nil
}

func (m *BaseAccess) Validate() error {

	if m.Owner == "" {
		return fmt.Errorf("Owner is required")
	}

	if m.Clients == nil {
		return fmt.Errorf("Clients is required")
	}

	// ali

	return nil
}

func (m *BaseProperties) Validate() error {

	return nil
}

func (m *BaseDates) Validate() error {

	return nil
}

func (m *BaseAccessUpdate) Validate() error {

	return nil
}

func (m *BasePropertiesUpdate) Validate() error {

	return nil
}

func (m *BaseDatesUpdate) Validate() error {

	return nil
}
