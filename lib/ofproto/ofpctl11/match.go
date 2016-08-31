package ofp11

import (
	"bytes"
	"encoding/binary"
)

func NewMatch() (m *Match) {
	m = new(Match)
	return
}

func (m *Match) Len() (l int) {
	l = OFPMTStandardLen
	return
}

func (m *Match) PackBinary() (data []byte, err error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	binary.Write(buf, binary.BigEndian, m.Type)
	binary.Write(buf, binary.BigEndian, m.Length)
	binary.Write(buf, binary.BigEndian, m.InPort)
	binary.Write(buf, binary.BigEndian, m.Wildcards)
	binary.Write(buf, binary.BigEndian, m.DLSrc)
	binary.Write(buf, binary.BigEndian, m.DLSrcMask)
	binary.Write(buf, binary.BigEndian, m.DLDst)
	binary.Write(buf, binary.BigEndian, m.DLDstMask)
	binary.Write(buf, binary.BigEndian, m.DLVLAN)
	binary.Write(buf, binary.BigEndian, m.DLVLANPCP)
	binary.Write(buf, binary.BigEndian, m.pad)
	binary.Write(buf, binary.BigEndian, m.DLType)
	binary.Write(buf, binary.BigEndian, m.NWTos)
	binary.Write(buf, binary.BigEndian, m.NWProto)
	binary.Write(buf, binary.BigEndian, m.NWSrc)
	binary.Write(buf, binary.BigEndian, m.NWSrcMask)
	binary.Write(buf, binary.BigEndian, m.NWDst)
	binary.Write(buf, binary.BigEndian, m.NWDstMask)
	binary.Write(buf, binary.BigEndian, m.TPSrc)
	binary.Write(buf, binary.BigEndian, m.TPDst)
	binary.Write(buf, binary.BigEndian, m.MPLSLabel)
	binary.Write(buf, binary.BigEndian, m.MPLSTC)
	binary.Write(buf, binary.BigEndian, m.pad2)
	binary.Write(buf, binary.BigEndian, m.Metadata)
	binary.Write(buf, binary.BigEndian, m.MetadataMask)
	data = buf.Bytes()
	return
}

func (m *Match) UnpackBinary(data []byte) (err error) {
	buf := bytes.NewBuffer(data)
	binary.Read(buf, binary.BigEndian, &m.Type)
	binary.Read(buf, binary.BigEndian, &m.Length)
	binary.Read(buf, binary.BigEndian, &m.InPort)
	binary.Read(buf, binary.BigEndian, &m.Wildcards)
	binary.Read(buf, binary.BigEndian, &m.DLSrc)
	binary.Read(buf, binary.BigEndian, &m.DLSrcMask)
	binary.Read(buf, binary.BigEndian, &m.DLDst)
	binary.Read(buf, binary.BigEndian, &m.DLDstMask)
	binary.Read(buf, binary.BigEndian, &m.DLVLAN)
	binary.Read(buf, binary.BigEndian, &m.DLVLANPCP)
	binary.Read(buf, binary.BigEndian, &m.pad)
	binary.Read(buf, binary.BigEndian, &m.DLType)
	binary.Read(buf, binary.BigEndian, &m.NWTos)
	binary.Read(buf, binary.BigEndian, &m.NWProto)
	binary.Read(buf, binary.BigEndian, &m.NWSrc)
	binary.Read(buf, binary.BigEndian, &m.NWSrcMask)
	binary.Read(buf, binary.BigEndian, &m.NWDst)
	binary.Read(buf, binary.BigEndian, &m.NWDstMask)
	binary.Read(buf, binary.BigEndian, &m.TPSrc)
	binary.Read(buf, binary.BigEndian, &m.TPDst)
	binary.Read(buf, binary.BigEndian, &m.MPLSLabel)
	binary.Read(buf, binary.BigEndian, &m.MPLSTC)
	binary.Read(buf, binary.BigEndian, &m.pad2)
	binary.Read(buf, binary.BigEndian, &m.Metadata)
	binary.Read(buf, binary.BigEndian, &m.MetadataMask)
	return
}
