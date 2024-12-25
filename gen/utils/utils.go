package utils

import (
	"encoding/json"
)

func MarshalUnmarshalTo(in, newIn any) error {
	inBytes, err := json.Marshal(in)
	if err != nil {
		return err
	}

	err = json.Unmarshal(inBytes, &newIn)
	if err != nil {
		return err
	}

	return nil
}
