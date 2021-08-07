package mbinary

// VarInt decodes an int64 from buf and returns that value and the
// number of bytes read (> 0). If an error occurred, the value is 0
// and the number of bytes n is <= 0 with the following meaning:
//
// 	n == 0: buf too small
// 	n  < 0: value larger than 64 bits (overflow)
// 	        and -n is the number of bytes read
//
// Important: Minecraft is coding bytes in a little bit different way than
// protocol buffer so binary is useless.
func VarInt(b []byte) (int64, int) {
	var value = 0
	var bitOffset byte = 0
	var currIndx = 0
	var currentByte byte

	for {
		if bitOffset == 35 {
			return 0, 0
		}

		currentByte = b[currIndx]
		value |= int((currentByte & 0b01111111) << bitOffset)

		currIndx++
		bitOffset += 7

		if currentByte & 0b10000000 == 0 {
			break
		}
	}

	return int64(value), currIndx
}
