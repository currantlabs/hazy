package hazy

import "database/sql/driver"

func (id *ID) Scan(value interface{}) error {
	if value == nil {
		return nil
	}
	switch v := value.(type) {
	case uint64:
		id.Clear = v
		id.Hazy = obscure(v)
	case int64:
		id.Clear = uint64(v)
		id.Hazy = obscure(id.Clear)
	default:
		return ErrInvalidDBValue
	}
	return nil
}

func (id ID) Value() (driver.Value, error) {
	if id.Clear == 0 {
		return nil, nil
	}
	return int64(id.Clear), nil
}
