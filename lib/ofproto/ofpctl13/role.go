package ofpctl13

import (
	"bytes"
	"encoding/binary"

	"jd.com/jdcontroller/protocol/ofp13"
)

func NewRoleRequest() (r *ofp13.RoleRequest) {
	r = new(ofp13.RoleRequest)
	return
}

func (r *ofp13.RoleRequest) Len() (l int) {
	l = 24
	return
}

func (r *ofp13.RoleRequest) PackBinary() (data []byte, err error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	hs := make([]byte, 0)
	hs, err = r.Header.PackBinary()
	if err != nil {
		return
	}
	binary.Write(buf, binary.BigEndian, hs)
	binary.Write(buf, binary.BigEndian, r.Role)
	binary.Write(buf, binary.BigEndian, r.pad)
	binary.Write(buf, binary.BigEndian, r.GenerationID)
	data = buf.Bytes()
	return
}

func (r *ofp13.RoleRequest) UnpackBinary(data []byte) (err error) {
	buf := bytes.NewBuffer(data)
	hs := make([]byte, r.Header.Len())
	binary.Read(buf, binary.BigEndian, hs)
	err = r.Header.UnpackBinary(hs)
	if err != nil {
		return
	}
	binary.Read(buf, binary.BigEndian, &r.Role)
	binary.Read(buf, binary.BigEndian, &r.pad)
	binary.Read(buf, binary.BigEndian, &r.GenerationID)
	return
}
