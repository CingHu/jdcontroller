package ofpctl10

import (
	"bytes"
	"encoding/binary"
	"errors"

	"jd.com/jdcontroller/protocol/ofp10"
)

func NewPacketOut() (p *ofp10.PacketOut) {
	p = new(ofp10.PacketOut)
	p.Header.Type = uint8(ofp10.OFPTPacketOut)
	return
}

func (p *ofp10.PacketOut) AddAction(a ofp10.Action) {
	p.Actions = append(p.Actions, a)
	p.ActionsLen += uint16(a.Len())
}

func (p *ofp10.PacketOut) Len() (l int) {
	l = 16
	return
}

func (p *ofp10.PacketOut) PackBinary() (data []byte, err error) {
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

	for _, a := range p.Actions {
		bs := make([]byte, 0)
		bs, err = EncodeAction(a)
		if err != nil {
			return
		}
		binary.Write(buf, binary.BigEndian, bs)
	}

	ds := make([]byte, p.Data.Len())
	ds, err = p.Data.PackBinary()
	if err != nil {
		return
	}
	binary.Write(buf, binary.BigEndian, ds)
	data = buf.Bytes()
	return
}

func (p *ofp10.PacketOut) UnpackBinary(data []byte) (err error) {
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

	n := p.Len()
	pos := p.Len() + int(p.ActionsLen)
	for n < pos {
		var a Action
		a, err = DecodeAction(data[n:])
		p.Actions = append(p.Actions, a)
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

func NewPacketIn() (p *ofp10.PacketIn) {
	p = new(ofp10.PacketIn)
	p.Header.Type = uint8(ofp10.OFPTPacketIn)
	return
}

func (p *ofp10.PacketIn) Len() (l int) {
	l = 20
	return
}

func (p *ofp10.PacketIn) PackBinary() (data []byte, err error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	hs := make([]byte, 0)
	hs, err = p.Header.PackBinary()
	if err != nil {
		return
	}
	binary.Write(buf, binary.BigEndian, hs)
	binary.Write(buf, binary.BigEndian, p.BufferID)
	binary.Write(buf, binary.BigEndian, p.TotalLen)
	binary.Write(buf, binary.BigEndian, p.InPort)
	binary.Write(buf, binary.BigEndian, p.Reason)
	binary.Write(buf, binary.BigEndian, p.pad)

	ds := make([]byte, 0)
	ds, err = p.Data.PackBinary()
	if err != nil {
		return
	}
	binary.Write(buf, binary.BigEndian, ds)
	data = buf.Bytes()
	return
}

func (p *ofp10.PacketIn) UnpackBinary(data []byte) (err error) {
	buf := bytes.NewBuffer(data)
	hs := make([]byte, p.Header.Len())
	binary.Read(buf, binary.BigEndian, hs)
	err = p.Header.UnpackBinary(hs)
	if err != nil {
		return
	}

	binary.Read(buf, binary.BigEndian, &p.BufferID)
	binary.Read(buf, binary.BigEndian, &p.TotalLen)
	binary.Read(buf, binary.BigEndian, &p.InPort)
	binary.Read(buf, binary.BigEndian, &p.Reason)
	binary.Read(buf, binary.BigEndian, &p.pad)
	if p.Len() < len(data) {
		ds := make([]byte, p.Data.Len())
		binary.Read(buf, binary.BigEndian, ds)
		err = p.Data.UnpackBinary(ds)
		if err != nil {
			return
		}
	}
	return
}

func NewVendorHeader() (v *ofp10.VendorHeader) {
	v = new(ofp10.VendorHeader)
	v.Header.Type = uint8(OFPTVendor)
	return
}

func (v *ofp10.VendorHeader) Len() (l int) {
	l = 12
	return
}

func (v *ofp10.VendorHeader) PackBinary() (data []byte, err error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	hs := make([]byte, 0)
	hs, err = v.Header.PackBinary()
	if err != nil {
		return
	}
	binary.Write(buf, binary.BigEndian, hs)
	binary.Write(buf, binary.BigEndian, v.Vendor)
	data = buf.Bytes()
	return
}

func (v *ofp10.VendorHeader) UnpackBinary(data []byte) (err error) {
	if len(data) < int(v.Len()) {
		return errors.New("The []byte the wrong size to Unpack an " +
			"VendorHeader message.")
	}
	buf := bytes.NewBuffer(data)
	hs := make([]byte, v.Header.Len())
	binary.Read(buf, binary.BigEndian, hs)
	err = v.Header.UnpackBinary(hs)
	if err != nil {
		return
	}

	binary.Read(buf, binary.BigEndian, &v.Vendor)
	return
}
