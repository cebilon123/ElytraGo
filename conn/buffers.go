package conn

import "strconv"

type Buffer interface {
	PullVarInt() int32
}

type buffer struct {
	iIndex int32
	oIndex int32

	bArray []byte
}

func NewBuffer() Buffer {
	return NewBufferWithBytes(make([]byte, 0))
}

func NewBufferWithBytes(b []byte) Buffer {
	return &buffer{bArray: b}
}

func (b *buffer) PullVarInt() int32 {
	return int32(b.pullVariable(5))
}

// internal

func (b *buffer) Len() int32 {
	return int32(len(b.bArray))
}

func (b *buffer) pullNext() byte {
	if b.iIndex >= b.Len() {
		return 0
		// panic("reached end of buffer")
	}

	next := b.bArray[b.iIndex]
	b.iIndex++

	if b.oIndex > 0 {
		b.oIndex--
	}

	return next
}

func (b *buffer) pullVariable(max int) int64 {
	var num int
	var res int64

	for {
		tmp := int64(b.pullNext())
		res |= (tmp & 0x7F) << uint(num*7)

		if num++; num > max {
			panic("VarInt > " + strconv.Itoa(max))
		}

		if tmp&0x80 != 0x80 {
			break
		}
	}

	return res
}