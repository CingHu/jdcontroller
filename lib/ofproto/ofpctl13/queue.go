package ofpctl13

import (
	"bytes"
	"encoding/binary"
	"errors"

	"jd.com/jdcontroller/protocol/ofp13"
)

func NewPacketQueue() (p *ofp13.PacketQueue) {
	p = new(ofp13.PacketQueue)
	return
}

func (p *ofp13.PacketQueue) Len() (l int) {
	l = 16
	return
}

func (p *ofp13.PacketQueue) PackBinary() (data []byte, err error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	binary.Write(buf, binary.BigEndian, p.QueueID)
	binary.Write(buf, binary.BigEndian, p.Port)
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

func (p *ofp13.PacketQueue) UnpackBinary(data []byte) (err error) {
	buf := bytes.NewBuffer(data)
	binary.Read(buf, binary.BigEndian, &p.QueueID)
	binary.Read(buf, binary.BigEndian, &p.Port)
	binary.Read(buf, binary.BigEndian, &p.Length)
	binary.Read(buf, binary.BigEndian, &p.pad)
	n := p.Len()
	for n < len(data) {
		var q ofp13.QueueProp
		q, err = DecodeQueue(data[n:])
		if err != nil {
			return
		}
		p.Properties = append(p.Properties, q)
		n += q.Len()
	}
	return
}

func NewQueuePropHeader() (q *ofp13.QueuePropHeader) {
	q = new(ofp13.QueuePropHeader)
	return
}

func (q *ofp13.QueuePropHeader) Len() (l int) {
	l = 8
	return
}

func (q *ofp13.QueuePropHeader) PackBinary() (data []byte, err error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	binary.Write(buf, binary.BigEndian, q.Property)
	binary.Write(buf, binary.BigEndian, q.Length)
	binary.Write(buf, binary.BigEndian, q.pad)
	data = buf.Bytes()
	return
}

func (q *ofp13.QueuePropHeader) UnpackBinary(data []byte) (err error) {
	buf := bytes.NewBuffer(data)
	binary.Read(buf, binary.BigEndian, &q.Property)
	binary.Read(buf, binary.BigEndian, &q.Length)
	binary.Read(buf, binary.BigEndian, &q.pad)
	return
}

func NewQueuePropMinRate() (q *ofp13.QueuePropMinRate) {
	q = new(ofp13.QueuePropMinRate)
	q.Header.Property = uint16(OFPQTMinRate)
	q.Header.Length = uint16(16)
	return
}

func (q *ofp13.QueuePropMinRate) Len() (l int) {
	l = 16
	return
}

func (q *ofp13.QueuePropMinRate) PackBinary() (data []byte, err error) {
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

func (q *ofp13.QueuePropMinRate) UnpackBinary(data []byte) (err error) {
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

func NewQueuePropMaxRate() (q *ofp13.QueuePropMaxRate) {
	q = new(ofp13.QueuePropMaxRate)
	q.Header.Property = uint16(OFPQTMaxRate)
	q.Header.Length = uint16(16)
	return
}

func (q *ofp13.QueuePropMaxRate) Len() (l int) {
	l = 16
	return
}

func (q *ofp13.QueuePropMaxRate) PackBinary() (data []byte, err error) {
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

func (q *ofp13.QueuePropMaxRate) UnpackBinary(data []byte) (err error) {
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

func NewQueuePropExperimenter() (q *ofp13.QueuePropExperimenter) {
	q = new(ofp13.QueuePropExperimenter)
	q.Header.Property = uint16(OFPQTExperimenter)
	q.Header.Length = uint16(16)
	return
}

func (q *ofp13.QueuePropExperimenter) Len() (l int) {
	l = 16
	return
}

func (q *ofp13.QueuePropExperimenter) PackBinary() (data []byte, err error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	hs := make([]byte, 0)
	hs, err = q.Header.PackBinary()
	if err != nil {
		return
	}
	binary.Write(buf, binary.BigEndian, hs)
	binary.Write(buf, binary.BigEndian, q.Experimenter)
	binary.Write(buf, binary.BigEndian, q.pad)
	for _, b := range q.Data {
		binary.Write(buf, binary.BigEndian, b)
	}
	data = buf.Bytes()
	return
}

func (q *ofp13.QueuePropExperimenter) UnpackBinary(data []byte) (err error) {
	buf := bytes.NewBuffer(data)
	hs := make([]byte, q.Header.Len())
	binary.Read(buf, binary.BigEndian, hs)
	err = q.Header.UnpackBinary(hs)
	if err != nil {
		return
	}
	binary.Read(buf, binary.BigEndian, &q.Experimenter)
	binary.Read(buf, binary.BigEndian, &q.pad)
	n := q.Len()
	for n < len(data) {
		b := new(uint8)
		binary.Read(buf, binary.BigEndian, b)
		q.Data = append(q.Data, *b)
		n += 1
	}
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
		p := new(ofp13.PacketQueue)
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

func EncodeQueue(q ofp13.QueueProp) (data []byte, err error) {
	data = make([]byte, 0)
	switch q.(type) {
	case *ofp13.QueuePropMinRate:
		data, err = q.(*ofp13.QueuePropMinRate).PackBinary()
	case *ofp13.QueuePropMaxRate:
		data, err = q.(*ofp13.QueuePropMaxRate).PackBinary()
	default:
		err = errors.New("Can not parse this queue request.")
		return
	}
	if err != nil {
		return
	}
	return
}

func DecodeQueue(data []byte) (q ofp13.QueueProp, err error) {
	buf := bytes.NewBuffer(data)
	switch binary.BigEndian.Uint16(data[:]) {
	case OFPQTMinRate:
		q = new(ofp13.QueuePropMinRate)
	case OFPQTMaxRate:
		q = new(ofp13.QueuePropMaxRate)
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
