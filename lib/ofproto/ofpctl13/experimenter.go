package ofpctl13

import (
	"bytes"
	"encoding/binary"
	
	"jd.com/jdcontroller/protocol/ofp13"
)

func NewExperimenterHeader() (h *ofp13.ExperimenterHeader) {
	h = new(ofp13.ExperimenterHeader)
	return
}

func (h *ofp13.ExperimenterHeader) Len() (l int) {
	l = 16
	return
}

func (h *ofp13.ExperimenterHeader) PackBinary() (data []byte, err error) {
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

func (h *ofp13.ExperimenterHeader) UnpackBinary(data []byte) (err error) {
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
