package ofpctl10

import (
	"bytes"
	"encoding/binary"
	
	"jd.com/jdcontroller/protocol/ofp10"
)

func NewPacketQueue() (p *ofp10.PacketQueue) {
	p = new(ofp10.PacketQueue)
	return
}

func (p *ofp10.PacketQueue) Len() (l int) {
	l = 8
	return
}

func (p *ofp10.PacketQueue) PackBinary() (data []byte, err error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	binary.Write(buf, binary.BigEndian, p.QueueID)
	binary.Write(buf, binary.BigEndian, p.Length)
	binary.Write(buf, binary.BigEndian, p.pad)
	for _, q := range p.Properties {
		bs := make([]byte, 0)
		switch q.(type) {
		case *ofp10.QueuePropMinRate:
			bs, err = q.(*ofp10.QueuePropMinRate).PackBinary()
		}
		if err != nil {
			return
		}
		binary.Write(buf, binary.BigEndian, bs)
	}
	data = buf.Bytes()
	return
}

func (p *ofp10.PacketQueue) UnpackBinary(data []byte) (err error) {
	buf := bytes.NewBuffer(data)
	binary.Read(buf, binary.BigEndian, &p.QueueID)
	binary.Read(buf, binary.BigEndian, &p.Length)
	binary.Read(buf, binary.BigEndian, &p.pad)
	n := p.Len()
	for n < len(data) {
		var q ofp10.QueueProp
		var l int
		switch binary.BigEndian.Uint16(data[n:]) {
		case ofp10.OFPQTMinRate:
			q = new(ofp10.QueuePropMinRate)
			r := q.(*ofp10.QueuePropMinRate)
			qs := make([]byte, r.Len())
			binary.Read(buf, binary.BigEndian, qs)
			r.UnpackBinary(qs)
			l = r.Len()
		}
		p.Properties = append(p.Properties, q)
		n += l
	}
	return
}

func NewQueuePropHeader() (q *ofp10.QueuePropHeader) {
	q = new(ofp10.QueuePropHeader)
	return
}

func (q *ofp10.QueuePropHeader) Len() (l int) {
	l = 8
	return
}

func (q *ofp10.QueuePropHeader) PackBinary() (data []byte, err error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	binary.Write(buf, binary.BigEndian, q.Property)
	binary.Write(buf, binary.BigEndian, q.Length)
	binary.Write(buf, binary.BigEndian, q.pad)
	data = buf.Bytes()
	return
}

func (q *ofp10.QueuePropHeader) UnpackBinary(data []byte) (err error) {
	buf := bytes.NewBuffer(data)
	binary.Read(buf, binary.BigEndian, &q.Property)
	binary.Read(buf, binary.BigEndian, &q.Length)
	binary.Read(buf, binary.BigEndian, &q.pad)
	return
}

func NewQueuePropMinRate() (q *ofp10.QueuePropMinRate) {
	q = new(ofp10.QueuePropMinRate)
	q.Header.Property = uint16(ofp10.OFPQTMinRate)
	q.Header.Length = uint16(16)
	return
}

func (q *ofp10.QueuePropMinRate) Len() (l int) {
	l = 16
	return
}

func (q *ofp10.QueuePropMinRate) PackBinary() (data []byte, err error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	hs := make([]byte, 0)
	hs, err = q.Header.PackBinary()
	if err != nil {
		return
	}
	binary.Write(buf, binary.BigEndian, hs)
	binary.Write(buf, binary.BigEndian, q.Rate)
	binary.Write(buf, binary.BigEndian, q.pad)
	data = buf.Bytes()
	return
}

func (q *ofp10.QueuePropMinRate) UnpackBinary(data []byte) (err error) {
	buf := bytes.NewBuffer(data)
	hs := make([]byte, q.Header.Len())
	binary.Read(buf, binary.BigEndian, hs)
	err = q.Header.UnpackBinary(hs)
	if err != nil {
		return
	}
	binary.Read(buf, binary.BigEndian, &q.Rate)
	binary.Read(buf, binary.BigEndian, &q.pad)
	return
}

func NewQueueGetConfigRequest() (q *ofp10.QueueGetConfigRequest) {
	q = new(ofp10.QueueGetConfigRequest)
	return
}

func (q *ofp10.QueueGetConfigRequest) Len() (l int) {
	l = 16
	return
}

func (q *ofp10.QueueGetConfigRequest) PackBinary() (data []byte, err error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	hs := make([]byte, 0)
	hs, err = q.Header.PackBinary()
	if err != nil {
		return
	}
	binary.Write(buf, binary.BigEndian, hs)
	binary.Write(buf, binary.BigEndian, q.Port)
	binary.Write(buf, binary.BigEndian, q.pad)
	data = buf.Bytes()
	return
}

func (q *ofp10.QueueGetConfigRequest) UnpackBinary(data []byte) (err error) {
	buf := bytes.NewBuffer(data)
	hs := make([]byte, q.Header.Len())
	binary.Read(buf, binary.BigEndian, hs)
	err = q.Header.UnpackBinary(hs)
	if err != nil {
		return
	}
	binary.Read(buf, binary.BigEndian, &q.Port)
	binary.Read(buf, binary.BigEndian, &q.pad)
	return
}

func NewQueueGetConfigReply() (q *ofp10.QueueGetConfigReply) {
	q = new(ofp10.QueueGetConfigReply)
	return
}

func (q *ofp10.QueueGetConfigReply) Len() (l int) {
	l = 16
	return
}

func (q *ofp10.QueueGetConfigReply) PackBinary() (data []byte, err error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	hs := make([]byte, 0)
	hs, err = q.Header.PackBinary()
	if err != nil {
		return
	}
	binary.Write(buf, binary.BigEndian, hs)
	binary.Write(buf, binary.BigEndian, q.Port)
	binary.Write(buf, binary.BigEndian, q.pad)
	for _, p := range q.Queues {
		bs := make([]byte, 0)
		bs, err = p.PackBinary()
		if err != nil {
			return
		}
		binary.Write(buf, binary.BigEndian, bs)
	}
	data = buf.Bytes()
	return
}

func (q *ofp10.QueueGetConfigReply) UnpackBinary(data []byte) (err error) {
	buf := bytes.NewBuffer(data)
	hs := make([]byte, q.Header.Len())
	binary.Read(buf, binary.BigEndian, hs)
	err = q.Header.UnpackBinary(hs)
	if err != nil {
		return
	}
	binary.Read(buf, binary.BigEndian, &q.Port)
	binary.Read(buf, binary.BigEndian, &q.pad)
	n := q.Len()
	for n < len(data) {
		p := new(ofp10.PacketQueue)
		ps := make([]byte, q.Header.Len())
		binary.Read(buf, binary.BigEndian, ps)
		err = p.UnpackBinary(ps)
		if err != nil {
			return
		}
		q.Queues = append(q.Queues, *p)
		n += p.Len()
	}
	return
}
