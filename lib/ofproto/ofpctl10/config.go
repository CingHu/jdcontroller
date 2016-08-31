package ofpctl10

import (
	"bytes"
	"encoding/binary"
	"jd.com/jdcontroller/protocol/ofp10"
)

func NewConfigRequest() (h *ofp10.Header) {
	h = NewOfp10Header()
	h.Type = OFPTGetConfigRequest
	return
}

func NewSetConfig() (c *SwitchConfig) {
	c = new(SwitchConfig)
	c.Header.Type = uint8(OFPTSetConfig)
	c.MissSendLen = 512
	return
}

func (c *SwitchConfig) Len() (l int) {
	l = 12
	return
}

func (c *SwitchConfig) PackBinary() (data []byte, err error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	hs := make([]byte, 0)
	hs, err = c.Header.PackBinary()
	if err != nil {
		return
	}
	binary.Write(buf, binary.BigEndian, hs)
	binary.Write(buf, binary.BigEndian, c.Flags)
	binary.Write(buf, binary.BigEndian, c.MissSendLen)
	data = buf.Bytes()
	return
}

func (c *SwitchConfig) UnpackBinary(data []byte) (err error) {
	buf := bytes.NewBuffer(data)
	hs := make([]byte, c.Header.Len())
	binary.Read(buf, binary.BigEndian, hs)
	err = c.Header.UnpackBinary(hs)
	if err != nil {
		return
	}
	binary.Read(buf, binary.BigEndian, &c.Flags)
	binary.Read(buf, binary.BigEndian, &c.MissSendLen)
	return
}
