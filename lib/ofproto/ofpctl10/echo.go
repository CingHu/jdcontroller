package ofpctl10

import (
	"jd.com/jdcontroller/protocol/ofp10"
)

// Echo request/reply messages can be sent from either the
// switch or the controller, and must return an echo reply. They
// can be used to indicate the latency, bandwidth, and/or
// liveness of a controller-switch connection.
func NewEchoRequest() (h *ofp10.Header) {
	h = NewOfp10Header()
	h.Type = ofp10.OFPTEchoRequest
	return
}

// Echo request/reply messages can be sent from either the
// switch or the controller, and must return an echo reply. They
// can be used to indicate the latency, bandwidth, and/or
// liveness of a controller-switch connection.
func NewEchoReply() (h *ofp10.Header) {
	h = NewOfp10Header()
	h.Type = ofp10.OFPTEchoReply
	return
}
