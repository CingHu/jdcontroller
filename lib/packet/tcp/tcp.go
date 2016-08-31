package tcp
//
//import (
//	"bytes"
//	"encoding/binary"
//	"io"
//)
//
//type TCP struct {
//	PortSrc    uint16
//	PortDst    uint16
//	SeqNum     uint32
//	AckNum     uint32
//	DataOffset uint8
//	FIN        bool
//	SYN        bool
//	RST        bool
//	PSH        bool
//	ACK        bool
//	URG        bool
//	WinSize    uint16
//	Checksum   uint16
//	UrgFlag    uint16
//
//	Data []byte
//}
//
//func (t *TCP) Len() (l int) {
//	if t.Data != nil {
//		return uint16(8 + len(t.Data))
//	}
//	return uint16(8)
//}
//
//func (t *TCP) Read(b []byte) (n int, err error) {
//	buf := new(bytes.Buffer)
//	binary.Write(buf, binary.BigEndian, t.PortSrc)
//	binary.Write(buf, binary.BigEndian, t.PortDst)
//	binary.Write(buf, binary.BigEndian, t.Length)
//	binary.Write(buf, binary.BigEndian, t.Checksum)
//	binary.Write(buf, binary.BigEndian, t.Data)
//	if n, err = buf.Read(b); n == 0 {
//		return
//	}
//	return n, io.EOF
//}
//
//func (t *TCP) ReadFrom(r io.Reader) (n int64, err error) {
//	if err = binary.Read(r, binary.BigEndian, &t.PortSrc); err != nil {
//		return
//	}
//	n += 2
//	if err = binary.Read(r, binary.BigEndian, &t.PortDst); err != nil {
//		return
//	}
//	n += 2
//	if err = binary.Read(r, binary.BigEndian, &t.Length); err != nil {
//		return
//	}
//	n += 2
//	if err = binary.Read(r, binary.BigEndian, &t.Checksum); err != nil {
//		return
//	}
//	n += 2
//	return
//}
//
//func (t *TCP) Write(b []byte) (n int, err error) {
//	buf := bytes.NewBuffer(b)
//	if err = binary.Read(buf, binary.BigEndian, &t.PortSrc); err != nil {
//		return
//	}
//	n += 2
//	if err = binary.Read(buf, binary.BigEndian, &t.PortDst); err != nil {
//		return
//	}
//	n += 2
//	if err = binary.Read(buf, binary.BigEndian, &t.Length); err != nil {
//		return
//	}
//	n += 2
//	if err = binary.Read(buf, binary.BigEndian, &t.Checksum); err != nil {
//		return
//	}
//	n += 2
//	t.Data = make([]byte, len(b)-n)
//	if err = binary.Read(buf, binary.BigEndian, &t.Data); err != nil {
//		return
//	}
//	n += len(t.Data)
//	return
//}
