package ofp12

import (
	"bytes"
	"encoding/binary"
)

func NewTableMod() (m *TableMod) {
	m = new(TableMod)
	return
}

func (m *TableMod) Len() (l int) {
	l = 16
	return
}

func (m *TableMod) PackBinary() (data []byte, err error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	hs := make([]byte, 0)
	hs, err = m.Header.PackBinary()
	if err != nil {
		return
	}
	binary.Write(buf, binary.BigEndian, hs)
	binary.Write(buf, binary.BigEndian, m.TableID)
	binary.Write(buf, binary.BigEndian, m.Config)
	data = buf.Bytes()
	return
}

func (m *TableMod) UnpackBinary(data []byte) (err error) {
	buf := bytes.NewBuffer(data)
	hs := make([]byte, m.Header.Len())
	binary.Read(buf, binary.BigEndian, hs)
	err = m.Header.UnpackBinary(hs)
	if err != nil {
		return
	}
	binary.Read(buf, binary.BigEndian, &m.TableID)
	binary.Read(buf, binary.BigEndian, &m.Config)
	return
}

func NewFlowMod() (m *FlowMod) {
	m = new(FlowMod)
	return
}

func (m *FlowMod) Len() (l int) {
	l = 56
	return
}

func (m *FlowMod) PackBinary() (data []byte, err error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	hs := make([]byte, 0)
	hs, err = m.Header.PackBinary()
	if err != nil {
		return
	}
	binary.Write(buf, binary.BigEndian, hs)
	binary.Write(buf, binary.BigEndian, m.Cookie)
	binary.Write(buf, binary.BigEndian, m.CookieMask)
	binary.Write(buf, binary.BigEndian, m.TableID)
	binary.Write(buf, binary.BigEndian, m.Command)
	binary.Write(buf, binary.BigEndian, m.IdleTimeout)
	binary.Write(buf, binary.BigEndian, m.HardTimeout)
	binary.Write(buf, binary.BigEndian, m.Priority)
	binary.Write(buf, binary.BigEndian, m.BufferID)
	binary.Write(buf, binary.BigEndian, m.OutPort)
	binary.Write(buf, binary.BigEndian, m.OutGroup)
	binary.Write(buf, binary.BigEndian, m.Flags)
	binary.Write(buf, binary.BigEndian, m.pad)

	ms := make([]byte, 0)
	ms, err = m.Match.PackBinary()
	if err != nil {
		return
	}
	binary.Write(buf, binary.BigEndian, ms)

	for _, i := range m.Instructions {
		bs := make([]byte, 0)
		bs, err = EncodeInstruction(i)
		if err != nil {
			return
		}
		binary.Write(buf, binary.BigEndian, bs)
	}
	data = buf.Bytes()
	return
}

func (m *FlowMod) UnpackBinary(data []byte) (err error) {
	buf := bytes.NewBuffer(data)
	hs := make([]byte, m.Header.Len())
	binary.Read(buf, binary.BigEndian, hs)
	err = m.Header.UnpackBinary(hs)
	if err != nil {
		return
	}
	binary.Read(buf, binary.BigEndian, &m.Cookie)
	binary.Read(buf, binary.BigEndian, &m.CookieMask)
	binary.Read(buf, binary.BigEndian, &m.TableID)
	binary.Read(buf, binary.BigEndian, &m.Command)
	binary.Read(buf, binary.BigEndian, &m.IdleTimeout)
	binary.Read(buf, binary.BigEndian, &m.HardTimeout)
	binary.Read(buf, binary.BigEndian, &m.Priority)
	binary.Read(buf, binary.BigEndian, &m.BufferID)
	binary.Read(buf, binary.BigEndian, &m.OutPort)
	binary.Read(buf, binary.BigEndian, &m.OutGroup)
	binary.Read(buf, binary.BigEndian, &m.Flags)

	ms := make([]byte, m.Match.Len())
	binary.Read(buf, binary.BigEndian, ms)
	err = m.Match.UnpackBinary(ms)
	if err != nil {
		return
	}

	n := m.Len()
	for n < len(data) {
		var i Instruction
		i, err = DecodeInstruction(data[n:])
		if err != nil {
			return
		}
		m.Instructions = append(m.Instructions, i)
		n += i.Len()
	}

	return
}

func NewGroupMod() (m *GroupMod) {
	m = new(GroupMod)
	return
}

func (m *GroupMod) Len() (l int) {
	l = 16
	return
}

func (m *GroupMod) PackBinary() (data []byte, err error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	hs := make([]byte, 0)
	hs, err = m.Header.PackBinary()
	if err != nil {
		return
	}
	binary.Write(buf, binary.BigEndian, hs)
	binary.Write(buf, binary.BigEndian, m.Command)
	binary.Write(buf, binary.BigEndian, m.Type)
	binary.Write(buf, binary.BigEndian, m.pad)
	binary.Write(buf, binary.BigEndian, m.GroupID)
	for _, b := range m.Buckets {
		bs := make([]byte, 0)
		bs, err = b.PackBinary()
		if err != nil {
			return
		}
		binary.Write(buf, binary.BigEndian, bs)
	}
	data = buf.Bytes()
	return
}

func (m *GroupMod) UnpackBinary(data []byte) (err error) {
	buf := bytes.NewBuffer(data)
	hs := make([]byte, m.Header.Len())
	binary.Read(buf, binary.BigEndian, hs)
	err = m.Header.UnpackBinary(hs)
	if err != nil {
		return
	}
	binary.Read(buf, binary.BigEndian, &m.Command)
	binary.Read(buf, binary.BigEndian, &m.Type)
	binary.Read(buf, binary.BigEndian, &m.pad)
	binary.Read(buf, binary.BigEndian, &m.GroupID)

	n := int(m.Len())
	for n < len(data) {
		b := new(Bucket)
		bs := make([]byte, b.Len())
		binary.Read(buf, binary.BigEndian, bs)
		b.UnpackBinary(bs)
		m.Buckets = append(m.Buckets, *b)
		n += b.Len()
	}

	return
}

func NewBucket() (b *Bucket) {
	b = new(Bucket)
	return
}

func (b *Bucket) Len() (l int) {
	l = 16
	return
}

func (b *Bucket) PackBinary() (data []byte, err error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	binary.Write(buf, binary.BigEndian, b.Length)
	binary.Write(buf, binary.BigEndian, b.Weight)
	binary.Write(buf, binary.BigEndian, b.WatchPort)
	binary.Write(buf, binary.BigEndian, b.WatchGroup)
	binary.Write(buf, binary.BigEndian, b.pad)
	for _, a := range b.Actions {
		bs := make([]byte, 0)
		bs, err = EncodeAction(a)
		if err != nil {
			return
		}
		binary.Write(buf, binary.BigEndian, bs)
	}

	data = buf.Bytes()
	return
}

func (b *Bucket) UnpackBinary(data []byte) (err error) {
	buf := bytes.NewBuffer(data)
	binary.Read(buf, binary.BigEndian, &b.Length)
	binary.Read(buf, binary.BigEndian, &b.Weight)
	binary.Read(buf, binary.BigEndian, &b.WatchPort)
	binary.Read(buf, binary.BigEndian, &b.WatchGroup)
	binary.Read(buf, binary.BigEndian, &b.pad)

	n := b.Len()
	for n < len(data) {
		var a Action
		a, err = DecodeAction(data[n:])
		b.Actions = append(b.Actions, a)
		n += a.Len()
	}
	return
}

func NewPortMod() (m *PortMod) {
	m = new(PortMod)
	return
}

func (m *PortMod) Len() (l int) {
	l = 40
	return
}

func (m *PortMod) PackBinary() (data []byte, err error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	hs := make([]byte, 0)
	hs, err = m.Header.PackBinary()
	if err != nil {
		return
	}
	binary.Write(buf, binary.BigEndian, hs)
	binary.Write(buf, binary.BigEndian, m.PortNO)
	binary.Write(buf, binary.BigEndian, m.pad)
	binary.Write(buf, binary.BigEndian, m.HWAddr)
	binary.Write(buf, binary.BigEndian, m.pad2)
	binary.Write(buf, binary.BigEndian, m.Config)
	binary.Write(buf, binary.BigEndian, m.Mask)
	binary.Write(buf, binary.BigEndian, m.Advertise)
	binary.Write(buf, binary.BigEndian, m.pad3)
	data = buf.Bytes()
	return
}

func (m *PortMod) UnpackBinary(data []byte) (err error) {
	buf := bytes.NewBuffer(data)
	hs := make([]byte, m.Header.Len())
	binary.Read(buf, binary.BigEndian, hs)
	err = m.Header.UnpackBinary(hs)
	if err != nil {
		return
	}
	binary.Read(buf, binary.BigEndian, &m.PortNO)
	binary.Read(buf, binary.BigEndian, &m.pad)
	binary.Read(buf, binary.BigEndian, &m.HWAddr)
	binary.Read(buf, binary.BigEndian, &m.pad2)
	binary.Read(buf, binary.BigEndian, &m.Config)
	binary.Read(buf, binary.BigEndian, &m.Mask)
	binary.Read(buf, binary.BigEndian, &m.Advertise)
	binary.Read(buf, binary.BigEndian, &m.pad3)
	return
}

func NewFlowRemoved() (f *FlowRemoved) {
	f = new(FlowRemoved)
	return
}

func (f *FlowRemoved) Len() (l int) {
	l = 56
	return
}

func (f *FlowRemoved) PackBinary() (data []byte, err error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	hs := make([]byte, 0)
	hs, err = f.Header.PackBinary()
	if err != nil {
		return
	}
	binary.Write(buf, binary.BigEndian, hs)
	binary.Write(buf, binary.BigEndian, f.Cookie)
	binary.Write(buf, binary.BigEndian, f.Priority)
	binary.Write(buf, binary.BigEndian, f.Reason)
	binary.Write(buf, binary.BigEndian, f.TableID)
	binary.Write(buf, binary.BigEndian, f.DurationSec)
	binary.Write(buf, binary.BigEndian, f.DurationNSec)
	binary.Write(buf, binary.BigEndian, f.IdleTimeout)
	binary.Write(buf, binary.BigEndian, f.HardTimeout)
	binary.Write(buf, binary.BigEndian, f.PacketCount)
	binary.Write(buf, binary.BigEndian, f.ByteCount)
	ms := make([]byte, 0)
	ms, err = f.Match.PackBinary()
	if err != nil {
		return
	}
	binary.Write(buf, binary.BigEndian, ms)
	data = buf.Bytes()
	return
}

func (f *FlowRemoved) UnpackBinary(data []byte) (err error) {
	buf := bytes.NewBuffer(data)
	hs := make([]byte, f.Header.Len())
	binary.Read(buf, binary.BigEndian, hs)
	err = f.Header.UnpackBinary(hs)
	if err != nil {
		return
	}
	binary.Read(buf, binary.BigEndian, &f.Cookie)
	binary.Read(buf, binary.BigEndian, &f.Priority)
	binary.Read(buf, binary.BigEndian, &f.Reason)
	binary.Read(buf, binary.BigEndian, &f.TableID)
	binary.Read(buf, binary.BigEndian, &f.DurationSec)
	binary.Read(buf, binary.BigEndian, &f.DurationNSec)
	binary.Read(buf, binary.BigEndian, &f.IdleTimeout)
	binary.Read(buf, binary.BigEndian, &f.HardTimeout)
	binary.Read(buf, binary.BigEndian, &f.PacketCount)
	binary.Read(buf, binary.BigEndian, &f.ByteCount)

	ms := make([]byte, f.Match.Len())
	binary.Read(buf, binary.BigEndian, ms)
	err = f.Match.UnpackBinary(ms)
	if err != nil {
		return
	}
	return
}
