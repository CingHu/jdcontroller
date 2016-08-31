package ofp

type Header struct {
	Version uint8
	Type    uint8
	Length  uint16
	XID     uint32
}
