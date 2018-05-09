package hazy

func (id ID) Bytes() []byte {
	b, _ := id.Marshal()
	return b
}

func (id ID) Marshal() ([]byte, error) {
	b := []byte{
		byte(id),
		byte(id >> 8),
		byte(id >> 16),
		byte(id >> 24),
		byte(id >> 32),
		byte(id >> 40),
		byte(id >> 48),
		byte(id >> 56),
	}
	return b, nil
}

func (id ID) MarshalTo(data []byte) (n int, err error) {
	if len(data) < 8 {
		return 0, ErrInvalidIDLength
	}
	data[0] = byte(id)
	data[1] = byte(id >> 8)
	data[2] = byte(id >> 16)
	data[3] = byte(id >> 24)
	data[4] = byte(id >> 32)
	data[5] = byte(id >> 40)
	data[6] = byte(id >> 48)
	data[7] = byte(id >> 56)
	return 8, nil
}

func (id *ID) Unmarshal(data []byte) error {
	if len(data) < 8 {
		*id = 0
		return ErrInvalidIDLength
	}
	*id = ID(uint64(data[0]) | (uint64(data[1]) << 8) | (uint64(data[2]) << 16) | (uint64(data[3]) << 24) | (uint64(data[4]) << 32) | (uint64(data[5]) << 40) | (uint64(data[6]) << 48) | (uint64(data[7]) << 56))
	return nil
}

func (id ID) Size() (n int) {
	return 8
}
