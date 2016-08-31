package ipv4

import (
	"encoding/binary"
	"errors"
	"net"
	"bytes"
	"fmt"

	"jd.com/jdcontroller/lib/buffer"
	"jd.com/jdcontroller/lib/packet/icmp"
	"jd.com/jdcontroller/lib/packet/udp"
)

const (
	ICMP     = 0x01
	TCP      = 0x06
	UDP      = 0x11
	IPv6     = 0x29
	IPv6ICMP = 0x3a
)

type IPv4 struct {
	Version        uint8 //4-bits
	IHL            uint8 //4-bits
	DSCP           uint8 //6-bits
	ECN            uint8 //2-bits
	Length         uint16
	Id             uint16
	Flags          uint16 //3-bits
	FragmentOffset uint16 //13-bits
	TTL            uint8
	Protocol       uint8
	Checksum       uint16
	NWSrc          net.IP
	NWDst          net.IP
	Options        buffer.Buffer
	Data           buffer.Message
}

func New() *IPv4 {
	ip := new(IPv4)
	ip.NWSrc = make([]byte, 4)
	ip.NWDst = make([]byte, 4)
	ip.Options = *new(buffer.Buffer)
	return ip
}

func (i *IPv4) Len() (l int) {
	i.IHL = 5
	if i.Data != nil {
		return int(i.IHL)*4 + i.Data.Len()
	}
	return int(i.IHL * 4)
}


func (i *IPv4) CheckSum() {
	var (
		sum uint32
		length int = 20
		index int
	)
	data, _ := i.PackBinary();
	for length > 1 {
		sum += uint32(data[index])<<8 + uint32(data[index + 1])
		index += 2
		length -= 2
	}
	if length > 0 {
		sum += uint32(data[index])
	}
	sum = (sum >> 16) + (sum & 0xffff)
	sum += (sum >> 16)

	i.Checksum = uint16(^sum)
}


func (i *IPv4) PackBinary() (data []byte, err error) {
	data = make([]byte, int(i.Len()))
	b := make([]byte, 0)
	n := 0

	var ihl uint8 = (i.Version << 4) + i.IHL
	data[n] = ihl
	n += 1
	var ecn uint8 = (i.DSCP << 2) + i.ECN
	data[n] = ecn
	n += 1
	binary.BigEndian.PutUint16(data[n:], i.Length)
	n += 2
	binary.BigEndian.PutUint16(data[n:], i.Id)
	n += 2
	var flg uint16 = (i.Flags << 13) + i.FragmentOffset
	binary.BigEndian.PutUint16(data[n:], flg)
	n += 2
	data[n] = i.TTL
	n += 1
	data[n] = i.Protocol
	n += 1
	binary.BigEndian.PutUint16(data[n:], i.Checksum)
	n += 2

	copy(data[n:], i.NWSrc.To4())
	n += 4 // Underlying rep can be 16 bytes.
	copy(data[n:], i.NWDst.To4())
	n += 4 // Underlying rep can be 16 bytes.

	b, err = i.Options.PackBinary()
	copy(data[n:], b)
	n += len(b)

	if i.Data != nil {
		b, err = i.Data.PackBinary()
		if err != nil {
			return
		}
		copy(data[n:], b)
		n += len(b)
	}
	return
}

func (i *IPv4) UnpackBinary(data []byte) (err error) {
	if len(data) < 20 {
		return errors.New("The []byte is too short to Unpack a full IPv4 message.")
	}
	n := 0

	var ihl uint8
	ihl = data[n]
	i.Version = ihl >> 4
	i.IHL = ihl & 0x0f
	n += 1

	var ecn uint8
	ecn = data[n]
	i.DSCP = ecn >> 2
	i.ECN = ecn & 0x03
	n += 1

	i.Length = binary.BigEndian.Uint16(data[n:])
	n += 2
	i.Id = binary.BigEndian.Uint16(data[n:])
	n += 2

	var flg uint16
	flg = binary.BigEndian.Uint16(data[n:])
	i.Flags = flg >> 13
	i.FragmentOffset = flg & 0x1fff
	n += 2

	i.TTL = data[n]
	n += 1
	i.Protocol = data[n]
	n += 1
	i.Checksum = binary.BigEndian.Uint16(data[n:])
	n += 2
	i.NWSrc = data[n : n+4]
	n += 4
	i.NWDst = data[n : n+4]
	n += 4

	if i.IHL * 4 > uint8(20) {
		i.Options.UnpackBinary(data[n:int(i.IHL * 4)])
		n += int(i.IHL * 4) - n
	}

	datalen := i.Length - uint16(n)
	switch i.Protocol {
	case ICMP:
		i.Data = icmp.New()
	case UDP:
		i.Data = udp.New()
		buf := bytes.NewBuffer(data[n:])

		bs := make([]byte, datalen)
		binary.Read(buf, binary.BigEndian, bs)
		fmt.Println("udp UnpackBinary", bs)
		i.Data.UnpackBinary(bs)

	default:
		i.Data = new(buffer.Buffer)
	}

	//return i.Data.UnpackBinary(data[n:])
	return err
}
