package ofp11

import (
	"testing"
)

func TestExperimenterHeader(t *testing.T) {
	e := NewExperimenterHeader()
	e.Experimenter = 1
	data, err := e.PackBinary()
	if err != nil {
		t.Error("Pack binary error:", err)
	}

	e2 := new(ExperimenterHeader)
	e2.UnpackBinary(data)
	if e.Experimenter != e2.Experimenter {
		t.Error("Encode / Decode - Experimenter:", e, e2)
	}
}
