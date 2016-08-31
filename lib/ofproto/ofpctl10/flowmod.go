package ofpctl10

import (
	"bytes"
	"encoding/binary"

	"jd.com/jdcontroller/protocol/ofp10"
)

func NewFlowMod() (f *ofp10.FlowMod) {
	f = new(ofp10.FlowMod)
	f.Header.Type = uint8(ofp10.OFPTFlowMod)
	return
}

func (f *ofp10.FlowMod) AddAction(a Action) {
	f.Actions = append(f.Actions, a)
}

func (f *ofp10.FlowMod) Len() (l int) {
	l = 72
	return
}

func (f *ofp10.FlowMod) PackBinary() (data []byte, err error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	hs := make([]byte, 0)
	hs, err = f.Header.PackBinary()
	if err != nil {
		return
	}
	binary.Write(buf, binary.BigEndian, hs)
	ms := make([]byte, 0)
	ms, err = f.Match.PackBinary()
	if err != nil {
		return
	}
	binary.Write(buf, binary.BigEndian, ms)
	binary.Write(buf, binary.BigEndian, f.Cookie)
	binary.Write(buf, binary.BigEndian, f.Command)
	binary.Write(buf, binary.BigEndian, f.IdleTimeout)
	binary.Write(buf, binary.BigEndian, f.HardTimeout)
	binary.Write(buf, binary.BigEndian, f.Priority)
	binary.Write(buf, binary.BigEndian, f.BufferID)
	binary.Write(buf, binary.BigEndian, f.OutPort)
	binary.Write(buf, binary.BigEndian, f.Flags)
	for _, a := range f.Actions {
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

func (f *ofp10.FlowMod) UnpackBinary(data []byte) (err error) {
	buf := bytes.NewBuffer(data)
	hs := make([]byte, f.Header.Len())
	binary.Read(buf, binary.BigEndian, hs)
	err = f.Header.UnpackBinary(hs)
	if err != nil {
		return
	}
	ms := make([]byte, f.Match.Len())
	binary.Read(buf, binary.BigEndian, ms)
	err = f.Match.UnpackBinary(ms)
	if err != nil {
		return
	}
	binary.Read(buf, binary.BigEndian, &f.Cookie)
	binary.Read(buf, binary.BigEndian, &f.Command)
	binary.Read(buf, binary.BigEndian, &f.IdleTimeout)
	binary.Read(buf, binary.BigEndian, &f.HardTimeout)
	binary.Read(buf, binary.BigEndian, &f.Priority)
	binary.Read(buf, binary.BigEndian, &f.BufferID)
	binary.Read(buf, binary.BigEndian, &f.OutPort)
	binary.Read(buf, binary.BigEndian, &f.Flags)
	n := f.Len()
	for n < len(data) {
		var a Action
		a, err = DecodeAction(data[n:])
		if err != nil {
			return
		}
		f.Actions = append(f.Actions, a)
		n += a.Len()
	}
	return
}

func NewFlowRemoved() (f *FlowRemoved) {
	f = new(FlowRemoved)
	return
}

func (f *FlowRemoved) Len() (l int) {
	l = 88
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
	ms := make([]byte, 0)
	ms, err = f.Match.PackBinary()
	if err != nil {
		return
	}
	binary.Write(buf, binary.BigEndian, ms)
	binary.Write(buf, binary.BigEndian, f.Cookie)
	binary.Write(buf, binary.BigEndian, f.Priority)
	binary.Write(buf, binary.BigEndian, f.Reason)
	binary.Write(buf, binary.BigEndian, f.pad)
	binary.Write(buf, binary.BigEndian, f.DurationSec)
	binary.Write(buf, binary.BigEndian, f.DurationNSec)
	binary.Write(buf, binary.BigEndian, f.IdleTimeout)
	binary.Write(buf, binary.BigEndian, f.pad2)
	binary.Write(buf, binary.BigEndian, f.PacketCount)
	binary.Write(buf, binary.BigEndian, f.ByteCount)
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
	ms := make([]byte, f.Match.Len())
	binary.Read(buf, binary.BigEndian, ms)
	err = f.Header.UnpackBinary(ms)
	if err != nil {
		return
	}
	binary.Read(buf, binary.BigEndian, &f.Cookie)
	binary.Read(buf, binary.BigEndian, &f.Priority)
	binary.Read(buf, binary.BigEndian, &f.Reason)
	binary.Read(buf, binary.BigEndian, &f.pad)
	binary.Read(buf, binary.BigEndian, &f.DurationSec)
	binary.Read(buf, binary.BigEndian, &f.DurationNSec)
	binary.Read(buf, binary.BigEndian, &f.IdleTimeout)
	binary.Read(buf, binary.BigEndian, &f.pad2)
	binary.Read(buf, binary.BigEndian, &f.PacketCount)
	binary.Read(buf, binary.BigEndian, &f.ByteCount)
	return
}
