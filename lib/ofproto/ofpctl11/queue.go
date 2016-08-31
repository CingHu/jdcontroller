package ofp11

import (
	"bytes"
	"encoding/binary"
	"errors"
)

func NewPacketQueue() (p *PacketQueue) {
	p = new(PacketQueue)
	return
}

func (p *PacketQueue) Len() (l int) {
	l = 8
	return
}

func (p *PacketQueue) PackBinary() (data []byte, err error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	binary.Write(buf, binary.BigEndian, p.QueueID)
	binary.Write(buf, binary.BigEndian, p.Length)
	binary.Write(buf, binary.BigEndian, p.pad)
	for _, q := range p.Properties {
		bs := make([]byte, 0)
		bs, err = EncodeQueue(q)
		if err != nil {
			return
		}
		binary.Write(buf, binary.BigEndian, bs)
	}
	data = buf.Bytes()
	return
}

func (p *PacketQueue) UnpackBinary(data []byte) (err error) {
	buf := bytes.NewBuffer(data)
	binary.Read(buf, binary.BigEndian, &p.QueueID)
	binary.Read(buf, binary.BigEndian, &p.Length)
	binary.Read(buf, binary.BigEndian, &p.pad)
	n := p.Len()
	for n < len(data) {
		var q QueueProp
		var l int
		switch binary.BigEndian.Uint16(data[n:]) {
		case OFPQTMinRate:
			q = new(QueuePropMinRate)
			r := q.(*QueuePropMinRate)
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

func NewQueuePropHeader() (q *QueuePropHeader) {
	q = new(QueuePropHeader)
	return
}

func (q *QueuePropHeader) Len() (l int) {
	l = 8
	return
}

func (q *QueuePropHeader) PackBinary() (data []byte, err error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	binary.Write(buf, binary.BigEndian, q.Property)
	binary.Write(buf, binary.BigEndian, q.Length)
	binary.Write(buf, binary.BigEndian, q.pad)
	data = buf.Bytes()
	return
}

func (q *QueuePropHeader) UnpackBinary(data []byte) (err error) {
	buf := bytes.NewBuffer(data)
	binary.Read(buf, binary.BigEndian, &q.Property)
	binary.Read(buf, binary.BigEndian, &q.Length)
	binary.Read(buf, binary.BigEndian, &q.pad)
	return
}

func NewQueuePropMinRate() (q *QueuePropMinRate) {
	q = new(QueuePropMinRate)
	q.Header.Property = uint16(OFPQTMinRate)
	q.Header.Length = uint16(16)
	return
}

func (q *QueuePropMinRate) Len() (l int) {
	l = 16
	return
}

func (q *QueuePropMinRate) PackBinary() (data []byte, err error) {
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

func (q *QueuePropMinRate) UnpackBinary(data []byte) (err error) {
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

func NewQueueGetConfigRequest() (q *QueueGetConfigRequest) {
	q = new(QueueGetConfigRequest)
	return
}

func (q *QueueGetConfigRequest) Len() (l int) {
	l = 16
	return
}

func (q *QueueGetConfigRequest) PackBinary() (data []byte, err error) {
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

func (q *QueueGetConfigRequest) UnpackBinary(data []byte) (err error) {
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

func NewQueueGetConfigReply() (q *QueueGetConfigReply) {
	q = new(QueueGetConfigReply)
	return
}

func (q *QueueGetConfigReply) Len() (l int) {
	l = 16
	return
}

func (q *QueueGetConfigReply) PackBinary() (data []byte, err error) {
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

func (q *QueueGetConfigReply) UnpackBinary(data []byte) (err error) {
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
		p := new(PacketQueue)
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

func EncodeQueue(q QueueProp) (data []byte, err error) {
	data = make([]byte, 0)
	switch q.(type) {
	case *QueuePropMinRate:
		data, err = q.(*QueuePropMinRate).PackBinary()
	default:
		err = errors.New("Can not parse this queue request.")
		return
	}
	if err != nil {
		return
	}
	return
}

func DecodeQueue(data []byte, t uint16) (q QueueProp, err error) {
	buf := bytes.NewBuffer(data)
	switch binary.BigEndian.Uint16(data[:]) {
	case OFPQTMinRate:
		q = new(QueuePropMinRate)
	default:
		err = errors.New("Can not parse this stats request.")
		return
	}
	qs := make([]byte, q.Len())
	binary.Read(buf, binary.BigEndian, qs)
	err = q.UnpackBinary(qs)
	if err != nil {
		return
	}
	return
}
