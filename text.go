package hazy

func (id ID) MarshalText() (text []byte, err error) {
	if id.IsZero() {
		return []byte(""), nil
	}
	return encode(id.Hazy), nil
}

func (id *ID) UnmarshalText(text []byte) error {
	if len(text) == 0 {
		id.Clear = 0
		id.Hazy = 0
		return nil
	}
	v, err := decode(text)
	if err != nil {
		return err
	}
	id.Hazy = v
	id.Clear = reveal(v)
	return nil
}
