package ofp12

import (
	"bytes"
	"encoding/binary"
	"jd.com/jdcontroller/protocol"
)

func NewHeader() (h *Header) {
	h = new(Header)
	return
}

func (h *Header) Len() (l int) {
	l = 8
	return
}

func (h *Header) PackBinary() (data []byte, err error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	binary.Write(buf, binary.BigEndian, h.Version)
	binary.Write(buf, binary.BigEndian, h.Type)
	binary.Write(buf, binary.BigEndian, h.Length)
	binary.Write(buf, binary.BigEndian, h.XID)
	data = buf.Bytes()
	return
}

func (h *Header) UnpackBinary(data []byte) (err error) {
	buf := bytes.NewBuffer(data)
	binary.Read(buf, binary.BigEndian, &h.Version)
	binary.Read(buf, binary.BigEndian, &h.Type)
	binary.Read(buf, binary.BigEndian, &h.Length)
	binary.Read(buf, binary.BigEndian, &h.XID)
	return
}
