package ofpctl13

import (
	"bytes"
	"encoding/binary"
	
	"jd.com/jdcontroller/protocol/ofp13"
)

func NewDesc() (d *ofp13.Desc) {
	d = new(ofp13.Desc)
	return
}

func (d *ofp13.Desc) Len() (l int) {
	l = ofp13.OFPDescStrLen*4 + ofp13.OFPSerialNumLen
	return
}

func (d *ofp13.Desc) PackBinary() (data []byte, err error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	binary.Write(buf, binary.BigEndian, d.Mfrofp13.Desc)
	binary.Write(buf, binary.BigEndian, d.HWofp13.Desc)
	binary.Write(buf, binary.BigEndian, d.SWofp13.Desc)
	binary.Write(buf, binary.BigEndian, d.SerialNum)
	binary.Write(buf, binary.BigEndian, d.DPofp13.Desc)
	data = buf.Bytes()
	return
}

func (d *ofp13.Desc) UnpackBinary(data []byte) (err error) {
	buf := bytes.NewBuffer(data)
	binary.Read(buf, binary.BigEndian, &d.Mfrofp13.Desc)
	binary.Read(buf, binary.BigEndian, &d.HWofp13.Desc)
	binary.Read(buf, binary.BigEndian, &d.SWofp13.Desc)
	binary.Read(buf, binary.BigEndian, &d.SerialNum)
	binary.Read(buf, binary.BigEndian, &d.DPofp13.Desc)
	return
}

func NewFlowStatsRequest() (s *ofp13.FlowStatsRequest) {
	s = new(ofp13.FlowStatsRequest)
	return
}

func (s *ofp13.FlowStatsRequest) Len() (l int) {
	l = 40
	return
}

func (s *ofp13.FlowStatsRequest) PackBinary() (data []byte, err error) {
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

func (s *ofp13.FlowStatsRequest) UnpackBinary(data []byte) (err error) {
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

func NewFlowStats() (s *ofp13.FlowStats) {
	s = new(ofp13.FlowStats)
	return
}

func (s *ofp13.FlowStats) Len() (l int) {
	l = 56
	return
}

func (s *ofp13.FlowStats) PackBinary() (data []byte, err error) {
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
	for _, i := range s.ofp13.Instructions {
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

func (s *ofp13.FlowStats) UnpackBinary(data []byte) (err error) {
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
		var i ofp13.Instruction
		i, err = DecodeInstruction(data[n:])
		if err != nil {
			return
		}
		s.ofp13.Instructions = append(s.ofp13.Instructions, i)
		n += i.Len()
	}
	return
}

func NewAggregateStatsRequest() (s *ofp13.AggregateStatsRequest) {
	s = new(ofp13.AggregateStatsRequest)
	return
}

func (s *ofp13.AggregateStatsRequest) Len() (l int) {
	l = 40
	return
}

func (s *ofp13.AggregateStatsRequest) PackBinary() (data []byte, err error) {
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

func (s *ofp13.AggregateStatsRequest) UnpackBinary(data []byte) (err error) {
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

func NewAggregateStatsReply() (s *ofp13.AggregateStatsReply) {
	s = new(ofp13.AggregateStatsReply)
	return
}

func (s *ofp13.AggregateStatsReply) Len() (l int) {
	l = 120
	return
}

func (s *ofp13.AggregateStatsReply) PackBinary() (data []byte, err error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	binary.Write(buf, binary.BigEndian, s.PacketCount)
	binary.Write(buf, binary.BigEndian, s.ByteCount)
	binary.Write(buf, binary.BigEndian, s.FlowCount)
	binary.Write(buf, binary.BigEndian, s.pad)
	data = buf.Bytes()
	return
}

func (s *ofp13.AggregateStatsReply) UnpackBinary(data []byte) (err error) {
	buf := bytes.NewBuffer(data)
	binary.Read(buf, binary.BigEndian, &s.PacketCount)
	binary.Read(buf, binary.BigEndian, &s.ByteCount)
	binary.Read(buf, binary.BigEndian, &s.FlowCount)
	binary.Read(buf, binary.BigEndian, &s.pad)
	return
}

func NewTableStats() (s *ofp13.TableStats) {
	s = new(ofp13.TableStats)
	return
}

func (s *ofp13.TableStats) Len() (l int) {
	l = 24
	return
}

func (s *ofp13.TableStats) PackBinary() (data []byte, err error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	binary.Write(buf, binary.BigEndian, s.TableID)
	binary.Write(buf, binary.BigEndian, s.pad)
	binary.Write(buf, binary.BigEndian, s.ActiveCount)
	binary.Write(buf, binary.BigEndian, s.LookupCount)
	binary.Write(buf, binary.BigEndian, s.MatchedCount)
	data = buf.Bytes()
	return
}

func (s *ofp13.TableStats) UnpackBinary(data []byte) (err error) {
	buf := bytes.NewBuffer(data)
	binary.Read(buf, binary.BigEndian, &s.TableID)
	binary.Read(buf, binary.BigEndian, &s.pad)
	binary.Read(buf, binary.BigEndian, &s.ActiveCount)
	binary.Read(buf, binary.BigEndian, &s.LookupCount)
	binary.Read(buf, binary.BigEndian, &s.MatchedCount)
	return
}

func NewTableFeatures() (f *ofp13.TableFeatures) {
	f = new(ofp13.TableFeatures)
	return
}

func (f *ofp13.TableFeatures) Len() (l int) {
	l = 24
	return
}

func (f *ofp13.TableFeatures) PackBinary() (data []byte, err error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	binary.Write(buf, binary.BigEndian, f.Length)
	binary.Write(buf, binary.BigEndian, f.TableID)
	binary.Write(buf, binary.BigEndian, f.pad)
	binary.Write(buf, binary.BigEndian, f.Name)
	binary.Write(buf, binary.BigEndian, f.MetadataMatch)
	binary.Write(buf, binary.BigEndian, f.MetadataWrite)
	binary.Write(buf, binary.BigEndian, f.Config)
	binary.Write(buf, binary.BigEndian, f.MaxEntries)
	for _, p := range f.Properties {
		bs := make([]byte, 0)
		bs, err = EncodeTableFeatureProp(p)
		if err != nil {
			return
		}
		binary.Write(buf, binary.BigEndian, bs)
	}
	data = buf.Bytes()
	return
}

func (f *ofp13.TableFeatures) UnpackBinary(data []byte) (err error) {
	buf := bytes.NewBuffer(data)
	binary.Read(buf, binary.BigEndian, &f.Length)
	binary.Read(buf, binary.BigEndian, &f.TableID)
	binary.Read(buf, binary.BigEndian, &f.pad)
	binary.Read(buf, binary.BigEndian, &f.Name)
	binary.Read(buf, binary.BigEndian, &f.MetadataMatch)
	binary.Read(buf, binary.BigEndian, &f.MetadataWrite)
	binary.Read(buf, binary.BigEndian, &f.Config)
	binary.Read(buf, binary.BigEndian, &f.MaxEntries)

	n := f.Len()
	for n < len(data) {
		var t ofp13.TableFeatureProp
		t, err = DecodeTableFeatureProp(data[n:])
		if err != nil {
			return
		}
		f.Properties = append(f.Properties, t)
		n += t.Len()
	}
	return
}

func NewTableFeaturesPropHeader() (f *ofp13.TableFeaturePropHeader) {
	f = new(ofp13.TableFeaturePropHeader)
	return
}

func (f *ofp13.TableFeaturePropHeader) Len() (l int) {
	l = 4
	return
}

func (f *ofp13.TableFeaturePropHeader) PackBinary() (data []byte, err error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	binary.Write(buf, binary.BigEndian, f.Type)
	binary.Write(buf, binary.BigEndian, f.Length)
	data = buf.Bytes()
	return
}

func (f *ofp13.TableFeaturePropHeader) UnpackBinary(data []byte) (err error) {
	buf := bytes.NewBuffer(data)
	binary.Read(buf, binary.BigEndian, &f.Type)
	binary.Read(buf, binary.BigEndian, &f.Length)
	return
}

func NewTableFeaturesPropInstructions() (f *ofp13.TableFeaturePropInstructions) {
	f = new(ofp13.TableFeaturePropInstructions)
	return
}

func (f *ofp13.TableFeaturePropInstructions) Len() (l int) {
	l = 4
	return
}

func (f *ofp13.TableFeaturePropInstructions) PackBinary() (data []byte, err error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	hs := make([]byte, 0)
	hs, err = f.Header.PackBinary()
	if err != nil {
		return
	}
	binary.Write(buf, binary.BigEndian, hs)
	for _, i := range f.ofp13.Instructions {
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

func (f *ofp13.TableFeaturePropInstructions) UnpackBinary(data []byte) (err error) {
	buf := bytes.NewBuffer(data)
	hs := make([]byte, f.Header.Len())
	binary.Read(buf, binary.BigEndian, hs)
	err = f.Header.UnpackBinary(hs)
	if err != nil {
		return
	}

	n := f.Len()
	for n < len(data) {
		var i ofp13.Instruction
		i, err = DecodeInstruction(data[n:])
		f.ofp13.Instructions = append(f.ofp13.Instructions, i)
		n += i.Len()
	}
	return
}

func NewTableFeaturesPropNextTables() (f *ofp13.TableFeaturePropNextTables) {
	f = new(ofp13.TableFeaturePropNextTables)
	return
}

func (f *ofp13.TableFeaturePropNextTables) Len() (l int) {
	l = 4
	return
}

func (f *ofp13.TableFeaturePropNextTables) PackBinary() (data []byte, err error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	hs := make([]byte, 0)
	hs, err = f.Header.PackBinary()
	if err != nil {
		return
	}
	binary.Write(buf, binary.BigEndian, hs)
	for _, t := range f.NextTable {
		binary.Write(buf, binary.BigEndian, t)
	}
	data = buf.Bytes()
	return
}

func (f *ofp13.TableFeaturePropNextTables) UnpackBinary(data []byte) (err error) {
	buf := bytes.NewBuffer(data)
	hs := make([]byte, f.Header.Len())
	binary.Read(buf, binary.BigEndian, hs)
	err = f.Header.UnpackBinary(hs)
	if err != nil {
		return
	}

	n := f.Len()
	for n < len(data) {
		b := new(uint8)
		binary.Read(buf, binary.BigEndian, b)
		f.NextTable = append(f.NextTable, *b)
		n += 1
	}
	return
}

func NewTableFeaturesPropActions() (f *ofp13.TableFeaturePropActions) {
	f = new(ofp13.TableFeaturePropActions)
	return
}

func (f *ofp13.TableFeaturePropActions) Len() (l int) {
	l = 4
	return
}

func (f *ofp13.TableFeaturePropActions) PackBinary() (data []byte, err error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	hs := make([]byte, 0)
	hs, err = f.Header.PackBinary()
	if err != nil {
		return
	}
	binary.Write(buf, binary.BigEndian, hs)
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

func (f *ofp13.TableFeaturePropActions) UnpackBinary(data []byte) (err error) {
	buf := bytes.NewBuffer(data)
	hs := make([]byte, f.Header.Len())
	binary.Read(buf, binary.BigEndian, hs)
	err = f.Header.UnpackBinary(hs)
	if err != nil {
		return
	}

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

func NewTableFeaturesPropOXM() (f *ofp13.TableFeaturePropOXM) {
	f = new(ofp13.TableFeaturePropOXM)
	return
}

func (f *ofp13.TableFeaturePropOXM) Len() (l int) {
	l = 4
	return
}

func (f *ofp13.TableFeaturePropOXM) PackBinary() (data []byte, err error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	hs := make([]byte, 0)
	hs, err = f.Header.PackBinary()
	if err != nil {
		return
	}
	binary.Write(buf, binary.BigEndian, hs)
	for _, o := range f.OXM {
		binary.Write(buf, binary.BigEndian, o)
	}
	data = buf.Bytes()
	return
}

func (f *ofp13.TableFeaturePropOXM) UnpackBinary(data []byte) (err error) {
	buf := bytes.NewBuffer(data)
	hs := make([]byte, f.Header.Len())
	binary.Read(buf, binary.BigEndian, hs)
	err = f.Header.UnpackBinary(hs)
	if err != nil {
		return
	}

	n := f.Len()
	for n < len(data) {
		b := new(uint32)
		binary.Read(buf, binary.BigEndian, b)
		f.OXM = append(f.OXM, *b)
		n += 1
	}
	return
}

func NewTableFeaturesPropExperimenter() (f *ofp13.TableFeaturePropExperimenter) {
	f = new(ofp13.TableFeaturePropExperimenter)
	return
}

func (f *ofp13.TableFeaturePropExperimenter) Len() (l int) {
	l = 4
	return
}

func (f *ofp13.TableFeaturePropExperimenter) PackBinary() (data []byte, err error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	hs := make([]byte, 0)
	hs, err = f.Header.PackBinary()
	if err != nil {
		return
	}
	binary.Write(buf, binary.BigEndian, hs)
	binary.Write(buf, binary.BigEndian, f.Experimenter)
	binary.Write(buf, binary.BigEndian, f.ExpType)
	for _, d := range f.Data {
		binary.Write(buf, binary.BigEndian, d)
	}
	data = buf.Bytes()
	return
}

func (f *ofp13.TableFeaturePropExperimenter) UnpackBinary(data []byte) (err error) {
	buf := bytes.NewBuffer(data)
	hs := make([]byte, f.Header.Len())
	binary.Read(buf, binary.BigEndian, hs)
	err = f.Header.UnpackBinary(hs)
	if err != nil {
		return
	}
	binary.Read(buf, binary.BigEndian, &f.Experimenter)
	binary.Read(buf, binary.BigEndian, &f.ExpType)

	n := f.Len()
	for n < len(data) {
		b := new(uint32)
		binary.Read(buf, binary.BigEndian, b)
		f.Data = append(f.Data, *b)
		n += 1
	}
	return
}

func NewPortStatsRequest(p uint32) (s *ofp13.PortStatsRequest) {
	s = new(PortStatsRequest)
	s.PortNO = p
	return
}

func (s *ofp13.PortStatsRequest) Len() (l int) {
	l = 8
	return
}

func (s *ofp13.PortStatsRequest) PackBinary() (data []byte, err error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	binary.Write(buf, binary.BigEndian, s.PortNO)
	binary.Write(buf, binary.BigEndian, s.pad)
	data = buf.Bytes()
	return
}

func (s *ofp13.PortStatsRequest) UnpackBinary(data []byte) (err error) {
	buf := bytes.NewBuffer(data)
	binary.Read(buf, binary.BigEndian, &s.PortNO)
	binary.Read(buf, binary.BigEndian, &s.pad)
	return
}

func NewPortStats() (s *ofp13.PortStats) {
	s = new(PortStats)
	return
}

func (s PortStats) Len() (l int) {
	return 112
}

func (s PortStats) PackBinary() (data []byte, err error) {
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
	binary.Write(buf, binary.BigEndian, s.DurationSec)
	binary.Write(buf, binary.BigEndian, s.DurationNSec)
	data = buf.Bytes()
	return
}

func (s *ofp13.PortStats) UnpackBinary(data []byte) (err error) {
	buf := bytes.NewBuffer(data)
	binary.Read(buf, binary.BigEndian, &s.PortNO)
	binary.Read(buf, binary.BigEndian, s.pad)
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
	binary.Read(buf, binary.BigEndian, &s.DurationSec)
	binary.Read(buf, binary.BigEndian, &s.DurationNSec)
	return
}

func NewQueueStatsRequest() (s *ofp13.QueueStatsRequest) {
	s = new(QueueStatsRequest)
	return
}

func (s *ofp13.QueueStatsRequest) Len() (l int) {
	l = 8
	return
}

func (s *ofp13.QueueStatsRequest) PackBinary() (data []byte, err error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	binary.Write(buf, binary.BigEndian, s.PortNO)
	binary.Write(buf, binary.BigEndian, s.QueueID)
	data = buf.Bytes()
	return
}

func (s *ofp13.QueueStatsRequest) UnpackBinary(data []byte) (err error) {
	buf := bytes.NewBuffer(data)
	binary.Read(buf, binary.BigEndian, &s.PortNO)
	binary.Read(buf, binary.BigEndian, &s.QueueID)
	return
}

func NewQueueStats() (s *ofp13.QueueStats) {
	s = new(QueueStats)
	return
}

func (s *ofp13.QueueStats) Len() (l int) {
	return 40
}

func (s *ofp13.QueueStats) PackBinary() (data []byte, err error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	binary.Write(buf, binary.BigEndian, s.PortNO)
	binary.Write(buf, binary.BigEndian, s.QueueID)
	binary.Write(buf, binary.BigEndian, s.TxBytes)
	binary.Write(buf, binary.BigEndian, s.TxPackets)
	binary.Write(buf, binary.BigEndian, s.TxErrors)
	binary.Write(buf, binary.BigEndian, s.DurationSec)
	binary.Write(buf, binary.BigEndian, s.DurationNSec)
	data = buf.Bytes()
	return
}

func (s *ofp13.QueueStats) UnpackBinary(data []byte) (err error) {
	buf := bytes.NewBuffer(data)
	binary.Read(buf, binary.BigEndian, &s.PortNO)
	binary.Read(buf, binary.BigEndian, &s.QueueID)
	binary.Read(buf, binary.BigEndian, &s.TxBytes)
	binary.Read(buf, binary.BigEndian, &s.TxPackets)
	binary.Read(buf, binary.BigEndian, &s.TxErrors)
	binary.Read(buf, binary.BigEndian, &s.DurationSec)
	binary.Read(buf, binary.BigEndian, &s.DurationNSec)
	return
}

func NewGroupStatsRequest() (s *ofp13.GroupStatsRequest) {
	s = new(GroupStatsRequest)
	return
}

func (s *ofp13.GroupStatsRequest) Len() (l int) {
	l = 8
	return
}

func (s *ofp13.GroupStatsRequest) PackBinary() (data []byte, err error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	binary.Write(buf, binary.BigEndian, s.GroupID)
	binary.Write(buf, binary.BigEndian, s.pad)
	data = buf.Bytes()
	return
}

func (s *ofp13.GroupStatsRequest) UnpackBinary(data []byte) (err error) {
	buf := bytes.NewBuffer(data)
	binary.Read(buf, binary.BigEndian, &s.GroupID)
	binary.Read(buf, binary.BigEndian, &s.pad)
	return
}

func NewGroupStats() (s *ofp13.GroupStats) {
	s = new(GroupStats)
	return
}

func (s *ofp13.GroupStats) Len() (l int) {
	return 40
}

func (s *ofp13.GroupStats) PackBinary() (data []byte, err error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	binary.Write(buf, binary.BigEndian, s.Length)
	binary.Write(buf, binary.BigEndian, s.pad)
	binary.Write(buf, binary.BigEndian, s.GroupID)
	binary.Write(buf, binary.BigEndian, s.RefCount)
	binary.Write(buf, binary.BigEndian, s.pad2)
	binary.Write(buf, binary.BigEndian, s.PacketCount)
	binary.Write(buf, binary.BigEndian, s.ByteCount)
	binary.Write(buf, binary.BigEndian, s.DurationSec)
	binary.Write(buf, binary.BigEndian, s.DurationNSec)
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

func (s *ofp13.GroupStats) UnpackBinary(data []byte) (err error) {
	buf := bytes.NewBuffer(data)
	binary.Read(buf, binary.BigEndian, &s.Length)
	binary.Read(buf, binary.BigEndian, &s.pad)
	binary.Read(buf, binary.BigEndian, &s.GroupID)
	binary.Read(buf, binary.BigEndian, &s.RefCount)
	binary.Read(buf, binary.BigEndian, &s.pad2)
	binary.Read(buf, binary.BigEndian, &s.PacketCount)
	binary.Read(buf, binary.BigEndian, &s.ByteCount)
	binary.Read(buf, binary.BigEndian, &s.DurationSec)
	binary.Read(buf, binary.BigEndian, &s.DurationNSec)
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

func NewBucketCounter() (s *ofp13.BucketCounter) {
	s = new(BucketCounter)
	return
}

func (s *ofp13.BucketCounter) Len() (l int) {
	return 16
}

func (s *ofp13.BucketCounter) PackBinary() (data []byte, err error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	binary.Write(buf, binary.BigEndian, s.PacketCount)
	binary.Write(buf, binary.BigEndian, s.ByteCount)
	data = buf.Bytes()
	return
}

func (s *ofp13.BucketCounter) UnpackBinary(data []byte) (err error) {
	buf := bytes.NewBuffer(data)
	binary.Read(buf, binary.BigEndian, &s.PacketCount)
	binary.Read(buf, binary.BigEndian, &s.ByteCount)
	return
}

func NewGroupDescStats() (s *ofp13.GroupDescStats) {
	s = new(GroupDescStats)
	return
}

func (s *ofp13.GroupDescStats) Len() (l int) {
	return 8
}

func (s *ofp13.GroupDescStats) PackBinary() (data []byte, err error) {
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

func (s *ofp13.GroupDescStats) UnpackBinary(data []byte) (err error) {
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

func NewGroupFeatures() (s *ofp13.GroupFeatures) {
	s = new(GroupFeatures)
	return
}

func (s *ofp13.GroupFeatures) Len() (l int) {
	return 40
}

func (s *ofp13.GroupFeatures) PackBinary() (data []byte, err error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	binary.Write(buf, binary.BigEndian, s.Types)
	binary.Write(buf, binary.BigEndian, s.Capabilities)
	binary.Write(buf, binary.BigEndian, s.MaxGroups)
	binary.Write(buf, binary.BigEndian, s.Actions)
	data = buf.Bytes()
	return
}

func (s *ofp13.GroupFeatures) UnpackBinary(data []byte) (err error) {
	buf := bytes.NewBuffer(data)
	binary.Read(buf, binary.BigEndian, &s.Types)
	binary.Read(buf, binary.BigEndian, &s.Capabilities)
	binary.Read(buf, binary.BigEndian, &s.MaxGroups)
	binary.Read(buf, binary.BigEndian, &s.Actions)
	return
}

func NewMeterStats() (s *ofp13.MeterStats) {
	s = new(MeterStats)
	return
}

func (s *ofp13.MeterStats) Len() (l int) {
	return 40
}

func (s *ofp13.MeterStats) PackBinary() (data []byte, err error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	binary.Write(buf, binary.BigEndian, s.MeterID)
	binary.Write(buf, binary.BigEndian, s.Length)
	binary.Write(buf, binary.BigEndian, s.pad)
	binary.Write(buf, binary.BigEndian, s.FlowCount)
	binary.Write(buf, binary.BigEndian, s.PacketInCount)
	binary.Write(buf, binary.BigEndian, s.ByteInCount)
	binary.Write(buf, binary.BigEndian, s.DurationSec)
	binary.Write(buf, binary.BigEndian, s.DurationNSec)

	for _, m := range s.BandStats {
		bs := make([]byte, 0)
		bs, err = m.PackBinary()
		if err != nil {
			return
		}
		binary.Write(buf, binary.BigEndian, bs)
	}
	data = buf.Bytes()
	return
}

func (s *ofp13.MeterStats) UnpackBinary(data []byte) (err error) {
	buf := bytes.NewBuffer(data)
	binary.Read(buf, binary.BigEndian, &s.MeterID)
	binary.Read(buf, binary.BigEndian, &s.Length)
	binary.Read(buf, binary.BigEndian, &s.pad)
	binary.Read(buf, binary.BigEndian, &s.FlowCount)
	binary.Read(buf, binary.BigEndian, &s.PacketInCount)
	binary.Read(buf, binary.BigEndian, &s.ByteInCount)
	binary.Read(buf, binary.BigEndian, &s.DurationSec)
	binary.Read(buf, binary.BigEndian, &s.DurationNSec)

	n := s.Len()
	for n < len(data) {
		b := new(MeterBandStats)
		binary.Read(buf, binary.BigEndian, b)
		s.BandStats = append(s.BandStats, *b)
		n += b.Len()
	}
	return
}

func NewMeterBandStats() (s *ofp13.MeterBandStats) {
	s = new(MeterBandStats)
	return
}

func (s *ofp13.MeterBandStats) Len() (l int) {
	return 16
}

func (s *ofp13.MeterBandStats) PackBinary() (data []byte, err error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	binary.Write(buf, binary.BigEndian, s.PacketBandCount)
	binary.Write(buf, binary.BigEndian, s.ByteBandCount)
	data = buf.Bytes()
	return
}

func (s *ofp13.MeterBandStats) UnpackBinary(data []byte) (err error) {
	buf := bytes.NewBuffer(data)
	binary.Read(buf, binary.BigEndian, &s.PacketBandCount)
	binary.Read(buf, binary.BigEndian, &s.ByteBandCount)
	return
}

func NewMeterFeatures() (s *ofp13.MeterFeatures) {
	s = new(MeterFeatures)
	return
}

func (s *ofp13.MeterFeatures) Len() (l int) {
	return 16
}

func (s *ofp13.MeterFeatures) PackBinary() (data []byte, err error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	binary.Write(buf, binary.BigEndian, s.MaxMeter)
	binary.Write(buf, binary.BigEndian, s.BandTypes)
	binary.Write(buf, binary.BigEndian, s.Capabilities)
	binary.Write(buf, binary.BigEndian, s.MaxBands)
	binary.Write(buf, binary.BigEndian, s.MaxColor)
	binary.Write(buf, binary.BigEndian, s.pad)
	data = buf.Bytes()
	return
}

func (s *ofp13.MeterFeatures) UnpackBinary(data []byte) (err error) {
	buf := bytes.NewBuffer(data)
	binary.Read(buf, binary.BigEndian, &s.MaxMeter)
	binary.Read(buf, binary.BigEndian, &s.BandTypes)
	binary.Read(buf, binary.BigEndian, &s.Capabilities)
	binary.Read(buf, binary.BigEndian, &s.MaxBands)
	binary.Read(buf, binary.BigEndian, &s.MaxColor)
	binary.Read(buf, binary.BigEndian, &s.pad)
	return
}

func EncodeTableFeatureProp(t ofp13.TableFeatureProp) (data []byte, err error) {
	data = make([]byte, 0)
	switch t.(type) {
	case *ofp13.TableFeaturePropInstructions:
		data, err = t.(*ofp13.TableFeaturePropInstructions).PackBinary()
	case *ofp13.TableFeaturePropNextTables:
		data, err = t.(*ofp13.TableFeaturePropNextTables).PackBinary()
	case *ofp13.TableFeaturePropActions:
		data, err = t.(*ofp13.TableFeaturePropActions).PackBinary()
	case *ofp13.TableFeaturePropOXM:
		data, err = t.(*ofp13.TableFeaturePropOXM).PackBinary()
	case *ofp13.TableFeaturePropExperimenter:
		data, err = t.(*ofp13.TableFeaturePropExperimenter).PackBinary()
	}
	if err != nil {
		return
	}
	return
}

func DecodeTableFeatureProp(data []byte) (t ofp13.TableFeatureProp, err error) {
	switch binary.BigEndian.Uint16(data[:]) {
	case ofp13.OFPTFPTInstructions:
		t = new(ofp13.TableFeaturePropInstructions)
	case ofp13.OFPTFPTInstructionsMiss:
		t = new(ofp13.TableFeaturePropInstructions)
	case ofp13.OFPTFPTNextTables:
		t = new(ofp13.TableFeaturePropNextTables)
	case ofp13.OFPTFPTNextTablesMiss:
		t = new(ofp13.TableFeaturePropNextTables)
	case ofp13.OFPTFPTWriteActions:
		t = new(ofp13.TableFeaturePropActions)
	case ofp13.OFPTFPTWriteActionsMiss:
		t = new(ofp13.TableFeaturePropActions)
	case ofp13.OFPTFPTApplyActions:
		t = new(ofp13.TableFeaturePropActions)
	case ofp13.OFPTFPTApplyActionsMiss:
		t = new(ofp13.TableFeaturePropActions)
	case ofp13.OFPTFPTMatch:
		t = new(ofp13.TableFeaturePropOXM)
	case ofp13.OFPTFPTWildcards:
		t = new(ofp13.TableFeaturePropOXM)
	case ofp13.OFPTFPTWriteSetField:
		t = new(ofp13.TableFeaturePropOXM)
	case ofp13.OFPTFPTWriteSetFieldMiss:
		t = new(ofp13.TableFeaturePropOXM)
	case ofp13.OFPTFPTApplySetField:
		t = new(ofp13.TableFeaturePropOXM)
	case ofp13.OFPTFPTApplySetFieldMiss:
		t = new(ofp13.TableFeaturePropOXM)
	case ofp13.OFPTFPTExperimenter:
		t = new(ofp13.TableFeaturePropExperimenter)
	case ofp13.OFPTFPTExperimenterMiss:
		t = new(ofp13.TableFeaturePropExperimenter)
	}
	buf := bytes.NewBuffer(data)
	ts := make([]byte, t.Len())
	binary.Read(buf, binary.BigEndian, ts)
	err = t.UnpackBinary(ts)
	if err != nil {
		return
	}
	return
}
