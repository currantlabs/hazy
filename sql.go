package hazy

import (
	"database/sql/driver"
	"strconv"
)

func (id *ID) Scan(value interface{}) error {
	if value == nil {
		return nil
	}
	switch v := value.(type) {
	case uint64:
		*id = ID(obscure(v))
	case int64:
		*id = ID(obscure(uint64(v)))
	case []byte:
		clear, err := strconv.ParseUint(string(v), 10, 64)
		if err != nil {
			return err
		}
		*id = ID(obscure(clear))
	default:
		return ErrInvalidDBValue
	}
	return nil
}

func (id ID) Value() (driver.Value, error) {
	if id.IsZero() {
		return nil, nil
	}
	return int64(reveal(uint64(id))), nil
}
