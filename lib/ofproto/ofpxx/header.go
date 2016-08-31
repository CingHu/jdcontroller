package ofpxx

import (
	"encoding/binary"
	"errors"
)

// Openflow message Xid
var messageXid uint32 = 1

// Returns a new OpenFlow header with version field set to v1.0.
func NewOFP10Header() (h *Header) {
	h = newHeader(1)
	return
}

// Returns a new OpenFlow header with version field set to v1.3.
func NewOFP13Header() (h *Header) {
	h = newHeader(4)
	return
}

func newHeader(ver int) (h *Header) {
	h = new(Header)
	h.Version = uint8(ver)
	h.Type = 0
	h.Length = 8

	messageXid += 1
	h.Xid = messageXid
	return
}

func (h *Header) Len() (l int) {
	l = 8
	return
}

func (h *Header) PackBinary() (data []byte, err error) {
	data = make([]byte, 8)
	data[0] = h.Version
	data[1] = h.Type
	binary.BigEndian.PutUint16(data[2:4], h.Length)
	binary.BigEndian.PutUint32(data[4:8], h.Xid)
	return
}

func (h *Header) UnpackBinary(data []byte) error {
	if len(data) < 4 {
		return errors.New("The []byte is too short to unmark a full HelloElemHeader.")
	}
	h.Version = data[0]
	h.Type = data[1]
	h.Length = binary.BigEndian.Uint16(data[2:4])
	h.Xid = binary.BigEndian.Uint32(data[4:8])
	return nil
}

func NewHelloElemHeader() *HelloElemHeader {
	h := new(HelloElemHeader)
	h.Type = HelloElemTypeVersionBitmap
	h.Length = 4
	return h
}

func (h *HelloElemHeader) Header() *HelloElemHeader {
	return h
}

func (h *HelloElemHeader) Len() (l int) {
	l = 4
	return
}

func (h *HelloElemHeader) PackBinary() (data []byte, err error) {
	data = make([]byte, 4)
	binary.BigEndian.PutUint16(data[:2], h.Type)
	binary.BigEndian.PutUint16(data[2:4], h.Length)
	return
}

func (h *HelloElemHeader) UnpackBinary(data []byte) error {
	if len(data) < 4 {
		return errors.New("The []byte is too short to Unpack a full HelloElemHeader.")
	}
	h.Type = binary.BigEndian.Uint16(data[:2])
	h.Length = binary.BigEndian.Uint16(data[2:4])
	return nil
}

func NewHelloElemVersionBitmap() (h *HelloElemVersionBitmap) {
	h = new(HelloElemVersionBitmap)
	h.HelloElemHeader = *NewHelloElemHeader()
	h.Bitmaps = make([]uint32, 0)
	// 1001
	// h.Bitmaps = append(h.Bitmaps, uint32(8) | uint32(1))
	h.Bitmaps = append(h.Bitmaps, uint32(1))
	h.Length = h.Length + uint16(len(h.Bitmaps)*4)
	return
}

func (h *HelloElemVersionBitmap) Header() *HelloElemHeader {
	return &h.HelloElemHeader
}

func (h *HelloElemVersionBitmap) Len() (l int) {
	l = h.HelloElemHeader.Len()
	l += len(h.Bitmaps) * 4
	return
}

func (h *HelloElemVersionBitmap) PackBinary() (data []byte, err error) {
	data = make([]byte, int(h.Len()))
	bytes := make([]byte, 0)
	next := 0

	bytes, err = h.HelloElemHeader.PackBinary()
	copy(data[next:], bytes)
	next += len(bytes)

	for _, m := range h.Bitmaps {
		binary.BigEndian.PutUint32(data[next:], m)
		next += 4
	}
	return
}

func (h *HelloElemVersionBitmap) UnpackBinary(data []byte) error {
	length := len(data)
	read := 0
	if err := h.HelloElemHeader.UnpackBinary(data[:4]); err != nil {
		return err
	}
	read += int(h.HelloElemHeader.Len())

	h.Bitmaps = make([]uint32, 0)
	for read < length {
		h.Bitmaps = append(h.Bitmaps, binary.BigEndian.Uint32(data[read:read+4]))
		read += 4
	}
	return nil
}

func NewHello(ver int) (h *Hello, err error) {
	h = new(Hello)
	//	h.Elements = make([]HelloElem, 0)

	if ver == 1 {
		h.Header = *NewOFP10Header()
	} else if ver == 4 {
		h.Header = *NewOFP13Header()
	} else {
		err = errors.New("New hello message with unsupported verion was attempted to be created.")
	}
	//	h.Elements = append(h.Elements, NewHelloElemVersionBitmap())
	return
}

func (h *Hello) Len() (l int) {
	l = h.Header.Len()
	//	for _, e := range h.Elements {
	//		n += e.Len()
	//	}
	return
}

func (h *Hello) PackBinary() (data []byte, err error) {
	data = make([]byte, int(h.Len()))
	bytes := make([]byte, 0)
	next := 0

	h.Header.Length = uint16(h.Len())
	bytes, err = h.Header.PackBinary()
	copy(data[next:], bytes)
	//	next += len(bytes)
	//
	//	for _, e := range h.Elements {
	//		bytes, err = e.PackBinary()
	//		copy(data[next:], bytes)
	//		next += len(bytes)
	//	}
	return
}

func (h *Hello) UnpackBinary(data []byte) error {
	next := 0
	err := h.Header.UnpackBinary(data[next:])
	//	next += int(h.Header.Len())
	//
	//	h.Elements = make([]HelloElem, 0)
	//	for next < len(data) {
	//		e := NewHelloElemHeader()
	//		e.UnpackBinary(data[next:])
	//
	//		switch e.Type {
	//		case HelloElemTypeVersionBitmap:
	//			v := NewHelloElemVersionBitmap()
	//			err = v.UnpackBinary(data[next:])
	//			next += int(v.Len())
	//			h.Elements = append(h.Elements, v)
	//		}
	//	}
	return err
}
