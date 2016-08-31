package ofp12

import (
	"bytes"
	"encoding/binary"
)

func NewRoleRequest() (r *RoleRequest) {
	r = new(RoleRequest)
	return
}

func (r *RoleRequest) Len() (l int) {
	l = 24
	return
}

func (r *RoleRequest) PackBinary() (data []byte, err error) {
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

func (r *RoleRequest) UnpackBinary(data []byte) (err error) {
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
