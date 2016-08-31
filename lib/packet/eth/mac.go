package eth

import (
	"bytes"
	"encoding/binary"
)

type HardwareAddr struct {
	Addr uint64
}

func NewHardwareAddr(i interface{}) (h *HardwareAddr) {
	h = new(HardwareAddr)
	switch i.(type) {
	case uint64:
		h.Addr = i.(uint64)
	case [6]byte:
		bs := make([]byte, 2)
		b := i.([6]byte)
		for _, bb := range b[:] {
			bs = append(bs, bb)
		}
		h.Addr = binary.BigEndian.Uint64(bs)
	}
	return
}

func (h *HardwareAddr) String() (s string) {
	const hexDigit = "0123456789abcdef"
	bs := make([]byte, 8)
	binary.BigEndian.PutUint64(bs, h.Addr)
	buf := make([]byte, 0, len(bs[2:])*3-1)
	for i, b := range bs[2:] {
		if i > 0 {
			buf = append(buf, ':')
		}
		buf = append(buf, hexDigit[b>>4])
		buf = append(buf, hexDigit[b&0xF])
	}
	s = string(buf)
	return
}

func (h *HardwareAddr) Array() (a [6]byte) {
	buf := bytes.NewBuffer(make([]byte, 0))
	binary.Write(buf, binary.BigEndian, h.Addr)
	bs := buf.Bytes()
	a = [...]byte{bs[2], bs[3], bs[4], bs[5], bs[6], bs[7]}
	return
}

func (h *HardwareAddr) Int() (i uint64) {
	i = h.Addr
	return
}
