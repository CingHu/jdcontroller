package ofpctl13

import (
	"bytes"
	"encoding/binary"

	"jd.com/jdcontroller/lib/ofproto/ofp"
	"jd.com/jdcontroller/protocol/ofp13"
)

func NewHello() (h *ofp13.Hello) {
	h = new(ofp13.Hello)
	h.Header.Version = ofp13.Version
	h.Header.Type = uint8(ofp13.OFPTHello)
	h.Header.Length = uint16(h.Len())
	h.Header.XID = <-ofp.XID
	return
}

func (h *ofp13.Hello) Len() (l int) {
	l = 8
	return
}

func (h *ofp13.Hello) PackBinary() (data []byte, err error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	hs := make([]byte, 0)
	hs, err = h.Header.PackBinary()
	if err != nil {
		return
	}
	binary.Write(buf, binary.BigEndian, hs)

	for _, e := range h.Elements {
		bs := make([]byte, 0)
		switch e.(type) {
		case *ofp13.HelloElemVersionBitmap:
			bs, err = e.(*ofp13.HelloElemVersionBitmap).PackBinary()
		}
		if err != nil {
			return
		}
		binary.Write(buf, binary.BigEndian, bs)
	}
	data = buf.Bytes()
	return
}

func (h *ofp13.Hello) UnpackBinary(data []byte) (err error) {
	buf := bytes.NewBuffer(data)
	hs := make([]byte, h.Header.Len())
	binary.Read(buf, binary.BigEndian, hs)
	err = h.Header.UnpackBinary(hs)
	if err != nil {
		return
	}
	n := h.Len()
	for n < len(data) {
		var e ofp13.HelloElem
		var l int
		switch binary.BigEndian.Uint16(data[n:]) {
		case OFPHETVersionBitmap:
			e = new(ofp13.HelloElemVersionBitmap)
			r := e.(*ofp13.HelloElemVersionBitmap)
			bs := make([]byte, r.Len())
			binary.Read(buf, binary.BigEndian, bs)
			r.UnpackBinary(bs)
			l = r.Len()
		}
		h.Elements = append(h.Elements, e)
		n += l
	}
	return
}

func NewHelloElemHeader() (h *ofp13.HelloElemHeader) {
	h = new(ofp13.HelloElemHeader)
	return
}

func (h *ofp13.HelloElemHeader) Len() (l int) {
	l = 4
	return
}

func (h *ofp13.HelloElemHeader) PackBinary() (data []byte, err error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	binary.Write(buf, binary.BigEndian, h.Type)
	binary.Write(buf, binary.BigEndian, h.Length)
	data = buf.Bytes()
	return
}

func (h *ofp13.HelloElemHeader) UnpackBinary(data []byte) (err error) {
	buf := bytes.NewBuffer(data)
	binary.Read(buf, binary.BigEndian, &h.Type)
	binary.Read(buf, binary.BigEndian, &h.Length)
	return
}

func NewHelloElemVersionBitmap() (h *ofp13.HelloElemVersionBitmap) {
	h = new(ofp13.HelloElemVersionBitmap)
	return
}

func (h *ofp13.HelloElemVersionBitmap) Len() (l int) {
	l = 4
	return
}

func (h *ofp13.HelloElemVersionBitmap) PackBinary() (data []byte, err error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	hs := make([]byte, 0)
	hs, err = h.Header.PackBinary()
	if err != nil {
		return
	}
	binary.Write(buf, binary.BigEndian, hs)

	for _, b := range h.Bitmaps {
		binary.Write(buf, binary.BigEndian, b)
	}
	data = buf.Bytes()
	return
}

func (h *ofp13.HelloElemVersionBitmap) UnpackBinary(data []byte) (err error) {
	buf := bytes.NewBuffer(data)
	hs := make([]byte, h.Header.Len())
	binary.Read(buf, binary.BigEndian, hs)
	err = h.Header.UnpackBinary(hs)
	if err != nil {
		return
	}
	n := h.Len()
	for n < len(data) {
		b := new(uint32)
		binary.Read(buf, binary.BigEndian, b)
		h.Bitmaps = append(h.Bitmaps, *b)
		n += 4
	}
	return
}
