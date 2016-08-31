package arp

import (
	"encoding/binary"
	"errors"
	"bytes"
)

const (
	Request = 1
	Reply   = 2
)

type ARP struct {
	HWType      uint16
	ProtoType   uint16
	HWLength    uint8
	ProtoLength uint8
	Operation   uint16
	HWSrc       [6]byte
	IPSrc       [4]byte
	HWDst       [6]byte
	IPDst       [4]byte
}

func New(opt int) (a *ARP, err error) {
	if opt != Request && opt != Reply {
		return nil, errors.New("Invalid ARP Operation.")
	}
	a = new(ARP)
	a.HWType = 1
	a.ProtoType = 0x800
	a.HWLength = 6
	a.ProtoLength = 4
	a.Operation = uint16(opt)
	return
}

func (a *ARP) Len() (l int) {
	l = 28
//	l += int(a.HWLength)*2 + int(a.ProtoLength)*2
	return
}

func (a *ARP) PackBinary() (data []byte, err error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	binary.Write(buf, binary.BigEndian, a.HWType)
	binary.Write(buf, binary.BigEndian, a.ProtoType)
	binary.Write(buf, binary.BigEndian, a.HWLength)
	binary.Write(buf, binary.BigEndian, a.ProtoLength)
	binary.Write(buf, binary.BigEndian, a.Operation)
	binary.Write(buf, binary.BigEndian, a.HWSrc)
	binary.Write(buf, binary.BigEndian, a.IPSrc)
	binary.Write(buf, binary.BigEndian, a.HWDst)
	binary.Write(buf, binary.BigEndian, a.IPDst)
	data = buf.Bytes()
	return
}

func (a *ARP) UnpackBinary(data []byte) error {
	if len(data) < 8 {
		return errors.New("The []byte is too short to Unpack a full ARP message.")
	}
	buf := bytes.NewBuffer(data)
	binary.Read(buf, binary.BigEndian, &a.HWType)
	binary.Read(buf, binary.BigEndian, &a.ProtoType)
	binary.Read(buf, binary.BigEndian, &a.HWLength)
	binary.Read(buf, binary.BigEndian, &a.ProtoLength)
	binary.Read(buf, binary.BigEndian, &a.Operation)
	binary.Read(buf, binary.BigEndian, &a.HWSrc)
	binary.Read(buf, binary.BigEndian, &a.IPSrc)
	binary.Read(buf, binary.BigEndian, &a.HWDst)
	binary.Read(buf, binary.BigEndian, &a.IPDst)
	return nil
}
