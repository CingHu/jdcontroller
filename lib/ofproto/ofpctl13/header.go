package ofpctl13

import (
	"bytes"
	"encoding/binary"

	"jd.com/jdcontroller/protocol/ofp13"
)

var messageXid uint32 = 1

func NewHeader() (h *ofp13.Header) {
	h = new(ofp13.Header)
	h.Version = ofp13.Version
	h.Type = 0
	h.Length = uint16(h.Len())
	messageXid += 1
	h.XID = messageXid
	return
}

func (h *ofp13.Header) Len() (l int) {
	l = 8
	return
}

func (h *ofp13.Header) PackBinary() (data []byte, err error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	binary.Write(buf, binary.BigEndian, h.Version)
	binary.Write(buf, binary.BigEndian, h.Type)
	binary.Write(buf, binary.BigEndian, h.Length)
	binary.Write(buf, binary.BigEndian, h.XID)
	data = buf.Bytes()
	return
}

func (h *ofp13.Header) UnpackBinary(data []byte) (err error) {
	buf := bytes.NewBuffer(data)
	binary.Read(buf, binary.BigEndian, &h.Version)
	binary.Read(buf, binary.BigEndian, &h.Type)
	binary.Read(buf, binary.BigEndian, &h.Length)
	binary.Read(buf, binary.BigEndian, &h.XID)
	return
}
