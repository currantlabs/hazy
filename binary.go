package hazy

func (id ID) Bytes() []byte {
	b, _ := id.Marshal()
	return b
}

func (id ID) Marshal() ([]byte, error) {
	b := []byte{
		byte(id.Hazy),
		byte(id.Hazy >> 8),
		byte(id.Hazy >> 16),
		byte(id.Hazy >> 24),
		byte(id.Hazy >> 32),
		byte(id.Hazy >> 40),
		byte(id.Hazy >> 48),
		byte(id.Hazy >> 56),
	}
	return b, nil
}

func (id ID) MarshalTo(data []byte) (n int, err error) {
	if len(data) < 8 {
		return 0, ErrInvalidIDLength
	}
	data[0] = byte(id.Hazy)
	data[1] = byte(id.Hazy >> 8)
	data[2] = byte(id.Hazy >> 16)
	data[3] = byte(id.Hazy >> 24)
	data[4] = byte(id.Hazy >> 32)
	data[5] = byte(id.Hazy >> 40)
	data[6] = byte(id.Hazy >> 48)
	data[7] = byte(id.Hazy >> 56)
	return 8, nil
}

func (id *ID) Unmarshal(data []byte) error {
	if len(data) < 8 {
		id.Clear = 0
		id.Hazy = 0
		return ErrInvalidIDLength
	}
	id.Hazy = uint64(data[0]) | (uint64(data[1]) << 8) | (uint64(data[2]) << 16) | (uint64(data[3]) << 24) | (uint64(data[4]) << 32) | (uint64(data[5]) << 40) | (uint64(data[6]) << 48) | (uint64(data[7]) << 56)
	id.Clear = reveal(id.Hazy)
	return nil
}

func (id ID) Size() (n int) {
	return 8
}
