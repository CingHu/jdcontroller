package ofpctl13

import (
	"bytes"
	_"fmt"
	"encoding/binary"

	"jd.com/jdcontroller/config"
	"jd.com/jdcontroller/protocol/ofp13"
)

func NewTableMod() (m *ofp13.TableMod) {
	m = new(ofp13.TableMod)
	return
}

func (m *ofp13.TableMod) Len() (l int) {
	l = 16
	return
}

func (m *ofp13.TableMod) PackBinary() (data []byte, err error) {
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

func (m *ofp13.TableMod) UnpackBinary(data []byte) (err error) {
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

func NewFlowMod(priority uint16) (m *FlowMod) {
	m = new(FlowMod)
	m.Header.Version = Version
	m.Header.Type = OFPTFlowMod
	m.Header.Length = 52 //before match and match's pad[4]
	//m.IdleTimeout = idletime
	if priority == 0 || priority == 1 {
		m.IdleTimeout = 0
	} else {
		m.IdleTimeout = uint16(config.GetConfig().FlowmodIdleTimeout) //设置单位秒
	}

	m.Priority = priority
	//m.Flags = 0x1 //ovs report flowremoved
	m.Match.Type = OFPMTOXM
	m.Match.Length = 4
	return
}

func (m *FlowMod) Len() (l int) {
	l = 56
	return
}

func (m *FlowMod) AddMatch(match *MatchField) (err error) {
	OxmField := match.XMFields
	fieldlen := match.GetLen()
	length := fieldlen + OFPXMHEADERLEN

	m.Match.Length += length
	m.Match.OXMFields = append(m.Match.OXMFields, OxmField...)
	m.Header.Length += length
	return
}

func (m *FlowMod) AddInsAction(inacts *InstructionActions) (err error) {
	m.Instructions = append(m.Instructions, inacts)
	m.Header.Length += inacts.Header.Length
	return
}

func (m *FlowMod) PackBinary() (data []byte, err error) {
	matchpadLen := uint16((m.Match.Length + 7) / 8 * 8) - m.Match.Length
	m.Header.Length += matchpadLen

	padLen := uint16((m.Header.Length + 7) / 8 * 8) - m.Header.Length
	m.Header.Length += padLen

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
		switch i.(type) {
		case *InstructionGotoTable:
			bs, err = i.(*InstructionGotoTable).PackBinary()
		case *InstructionWriteMetadata:
			bs, err = i.(*InstructionWriteMetadata).PackBinary()
		case *InstructionActions:
			bs, err = i.(*InstructionActions).PackBinary()
		}
		if err != nil {
			return
		}
		binary.Write(buf, binary.BigEndian, bs)
	}

	for i := 0; i < int(padLen); i++ {
		binary.Write(buf, binary.BigEndian, uint8(0))
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
		if (len(data) - n) < 8 {
			break
		}
		var i Instruction
		var l int
		switch binary.BigEndian.Uint16(data[n:]) {
		case OFPITGotoTable:
			i = new(InstructionGotoTable)
			r := i.(*InstructionGotoTable)
			is := make([]byte, r.Len())
			binary.Read(buf, binary.BigEndian, is)
			r.UnpackBinary(is)
			l = r.Len()
		case OFPITWriteActions:
			i = new(InstructionActions)
			r := i.(*InstructionActions)
			is := make([]byte, r.Len())
			binary.Read(buf, binary.BigEndian, is)
			r.UnpackBinary(is)
			l = r.Len()
		case OFPITApplyActions:
			i = new(InstructionActions)
			r := i.(*InstructionActions)
			is := make([]byte, r.Len())
			binary.Read(buf, binary.BigEndian, is)
			r.UnpackBinary(is)
			l = r.Len()
		case OFPITClearActions:
			i = new(InstructionActions)
			r := i.(*InstructionActions)
			is := make([]byte, r.Len())
			binary.Read(buf, binary.BigEndian, is)
			r.UnpackBinary(is)
			l = r.Len()
		case OFPITWriteMetadata:
			i = new(InstructionWriteMetadata)
			r := i.(*InstructionWriteMetadata)
			is := make([]byte, r.Len())
			binary.Read(buf, binary.BigEndian, is)
			r.UnpackBinary(is)
			l = r.Len()
		}
		m.Instructions = append(m.Instructions, i)
		n += l
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
		if err != nil {
			return
		}
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

func NewMeterMod() (m *MeterMod) {
	m = new(MeterMod)
	return
}

func (m *MeterMod) Len() (l int) {
	l = 16
	return
}

func (m *MeterMod) PackBinary() (data []byte, err error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	hs := make([]byte, 0)
	hs, err = m.Header.PackBinary()
	if err != nil {
		return
	}
	binary.Write(buf, binary.BigEndian, hs)
	binary.Write(buf, binary.BigEndian, m.Command)
	binary.Write(buf, binary.BigEndian, m.Flags)
	binary.Write(buf, binary.BigEndian, m.MeterID)
	for _, b := range m.Bands {
		bs := make([]byte, 0)
		switch b.(type) {
		case *MeterBandDrop:
			bs, err = b.(*MeterBandDrop).PackBinary()
		case *MeterBandDscpRemark:
			bs, err = b.(*MeterBandDscpRemark).PackBinary()
		case *MeterBandExperimenter:
			bs, err = b.(*MeterBandExperimenter).PackBinary()
		}
		if err != nil {
			return
		}
		binary.Write(buf, binary.BigEndian, bs)
	}

	data = buf.Bytes()
	return
}

func (m *MeterMod) UnpackBinary(data []byte) (err error) {
	buf := bytes.NewBuffer(data)
	hs := make([]byte, m.Header.Len())
	binary.Read(buf, binary.BigEndian, hs)
	err = m.Header.UnpackBinary(hs)
	if err != nil {
		return
	}
	binary.Read(buf, binary.BigEndian, &m.Command)
	binary.Read(buf, binary.BigEndian, &m.Flags)
	binary.Read(buf, binary.BigEndian, &m.MeterID)

	n := m.Len()
	for n < len(data) {
		var b MeterBand
		var l int
		switch binary.BigEndian.Uint16(data) {
		case OFPMBTDrop:
			b = new(MeterBandDrop)
			r := b.(*MeterBandDrop)
			bs := make([]byte, r.Len())
			binary.Read(buf, binary.BigEndian, bs)
			r.UnpackBinary(bs)
			l = r.Len()
		case OFPMBTDscpRemark:
			b = new(MeterBandDscpRemark)
			r := b.(*MeterBandDscpRemark)
			bs := make([]byte, r.Len())
			binary.Read(buf, binary.BigEndian, bs)
			r.UnpackBinary(bs)
			l = r.Len()
		case OFPMBTExperimenter:
			b = new(MeterBandExperimenter)
			r := b.(*MeterBandExperimenter)
			bs := make([]byte, r.Len())
			binary.Read(buf, binary.BigEndian, bs)
			r.UnpackBinary(bs)
			l = r.Len()
		}
		m.Bands = append(m.Bands, b)
		n += l
	}
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
	binary.Read(buf, binary.BigEndian, &f.Match.Type)
	binary.Read(buf, binary.BigEndian, &f.Match.Length)

	ms := make([]byte, f.Match.Length - 2)
	binary.Read(buf, binary.BigEndian, ms)
	buf = bytes.NewBuffer(ms)

	for n := 0; n < len(ms) - 2; n++ {
		b := new(uint8)
		binary.Read(buf, binary.BigEndian, b)
		f.Match.OXMFields = append(f.Match.OXMFields, *b)
	}

	binary.Read(buf, binary.BigEndian, &f.Match.pad)

	return
}
