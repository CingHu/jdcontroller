package ofp12

import (
	"bytes"
	"encoding/binary"
)

func NewPort() (p *Port) {
	p = new(Port)
	return
}

func (p *Port) Len() (l int) {
	l = 64
	return
}

func (p *Port) PackBinary() (data []byte, err error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	binary.Write(buf, binary.BigEndian, p.PortNO)
	binary.Write(buf, binary.BigEndian, p.pad)
	binary.Write(buf, binary.BigEndian, p.HWAddr)
	binary.Write(buf, binary.BigEndian, p.pad2)
	binary.Write(buf, binary.BigEndian, p.Name)
	binary.Write(buf, binary.BigEndian, p.Config)
	binary.Write(buf, binary.BigEndian, p.State)
	binary.Write(buf, binary.BigEndian, p.Curr)
	binary.Write(buf, binary.BigEndian, p.Advertised)
	binary.Write(buf, binary.BigEndian, p.Supported)
	binary.Write(buf, binary.BigEndian, p.Peer)
	binary.Write(buf, binary.BigEndian, p.CurrSpeed)
	binary.Write(buf, binary.BigEndian, p.MaxSpeed)
	data = buf.Bytes()
	return
}

func (p *Port) UnpackBinary(data []byte) (err error) {
	buf := bytes.NewBuffer(data)
	binary.Read(buf, binary.BigEndian, &p.PortNO)
	binary.Read(buf, binary.BigEndian, &p.pad)
	binary.Read(buf, binary.BigEndian, &p.HWAddr)
	binary.Read(buf, binary.BigEndian, &p.pad2)
	binary.Read(buf, binary.BigEndian, &p.Name)
	binary.Read(buf, binary.BigEndian, &p.Config)
	binary.Read(buf, binary.BigEndian, &p.State)
	binary.Read(buf, binary.BigEndian, &p.Curr)
	binary.Read(buf, binary.BigEndian, &p.Advertised)
	binary.Read(buf, binary.BigEndian, &p.Supported)
	binary.Read(buf, binary.BigEndian, &p.Peer)
	binary.Read(buf, binary.BigEndian, &p.CurrSpeed)
	binary.Read(buf, binary.BigEndian, &p.MaxSpeed)
	return
}

func NewPortStatus() (p *PortStatus) {
	p = new(PortStatus)
	return
}

func (p *PortStatus) Len() (l int) {
	l = 80
	return
}

func (p *PortStatus) PackBinary() (data []byte, err error) {
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

func (p *PortStatus) UnpackBinary(data []byte) (err error) {
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
