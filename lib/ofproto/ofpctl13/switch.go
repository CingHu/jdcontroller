package ofpctl13

import (
	"bytes"
	"encoding/binary"

	"jd.com/jdcontroller/protocol/ofp13"
)

func NewFeaturesRequest() (h *ofp13.Header) {
	h = NewHeader()
	h.Type = uint8(ofp13.OFPTFeaturesRequest)
	return
}

func NewFeaturesReply() (s *ofp13.SwitchFeatures) {
	s = new(ofp13.SwitchFeatures)
	s.Header.Type = uint8(ofp13.OFPTFeaturesReply)
	return
}

func NewSwitchFeatures() (s *ofp13.SwitchFeatures) {
	s = new(ofp13.SwitchFeatures)
	return
}

func (s *ofp13.SwitchFeatures) Len() (l int) {
	l = 32
	return
}

func (s *ofp13.SwitchFeatures) PackBinary() (data []byte, err error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	hs := make([]byte, 0)
	hs, err = s.Header.PackBinary()
	if err != nil {
		return
	}
	binary.Write(buf, binary.BigEndian, hs)
	binary.Write(buf, binary.BigEndian, s.DPID)
	binary.Write(buf, binary.BigEndian, s.Buffers)
	binary.Write(buf, binary.BigEndian, s.Tables)
	binary.Write(buf, binary.BigEndian, s.pad)
	binary.Write(buf, binary.BigEndian, s.Capabilities)
	binary.Write(buf, binary.BigEndian, s.Reserved)
	data = buf.Bytes()
	return
}

func (s *ofp13.SwitchFeatures) UnpackBinary(data []byte) (err error) {
	buf := bytes.NewBuffer(data)
	hs := make([]byte, s.Header.Len())
	binary.Read(buf, binary.BigEndian, hs)
	err = s.Header.UnpackBinary(hs)
	if err != nil {
		return
	}
	binary.Read(buf, binary.BigEndian, &s.DPID)
	binary.Read(buf, binary.BigEndian, &s.Buffers)
	binary.Read(buf, binary.BigEndian, &s.Tables)
	binary.Read(buf, binary.BigEndian, &s.pad)
	binary.Read(buf, binary.BigEndian, &s.Capabilities)
	binary.Read(buf, binary.BigEndian, &s.Reserved)
	return
}

func NewSwitchConfig() (s *ofp13.SwitchConfig) {
	s = new(ofp13.SwitchConfig)
	s.Header = *NewHeader()
	s.Header.Length = uint16(s.Len())
	s.Header.Type = uint8(OFPTSetConfig)
	s.MissSendLen = 512

	return
}

func (s *ofp13.SwitchConfig) Len() (l int) {
	l = 12
	return
}

func (s *ofp13.SwitchConfig) PackBinary() (data []byte, err error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	header := make([]byte, 0)
	header, err = s.Header.PackBinary()
	if err != nil {
		return
	}
	binary.Write(buf, binary.BigEndian, header)
	binary.Write(buf, binary.BigEndian, s.Flags)
	binary.Write(buf, binary.BigEndian, s.MissSendLen)
	data = buf.Bytes()
	return
}

func (s *ofp13.SwitchConfig) UnpackBinary(data []byte) (err error) {
	buf := bytes.NewBuffer(data)
	hs := make([]byte, s.Header.Len())
	binary.Read(buf, binary.BigEndian, hs)
	err = s.Header.UnpackBinary(hs)
	if err != nil {
		return
	}
	binary.Read(buf, binary.BigEndian, &s.Flags)
	binary.Read(buf, binary.BigEndian, &s.MissSendLen)
	return
}
