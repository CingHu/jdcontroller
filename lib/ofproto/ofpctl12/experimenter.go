package ofp12

import (
	"bytes"
	"encoding/binary"
)

func NewExperimenterHeader() (h *ExperimenterHeader) {
	h = new(ExperimenterHeader)
	return
}

func (h *ExperimenterHeader) Len() (l int) {
	l = 16
	return
}

func (h *ExperimenterHeader) PackBinary() (data []byte, err error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	hs := make([]byte, 0)
	hs, err = h.Header.PackBinary()
	if err != nil {
		return
	}
	binary.Write(buf, binary.BigEndian, hs)
	binary.Write(buf, binary.BigEndian, h.Experimenter)
	binary.Write(buf, binary.BigEndian, h.ExpType)
	data = buf.Bytes()
	return
}

func (h *ExperimenterHeader) UnpackBinary(data []byte) (err error) {
	buf := bytes.NewBuffer(data)
	hs := make([]byte, h.Header.Len())
	binary.Read(buf, binary.BigEndian, hs)
	err = h.Header.UnpackBinary(hs)
	if err != nil {
		return
	}
	binary.Read(buf, binary.BigEndian, &h.Experimenter)
	binary.Read(buf, binary.BigEndian, &h.ExpType)
	return
}
