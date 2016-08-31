package ofp

import (
	"errors"
	"jd.com/jdcontroller/lib/buffer"
	"jd.com/jdcontroller/lib/ofproto/ofp10"
	"jd.com/jdcontroller/lib/ofproto/ofp11"
	"jd.com/jdcontroller/lib/ofproto/ofp12"
	"jd.com/jdcontroller/lib/ofproto/ofp13"
)

func Parse(data []byte) (msg buffer.Message, err error) {
	switch data[0] {
	case 1:
		// ofp 1.0
		msg, err = ofp10.Parse(data)
	case 2:
		// ofp 1.1
		msg, err = ofp11.Parse(data)
	case 3:
		// ofp 1.2
		msg, err = ofp12.Parse(data)
	case 4:
		// ofp 1.3
		msg, err = ofp13.Parse(data)
	default:
		err = errors.New("An unknown OpenFlow version was received.")
	}
	return
}
