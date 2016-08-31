package ofpctl10

import (
	"bytes"
	"encoding/binary"

	"jd.com/jdcontroller/lib/buffer"
	"jd.com/jdcontroller/protocol/ofp10"
)

func NewStatsRequest() (s *ofp10.StatsRequest) {
	s = new(ofp10.StatsRequest)
	s.Type = uint16(ofp10.OFPTStatsRequest)
	return
}

func (s *ofp10.StatsRequest) Len() (l int) {
	l = 12
	return
}

func (s *ofp10.StatsRequest) PackBinary() (data []byte, err error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	hs := make([]byte, 0)
	hs, err = s.Header.PackBinary()
	if err != nil {
		return
	}
	binary.Write(buf, binary.BigEndian, hs)
	binary.Write(buf, binary.BigEndian, s.Type)
	binary.Write(buf, binary.BigEndian, s.Flags)
	if s.Body != nil {
		bs := make([]byte, 0)
		bs, err = EncodeStats(s.Body)
		if err != nil {
			return
		}
		binary.Write(buf, binary.BigEndian, bs)
	}
	data = buf.Bytes()
	return
}

func (s *ofp10.StatsRequest) UnpackBinary(data []byte) (err error) {
	buf := bytes.NewBuffer(data)
	hs := make([]byte, s.Header.Len())
	binary.Read(buf, binary.BigEndian, hs)
	err = s.Header.UnpackBinary(hs)
	if err != nil {
		return
	}
	binary.Read(buf, binary.BigEndian, &s.Type)
	binary.Read(buf, binary.BigEndian, &s.Flags)
	n := s.Len()
	if n < len(data) {
		s.Body, err = DecodeStats(data[n:], s.Type)
		if err != nil {
			return
		}
	}
	return
}

func NewStatsReply() (s *ofp10.StatsReply) {
	s = new(ofp10.StatsReply)
	return
}

func (s *ofp10.StatsReply) Len() (l int) {
	l = 12
	return
}

func (s *ofp10.StatsReply) PackBinary() (data []byte, err error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	hs := make([]byte, 0)
	hs, err = s.Header.PackBinary()
	if err != nil {
		return
	}
	binary.Write(buf, binary.BigEndian, hs)
	binary.Write(buf, binary.BigEndian, s.Type)
	binary.Write(buf, binary.BigEndian, s.Flags)
	if s.Body != nil {
		bs := make([]byte, 0)
		bs, err = EncodeStats(s.Body)
		if err != nil {
			return
		}
		binary.Write(buf, binary.BigEndian, bs)
	}
	data = buf.Bytes()
	return
}

func (s *ofp10.StatsReply) UnpackBinary(data []byte) (err error) {
	buf := bytes.NewBuffer(data)
	hs := make([]byte, s.Header.Len())
	binary.Read(buf, binary.BigEndian, hs)
	err = s.Header.UnpackBinary(hs)
	if err != nil {
		return
	}
	binary.Read(buf, binary.BigEndian, &s.Type)
	binary.Read(buf, binary.BigEndian, &s.Flags)
	n := s.Len()
	if n < len(data) {
		s.Body, err = DecodeStats(data[n:], s.Type)
		if err != nil {
			return
		}
	}
	return
}

func NewDescStats() (s *ofp10.DescStats) {
	s = new(ofp10.DescStats)
	return
}

func (s *ofp10.DescStats) Len() (l int) {
	l = OFPDescStrLen*4 + OFPSerialNumLen
	return
}

func (s *ofp10.DescStats) PackBinary() (data []byte, err error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	binary.Write(buf, binary.BigEndian, s.MfrDesc)
	binary.Write(buf, binary.BigEndian, s.HWDesc)
	binary.Write(buf, binary.BigEndian, s.SWDesc)
	binary.Write(buf, binary.BigEndian, s.SerialNum)
	binary.Write(buf, binary.BigEndian, s.DPDesc)
	data = buf.Bytes()
	return
}

func (s *ofp10.DescStats) UnpackBinary(data []byte) (err error) {
	buf := bytes.NewBuffer(data)
	binary.Read(buf, binary.BigEndian, &s.MfrDesc)
	binary.Read(buf, binary.BigEndian, &s.HWDesc)
	binary.Read(buf, binary.BigEndian, &s.SWDesc)
	binary.Read(buf, binary.BigEndian, &s.SerialNum)
	binary.Read(buf, binary.BigEndian, &s.DPDesc)
	return
}

func NewFlowStatsRequest() (s *ofp10.FlowStatsRequest) {
	s = new(ofp10.FlowStatsRequest)
	return
}

func (s *ofp10.FlowStatsRequest) Len() (l int) {
	l = 44
	return
}

func (s *ofp10.FlowStatsRequest) PackBinary() (data []byte, err error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	ms := make([]byte, 0)
	ms, err = s.Match.PackBinary()
	if err != nil {
		return
	}
	binary.Write(buf, binary.BigEndian, ms)
	binary.Write(buf, binary.BigEndian, s.TableID)
	binary.Write(buf, binary.BigEndian, s.pad)
	binary.Write(buf, binary.BigEndian, s.OutPort)
	data = buf.Bytes()
	return
}

func (s *ofp10.FlowStatsRequest) UnpackBinary(data []byte) (err error) {
	buf := bytes.NewBuffer(data)
	ms := make([]byte, s.Match.Len())
	binary.Read(buf, binary.BigEndian, ms)
	err = s.Match.UnpackBinary(ms)
	if err != nil {
		return
	}
	binary.Read(buf, binary.BigEndian, &s.TableID)
	binary.Read(buf, binary.BigEndian, &s.pad)
	binary.Read(buf, binary.BigEndian, &s.OutPort)
	return
}

func NewFlowStats() (s *ofp10.FlowStats) {
	s = new(ofp10.FlowStats)
	return
}

func (s *ofp10.FlowStats) Len() (l int) {
	l = 88
	return
}

func (s *ofp10.FlowStats) PackBinary() (data []byte, err error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	binary.Write(buf, binary.BigEndian, s.Length)
	binary.Write(buf, binary.BigEndian, s.TableID)
	binary.Write(buf, binary.BigEndian, s.pad)
	ms := make([]byte, 0)
	ms, err = s.Match.PackBinary()
	if err != nil {
		return
	}
	binary.Write(buf, binary.BigEndian, ms)
	binary.Write(buf, binary.BigEndian, s.DurationSec)
	binary.Write(buf, binary.BigEndian, s.DurationNSec)
	binary.Write(buf, binary.BigEndian, s.Priority)
	binary.Write(buf, binary.BigEndian, s.IdleTimeout)
	binary.Write(buf, binary.BigEndian, s.HardTimeout)
	binary.Write(buf, binary.BigEndian, s.pad2)
	binary.Write(buf, binary.BigEndian, s.Cookie)
	binary.Write(buf, binary.BigEndian, s.PacketCount)
	binary.Write(buf, binary.BigEndian, s.ByteCount)

	for _, a := range s.Actions {
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

func (s *ofp10.FlowStats) UnpackBinary(data []byte) (err error) {
	buf := bytes.NewBuffer(data)
	binary.Read(buf, binary.BigEndian, &s.Length)
	binary.Read(buf, binary.BigEndian, &s.TableID)
	binary.Read(buf, binary.BigEndian, &s.pad)
	ms := make([]byte, s.Match.Len())
	binary.Read(buf, binary.BigEndian, ms)
	err = s.Match.UnpackBinary(ms)
	if err != nil {
		return
	}
	binary.Read(buf, binary.BigEndian, &s.DurationSec)
	binary.Read(buf, binary.BigEndian, &s.DurationNSec)
	binary.Read(buf, binary.BigEndian, &s.Priority)
	binary.Read(buf, binary.BigEndian, &s.IdleTimeout)
	binary.Read(buf, binary.BigEndian, &s.HardTimeout)
	binary.Read(buf, binary.BigEndian, &s.pad2)
	binary.Read(buf, binary.BigEndian, &s.Cookie)
	binary.Read(buf, binary.BigEndian, &s.PacketCount)
	binary.Read(buf, binary.BigEndian, &s.ByteCount)

	n := s.Len()
	for n < len(data) {
		var a ofp10.Action
		a, err = DecodeAction(data[n:])
		if err != nil {
			return
		}
		s.Actions = append(s.Actions, a)
		n += a.Len()
	}
	return
}

func NewAggregateStatsRequest() (s *ofp10.AggregateStatsRequest) {
	s = new(ofp10.AggregateStatsRequest)
	return
}

func (s *ofp10.AggregateStatsRequest) Len() (l int) {
	l = 44
	return
}

func (s *ofp10.AggregateStatsRequest) PackBinary() (data []byte, err error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	ms := make([]byte, 0)
	ms, err = s.Match.PackBinary()
	if err != nil {
		return
	}
	binary.Write(buf, binary.BigEndian, ms)
	binary.Write(buf, binary.BigEndian, s.TableID)
	binary.Write(buf, binary.BigEndian, s.pad)
	binary.Write(buf, binary.BigEndian, s.OutPort)

	data = buf.Bytes()
	return
}

func (s *ofp10.AggregateStatsRequest) UnpackBinary(data []byte) (err error) {
	buf := bytes.NewBuffer(data)
	ms := make([]byte, s.Match.Len())
	binary.Read(buf, binary.BigEndian, ms)
	err = s.Match.UnpackBinary(ms)
	if err != nil {
		return
	}
	binary.Read(buf, binary.BigEndian, &s.TableID)
	binary.Read(buf, binary.BigEndian, &s.pad)
	binary.Read(buf, binary.BigEndian, &s.OutPort)
	return
}

func NewAggregateStatsReply() (s *ofp10.AggregateStatsReply) {
	s = new(ofp10.AggregateStatsReply)
	return
}

func (s *ofp10.AggregateStatsReply) Len() (l int) {
	l = 24
	return
}

func (s *ofp10.AggregateStatsReply) PackBinary() (data []byte, err error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	binary.Write(buf, binary.BigEndian, s.PacketCount)
	binary.Write(buf, binary.BigEndian, s.ByteCount)
	binary.Write(buf, binary.BigEndian, s.FlowCount)
	binary.Write(buf, binary.BigEndian, s.pad)
	data = buf.Bytes()
	return
}

func (s *ofp10.AggregateStatsReply) UnpackBinary(data []byte) (err error) {
	buf := bytes.NewBuffer(data)
	binary.Read(buf, binary.BigEndian, &s.PacketCount)
	binary.Read(buf, binary.BigEndian, &s.ByteCount)
	binary.Read(buf, binary.BigEndian, &s.FlowCount)
	binary.Read(buf, binary.BigEndian, &s.pad)
	return
}

func NewTableStats() (s *ofp10.TableStats) {
	s = new(ofp10.TableStats)
	return
}

func (s *ofp10.TableStats) Len() (l int) {
	l = 64
	return
}

func (s *ofp10.TableStats) PackBinary() (data []byte, err error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	binary.Write(buf, binary.BigEndian, s.TableID)
	binary.Write(buf, binary.BigEndian, s.pad)
	binary.Write(buf, binary.BigEndian, s.Name)
	binary.Write(buf, binary.BigEndian, s.Wildcards)
	binary.Write(buf, binary.BigEndian, s.MaxEntries)
	binary.Write(buf, binary.BigEndian, s.ActiveCount)
	binary.Write(buf, binary.BigEndian, s.LookupCount)
	binary.Write(buf, binary.BigEndian, s.MatchedCount)
	data = buf.Bytes()
	return
}

func (s *ofp10.TableStats) UnpackBinary(data []byte) (err error) {
	buf := bytes.NewBuffer(data)
	binary.Read(buf, binary.BigEndian, &s.TableID)
	binary.Read(buf, binary.BigEndian, &s.pad)
	binary.Read(buf, binary.BigEndian, &s.Name)
	binary.Read(buf, binary.BigEndian, &s.Wildcards)
	binary.Read(buf, binary.BigEndian, &s.MaxEntries)
	binary.Read(buf, binary.BigEndian, &s.ActiveCount)
	binary.Read(buf, binary.BigEndian, &s.LookupCount)
	binary.Read(buf, binary.BigEndian, &s.MatchedCount)
	return
}

func NewPortStatsRequest() (s *ofp10.PortStatsRequest) {
	s = new(ofp10.PortStatsRequest)
	return
}

func (s *ofp10.PortStatsRequest) Len() (l int) {
	l = 8
	return
}

func (s *ofp10.PortStatsRequest) PackBinary() (data []byte, err error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	binary.Write(buf, binary.BigEndian, s.PortNO)
	binary.Write(buf, binary.BigEndian, s.pad)
	data = buf.Bytes()
	return
}

func (s *ofp10.PortStatsRequest) UnpackBinary(data []byte) (err error) {
	buf := bytes.NewBuffer(data)
	binary.Read(buf, binary.BigEndian, &s.PortNO)
	binary.Read(buf, binary.BigEndian, &s.pad)
	return
}

func NewPortStats() (s *PortStats) {
	s = new(PortStats)
	return
}

func (s *PortStats) Len() (l int) {
	l = 104
	return
}

func (s *PortStats) PackBinary() (data []byte, err error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	binary.Write(buf, binary.BigEndian, s.PortNO)
	binary.Write(buf, binary.BigEndian, s.pad)
	binary.Write(buf, binary.BigEndian, s.RxPackets)
	binary.Write(buf, binary.BigEndian, s.TxPackets)
	binary.Write(buf, binary.BigEndian, s.RxBytes)
	binary.Write(buf, binary.BigEndian, s.TxBytes)
	binary.Write(buf, binary.BigEndian, s.RxDropped)
	binary.Write(buf, binary.BigEndian, s.TxDropped)
	binary.Write(buf, binary.BigEndian, s.RxErrors)
	binary.Write(buf, binary.BigEndian, s.TxErrors)
	binary.Write(buf, binary.BigEndian, s.RxFrameErr)
	binary.Write(buf, binary.BigEndian, s.RxOverErr)
	binary.Write(buf, binary.BigEndian, s.RxCRCErr)
	binary.Write(buf, binary.BigEndian, s.Collisions)
	data = buf.Bytes()
	return
}

func (s *PortStats) UnpackBinary(data []byte) (err error) {
	buf := bytes.NewBuffer(data)
	binary.Read(buf, binary.BigEndian, &s.PortNO)
	binary.Read(buf, binary.BigEndian, &s.pad)
	binary.Read(buf, binary.BigEndian, &s.RxPackets)
	binary.Read(buf, binary.BigEndian, &s.TxPackets)
	binary.Read(buf, binary.BigEndian, &s.RxBytes)
	binary.Read(buf, binary.BigEndian, &s.TxBytes)
	binary.Read(buf, binary.BigEndian, &s.RxDropped)
	binary.Read(buf, binary.BigEndian, &s.TxDropped)
	binary.Read(buf, binary.BigEndian, &s.RxErrors)
	binary.Read(buf, binary.BigEndian, &s.TxErrors)
	binary.Read(buf, binary.BigEndian, &s.RxFrameErr)
	binary.Read(buf, binary.BigEndian, &s.RxOverErr)
	binary.Read(buf, binary.BigEndian, &s.RxCRCErr)
	binary.Read(buf, binary.BigEndian, &s.Collisions)
	return
}

func NewQueueStatsRequest() (s *ofp10.QueueStatsRequest) {
	s = new(ofp10.QueueStatsRequest)
	return
}

func (s *ofp10.QueueStatsRequest) Len() (l int) {
	l = 8
	return
}

func (s *ofp10.QueueStatsRequest) PackBinary() (data []byte, err error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	binary.Write(buf, binary.BigEndian, s.PortNO)
	binary.Write(buf, binary.BigEndian, s.pad)
	binary.Write(buf, binary.BigEndian, s.QueueID)
	data = buf.Bytes()
	return
}

func (s *ofp10.QueueStatsRequest) UnpackBinary(data []byte) (err error) {
	buf := bytes.NewBuffer(data)
	binary.Read(buf, binary.BigEndian, &s.PortNO)
	binary.Read(buf, binary.BigEndian, &s.pad)
	binary.Read(buf, binary.BigEndian, &s.QueueID)
	return
}

func NewQueueStats() (s *ofp10.QueueStats) {
	s = new(ofp10.QueueStats)
	return
}

func (s *ofp10.QueueStats) Len() (l int) {
	l = 32
	return
}

func (s *ofp10.QueueStats) PackBinary() (data []byte, err error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	binary.Write(buf, binary.BigEndian, s.PortNO)
	binary.Write(buf, binary.BigEndian, s.pad)
	binary.Write(buf, binary.BigEndian, s.QueueID)
	binary.Write(buf, binary.BigEndian, s.TxBytes)
	binary.Write(buf, binary.BigEndian, s.TxPackets)
	binary.Write(buf, binary.BigEndian, s.TxErrors)
	data = buf.Bytes()
	return
}

func (s *ofp10.QueueStats) UnpackBinary(data []byte) (err error) {
	buf := bytes.NewBuffer(data)
	binary.Read(buf, binary.BigEndian, &s.PortNO)
	binary.Read(buf, binary.BigEndian, &s.pad)
	binary.Read(buf, binary.BigEndian, &s.QueueID)
	binary.Read(buf, binary.BigEndian, &s.TxBytes)
	binary.Read(buf, binary.BigEndian, &s.TxPackets)
	binary.Read(buf, binary.BigEndian, &s.TxErrors)
	return
}

func EncodeStats(m buffer.Message) (data []byte, err error) {
	data = make([]byte, 0)
	switch m.(type) {
	case *ofp10.DescStats:
		data, err = m.(*ofp10.DescStats).PackBinary()
	case *ofp10.FlowStatsRequest:
		data, err = m.(*ofp10.FlowStatsRequest).PackBinary()
	case *ofp10.AggregateStatsRequest:
		data, err = m.(*ofp10.AggregateStatsRequest).PackBinary()
	case *ofp10.TableStats:
		data, err = m.(*ofp10.TableStats).PackBinary()
	case *ofp10.PortStatsRequest:
		data, err = m.(*ofp10.PortStatsRequest).PackBinary()
	case *ofp10.QueueStatsRequest:
		data, err = m.(*ofp10.QueueStatsRequest).PackBinary()
	case *VendorHeader:
		data, err = m.(*VendorHeader).PackBinary()
	}
	if err != nil {
		return
	}
	return
}

func DecodeStats(data []byte, t uint16) (m buffer.Message, err error) {
	buf := bytes.NewBuffer(data)
	switch t {
	case OFPSTDesc:
		m = new(ofp10.DescStats)
	case OFPSTFlow:
		m = new(ofp10.FlowStatsRequest)
	case OFPSTAggregate:
		m = new(ofp10.AggregateStatsRequest)
	case OFPSTTable:
		m = new(ofp10.TableStats)
	case OFPSTPort:
		m = new(ofp10.PortStatsRequest)
	case OFPSTQueue:
		m = new(ofp10.QueueStatsRequest)
	case OFPSTVendor:
		m = new(VendorHeader)
	}
	ss := make([]byte, m.Len())
	binary.Read(buf, binary.BigEndian, ss)
	err = m.UnpackBinary(ss)
	if err != nil {
		return
	}
	return
}
