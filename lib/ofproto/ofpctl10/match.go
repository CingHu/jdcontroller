package ofpctl10

import (
	"bytes"
	"encoding/binary"

	"jd.com/jdcontroller/protocol/ofp10"
)

func NewMatch() (m *ofp10.Match) {
	m = new(ofp10.Match)
	// By default wildcard all fields
	m.Wildcards = OFPFWAll
	return
}

func (m *ofp10.Match) Len() (l int) {
	return 40
}

func (m *ofp10.Match) PackBinary() (data []byte, err error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	binary.Write(buf, binary.BigEndian, m.Wildcards)
	binary.Write(buf, binary.BigEndian, m.InPort)
	binary.Write(buf, binary.BigEndian, m.DLSrc)
	binary.Write(buf, binary.BigEndian, m.DLDst)
	binary.Write(buf, binary.BigEndian, m.DLVLAN)
	binary.Write(buf, binary.BigEndian, m.DLVLANPCP)
	binary.Write(buf, binary.BigEndian, m.pad)
	binary.Write(buf, binary.BigEndian, m.DLType)
	binary.Write(buf, binary.BigEndian, m.NWTos)
	binary.Write(buf, binary.BigEndian, m.NWProto)
	binary.Write(buf, binary.BigEndian, m.pad2)
	binary.Write(buf, binary.BigEndian, m.NWSrc)
	binary.Write(buf, binary.BigEndian, m.NWDst)
	binary.Write(buf, binary.BigEndian, m.TPSrc)
	binary.Write(buf, binary.BigEndian, m.TPDst)
	data = buf.Bytes()
	return
}

func (m *ofp10.Match) UnpackBinary(data []byte) (err error) {
	buf := bytes.NewBuffer(data)
	binary.Read(buf, binary.BigEndian, &m.Wildcards)
	binary.Read(buf, binary.BigEndian, &m.InPort)
	binary.Read(buf, binary.BigEndian, &m.DLSrc)
	binary.Read(buf, binary.BigEndian, &m.DLDst)
	binary.Read(buf, binary.BigEndian, &m.DLVLAN)
	binary.Read(buf, binary.BigEndian, &m.DLVLANPCP)
	binary.Read(buf, binary.BigEndian, &m.pad)
	binary.Read(buf, binary.BigEndian, &m.DLType)
	binary.Read(buf, binary.BigEndian, &m.NWTos)
	binary.Read(buf, binary.BigEndian, &m.NWProto)
	binary.Read(buf, binary.BigEndian, &m.pad2)
	binary.Read(buf, binary.BigEndian, &m.NWSrc)
	binary.Read(buf, binary.BigEndian, &m.NWDst)
	binary.Read(buf, binary.BigEndian, &m.TPSrc)
	binary.Read(buf, binary.BigEndian, &m.TPDst)
	return
}
