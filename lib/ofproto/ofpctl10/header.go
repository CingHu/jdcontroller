package ofpctl10

import (
	"bytes"
	"encoding/binary"
	"errors"

	"jd.com/jdcontroller/protocol/ofp10"
)

// Openflow message Xid
var messageXid uint32 = 1

//func NewHeader() (h *Header) {
	//h = new(Header)
	//return
//}

func NewOfp10Header() (h *ofp10.Header) {
	h = new(ofp10.Header)
	h.Version = uint8(1)
	h.Type = 0
	h.Length = 8

	messageXid += 1
	h.XID = messageXid
	return
}

func (h *ofp10.Header) Len() (l int) {
	l = 8
	return
}

func (h *ofp10.Header) PackBinary() (data []byte, err error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	binary.Write(buf, binary.BigEndian, h.Version)
	binary.Write(buf, binary.BigEndian, h.Type)
	binary.Write(buf, binary.BigEndian, h.Length)
	binary.Write(buf, binary.BigEndian, h.XID)
	data = buf.Bytes()
	return
}

func (h *ofp10.Header) UnpackBinary(data []byte) (err error) {
	if len(data) < 4 {
		return errors.New("The []byte is too short to unmark a full HelloElemHeader.")
	}
	buf := bytes.NewBuffer(data)
	binary.Read(buf, binary.BigEndian, &h.Version)
	binary.Read(buf, binary.BigEndian, &h.Type)
	binary.Read(buf, binary.BigEndian, &h.Length)
	binary.Read(buf, binary.BigEndian, &h.XID)
	return
}

func NewHello() (h *ofp10.Header) {
	h = new(ofp10.Header)
	return
}
