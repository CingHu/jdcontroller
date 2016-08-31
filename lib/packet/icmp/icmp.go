package icmp

import (
	"encoding/binary"
	"bytes"
	"errors"
)

type ICMP struct {
	Type     uint8
	Code     uint8
	Checksum uint16
	Data     []byte
}

func New() (i *ICMP) {
	i = new(ICMP)
	i.Data = make([]byte, 0)
	return
}

func (i *ICMP) Len() (l int) {
//	l = 4 + len(i.Data)
	l = 4
	return
}

func (i *ICMP) PackBinary() (data []byte, err error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	binary.Write(buf, binary.BigEndian, i.Type)
	binary.Write(buf, binary.BigEndian, i.Code)
	binary.Write(buf, binary.BigEndian, i.Checksum)
	for _, b := range i.Data {
		binary.Write(buf, binary.BigEndian, b)
	}
	return
}

func (i *ICMP) UnpackBinary(data []byte) error {
	if len(data) < 4 {
		return errors.New("The []byte is too short to Unpack a full ARP message.")
	}
	buf := bytes.NewBuffer(data)
	binary.Read(buf, binary.BigEndian, &i.Type)
	binary.Read(buf, binary.BigEndian, &i.Code)
	binary.Read(buf, binary.BigEndian, &i.Checksum)
	n := i.Len()
	for n < len(data) {
		b := new(byte)
		binary.Read(buf, binary.BigEndian, b)
		i.Data = append(i.Data, *b)
		n += 1
	}
	return nil
}
