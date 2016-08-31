package ofpctl13

import (
	"jd.com/jdcontroller/protocol/ofp13"
)

func NewEchoRequest() (h *ofp13.Header) {
	h = NewHeader()
	h.Type = ofp13.OFPTEchoRequest
	return
}

func NewEchoReply() (h *ofp13.Header) {
	h = NewHeader()
	h.Type = ofp13.OFPTEchoReply
	return
}
