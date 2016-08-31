package ofpctl13

import (
	"bytes"
	"encoding/binary"
	"errors"

	"jd.com/jdcontroller/protocol/ofp13"
)

func NewMeterMultipartRequest() (m *ofp13.MeterMultipartRequest) {
	m = new(ofp13.MeterMultipartRequest)
	return
}

func (m *ofp13.MeterMultipartRequest) Len() (l int) {
	l = 8
	return
}

func (m *ofp13.MeterMultipartRequest) PackBinary() (data []byte, err error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	binary.Write(buf, binary.BigEndian, m.MeterID)
	binary.Write(buf, binary.BigEndian, m.pad)
	data = buf.Bytes()
	return
}

func (m *ofp13.MeterMultipartRequest) UnpackBinary(data []byte) (err error) {
	buf := bytes.NewBuffer(data)
	binary.Read(buf, binary.BigEndian, &m.MeterID)
	binary.Read(buf, binary.BigEndian, &m.pad)
	return
}

func NewMeterConfig() (c *MeterConfig) {
	c = new(MeterConfig)
	return
}

func (c *MeterConfig) Len() (l int) {
	l = 8
	return
}

func (c *MeterConfig) PackBinary() (data []byte, err error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	binary.Write(buf, binary.BigEndian, c.Length)
	binary.Write(buf, binary.BigEndian, c.Flags)
	binary.Write(buf, binary.BigEndian, c.MeterID)
	for _, m := range c.Bands {
		bs := make([]byte, 0)
		bs, err = EncodeMeterBand(m)
		if err != nil {
			return
		}
		binary.Write(buf, binary.BigEndian, bs)
	}

	data = buf.Bytes()
	return
}

func (c *MeterConfig) UnpackBinary(data []byte) (err error) {
	buf := bytes.NewBuffer(data)
	binary.Read(buf, binary.BigEndian, &c.Length)
	binary.Read(buf, binary.BigEndian, &c.Flags)
	binary.Read(buf, binary.BigEndian, &c.MeterID)
	n := c.Len()
	for n < len(data) {
		var m MeterBand
		m, err = DecodeMeterBand(data[n:])
		if err != nil {
			return
		}
		c.Bands = append(c.Bands, m)
	}
	return
}

func NewMeterBandHeader() (m *MeterBandHeader) {
	m = new(MeterBandHeader)
	return
}

func (m *MeterBandHeader) Len() (l int) {
	l = 12
	return
}

func (m *MeterBandHeader) PackBinary() (data []byte, err error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	binary.Write(buf, binary.BigEndian, m.Type)
	binary.Write(buf, binary.BigEndian, m.Length)
	binary.Write(buf, binary.BigEndian, m.Rate)
	binary.Write(buf, binary.BigEndian, m.BurstSize)
	data = buf.Bytes()
	return
}

func (m *MeterBandHeader) UnpackBinary(data []byte) (err error) {
	buf := bytes.NewBuffer(data)
	binary.Read(buf, binary.BigEndian, &m.Type)
	binary.Read(buf, binary.BigEndian, &m.Length)
	binary.Read(buf, binary.BigEndian, &m.Rate)
	binary.Read(buf, binary.BigEndian, &m.BurstSize)
	return
}

func NewMeterBandDrop() (m *MeterBandDrop) {
	m = new(MeterBandDrop)
	return
}

func (m *MeterBandDrop) Len() (l int) {
	l = 16
	return
}

func (m *MeterBandDrop) PackBinary() (data []byte, err error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	hs := make([]byte, 0)
	hs, err = m.Header.PackBinary()
	if err != nil {
		return
	}
	binary.Write(buf, binary.BigEndian, hs)
	binary.Write(buf, binary.BigEndian, m.pad)
	data = buf.Bytes()
	return
}

func (m *MeterBandDrop) UnpackBinary(data []byte) (err error) {
	buf := bytes.NewBuffer(data)
	hs := make([]byte, m.Header.Len())
	binary.Read(buf, binary.BigEndian, hs)
	err = m.Header.UnpackBinary(hs)
	if err != nil {
		return
	}
	binary.Read(buf, binary.BigEndian, &m.pad)
	return
}

func NewMeterBandDscpRemark() (m *MeterBandDscpRemark) {
	m = new(MeterBandDscpRemark)
	return
}

func (m *MeterBandDscpRemark) Len() (l int) {
	l = 16
	return
}

func (m *MeterBandDscpRemark) PackBinary() (data []byte, err error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	hs := make([]byte, 0)
	hs, err = m.Header.PackBinary()
	if err != nil {
		return
	}
	binary.Write(buf, binary.BigEndian, hs)
	binary.Write(buf, binary.BigEndian, m.PrecLevel)
	binary.Write(buf, binary.BigEndian, m.pad)
	data = buf.Bytes()
	return
}

func (m *MeterBandDscpRemark) UnpackBinary(data []byte) (err error) {
	buf := bytes.NewBuffer(data)
	hs := make([]byte, m.Header.Len())
	binary.Read(buf, binary.BigEndian, hs)
	err = m.Header.UnpackBinary(hs)
	if err != nil {
		return
	}
	binary.Read(buf, binary.BigEndian, &m.PrecLevel)
	binary.Read(buf, binary.BigEndian, &m.pad)
	return
}

func NewMeterBandExperimenter() (m *MeterBandExperimenter) {
	m = new(MeterBandExperimenter)
	return
}

func (m *MeterBandExperimenter) Len() (l int) {
	l = 16
	return
}

func (m *MeterBandExperimenter) PackBinary() (data []byte, err error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	hs := make([]byte, 0)
	hs, err = m.Header.PackBinary()
	if err != nil {
		return
	}
	binary.Write(buf, binary.BigEndian, hs)
	binary.Write(buf, binary.BigEndian, m.Experimenter)
	data = buf.Bytes()
	return
}

func (m *MeterBandExperimenter) UnpackBinary(data []byte) (err error) {
	buf := bytes.NewBuffer(data)
	hs := make([]byte, m.Header.Len())
	binary.Read(buf, binary.BigEndian, hs)
	err = m.Header.UnpackBinary(hs)
	if err != nil {
		return
	}
	binary.Read(buf, binary.BigEndian, &m.Experimenter)
	return
}

func EncodeMeterBand(m MeterBand) (data []byte, err error) {
	data = make([]byte, 0)
	switch m.(type) {
	case *MeterBandDrop:
		data, err = m.(*MeterBandDrop).PackBinary()
	case *MeterBandDscpRemark:
		data, err = m.(*MeterBandDscpRemark).PackBinary()
	case *MeterBandExperimenter:
		data, err = m.(*MeterBandExperimenter).PackBinary()
	default:
		err = errors.New("Can not parse this meter band request.")
		return
	}
	if err != nil {
		return
	}
	return
}

func DecodeMeterBand(data []byte) (m MeterBand, err error) {
	buf := bytes.NewBuffer(data)
	switch binary.BigEndian.Uint16(data[:]) {
	case OFPMBTDrop:
		m = new(MeterBandDrop)
	case OFPMBTDscpRemark:
		m = new(MeterBandDscpRemark)
	case OFPMBTExperimenter:
		m = new(MeterBandExperimenter)
	default:
		err = errors.New("Can not parse this meter band request.")
		return
	}
	ms := make([]byte, m.Len())
	binary.Read(buf, binary.BigEndian, ms)
	err = m.UnpackBinary(ms)
	if err != nil {
		return
	}
	return
}
