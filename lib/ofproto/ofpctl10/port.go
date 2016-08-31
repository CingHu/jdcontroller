package ofpctl10

import (
	"bytes"
	"encoding/binary"

	"jd.com/jdcontroller/protocol/ofp10"
)

func NewPhyPort() (p *ofp10.Port) {
	p = new(ofp10.Port)
	return
}

func (p *ofp10.Port) Len() (l int) {
	l = 48
	return
}

func (p *ofp10.Port) PackBinary() (data []byte, err error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	binary.Write(buf, binary.BigEndian, p.PortNO)
	binary.Write(buf, binary.BigEndian, p.HWAddr)
	binary.Write(buf, binary.BigEndian, p.Name)
	binary.Write(buf, binary.BigEndian, p.Config)
	binary.Write(buf, binary.BigEndian, p.State)
	binary.Write(buf, binary.BigEndian, p.Curr)
	binary.Write(buf, binary.BigEndian, p.Advertised)
	binary.Write(buf, binary.BigEndian, p.Supported)
	binary.Write(buf, binary.BigEndian, p.Peer)

	data = buf.Bytes()
	return
}

func (p *Port) UnpackBinary(data []byte) (err error) {
	buf := bytes.NewBuffer(data)
	binary.Read(buf, binary.BigEndian, &p.PortNO)
	binary.Read(buf, binary.BigEndian, &p.HWAddr)
	binary.Read(buf, binary.BigEndian, &p.Name)
	binary.Read(buf, binary.BigEndian, &p.Config)
	binary.Read(buf, binary.BigEndian, &p.State)
	binary.Read(buf, binary.BigEndian, &p.Curr)
	binary.Read(buf, binary.BigEndian, &p.Advertised)
	binary.Read(buf, binary.BigEndian, &p.Supported)
	binary.Read(buf, binary.BigEndian, &p.Peer)
	return
}

func NewPortMod() (p *ofp10.PortMod) {
	p = new(ofp10.PortMod)
	p.Header.Type = uint8(ofp10.OFPTPortMod)
	return
}

func (p *ofp10.PortMod) Len() (l int) {
	l = 32
	return
}

func (p *ofp10.PortMod) PackBinary() (data []byte, err error) {
	buf := bytes.NewBuffer(data)
	hs := make([]byte, 0)
	hs, err = p.Header.PackBinary()
	if err != nil {
		return
	}
	binary.Write(buf, binary.BigEndian, hs)
	binary.Write(buf, binary.BigEndian, p.PortNO)
	binary.Write(buf, binary.BigEndian, p.HWAddr)
	binary.Write(buf, binary.BigEndian, p.Config)
	binary.Write(buf, binary.BigEndian, p.Mask)
	binary.Write(buf, binary.BigEndian, p.Advertise)
	binary.Write(buf, binary.BigEndian, p.pad)
	data = buf.Bytes()
	return
}

func (p *ofp10.PortMod) UnpackBinary(data []byte) (err error) {
	buf := bytes.NewBuffer(data)
	hs := make([]byte, p.Header.Len())
	binary.Read(buf, binary.BigEndian, hs)
	err = p.Header.UnpackBinary(hs)
	if err != nil {
		return
	}

	binary.Read(buf, binary.BigEndian, &p.PortNO)
	binary.Read(buf, binary.BigEndian, &p.HWAddr)
	binary.Read(buf, binary.BigEndian, &p.Config)
	binary.Read(buf, binary.BigEndian, &p.Mask)
	binary.Read(buf, binary.BigEndian, &p.Advertise)
	binary.Read(buf, binary.BigEndian, &p.pad)
	return
}

func NewPortStatus() (p *ofp10.PortStatus) {
	p = new(ofp10.PortStatus)
	return
}

func (p *ofp10.PortStatus) Len() (l int) {
	l = 64
	return
}

func (p *ofp10.PortStatus) PackBinary() (data []byte, err error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	hs := make([]byte, 0)
	hs, err = p.Header.PackBinary()
	if err != nil {
		return
	}
	binary.Write(buf, binary.BigEndian, hs)
	binary.Write(buf, binary.BigEndian, p.Reason)
	binary.Write(buf, binary.BigEndian, p.pad)

	ds := make([]byte, 0)
	ds, err = p.Desc.PackBinary()
	if err != nil {
		return
	}
	binary.Write(buf, binary.BigEndian, ds)

	data = buf.Bytes()
	return
}

func (p *ofp10.PortStatus) UnpackBinary(data []byte) (err error) {
	buf := bytes.NewBuffer(data)
	hs := make([]byte, p.Header.Len())
	binary.Read(buf, binary.BigEndian, hs)
	err = p.Header.UnpackBinary(hs)
	if err != nil {
		return
	}
	binary.Read(buf, binary.BigEndian, &p.Reason)
	binary.Read(buf, binary.BigEndian, &p.pad)
	ds := make([]byte, p.Desc.Len())
	binary.Read(buf, binary.BigEndian, ds)
	err = p.Desc.UnpackBinary(ds)
	if err != nil {
		return
	}
	return
}
