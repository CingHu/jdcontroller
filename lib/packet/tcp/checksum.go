package tcp

func Checksum(b []byte) uint16 {
	csum := len(b) - 1 // checksum coverage
	s := uint32(0)
	for i := 0; i < csum; i += 2 {
		s += uint32(b[i+1])<<8 | uint32(b[i])
	}
	if csum&1 == 0 {
		s += uint32(b[csum])
	}
	s = s>>16 + s&0xffff
	s = s + s>>16
	s = ^s & 0xffff
	return uint16(s<<8 | s>>(16-8))
}
