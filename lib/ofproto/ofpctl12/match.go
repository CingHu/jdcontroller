package ofp12

import (
	"bytes"
	"encoding/binary"
)

func NewMatch() (m *Match) {
	m = new(Match)
	return
}

func (m *Match) Len() (l int) {
	l = 8
	return
}

func (m *Match) PackBinary() (data []byte, err error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	binary.Write(buf, binary.BigEndian, m.Type)
	binary.Write(buf, binary.BigEndian, m.Length)
	binary.Write(buf, binary.BigEndian, m.OXMFields)
	data = buf.Bytes()
	return
}

func (m *Match) UnpackBinary(data []byte) (err error) {
	buf := bytes.NewBuffer(data)
	binary.Read(buf, binary.BigEndian, &m.Type)
	binary.Read(buf, binary.BigEndian, &m.Length)
	binary.Read(buf, binary.BigEndian, &m.OXMFields)
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
