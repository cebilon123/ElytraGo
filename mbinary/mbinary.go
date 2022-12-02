package mbinary

import (
	"errors"
	"log"
)

const SEGMENT_BITS = 0x7f
const CONTINUE_BIT = 0x80

// VarInt decodes an int64 from buf and returns that value and the
// number of bytes read (> 0). If an error occurred, the value is -1
// and the number of bytes n is -1
// Important: Minecraft is coding bytes in a little bit different way than
// protocol buffer so binary is useless.
func VarInt(b []byte) (int64, int) {
	value := int64(0)
	position := 0
	var currentByte byte

	for {
		currentByte = b[position]
		value |= (int64(currentByte) & SEGMENT_BITS) << int64(position)
		if (currentByte & CONTINUE_BIT) == 0 {
			break
		}
		position += 7
		if position >= 32 {
			err := errors.New("VarInt to big")
			log.Println(err)
			return -1, -1
		}
	}
	return value, position
}

// VarText pulls text from given array
func VarText(b []byte) (string, int64) {
	size, readIndx := VarInt(b)
	txt := b[readIndx : int64(readIndx)+size]

	return string(txt), size
}
