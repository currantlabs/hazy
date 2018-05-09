package hazy

import (
	"bytes"
	"encoding/json"
)

var jsonNull = []byte("null")

func (id ID) MarshalJSON() ([]byte, error) {
	if id.IsZero() {
		return jsonNull, nil
	}
	return json.Marshal(string(Base32Encode(uint64(id))))
}

func (id *ID) UnmarshalJSON(input []byte) error {
	if bytes.Equal(input, jsonNull) {
		*id = Zero
		return nil
	}
	var s string
	err := json.Unmarshal(input, &s)
	if err != nil {
		return err
	}
	val, err := Base32Decode([]byte(s))
	if err != nil {
		return err
	}
	*id = ID(val)
	return nil
}
