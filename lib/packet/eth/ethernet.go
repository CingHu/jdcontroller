package eth

import (
	"encoding/binary"
	"errors"
	"bytes"
	"fmt"

	"jd.com/jdcontroller/lib/buffer"
	_"jd.com/jdcontroller/lib/packet/arp"
	"jd.com/jdcontroller/lib/packet/ipv4"
)

// see http://en.wikipedia.org/wiki/EtherType

const BroadcastAddr = "ff:ff:ff:ff:ff:ff"

const (
	IPv4Msg = 0x0800
	ARPMsg  = 0x0806
	LLDPMsg = 0x88cc
	WOLMsg  = 0x0842
	RARPMsg = 0x8035
	VLANMsg = 0x8100

	IPv6Msg    = 0x86DD
	STPMsg     = 0x4242
	STPBPDUMsg = 0xAAAA
)

type Ethernet struct {
	HWDst     [6]byte
	HWSrc     [6]byte
	Ethertype uint16
	Data      buffer.Message
}

func New() (e *Ethernet) {
	e = new(Ethernet)
	e.Ethertype = 0x800
	//e.Data = nil
	return
}

func (e *Ethernet) Len() (l int) {
	l = 14
	return
}

func (e *Ethernet) PackBinary() (data []byte, err error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	binary.Write(buf, binary.BigEndian, e.HWDst)
	binary.Write(buf, binary.BigEndian, e.HWSrc)
	binary.Write(buf, binary.BigEndian, e.Ethertype)

	if e.Data != nil {
		datalen := uint16(e.Len()) + e.Data.(*ipv4.IPv4).Length
		bs := make([]byte, datalen)
		bs, err = e.Data.PackBinary()
		if err != nil {
			return
		}
		binary.Write(buf, binary.BigEndian, bs)
	}

	data = buf.Bytes()
	return
}

func (e *Ethernet) UnpackBinary(data []byte) (err error) {
	if len(data) < 14 {
		return errors.New("The []byte is too short to Unpack a full Ethernet message.")
	}
	buf := bytes.NewBuffer(data)
	binary.Read(buf, binary.BigEndian, &e.HWDst)
	binary.Read(buf, binary.BigEndian, &e.HWSrc)
	binary.Read(buf, binary.BigEndian, &e.Ethertype)

	datalen := len(data) - e.Len()
	switch e.Ethertype {
	case IPv4Msg:
		e.Data = new(ipv4.IPv4)

		bs := make([]byte, datalen)
		binary.Read(buf, binary.BigEndian, bs)
		e.Data.UnpackBinary(bs)
		//case ARPMsg:
		//e.Data = new(arp.ARP)
	default:
		e.Data = new(buffer.Buffer)
	}

	return
}

func Ethaddriszero(ea[]byte) bool {
	i := ea[0] | ea[1] | ea[2] | ea[3] | ea[4] | ea[5]
	if i == 0 {
		return true
	} else {
		return false
	}
}

func Ethmaskisexact(ea[]byte) bool {
	i := (ea[0] & ea[1] & ea[2] & ea[3] & ea[4] & ea[5])
	if i == 0xff {
		return true
	} else {
		return false
	}
}

func GetMacAddr(hw [6]byte) string {
	mac := ""
	for i, _ := range(hw) {
		var macAlpha1 byte = hw[i] & 0xF0 >> 4
		var macAlpha2 byte = hw[i] & 0x0F
		mac += fmt.Sprintf("%x", uint8(macAlpha1))
		mac += fmt.Sprintf("%x", uint8(macAlpha2))
		if i < 5 {
			mac += ":"
		}
	}
	return mac
}
