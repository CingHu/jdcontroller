package ofpctl13

import (
	"fmt"
	"bytes"
	"encoding/binary"
	
	"jd.com/jdcontroller/lib/buffer"
	"jd.com/jdcontroller/protocol/ofp13"
)

var DPMultipartInfo map[uint64]*MultipartInfo

func NewMultipartInfo() {
	DPMultipartInfo = make(map[uint64]*ofp13.MultipartInfo)
	return
}

func GetMultipartInfo(dpid uint64) (info *ofp13.MultipartInfo) {
	info = DPMultipartInfo[dpid]
	return
}

func NewMultipartPortStatsRequest(p uint32) (m *MultipartRequest) {
	m = NewMultipartRequest()
	m.Type = uint16(OFPMPPortStats)
	portStats := NewPortStatsRequest(p)
	m.Body = append(m.Body, portStats)
	m.Header.Length += uint16(portStats.Len())
	//m.Flags = uint16(OFPMPFRequestMore)
	return
}

func NewMultipartPortDescRequest() (m *MultipartRequest) {
	m = NewMultipartRequest()
	m.Type = uint16(OFPMPPortDesc)
	//	m.Flags = uint16(OFPMPFRequestMore)
	return
}

func NewMultipartFlowStatsRequest() (m *MultipartRequest) {
	m = NewMultipartRequest()
	m.Type = uint16(OFPMPFlow)
	FlowStats := NewFlowStatsRequest()
	FlowStats.TableID = OFPTTAll
	FlowStats.OutPort = OFPPAny
	FlowStats.OutGroup = OFPGAny
	FlowStats.Match.Type = OFPMTOXM
	FlowStats.Match.Length = 4
	m.Body = append(m.Body, FlowStats)
	m.Header.Length += uint16(FlowStats.Len())
	//m.Flags = uint16(OFPMPFRequestMore)
	return
}

func NewMultipartAggregateStatsRequest() (m *MultipartRequest) {
	m = NewMultipartRequest()
	m.Type = uint16(OFPMPAggregate)
	AggregateStats := NewAggregateStatsRequest()
	AggregateStats.TableID = OFPTTAll
	AggregateStats.OutPort = OFPPAny
	AggregateStats.OutGroup = OFPGAny
	AggregateStats.Match.Type = OFPMTOXM
	AggregateStats.Match.Length = 4
	m.Body = append(m.Body, AggregateStats)
	m.Header.Length += uint16(AggregateStats.Len())
	//m.Flags = uint16(OFPMPFRequestMore)
	return
}

func NewMultipartRequest() (m *MultipartRequest) {
	m = new(MultipartRequest)
	m.Header = *NewHeader()
	m.Header.Type = uint8(OFPTMultipartRequest)
	m.Header.Length = uint16(m.Len())
	return
}

func (m *MultipartRequest) Len() (l int) {
	l = 16
	return
}

func (m *MultipartRequest) PackBinary() (data []byte, err error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	hs := make([]byte, 0)
	hs, err = m.Header.PackBinary()
	if err != nil {
		return
	}
	binary.Write(buf, binary.BigEndian, hs)
	binary.Write(buf, binary.BigEndian, m.Type)
	binary.Write(buf, binary.BigEndian, m.Flags)
	binary.Write(buf, binary.BigEndian, m.pad)

	if m.Body != nil {
		bs := make([]byte, 0)
		bs, err = m.EncodeMultipartBody(m.Type, m.Body)
		if err != nil {
			return
		}
		binary.Write(buf, binary.BigEndian, bs)
	}
	data = buf.Bytes()
	return
}

func (m *MultipartRequest) UnpackBinary(data []byte) (err error) {
	buf := bytes.NewBuffer(data)
	hs := make([]byte, m.Header.Len())
	binary.Read(buf, binary.BigEndian, hs)
	err = m.Header.UnpackBinary(hs)
	if err != nil {
		return
	}

	binary.Read(buf, binary.BigEndian, &m.Type)
	binary.Read(buf, binary.BigEndian, &m.Flags)
	binary.Read(buf, binary.BigEndian, &m.pad)

	n := m.Len()
	if n < len(data) {
		m.Body, err = m.DecodeMultipartBody(m.Type, data[n:])
		if err != nil {
			return
		}
	}
	return
}

func NewMultipartReply() (m *MultipartReply) {
	m = new(MultipartReply)
	m.Header.Type = uint8(OFPTMultipartReply)
	return
}

func (m *MultipartReply) Len() (l int) {
	l = 16
	return
}

func ParseMultipartReply(dpid uint64, pkt *MultipartReply) {
	Info, ok := DPMultipartInfo[dpid]
	if !ok {
		DPMultipartInfo[dpid] = new(ofp13.MultipartInfo)
		Info = DPMultipartInfo[dpid]
		Info.PortStats = make(map[uint32]*PortStats)
	}
	switch pkt.Type {
	case OFPMPDesc:
		fmt.Println("OFPMPDesc")
	case OFPMPFlow:
		fmt.Println("OFPMPFlow")
	case OFPMPAggregate:
		fmt.Println("OFPMPAggregate")
	case OFPMPTable:
		fmt.Println("OFPMPTable")
	case OFPMPPortStats:
		for _, v := range pkt.Body {
			PortNo := v.(*PortStats).PortNO
			Info.PortStats[PortNo] = v.(*PortStats)
		}
		for _, v := range DPMultipartInfo[dpid].PortStats {
			fmt.Println("portstats ", v)
		}
	case OFPMPQueue:
		fmt.Println("OFPMPQueue")
	case OFPMPGroup:
		fmt.Println("OFPMPGroup")
	case OFPMPGroupDesc:
		fmt.Println("OFPMPGroupDesc")
	case OFPMPGroupFeatures:
		fmt.Println("OFPMPGroupFeatures")
	case OFPMPMeter:
		fmt.Println("OFPMPMeter")
	case OFPMPMeterConfig:
		fmt.Println("OFPMPMeterConfig")
	case OFPMPMeterFeatures:
		fmt.Println("OFPMPMeterFeatures")
	case OFPMPTableFeatures:
		fmt.Println("OFPMPTableFeatures")
	case OFPMPPortDesc:
		fmt.Println("OFPMPPortDesc")
	case OFPMPExperimenter:
		fmt.Println("OFPMPExperimenter")
	}
}

func (m *MultipartReply) PackBinary() (data []byte, err error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	hs := make([]byte, 0)
	hs, err = m.Header.PackBinary()
	if err != nil {
		return
	}
	binary.Write(buf, binary.BigEndian, hs)
	binary.Write(buf, binary.BigEndian, m.Type)
	binary.Write(buf, binary.BigEndian, m.Flags)
	binary.Write(buf, binary.BigEndian, m.pad)

	if m.Body != nil {
		bs := make([]byte, 0)
		bs, err = m.EncodeMultipartBody(m.Type, m.Body)
		if err != nil {
			return
		}
		binary.Write(buf, binary.BigEndian, bs)
	}

	data = buf.Bytes()
	return
}

func (m *MultipartReply) UnpackBinary(data []byte) (err error) {
	buf := bytes.NewBuffer(data)
	hs := make([]byte, m.Header.Len())
	binary.Read(buf, binary.BigEndian, hs)
	err = m.Header.UnpackBinary(hs)
	if err != nil {
		return
	}

	binary.Read(buf, binary.BigEndian, &m.Type)
	binary.Read(buf, binary.BigEndian, &m.Flags)
	binary.Read(buf, binary.BigEndian, &m.pad)

	n := m.Len()
	if n < len(data) {
		m.Body, err = m.DecodeMultipartBody(m.Type, data[n:])
		if err != nil {
			return
		}
	}
	return
}

func NewExperimenterMultipartHeader() (e *ExperimenterMultipartHeader) {
	e = new(ExperimenterMultipartHeader)
	return
}

func (e *ExperimenterMultipartHeader) Len() (l int) {
	l = 8
	return
}

func (e *ExperimenterMultipartHeader) PackBinary() (data []byte, err error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	binary.Write(buf, binary.BigEndian, e.Experimenter)
	binary.Write(buf, binary.BigEndian, e.ExpType)
	data = buf.Bytes()
	return
}

func (e *ExperimenterMultipartHeader) UnpackBinary(data []byte) (err error) {
	buf := bytes.NewBuffer(data)
	binary.Read(buf, binary.BigEndian, &e.Experimenter)
	binary.Read(buf, binary.BigEndian, &e.ExpType)
	return
}

func (m *MultipartRequest) EncodeMultipartBody(t uint16, body []buffer.Message) (data []byte, err error) {
	for _, b := range body {
		var i buffer.Message
		switch t {
		case OFPMPDesc:
			//b = new(Desc)
		case OFPMPFlow:
			i = b.(*FlowStatsRequest)
		case OFPMPAggregate:
			i = b.(*AggregateStatsRequest)
		case OFPMPTable:
			//i = b.(*TableStats)
		case OFPMPPortStats:
			i = b.(*PortStatsRequest)
		case OFPMPQueue:
			i = b.(*QueueStatsRequest)
		case OFPMPGroup:
			i = b.(*GroupStatsRequest)
		case OFPMPGroupDesc:
			//i = b.(*GroupDescStats)
		case OFPMPGroupFeatures:
			//i = b.(*GroupFeatures)
		case OFPMPMeter:
			//i = b.(*MeterStats)
		case OFPMPMeterConfig:
			//i = b.(*MeterConfig)
		case OFPMPMeterFeatures:
			//i = b.(*MeterFeatures)
		case OFPMPTableFeatures:
			//i = b.(*TableFeatures)
		case OFPMPPortDesc:
			//i = b.(*Port)
		case OFPMPExperimenter:
		}
		bs := make([]byte, 0)
		bs, err = i.PackBinary()
		if err != nil {
			return
		}

		data = append(data, bs...)
	}
	return
}

func (m *MultipartRequest) DecodeMultipartBody(t uint16, data []byte) (body []buffer.Message, err error) {
	var b buffer.Message
	switch t {
	case OFPMPDesc:
		//b = new(Desc)
	case OFPMPFlow:
		b = new(FlowStatsRequest)
	case OFPMPAggregate:
	case OFPMPTable:
		//b = new(TableStats)
	case OFPMPPortStats:
		b = new(PortStatsRequest)
	case OFPMPQueue:
		b = new(QueueStatsRequest)
	case OFPMPGroup:
		b = new(GroupStatsRequest)
	case OFPMPGroupDesc:
		//b = new(GroupDescStats)
	case OFPMPGroupFeatures:
		//b = new(GroupFeatures)
	case OFPMPMeter:
		//b = new(MeterStats)
	case OFPMPMeterConfig:
		//b = new(MeterConfig)
	case OFPMPMeterFeatures:
		//b = new(MeterFeatures)
	case OFPMPTableFeatures:
		//b = new(TableFeatures)
	case OFPMPPortDesc:
		//b = new(Port)
	case OFPMPExperimenter:
	}
	n := 0
	for n < len(data) {
		i := b
		buf := bytes.NewBuffer(data)
		bs := make([]byte, i.Len())
		binary.Read(buf, binary.BigEndian, bs)
		err = i.UnpackBinary(bs)
		if err != nil {
			return
		}
		body = append(body, i)
		n += i.Len()
	}
	return
}

func (m *MultipartReply) EncodeMultipartBody(t uint16, body []buffer.Message) (data []byte, err error) {
	for _, b := range body {
		var i buffer.Message
		switch t {
		case OFPMPDesc:
			i = b.(*Desc)
		case OFPMPFlow:
			i = b.(*FlowStats)
		case OFPMPAggregate:
		case OFPMPTable:
			i = b.(*TableStats)
		case OFPMPPortStats:
			i = b.(*PortStats)
		case OFPMPQueue:
			i = b.(*QueueStats)
		case OFPMPGroup:
			i = b.(*GroupStats)
		case OFPMPGroupDesc:
			i = b.(*GroupDescStats)
		case OFPMPGroupFeatures:
			i = b.(*GroupFeatures)
		case OFPMPMeter:
			i = b.(*MeterStats)
		case OFPMPMeterConfig:
			i = b.(*MeterConfig)
		case OFPMPMeterFeatures:
			i = b.(*MeterFeatures)
		case OFPMPTableFeatures:
			i = b.(*TableFeatures)
		case OFPMPPortDesc:
			i = b.(*Port)
		case OFPMPExperimenter:
		}
		bs := make([]byte, 0)
		bs, err = i.PackBinary()
		if err != nil {
			return
		}
		data = append(data, bs...)
	}
	return
}

func (m *MultipartReply) DecodeMultipartBody(t uint16, data []byte) (body []buffer.Message, err error) {
	//var b buffer.Message
	switch t {
	case OFPMPDesc:
		//b = new(Desc)
	case OFPMPFlow:
		//b = new(FlowStats)
	case OFPMPAggregate:
	case OFPMPTable:
		//b = new(TableStats)
	case OFPMPPortStats:
		//b = new(PortStats)
	case OFPMPQueue:
		//b = new(QueueStats)
	case OFPMPGroup:
		//b = new(GroupStats)
	case OFPMPGroupDesc:
		//b = new(GroupDescStats)
	case OFPMPGroupFeatures:
		//b = new(GroupFeatures)
	case OFPMPMeter:
		//b = new(MeterStats)
	case OFPMPMeterConfig:
		//b = new(MeterConfig)
	case OFPMPMeterFeatures:
		//b = new(MeterFeatures)
	case OFPMPTableFeatures:
		//b = new(TableFeatures)
	case OFPMPPortDesc:
		//b = new(Port)
	case OFPMPExperimenter:
	}
	n := 0
	for n < len(data) {
		i := new(PortStats)
		buf := bytes.NewBuffer(data[n:])
		bs := make([]byte, i.Len())
		binary.Read(buf, binary.BigEndian, bs)
		err = i.UnpackBinary(bs)
		if err != nil {
			return
		}
		body = append(body, i)
		n += i.Len()
	}
	return
}
