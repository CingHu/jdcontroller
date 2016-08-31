package ofpctl10

import (
	"bytes"
	"encoding/binary"

	"jd.com/jdcontroller/protocol/ofp10"
)

// FeaturesRequest constructor
func NewFeaturesRequest() (h *ofp10.Header) {
	h = NewOfp10Header()
	h.Type = ofp10.OFPTFeaturesRequest
	return
}

// FeaturesReply constructor
func NewFeaturesReply() (r *ofp10.SwitchFeatures) {
	r = new(ofp10.SwitchFeatures)
	return
}

func (s *ofp10.SwitchFeatures) Len() (l int) {
	l = 32
	return
}

func (s *ofp10.SwitchFeatures) PackBinary() (data []byte, err error) {
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
	binary.Write(buf, binary.BigEndian, s.Actions)
	for _, p := range s.Ports {
		bs := make([]byte, 0)
		bs, err = p.PackBinary()
		if err != nil {
			return
		}
		binary.Write(buf, binary.BigEndian, bs)
	}

	data = buf.Bytes()
	return
}

func (s *ofp10.SwitchFeatures) UnpackBinary(data []byte) (err error) {
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
	binary.Read(buf, binary.BigEndian, &s.Actions)

	n := s.Len()
	for n < len(data) {
		b := new(Port)
		binary.Read(buf, binary.BigEndian, b)
		s.Ports = append(s.Ports, *b)
		n += b.Len()
	}
	return
}
