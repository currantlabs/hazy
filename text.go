package hazy

func (id ID) MarshalText() (text []byte, err error) {
	if id.IsZero() {
		return []byte(""), nil
	}
	return Base32Encode(uint64(id)), nil
}

func (id *ID) UnmarshalText(text []byte) error {
	if len(text) == 0 {
		*id = Zero
		return nil
	}
	v, err := Base32Decode(text)
	if err != nil {
		return err
	}
	*id = ID(v)
	return nil
}
