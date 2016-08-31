package ofp11

import (
	"bytes"
	"encoding/binary"
)

func NewExperimenterHeader() (e *ExperimenterHeader) {
	e = new(ExperimenterHeader)
	return
}

func (e *ExperimenterHeader) Len() (l int) {
	l = 16
	return
}

func (e *ExperimenterHeader) PackBinary() (data []byte, err error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	hs := make([]byte, 0)
	hs, err = e.Header.PackBinary()
	if err != nil {
		return
	}
	binary.Write(buf, binary.BigEndian, hs)
	binary.Write(buf, binary.BigEndian, e.Experimenter)
	binary.Write(buf, binary.BigEndian, e.pad)

	data = buf.Bytes()
	return
}

func (e *ExperimenterHeader) UnpackBinary(data []byte) (err error) {
	buf := bytes.NewBuffer(data)
	hs := make([]byte, e.Header.Len())
	binary.Read(buf, binary.BigEndian, hs)
	err = e.Header.UnpackBinary(hs)
	if err != nil {
		return
	}
	binary.Read(buf, binary.BigEndian, &e.Experimenter)
	binary.Read(buf, binary.BigEndian, &e.pad)
	return
}
