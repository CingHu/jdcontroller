package ofp11

import (
	"bytes"
	"encoding/binary"
	"errors"
	"jd.com/jdcontroller/lib/buffer"
)

func NewStatsRequest() (s *StatsRequest) {
	s = new(StatsRequest)
	return
}

func (s *StatsRequest) Len() (l int) {
	l = 16
	return
}

func (s *StatsRequest) PackBinary() (data []byte, err error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	hs := make([]byte, 0)
	hs, err = s.Header.PackBinary()
	if err != nil {
		return
	}
	binary.Write(buf, binary.BigEndian, hs)
	binary.Write(buf, binary.BigEndian, s.Type)
	binary.Write(buf, binary.BigEndian, s.Flags)
	binary.Write(buf, binary.BigEndian, s.pad)
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

func (s *StatsRequest) UnpackBinary(data []byte) (err error) {
	buf := bytes.NewBuffer(data)
	hs := make([]byte, s.Header.Len())
	binary.Read(buf, binary.BigEndian, hs)
	err = s.Header.UnpackBinary(hs)
	if err != nil {
		return
	}
	binary.Read(buf, binary.BigEndian, &s.Type)
	binary.Read(buf, binary.BigEndian, &s.Flags)
	binary.Read(buf, binary.BigEndian, &s.pad)
	n := s.Len()
	if n < len(data) {
		s.Body, err = DecodeStats(data[n:], s.Type)
		if err != nil {
			return
		}
	}
	return
}

func NewStatsReply() (s *StatsReply) {
	s = new(StatsReply)
	return
}

func (s *StatsReply) Len() (l int) {
	l = 16
	return
}

func (s *StatsReply) PackBinary() (data []byte, err error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	hs := make([]byte, 0)
	hs, err = s.Header.PackBinary()
	if err != nil {
		return
	}
	binary.Write(buf, binary.BigEndian, hs)
	binary.Write(buf, binary.BigEndian, s.Type)
	binary.Write(buf, binary.BigEndian, s.Flags)
	binary.Write(buf, binary.BigEndian, s.pad)
	if s.Body != nil {
		bs := make([]byte, 0)
		switch s.Type {
		case OFPSTDesc:
			bs, err = s.Body.(*DescStats).PackBinary()
		case OFPSTFlow:
			bs, err = s.Body.(*FlowStatsRequest).PackBinary()
		case OFPSTAggregate:
			bs, err = s.Body.(*AggregateStatsRequest).PackBinary()
		case OFPSTTable:
			bs, err = s.Body.(*TableStats).PackBinary()
		case OFPSTPort:
			bs, err = s.Body.(*PortStatsRequest).PackBinary()
		case OFPSTQueue:
			bs, err = s.Body.(*QueueStatsRequest).PackBinary()
		case OFPSTGroup:
			bs, err = s.Body.(*GroupStatsRequest).PackBinary()
		case OFPSTGroupDesc:
			bs, err = s.Body.(*GroupDescStats).PackBinary()
		case OFPSTExperimenter:
			return
		}
		if err != nil {
			return
		}
		binary.Write(buf, binary.BigEndian, bs)
	}
	data = buf.Bytes()
	return
}

func (s *StatsReply) UnpackBinary(data []byte) (err error) {
	buf := bytes.NewBuffer(data)
	hs := make([]byte, s.Header.Len())
	binary.Read(buf, binary.BigEndian, hs)
	err = s.Header.UnpackBinary(hs)
	if err != nil {
		return
	}
	binary.Read(buf, binary.BigEndian, &s.Type)
	binary.Read(buf, binary.BigEndian, &s.Flags)
	binary.Read(buf, binary.BigEndian, &s.pad)
	n := s.Len()
	if n < len(data) {
		var r buffer.Message
		switch s.Type {
		case OFPSTDesc:
			r = s.Body.(*DescStats)
		case OFPSTFlow:
			r = s.Body.(*FlowStatsRequest)
		case OFPSTAggregate:
			r = s.Body.(*AggregateStatsRequest)
		case OFPSTTable:
			r = s.Body.(*TableStats)
		case OFPSTPort:
			r = s.Body.(*PortStatsRequest)
		case OFPSTQueue:
			r = s.Body.(*QueueStatsRequest)
		case OFPSTGroup:
			r = s.Body.(*GroupStatsRequest)
		case OFPSTGroupDesc:
			r = s.Body.(*GroupDescStats)
		case OFPSTExperimenter:
			return
		}
		ss := make([]byte, r.Len())
		binary.Read(buf, binary.BigEndian, ss)
		err = r.UnpackBinary(ss)
		if err != nil {
			return
		}
	}
	return
}

func NewDescStats() (s *DescStats) {
	s = new(DescStats)
	return
}

func (s *DescStats) Len() (l int) {
	l = OFPDescStrLen*4 + OFPSerialNumLen
	return
}

func (s *DescStats) PackBinary() (data []byte, err error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	binary.Write(buf, binary.BigEndian, s.MfrDesc)
	binary.Write(buf, binary.BigEndian, s.HWDesc)
	binary.Write(buf, binary.BigEndian, s.SWDesc)
	binary.Write(buf, binary.BigEndian, s.SerialNum)
	binary.Write(buf, binary.BigEndian, s.DPDesc)
	data = buf.Bytes()
	return
}

func (s *DescStats) UnpackBinary(data []byte) (err error) {
	buf := bytes.NewBuffer(data)
	binary.Read(buf, binary.BigEndian, &s.MfrDesc)
	binary.Read(buf, binary.BigEndian, &s.HWDesc)
	binary.Read(buf, binary.BigEndian, &s.SWDesc)
	binary.Read(buf, binary.BigEndian, &s.SerialNum)
	binary.Read(buf, binary.BigEndian, &s.DPDesc)
	return
}

func NewFlowStatsRequest() (s *FlowStatsRequest) {
	s = new(FlowStatsRequest)
	return
}

func (s *FlowStatsRequest) Len() (l int) {
	l = 120
	return
}

func (s *FlowStatsRequest) PackBinary() (data []byte, err error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	binary.Write(buf, binary.BigEndian, s.TableID)
	binary.Write(buf, binary.BigEndian, s.pad)
	binary.Write(buf, binary.BigEndian, s.OutPort)
	binary.Write(buf, binary.BigEndian, s.OutGroup)
	binary.Write(buf, binary.BigEndian, s.pad2)
	binary.Write(buf, binary.BigEndian, s.Cookie)
	binary.Write(buf, binary.BigEndian, s.CookieMask)

	ms := make([]byte, 0)
	ms, err = s.Match.PackBinary()
	if err != nil {
		return
	}
	binary.Write(buf, binary.BigEndian, ms)
	data = buf.Bytes()
	return
}

func (s *FlowStatsRequest) UnpackBinary(data []byte) (err error) {
	buf := bytes.NewBuffer(data)
	binary.Read(buf, binary.BigEndian, &s.TableID)
	binary.Read(buf, binary.BigEndian, &s.pad)
	binary.Read(buf, binary.BigEndian, &s.OutPort)
	binary.Read(buf, binary.BigEndian, &s.OutGroup)
	binary.Read(buf, binary.BigEndian, &s.pad2)
	binary.Read(buf, binary.BigEndian, &s.Cookie)
	binary.Read(buf, binary.BigEndian, &s.CookieMask)

	ms := make([]byte, s.Match.Len())
	binary.Read(buf, binary.BigEndian, ms)
	err = s.Match.UnpackBinary(ms)
	if err != nil {
		return
	}
	return
}

func NewFlowStats() (s *FlowStats) {
	s = new(FlowStats)
	return
}

func (s *FlowStats) Len() (l int) {
	l = 136
	return
}

func (s *FlowStats) PackBinary() (data []byte, err error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	binary.Write(buf, binary.BigEndian, s.Length)
	binary.Write(buf, binary.BigEndian, s.TableID)
	binary.Write(buf, binary.BigEndian, s.pad)
	binary.Write(buf, binary.BigEndian, s.DurationSec)
	binary.Write(buf, binary.BigEndian, s.DurationNSec)
	binary.Write(buf, binary.BigEndian, s.Priority)
	binary.Write(buf, binary.BigEndian, s.IdleTimeout)
	binary.Write(buf, binary.BigEndian, s.HardTimeout)
	binary.Write(buf, binary.BigEndian, s.pad2)
	binary.Write(buf, binary.BigEndian, s.Cookie)
	binary.Write(buf, binary.BigEndian, s.PacketCount)
	binary.Write(buf, binary.BigEndian, s.ByteCount)

	ms := make([]byte, 0)
	ms, err = s.Match.PackBinary()
	if err != nil {
		return
	}
	binary.Write(buf, binary.BigEndian, ms)
	for _, i := range s.Instructions {
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

func (s *FlowStats) UnpackBinary(data []byte) (err error) {
	buf := bytes.NewBuffer(data)
	binary.Read(buf, binary.BigEndian, &s.Length)
	binary.Read(buf, binary.BigEndian, &s.TableID)
	binary.Read(buf, binary.BigEndian, &s.pad)
	binary.Read(buf, binary.BigEndian, &s.DurationSec)
	binary.Read(buf, binary.BigEndian, &s.DurationNSec)
	binary.Read(buf, binary.BigEndian, &s.Priority)
	binary.Read(buf, binary.BigEndian, &s.IdleTimeout)
	binary.Read(buf, binary.BigEndian, &s.HardTimeout)
	binary.Read(buf, binary.BigEndian, &s.pad2)
	binary.Read(buf, binary.BigEndian, &s.Cookie)
	binary.Read(buf, binary.BigEndian, &s.PacketCount)
	binary.Read(buf, binary.BigEndian, &s.ByteCount)

	ms := make([]byte, s.Match.Len())
	binary.Read(buf, binary.BigEndian, ms)
	err = s.Match.UnpackBinary(ms)
	if err != nil {
		return
	}

	n := s.Len()
	for n < len(data) {
		var i Instruction
		i, err = DecodeInstruction(data[n:])
		if err != nil {
			return
		}
		s.Instructions = append(s.Instructions, i)
		n += i.Len()
	}
	return
}

func NewAggregateStatsRequest() (s *AggregateStatsRequest) {
	s = new(AggregateStatsRequest)
	return
}

func (s *AggregateStatsRequest) Len() (l int) {
	l = 120
	return
}

func (s *AggregateStatsRequest) PackBinary() (data []byte, err error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	binary.Write(buf, binary.BigEndian, s.TableID)
	binary.Write(buf, binary.BigEndian, s.pad)
	binary.Write(buf, binary.BigEndian, s.OutPort)
	binary.Write(buf, binary.BigEndian, s.OutGroup)
	binary.Write(buf, binary.BigEndian, s.pad2)
	binary.Write(buf, binary.BigEndian, s.Cookie)
	binary.Write(buf, binary.BigEndian, s.CookieMask)

	ms := make([]byte, 0)
	ms, err = s.Match.PackBinary()
	if err != nil {
		return
	}
	binary.Write(buf, binary.BigEndian, ms)
	data = buf.Bytes()
	return
}

func (s *AggregateStatsRequest) UnpackBinary(data []byte) (err error) {
	buf := bytes.NewBuffer(data)
	binary.Read(buf, binary.BigEndian, &s.TableID)
	binary.Read(buf, binary.BigEndian, &s.pad)
	binary.Read(buf, binary.BigEndian, &s.OutPort)
	binary.Read(buf, binary.BigEndian, &s.OutGroup)
	binary.Read(buf, binary.BigEndian, &s.pad2)
	binary.Read(buf, binary.BigEndian, &s.Cookie)
	binary.Read(buf, binary.BigEndian, &s.CookieMask)

	ms := make([]byte, s.Match.Len())
	binary.Read(buf, binary.BigEndian, ms)
	err = s.Match.UnpackBinary(ms)
	if err != nil {
		return
	}
	return
}

func NewAggregateStatsReply() (s *AggregateStatsReply) {
	s = new(AggregateStatsReply)
	return
}

func (s *AggregateStatsReply) Len() (l int) {
	l = 120
	return
}

func (s *AggregateStatsReply) PackBinary() (data []byte, err error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	binary.Write(buf, binary.BigEndian, s.PacketCount)
	binary.Write(buf, binary.BigEndian, s.ByteCount)
	binary.Write(buf, binary.BigEndian, s.FlowCount)
	binary.Write(buf, binary.BigEndian, s.pad)
	data = buf.Bytes()
	return
}

func (s *AggregateStatsReply) UnpackBinary(data []byte) (err error) {
	buf := bytes.NewBuffer(data)
	binary.Read(buf, binary.BigEndian, &s.PacketCount)
	binary.Read(buf, binary.BigEndian, &s.ByteCount)
	binary.Read(buf, binary.BigEndian, &s.FlowCount)
	binary.Read(buf, binary.BigEndian, &s.pad)
	return
}

func NewTableStats() (s *TableStats) {
	s = new(TableStats)
	return
}

func (s *TableStats) Len() (l int) {
	l = 88
	return
}

func (s *TableStats) PackBinary() (data []byte, err error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	binary.Write(buf, binary.BigEndian, s.TableID)
	binary.Write(buf, binary.BigEndian, s.pad)
	binary.Write(buf, binary.BigEndian, s.Name)
	binary.Write(buf, binary.BigEndian, s.Wildcards)
	binary.Write(buf, binary.BigEndian, s.Match)
	binary.Write(buf, binary.BigEndian, s.Instructions)
	binary.Write(buf, binary.BigEndian, s.WriteActions)
	binary.Write(buf, binary.BigEndian, s.ApplyActions)
	binary.Write(buf, binary.BigEndian, s.Config)
	binary.Write(buf, binary.BigEndian, s.MaxEntries)
	binary.Write(buf, binary.BigEndian, s.ActiveCount)
	binary.Write(buf, binary.BigEndian, s.LookupCount)
	binary.Write(buf, binary.BigEndian, s.MatchedCount)
	data = buf.Bytes()
	return
}

func (s *TableStats) UnpackBinary(data []byte) (err error) {
	buf := bytes.NewBuffer(data)
	binary.Read(buf, binary.BigEndian, &s.TableID)
	binary.Read(buf, binary.BigEndian, &s.pad)
	binary.Read(buf, binary.BigEndian, &s.Name)
	binary.Read(buf, binary.BigEndian, &s.Wildcards)
	binary.Read(buf, binary.BigEndian, &s.Match)
	binary.Read(buf, binary.BigEndian, &s.Instructions)
	binary.Read(buf, binary.BigEndian, &s.WriteActions)
	binary.Read(buf, binary.BigEndian, &s.ApplyActions)
	binary.Read(buf, binary.BigEndian, &s.Config)
	binary.Read(buf, binary.BigEndian, &s.MaxEntries)
	binary.Read(buf, binary.BigEndian, &s.ActiveCount)
	binary.Read(buf, binary.BigEndian, &s.LookupCount)
	binary.Read(buf, binary.BigEndian, &s.MatchedCount)
	return
}

func NewPortStatsRequest() (s *PortStatsRequest) {
	s = new(PortStatsRequest)
	return
}

func (s *PortStatsRequest) Len() (l int) {
	l = 8
	return
}

func (s *PortStatsRequest) PackBinary() (data []byte, err error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	binary.Write(buf, binary.BigEndian, s.PortNO)
	binary.Write(buf, binary.BigEndian, s.pad)
	data = buf.Bytes()
	return
}

func (s *PortStatsRequest) UnpackBinary(data []byte) (err error) {
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
	return 104
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

func NewQueueStatsRequest() (s *QueueStatsRequest) {
	s = new(QueueStatsRequest)
	return
}

func (s *QueueStatsRequest) Len() (l int) {
	l = 8
	return
}

func (s *QueueStatsRequest) PackBinary() (data []byte, err error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	binary.Write(buf, binary.BigEndian, s.PortNO)
	binary.Write(buf, binary.BigEndian, s.QueueID)
	data = buf.Bytes()
	return
}

func (s *QueueStatsRequest) UnpackBinary(data []byte) (err error) {
	buf := bytes.NewBuffer(data)
	binary.Read(buf, binary.BigEndian, &s.PortNO)
	binary.Read(buf, binary.BigEndian, &s.QueueID)
	return
}

func NewQueueStats() (s *QueueStats) {
	s = new(QueueStats)
	return
}

func (s *QueueStats) Len() (l int) {
	return 32
}

func (s *QueueStats) PackBinary() (data []byte, err error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	binary.Write(buf, binary.BigEndian, s.PortNO)
	binary.Write(buf, binary.BigEndian, s.QueueID)
	binary.Write(buf, binary.BigEndian, s.TxBytes)
	binary.Write(buf, binary.BigEndian, s.TxPackets)
	binary.Write(buf, binary.BigEndian, s.TxErrors)
	data = buf.Bytes()
	return
}

func (s *QueueStats) UnpackBinary(data []byte) (err error) {
	buf := bytes.NewBuffer(data)
	binary.Read(buf, binary.BigEndian, &s.PortNO)
	binary.Read(buf, binary.BigEndian, &s.QueueID)
	binary.Read(buf, binary.BigEndian, &s.TxBytes)
	binary.Read(buf, binary.BigEndian, &s.TxPackets)
	binary.Read(buf, binary.BigEndian, &s.TxErrors)
	return
}

func NewGroupStatsRequest() (s *GroupStatsRequest) {
	s = new(GroupStatsRequest)
	return
}

func (s *GroupStatsRequest) Len() (l int) {
	l = 8
	return
}

func (s *GroupStatsRequest) PackBinary() (data []byte, err error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	binary.Write(buf, binary.BigEndian, s.GroupID)
	binary.Write(buf, binary.BigEndian, s.pad)
	data = buf.Bytes()
	return
}

func (s *GroupStatsRequest) UnpackBinary(data []byte) (err error) {
	buf := bytes.NewBuffer(data)
	binary.Read(buf, binary.BigEndian, &s.GroupID)
	binary.Read(buf, binary.BigEndian, &s.pad)
	return
}

func NewGroupStats() (s *GroupStats) {
	s = new(GroupStats)
	return
}

func (s *GroupStats) Len() (l int) {
	return 32
}

func (s *GroupStats) PackBinary() (data []byte, err error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	binary.Write(buf, binary.BigEndian, s.Length)
	binary.Write(buf, binary.BigEndian, s.pad)
	binary.Write(buf, binary.BigEndian, s.GroupID)
	binary.Write(buf, binary.BigEndian, s.RefCount)
	binary.Write(buf, binary.BigEndian, s.pad2)
	binary.Write(buf, binary.BigEndian, s.PacketCount)
	binary.Write(buf, binary.BigEndian, s.ByteCount)
	for _, b := range s.BucketStats {
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

func (s *GroupStats) UnpackBinary(data []byte) (err error) {
	buf := bytes.NewBuffer(data)
	binary.Read(buf, binary.BigEndian, &s.Length)
	binary.Read(buf, binary.BigEndian, &s.pad)
	binary.Read(buf, binary.BigEndian, &s.GroupID)
	binary.Read(buf, binary.BigEndian, &s.RefCount)
	binary.Read(buf, binary.BigEndian, &s.pad2)
	binary.Read(buf, binary.BigEndian, &s.PacketCount)
	binary.Read(buf, binary.BigEndian, &s.ByteCount)
	n := int(s.Len())
	for n < len(data) {
		b := new(BucketCounter)
		bs := make([]byte, b.Len())
		binary.Read(buf, binary.BigEndian, bs)
		err = b.UnpackBinary(bs)
		if err != nil {
			return
		}
		s.BucketStats = append(s.BucketStats, *b)
		n += b.Len()
	}
	return
}

func NewBucketCounter() (s *BucketCounter) {
	s = new(BucketCounter)
	return
}

func (s *BucketCounter) Len() (l int) {
	return 16
}

func (s *BucketCounter) PackBinary() (data []byte, err error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	binary.Write(buf, binary.BigEndian, s.PacketCount)
	binary.Write(buf, binary.BigEndian, s.ByteCount)
	data = buf.Bytes()
	return
}

func (s *BucketCounter) UnpackBinary(data []byte) (err error) {
	buf := bytes.NewBuffer(data)
	binary.Read(buf, binary.BigEndian, &s.PacketCount)
	binary.Read(buf, binary.BigEndian, &s.ByteCount)
	return
}

func NewGroupDescStats() (s *GroupDescStats) {
	s = new(GroupDescStats)
	return
}

func (s *GroupDescStats) Len() (l int) {
	return 16
}

func (s *GroupDescStats) PackBinary() (data []byte, err error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	binary.Write(buf, binary.BigEndian, s.Length)
	binary.Write(buf, binary.BigEndian, s.Type)
	binary.Write(buf, binary.BigEndian, s.pad)
	binary.Write(buf, binary.BigEndian, s.GroupID)
	for _, b := range s.Buckets {
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

func (s *GroupDescStats) UnpackBinary(data []byte) (err error) {
	buf := bytes.NewBuffer(data)
	binary.Read(buf, binary.BigEndian, &s.Length)
	binary.Read(buf, binary.BigEndian, &s.Type)
	binary.Read(buf, binary.BigEndian, &s.pad)
	binary.Read(buf, binary.BigEndian, &s.GroupID)
	n := int(s.Len())
	for n < len(data) {
		b := new(Bucket)
		bs := make([]byte, b.Len())
		binary.Read(buf, binary.BigEndian, bs)
		err = b.UnpackBinary(bs)
		if err != nil {
			return
		}
		s.Buckets = append(s.Buckets, *b)
		n += b.Len()
	}
	return
}

func EncodeStats(m buffer.Message) (data []byte, err error) {
	data = make([]byte, 0)
	switch m.(type) {
	case *DescStats:
		data, err = m.(*DescStats).PackBinary()
	case *FlowStatsRequest:
		data, err = m.(*FlowStatsRequest).PackBinary()
	case *AggregateStatsRequest:
		data, err = m.(*AggregateStatsRequest).PackBinary()
	case *TableStats:
		data, err = m.(*TableStats).PackBinary()
	case *PortStatsRequest:
		data, err = m.(*PortStatsRequest).PackBinary()
	case *QueueStatsRequest:
		data, err = m.(*QueueStatsRequest).PackBinary()
	case *GroupStatsRequest:
		data, err = m.(*GroupStatsRequest).PackBinary()
	case *GroupDescStats:
		data, err = m.(*GroupDescStats).PackBinary()
	default:
		err = errors.New("Can not parse this stats request.")
		return
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
		m = new(DescStats)
	case OFPSTFlow:
		m = new(FlowStatsRequest)
	case OFPSTAggregate:
		m = new(AggregateStatsRequest)
	case OFPSTTable:
		m = new(TableStats)
	case OFPSTPort:
		m = new(PortStatsRequest)
	case OFPSTQueue:
		m = new(QueueStatsRequest)
	case OFPSTGroup:
		m = new(GroupStatsRequest)
	case OFPSTGroupDesc:
		m = new(GroupDescStats)
	default:
		err = errors.New("Can not parse this stats request.")
		return
	}
	ss := make([]byte, m.Len())
	binary.Read(buf, binary.BigEndian, ss)
	err = m.UnpackBinary(ss)
	if err != nil {
		return
	}

	return
}
