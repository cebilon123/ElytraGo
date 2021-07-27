package tool

import (
	"bytes"
	"encoding/gob"
	"sync"
)

var lock = &sync.Mutex{}

// IStructToPayloadEncoder enables struct to be payload encoder.
// 	Encode returns byte array, and error (which can be nil in case if there is no error)
type IStructToPayloadEncoder interface {
	Encode() ([]byte, error)
	SetPayload(p interface{})
}

type structToPayloadEncoder struct {
	payload interface{}
	buf *bytes.Buffer
}

var instance IStructToPayloadEncoder

// GetPayloadEncoder returns encoder (as singleton) which can encode given object
func GetPayloadEncoder(p interface{}) *IStructToPayloadEncoder {
	if instance == nil {
		// goroutine safe
		lock.Lock()
		defer lock.Unlock()
		instance = &structToPayloadEncoder{buf: &bytes.Buffer{}}
	}

	instance.SetPayload(p)
	return &instance
}

func (s structToPayloadEncoder) SetPayload(p interface{}) {
	s.payload = p
}

func (s structToPayloadEncoder) Encode() ([]byte, error) {
	if s.buf == nil {
		s.buf = &bytes.Buffer{}
	}

	enc := gob.NewEncoder(s.buf)

	err := enc.Encode(s.payload)
	if err != nil {
		return nil, err
	}

	return s.buf.Bytes(), nil
}

