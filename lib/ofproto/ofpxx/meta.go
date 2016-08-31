package ofpxx

import (
	"jd.com/jdcontroller/lib/buffer"
)

const (
	reserved = iota
	HelloElemTypeVersionBitmap
)

type Header struct {
	Version uint8
	Type    uint8
	Length  uint16
	Xid     uint32
}

type Hello struct {
	Header
	//	Elements []HelloElem
}

type HelloElem interface {
	Header() *HelloElemHeader
	buffer.Message
}

type HelloElemHeader struct {
	Type   uint16
	Length uint16
}

type HelloElemVersionBitmap struct {
	HelloElemHeader
	Bitmaps []uint32
}
