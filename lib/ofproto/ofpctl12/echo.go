package ofp12

func NewEchoRequest() (h *Header) {
	h = new(Header)
	h.Type = OFPTEchoRequest
	return
}

func NewEchoReply() (h *Header) {
	h = new(Header)
	h.Type = OFPTEchoReply
	return
}
