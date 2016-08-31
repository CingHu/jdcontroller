package ofp12

import (
	"bytes"
	"encoding/binary"
)

func NewPacketOut() (p *PacketOut) {
	p = new(PacketOut)
	p.Header.Type = uint8(OFPTPacketOut)
	return
}

func (p *PacketOut) AddAction(a Action) {
	p.Actions = append(p.Actions, a)
	p.ActionsLen += uint16(a.Len())
}

func (p *PacketOut) Len() (l int) {
	l = 24
	return
}

func (p *PacketOut) PackBinary() (data []byte, err error) {
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
	if p.Len()+int(p.ActionsLen) < len(data) {
		ds := make([]byte, p.Data.Len())
		ds, err = p.Data.PackBinary()
		if err != nil {
			return
		}
		binary.Write(buf, binary.BigEndian, ds)
	}
	data = buf.Bytes()
	return
}

func (p *PacketOut) UnpackBinary(data []byte) (err error) {
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
	l = 26
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
	ms := make([]byte, p.Match.Len())
	binary.Read(buf, binary.BigEndian, ms)
	err = p.Match.UnpackBinary(ms)
	if err != nil {
		return
	}

	binary.Read(buf, binary.BigEndian, &p.pad)

	if p.Len() < len(data) {
		bs := make([]byte, p.Data.Len())
		binary.Read(buf, binary.BigEndian, bs)
		err = p.Data.UnpackBinary(bs)
		if err != nil {
			return
		}
	}
	return
}
