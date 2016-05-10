package hazy

import (
	"bytes"
	"encoding/json"
)

var jsonNull = []byte("null")

func (id ID) MarshalJSON() ([]byte, error) {
	if id.Clear == 0 {
		return jsonNull, nil
	}
	return json.Marshal(string(Base32Encode(id.Hazy)))
}

func (id *ID) UnmarshalJSON(input []byte) error {
	if bytes.Equal(input, jsonNull) {
		id.Clear = 0
		id.Hazy = 0
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
	id.Hazy = val
	id.Clear = reveal(val)
	return nil
}
