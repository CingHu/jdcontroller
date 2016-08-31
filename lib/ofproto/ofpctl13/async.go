package ofpctl13

import (
	"bytes"
	"encoding/binary"

	"jd.com/jdcontroller/protocol/ofp13"
)

func NewAsyncConfig() (c *ofp13.AsyncConfig) {
	c = new(ofp13.AsyncConfig)
	return
}

func (c *ofp13.AsyncConfig) Len() (l int) {
	l = 32
	return
}

func (c *ofp13.AsyncConfig) PackBinary() (data []byte, err error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	hs := make([]byte, 0)
	hs, err = c.Header.PackBinary()
	if err != nil {
		return
	}
	binary.Write(buf, binary.BigEndian, hs)
	binary.Write(buf, binary.BigEndian, c.PacketInMask)
	binary.Write(buf, binary.BigEndian, c.PortStatusMask)
	binary.Write(buf, binary.BigEndian, c.FlowRemovedMask)
	data = buf.Bytes()
	return
}

func (c *ofp13.AsyncConfig) UnpackBinary(data []byte) (err error) {
	buf := bytes.NewBuffer(data)
	hs := make([]byte, c.Header.Len())
	binary.Read(buf, binary.BigEndian, hs)
	err = c.Header.UnpackBinary(hs)
	if err != nil {
		return
	}
	binary.Read(buf, binary.BigEndian, &c.PacketInMask)
	binary.Read(buf, binary.BigEndian, &c.PortStatusMask)
	binary.Read(buf, binary.BigEndian, &c.FlowRemovedMask)
	return
}
