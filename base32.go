package hazy

const null = 0xFF

// Implementation of Crockford's Base32 encoding scheme (Hi Doug!)
// See http://www.crockford.com/wrmg/base32.html

// I, L, O and U excluded for confusion and decency reasons
var encodeChars = [32]byte{
	'0', '1', '2', '3', '4', '5', '6', '7',
	'8', '9', 'a', 'b', 'c', 'd', 'e', 'f',
	'g', 'h', 'j', 'k', 'm', 'n', 'p', 'q',
	'r', 's', 't', 'v', 'w', 'x', 'y', 'z',
}

// I, L and O are cast to their confusion doppelgangers, and all other invalid characters are rejected
var decodeVals = [256]byte{
	null, null, null, null, null, null, null, null,
	null, null, null, null, null, null, null, null,
	null, null, null, null, null, null, null, null,
	null, null, null, null, null, null, null, null,
	null, null, null, null, null, null, null, null,
	null, null, null, null, null, null, null, null,
	0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, // '0', '1', '2', '3', '4', '5', '6', '7'
	0x08, 0x09, null, null, null, null, null, null, // '8', '9'
	null, 0x0A, 0x0B, 0x0C, 0x0D, 0x0E, 0x0F, 0x10, //      'A', 'B', 'C', 'D', 'E', 'F', 'G'
	0x11, 0x01, 0x12, 0x13, 0x01, 0x14, 0x15, 0x00, // 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O' // I, L => 1, O => 0
	0x16, 0x17, 0x18, 0x19, 0x1A, null, 0x1B, 0x1C, // 'P', 'Q', 'R', 'S', 'T',      'V', 'W'
	0x1D, 0x1E, 0x1F, null, null, null, null, null, // 'X', 'Y', 'Z'
	null, 0x0A, 0x0B, 0x0C, 0x0D, 0x0E, 0x0F, 0x10, //      'a', 'b', 'c', 'd', 'e', 'f', 'g'
	0x11, 0x01, 0x12, 0x13, 0x01, 0x14, 0x15, 0x00, // 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o' // i, l => 1, o => 0
	0x16, 0x17, 0x18, 0x19, 0x1A, null, 0x1B, 0x1C, // 'p', 'q', 'r', 's', 't',      'v', 'w'
	0x1D, 0x1E, 0x1F, null, null, null, null, null, // 'x', 'y', 'z'
	null, null, null, null, null, null, null, null,
	null, null, null, null, null, null, null, null,
	null, null, null, null, null, null, null, null,
	null, null, null, null, null, null, null, null,
	null, null, null, null, null, null, null, null,
	null, null, null, null, null, null, null, null,
	null, null, null, null, null, null, null, null,
	null, null, null, null, null, null, null, null,
	null, null, null, null, null, null, null, null,
	null, null, null, null, null, null, null, null,
	null, null, null, null, null, null, null, null,
	null, null, null, null, null, null, null, null,
	null, null, null, null, null, null, null, null,
	null, null, null, null, null, null, null, null,
	null, null, null, null, null, null, null, null,
	null, null, null, null, null, null, null, null,
}

func Base32Encode(id uint64) []byte {
	const mask = 0x1F
	return []byte{
		encodeChars[byte(id>>0&mask)],
		encodeChars[byte(id>>5&mask)],
		encodeChars[byte(id>>10&mask)],
		encodeChars[byte(id>>15&mask)],
		encodeChars[byte(id>>20&mask)],
		encodeChars[byte(id>>25&mask)],
		encodeChars[byte(id>>30&mask)],
		encodeChars[byte(id>>35&mask)],
		encodeChars[byte(id>>40&mask)],
		encodeChars[byte(id>>45&mask)],
		encodeChars[byte(id>>50&mask)],
		encodeChars[byte(id>>55&mask)],
		encodeChars[byte(id>>60&mask)],
	}
}

func Base32Decode(s []byte) (uint64, error) {
	if len(s) != IDLength {
		return 0, ErrInvalidIDLength
	}
	var val uint64
	var shift uint64
	var b byte
	for i := 0; i < IDLength; i++ {
		b = decodeVals[s[i]]
		if b == null {
			return 0, ErrInvalidID
		}
		val |= uint64(b) << shift
		shift += 5
	}
	return val, nil
}
