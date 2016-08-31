package ofp11

import (
	"bytes"
	"encoding/binary"
)

func NewErrorMsg() (e *ErrorMsg) {
	e = new(ErrorMsg)
	return
}

func (e *ErrorMsg) Len() (l int) {
	l = 12
	return
}

func (e *ErrorMsg) PackBinary() (data []byte, err error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	hs := make([]byte, 0)
	hs, err = e.Header.PackBinary()
	if err != nil {
		return
	}
	binary.Write(buf, binary.BigEndian, hs)
	binary.Write(buf, binary.BigEndian, e.Type)
	binary.Write(buf, binary.BigEndian, e.Code)

	for _, d := range e.Data {
		binary.Write(buf, binary.BigEndian, d)
	}
	data = buf.Bytes()
	return
}

func (e *ErrorMsg) UnpackBinary(data []byte) (err error) {
	buf := bytes.NewBuffer(data)
	hs := make([]byte, e.Header.Len())
	binary.Read(buf, binary.BigEndian, hs)
	err = e.Header.UnpackBinary(hs)
	if err != nil {
		return
	}
	binary.Read(buf, binary.BigEndian, &e.Type)
	binary.Read(buf, binary.BigEndian, &e.Code)
	n := e.Len()
	for n < len(data) {
		d := new(uint8)
		binary.Read(buf, binary.BigEndian, d)
		e.Data = append(e.Data, *d)
		n += 1
	}
	return
}
