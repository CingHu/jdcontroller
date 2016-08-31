package ofp11

import (
	"bytes"
	"encoding/binary"
)

func NewFeaturesRequest() (h *Header) {
	h = new(Header)
	h.Type = uint8(OFPTFeaturesRequest)
	return
}

func NewFeaturesReply() (s *SwitchFeatures) {
	s = new(SwitchFeatures)
	s.Header.Type = uint8(OFPTFeaturesReply)
	return
}

func NewSwitchFeatures() (s *SwitchFeatures) {
	s = new(SwitchFeatures)
	return
}

func (s *SwitchFeatures) Len() (l int) {
	l = 32
	return
}

func (s *SwitchFeatures) PackBinary() (data []byte, err error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	header := make([]byte, 0)
	header, err = s.Header.PackBinary()
	if err != nil {
		return
	}
	binary.Write(buf, binary.BigEndian, header)
	binary.Write(buf, binary.BigEndian, s.DPID)
	binary.Write(buf, binary.BigEndian, s.Buffers)
	binary.Write(buf, binary.BigEndian, s.Tables)
	binary.Write(buf, binary.BigEndian, s.pad2)
	binary.Write(buf, binary.BigEndian, s.Capabilities)
	binary.Write(buf, binary.BigEndian, s.Reserved)

	for _, p := range s.Ports {
		data, err = p.PackBinary()
		if err != nil {
			return
		}
		binary.Write(buf, binary.BigEndian, data)
	}
	data = buf.Bytes()
	return
}

func (s *SwitchFeatures) UnpackBinary(data []byte) (err error) {
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
	binary.Read(buf, binary.BigEndian, &s.pad2)
	binary.Read(buf, binary.BigEndian, &s.Capabilities)
	binary.Read(buf, binary.BigEndian, &s.Reserved)
	n := s.Len()
	for n < len(data) {
		p := new(Port)
		ps := make([]byte, p.Len())
		binary.Read(buf, binary.BigEndian, ps)
		p.UnpackBinary(ps)
		s.Ports = append(s.Ports, *p)
		n += p.Len()
	}
	return
}

func NewSwitchConfig() (s *SwitchConfig) {
	s = new(SwitchConfig)
	return
}

func (s *SwitchConfig) Len() (l int) {
	l = 12
	return
}

func (s *SwitchConfig) PackBinary() (data []byte, err error) {
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

func (s *SwitchConfig) UnpackBinary(data []byte) (err error) {
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
