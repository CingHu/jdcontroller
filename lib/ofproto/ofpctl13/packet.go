package ofpctl13

import (
	_"fmt"
	"bytes"
	"strconv"
	"strings"
	"net"
	"encoding/binary"
	
	"jd.com/jdcontroller/lib/packet/ipv4"
	"jd.com/jdcontroller/protocol/ofp13"
)

func NewPacketOut() (p *ofp13.PacketOut) {
	p = new(ofp13.PacketOut)
	p.Header.Version = ofp13.Version
	p.Header.Type = uint8(ofp13.OFPTPacketOut)
	p.Header.Length = uint16(p.Len())
	return
}

func (p *ofp13.PacketOut) AddSetEthSrcField(EthSrc string) {
	a := p.NewActionSetField()
	a.SetFieldInit(uint32(ofp13.OFPXMCOpenFlowBasic), uint8(ofp13.OFPXMTFOFBEthDst), uint8(ofp13.OFPXMTFOFBEthLEN))
	ethsrcfield, _ := net.ParseMAC(EthSrc)
	ethsrc := []byte(ethsrcfield)
	a.Field = append(a.Field, ethsrc...)
	a.Header.Length += (ofp13.OFPXMHEADERLEN + ofp13.OFPXMTFOFBEthLEN)

	p.AddSetField(a)
	return
}

func(p *ofp13.PacketOut) AddSetEthDstField(EthDst string) {
	a := p.NewActionSetField()
	a.SetFieldInit(uint32(ofp13.OFPXMCOpenFlowBasic), uint8(ofp13.OFPXMTFOFBEthDst), uint8(ofp13.OFPXMTFOFBEthLEN))
	ethdstfield, _ := net.ParseMAC(EthDst)
	ethdst := []byte(ethdstfield)
	a.Field = append(a.Field, ethdst...)
	a.Header.Length += (ofp13.OFPXMHEADERLEN + ofp13.OFPXMTFOFBEthLEN)

	p.AddSetField(a)
	return
}

func(p *ofp13.PacketOut) AddSetIpSrcField(ipaddr string) {
	a := p.NewActionSetField()
	var ipv4src []uint8
	a.SetFieldInit(uint32(ofp13.OFPXMCOpenFlowBasic), uint8(ofp13.OFPXMTFOFBIPv4Src), uint8(ofp13.OFPXMTFOFBIPv4LEN))

	addr := strings.Split(ipaddr, "/")
	ip := strings.Split(addr[0], ".")
	for _, ipItem := range(ip) {
		ipItemNum, _ := strconv.Atoi(ipItem)
		ipv4src = append(ipv4src, uint8(ipItemNum))
	}
	a.Header.Length += (ofp13.OFPXMHEADERLEN + ofp13.OFPXMTFOFBIPv4LEN)
	a.Field = append(a.Field, ipv4src...)
	p.AddSetField(a)
	return
}

func(p *ofp13.PacketOut) AddSetIpDstField(ipaddr string) {
	a := p.NewActionSetField()
	var ipv4dst []uint8
	a.SetFieldInit(uint32(ofp13.OFPXMCOpenFlowBasic), uint8(ofp13.OFPXMTFOFBIPv4Dst), uint8(ofp13.OFPXMTFOFBIPv4LEN))

	addr := strings.Split(ipaddr, "/")
	ip := strings.Split(addr[0], ".")
	for _, ipItem := range(ip) {
		ipItemNum, _ := strconv.Atoi(ipItem)
		ipv4dst = append(ipv4dst, uint8(ipItemNum))
	}
	a.Header.Length += (ofp13.OFPXMHEADERLEN + ofp13.OFPXMTFOFBIPv4LEN)
	a.Field = append(a.Field, ipv4dst...)
	p.AddSetField(a)
	return
}

func(p *ofp13.PacketOut) AddSetTunnelIpSrcField(ipaddr string) {
	a := p.NewActionSetField()
	var ipv4src []uint8
	a.SetFieldInit(uint32(ofp13.OFPXMCNXM1), uint8(NXMNXTUNIPV4SRC), uint8(NXMNXTUNIPV4LEN))

	addr := strings.Split(ipaddr, "/")
	ip := strings.Split(addr[0], ".")
	for _, ipItem := range(ip) {
		ipItemNum, _ := strconv.Atoi(ipItem)
		ipv4src = append(ipv4src, uint8(ipItemNum))
	}
	a.Header.Length += (ofp13.OFPXMHEADERLEN + NXMNXTUNIPV4LEN)
	a.Field = append(a.Field, ipv4src...)
	p.AddSetField(a)
	return
}

func(p *ofp13.PacketOut) AddSetTunnelIpDstField(ipaddr string) {
	a := p.NewActionSetField()
	var ipv4dst []uint8
	a.SetFieldInit(uint32(ofp13.OFPXMCNXM1), uint8(NXMNXTUNIPV4DST), uint8(NXMNXTUNIPV4LEN))

	addr := strings.Split(ipaddr, "/")
	ip := strings.Split(addr[0], ".")
	for _, ipItem := range(ip) {
		ipItemNum, _ := strconv.Atoi(ipItem)
		ipv4dst = append(ipv4dst, uint8(ipItemNum))
	}
	a.Header.Length += (ofp13.OFPXMHEADERLEN + NXMNXTUNIPV4LEN)
	a.Field = append(a.Field, ipv4dst...)
	p.AddSetField(a)
	return
}

func (p *ofp13.PacketOut)AddOutputAction(outport uint32) {
	o := new(ActionOutput)
	o.Header.Type = uint16(OFPATOutput)
	o.Header.Length = uint16(o.Len())
	o.Port = outport
	o.MaxLen = 0

	p.Actions = append(p.Actions, o)
	p.ActionsLen += uint16(o.Len())
	p.Header.Length += uint16(o.Len())
}

func(p *ofp13.PacketOut)NewActionSetField() (a *ActionSetField) {
	a = new(ActionSetField)
	a.Header.Type = uint16(OFPATSetField)
	a.Header.Length = uint16(a.Len())
	return
}

func(p *ofp13.PacketOut)AddSetField(setfieldact *ActionSetField) {
	//padLen := 8 - int(setfieldact.Header.Length) % 8
	padLen := ((setfieldact.Header.Length + 7) / 8 * 8) - setfieldact.Header.Length
	setfieldact.Header.Length += uint16(padLen)

	for i := 0; i < int(padLen); i++ {
		setfieldact.Field = append(setfieldact.Field, uint8(0))
	}

	p.Actions = append(p.Actions, setfieldact)
	p.ActionsLen += uint16(setfieldact.Header.Length)
	p.Header.Length += uint16(setfieldact.Header.Length)

	return
}

func (p *ofp13.PacketOut) Len() (l int) {
	l = 24
	return
}

func (p *ofp13.PacketOut) PackBinary() (data []byte, err error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	hs := make([]byte, 0)
	hs, err = p.Header.PackBinary()
	if err != nil {
		return
	}
	binary.Write(buf, binary.BigEndian, hs)
	binary.Write(buf, binary.BigEndian, p.BufferID)
	binary.Write(buf, binary.BigEndian, p.InPort)
	binary.Write(buf, binary.BigEndian, p.ActionsLen)
	binary.Write(buf, binary.BigEndian, p.pad)

	for _, a := range p.Actions {
		bs := make([]byte, 0)
		bs, err = EncodeAction(a)
		if err != nil {
			return
		}
		binary.Write(buf, binary.BigEndian, bs)
	}

	if p.BufferID == 0xffffffff {
		ethPkt := p.Data
		datalen := uint16(ethPkt.Len()) + ethPkt.Data.(*ipv4.IPv4).Length
		ds := make([]byte, datalen)
		ds, err = p.Data.PackBinary()
		if err != nil {
			return
		}
		binary.Write(buf, binary.BigEndian, ds)
	}
	data = buf.Bytes()
	return
}

func (p *ofp13.PacketOut) UnpackBinary(data []byte) (err error) {
	buf := bytes.NewBuffer(data)
	hs := make([]byte, p.Header.Len())
	binary.Read(buf, binary.BigEndian, hs)
	err = p.Header.UnpackBinary(hs)
	if err != nil {
		return
	}

	binary.Read(buf, binary.BigEndian, &p.BufferID)
	binary.Read(buf, binary.BigEndian, &p.InPort)
	binary.Read(buf, binary.BigEndian, &p.ActionsLen)
	binary.Read(buf, binary.BigEndian, &p.pad)

	n := p.Len()
	pos := p.Len() + int(p.ActionsLen)
	for n < pos {
		var a Action
		a, err = DecodeAction(data[n:])
		if err != nil {
			return
		}
		n += a.Len()
	}
	if pos < len(data) {
		ds := make([]byte, len(data)-pos)
		binary.Read(buf, binary.BigEndian, ds)
		err = p.Data.UnpackBinary(ds)
		if err != nil {
			return
		}
	}
	return
}

func NewPacketIn() (p *PacketIn) {
	p = new(PacketIn)
	p.Header.Type = uint8(OFPTPacketIn)
	return
}

func (p *PacketIn) Len() (l int) {
	l = 36
	return
}

func (p *PacketIn) PackBinary() (data []byte, err error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	hs := make([]byte, 0)
	hs, err = p.Header.PackBinary()
	if err != nil {
		return
	}
	binary.Write(buf, binary.BigEndian, hs)
	binary.Write(buf, binary.BigEndian, p.BufferID)
	binary.Write(buf, binary.BigEndian, p.TotalLen)
	binary.Write(buf, binary.BigEndian, p.Reason)
	binary.Write(buf, binary.BigEndian, p.TableID)
	binary.Write(buf, binary.BigEndian, p.Cookie)
	ms := make([]byte, 0)
	ms, err = p.Match.PackBinary()
	if err != nil {
		return
	}
	binary.Write(buf, binary.BigEndian, ms)
	binary.Write(buf, binary.BigEndian, p.pad)

	if p.Len() < len(data) {
		bs := make([]byte, 0)
		bs, err = p.Data.PackBinary()
		if err != nil {
			return
		}
		binary.Write(buf, binary.BigEndian, bs)
	}
	data = buf.Bytes()
	return
}

func (p *PacketIn) UnpackBinary(data []byte) (err error) {
	buf := bytes.NewBuffer(data)
	hs := make([]byte, p.Header.Len())
	binary.Read(buf, binary.BigEndian, hs)
	err = p.Header.UnpackBinary(hs)
	if err != nil {
		return
	}

	binary.Read(buf, binary.BigEndian, &p.BufferID)
	binary.Read(buf, binary.BigEndian, &p.TotalLen)
	binary.Read(buf, binary.BigEndian, &p.Reason)
	binary.Read(buf, binary.BigEndian, &p.TableID)
	binary.Read(buf, binary.BigEndian, &p.Cookie)
	binary.Read(buf, binary.BigEndian, &p.Match.Type)
	binary.Read(buf, binary.BigEndian, &p.Match.Length)
	ms := make([]byte, p.Match.Length)
	binary.Read(buf, binary.BigEndian, ms)
	err = p.Match.UnpackBinary(ms)
	if err != nil {
		return
	}

	binary.Read(buf, binary.BigEndian, &p.pad)

	if p.Len() < len(data) {
		bs := make([]byte, p.TotalLen)
		binary.Read(buf, binary.BigEndian, bs)
		err = p.Data.UnpackBinary(bs)
		if err != nil {
			return
		}
	}
	return
}

func (p *PacketIn) GetInport() (inport uint32) {
	match := p.Match
	port := match.OXMFields[4] << 24 | match.OXMFields[5] << 16 | match.OXMFields[6] << 8 | match.OXMFields[7]
	inport = uint32(port)
	return
}
