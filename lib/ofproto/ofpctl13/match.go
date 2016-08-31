package ofpctl13

import (
	_"fmt"
	"bytes"
	"encoding/binary"

	"jd.com/jdcontroller/protocol/ofp13"
)

func NewMatch() (m *ofp13.Match) {
	m = new(ofp13.Match)
	return
}

func (m *ofp13.Match) Len() (l int) {
	l = 8
	return
}

func (m *ofp13.Match) PackBinary() (data []byte, err error) {
	//padLen := 8 - int(m.Length) % 8
	padLen := ((m.Length + 7) / 8 * 8) - m.Length

	for i := 0; i < int(padLen); i++ {
		m.pad = append(m.pad, uint8(0))
	}

	buf := bytes.NewBuffer(make([]byte, 0))
	binary.Write(buf, binary.BigEndian, m.Type)
	binary.Write(buf, binary.BigEndian, m.Length)
	for _, o := range m.OXMFields {
		binary.Write(buf, binary.BigEndian, o)
	}
	binary.Write(buf, binary.BigEndian, m.pad)
	data = buf.Bytes()
	return
}

func (m *ofp13.Match) UnpackBinary(data []byte) (err error) {
	buf := bytes.NewBuffer(data)

	for n := 0; n < len(data) - 4; n++ {
		b := new(uint8)
		binary.Read(buf, binary.BigEndian, b)
		m.OXMFields = append(m.OXMFields, *b)
	}

	binary.Read(buf, binary.BigEndian, &m.pad)
	return
}

func NewOXMExperimenterHeader() (o *OXMExperimenterHeader) {
	o = new(OXMExperimenterHeader)
	return
}

func (o *OXMExperimenterHeader) Len() (l int) {
	l = 8
	return
}

func (o *OXMExperimenterHeader) PackBinary() (data []byte, err error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	binary.Write(buf, binary.BigEndian, o.OXMHeader)
	binary.Write(buf, binary.BigEndian, o.Experimenter)
	data = buf.Bytes()
	return
}

func (o *OXMExperimenterHeader) UnpackBinary(data []byte) (err error) {
	buf := bytes.NewBuffer(data)
	binary.Read(buf, binary.BigEndian, &o.OXMHeader)
	binary.Read(buf, binary.BigEndian, &o.Experimenter)
	return
}
