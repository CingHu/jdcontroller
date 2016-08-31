package ofpctl13

import (
	"bytes"
	"encoding/binary"

	"jd.com/jdcontroller/protocol/ofp13"
)

func NewInstructionHeader() (i *ofp13.InstructionHeader) {
	i = new(ofp13.InstructionHeader)
	return
}

func (i *ofp13.InstructionHeader) Len() (l int) {
	l = 4
	return
}

func (i *ofp13.InstructionHeader) PackBinary() (data []byte, err error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	binary.Write(buf, binary.BigEndian, i.Type)
	binary.Write(buf, binary.BigEndian, i.Length)
	data = buf.Bytes()
	return
}

func (i *ofp13.InstructionHeader) UnpackBinary(data []byte) (err error) {
	buf := bytes.NewBuffer(data)
	binary.Read(buf, binary.BigEndian, &i.Type)
	binary.Read(buf, binary.BigEndian, &i.Length)
	return
}

func NewInstructionGotoTable() (i *ofp13.InstructionGotoTable) {
	i = new(ofp13.InstructionGotoTable)
	i.Header.Type = uint16(OFPITGotoTable)
	i.Header.Length = uint16(i.Len())
	return
}

func (i *ofp13.InstructionGotoTable) Len() (l int) {
	l = 8
	return
}

func (i *ofp13.InstructionGotoTable) PackBinary() (data []byte, err error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	hs := make([]byte, 0)
	hs, err = i.Header.PackBinary()
	if err != nil {
		return
	}
	binary.Write(buf, binary.BigEndian, hs)
	binary.Write(buf, binary.BigEndian, i.TableID)
	binary.Write(buf, binary.BigEndian, i.pad)
	data = buf.Bytes()
	return
}

func (i *ofp13.InstructionGotoTable) UnpackBinary(data []byte) (err error) {
	buf := bytes.NewBuffer(data)
	hs := make([]byte, i.Header.Len())
	binary.Read(buf, binary.BigEndian, hs)
	err = i.Header.UnpackBinary(hs)
	if err != nil {
		return
	}
	binary.Read(buf, binary.BigEndian, &i.TableID)
	binary.Read(buf, binary.BigEndian, &i.pad)
	return
}

func NewInstructionWriteMetadata() (i *InstructionWriteMetadata) {
	i = new(InstructionWriteMetadata)
	i.Header.Type = uint16(OFPITWriteMetadata)
	i.Header.Length = uint16(i.Len())
	return
}

func (i *InstructionWriteMetadata) Len() (l int) {
	l = 24
	return
}

func (i *InstructionWriteMetadata) PackBinary() (data []byte, err error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	hs := make([]byte, 0)
	hs, err = i.Header.PackBinary()
	if err != nil {
		return
	}
	binary.Write(buf, binary.BigEndian, hs)
	binary.Write(buf, binary.BigEndian, i.pad)
	binary.Write(buf, binary.BigEndian, i.Metadata)
	binary.Write(buf, binary.BigEndian, i.MetadataMask)
	data = buf.Bytes()
	return
}

func (i *InstructionWriteMetadata) UnpackBinary(data []byte) (err error) {
	buf := bytes.NewBuffer(data)
	hs := make([]byte, i.Header.Len())
	binary.Read(buf, binary.BigEndian, hs)
	err = i.Header.UnpackBinary(hs)
	if err != nil {
		return
	}
	binary.Read(buf, binary.BigEndian, &i.pad)
	binary.Read(buf, binary.BigEndian, &i.Metadata)
	binary.Read(buf, binary.BigEndian, &i.MetadataMask)
	return
}

func NewInstructionActions() (i *ofp13.InstructionActions) {
	i = new(ofp13.InstructionActions)
	i.Header.Type = ofp13.OFPITApplyActions
	i.Header.Length = uint16(i.Len()) //ofp13.InstructionActions struct's length
	return
}

func (i *ofp13.InstructionActions) Len() (l int) {
	l = 8
	return
}

func (i *ofp13.InstructionActions) PackBinary() (data []byte, err error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	hs := make([]byte, 0)
	hs, err = i.Header.PackBinary()
	if err != nil {
		return
	}
	binary.Write(buf, binary.BigEndian, hs)
	binary.Write(buf, binary.BigEndian, i.pad)
	for _, a := range i.Actions {
		bs := make([]byte, 0)
		bs, err = EncodeAction(a)
		if err != nil {
			return
		}
		binary.Write(buf, binary.BigEndian, bs)
	}

	data = buf.Bytes()
	return
}

func (i *ofp13.InstructionActions) UnpackBinary(data []byte) (err error) {
	buf := bytes.NewBuffer(data)
	hs := make([]byte, i.Header.Len())
	binary.Read(buf, binary.BigEndian, hs)
	err = i.Header.UnpackBinary(hs)
	if err != nil {
		return
	}
	binary.Read(buf, binary.BigEndian, &i.pad)

	n := i.Len()
	for n < len(data) {
		var a Action
		a, err = DecodeAction(data[n:])
		if err != nil {
			return
		}
		i.Actions = append(i.Actions, a)
		n += a.Len()
	}
	return
}

func NewInstructionMeter() (i *InstructionMeter) {
	i = new(InstructionMeter)
	i.Header.Type = uint16(OFPITMeter)
	i.Header.Length = uint16(i.Len())
	return
}

func (i *InstructionMeter) Len() (l int) {
	l = 8
	return
}

func (i *InstructionMeter) PackBinary() (data []byte, err error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	hs := make([]byte, 0)
	hs, err = i.Header.PackBinary()
	if err != nil {
		return
	}
	binary.Write(buf, binary.BigEndian, hs)
	binary.Write(buf, binary.BigEndian, i.MeterID)
	data = buf.Bytes()
	return
}

func (i *InstructionMeter) UnpackBinary(data []byte) (err error) {
	buf := bytes.NewBuffer(data)
	hs := make([]byte, i.Header.Len())
	binary.Read(buf, binary.BigEndian, hs)
	err = i.Header.UnpackBinary(hs)
	if err != nil {
		return
	}
	binary.Read(buf, binary.BigEndian, &i.MeterID)
	return
}

func NewInstructionExperimenter() (i *InstructionExperimenter) {
	i = new(InstructionExperimenter)
	i.Header.Type = uint16(OFPITExperimenter)
	i.Header.Length = uint16(i.Len())
	return
}

func (i *InstructionExperimenter) Len() (l int) {
	l = 8
	return
}

func (i *InstructionExperimenter) PackBinary() (data []byte, err error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	hs := make([]byte, 0)
	hs, err = i.Header.PackBinary()
	if err != nil {
		return
	}
	binary.Write(buf, binary.BigEndian, hs)
	binary.Write(buf, binary.BigEndian, i.Experimenter)
	data = buf.Bytes()
	return
}

func (i *InstructionExperimenter) UnpackBinary(data []byte) (err error) {
	buf := bytes.NewBuffer(data)
	hs := make([]byte, i.Header.Len())
	binary.Read(buf, binary.BigEndian, hs)
	err = i.Header.UnpackBinary(hs)
	if err != nil {
		return
	}
	binary.Read(buf, binary.BigEndian, &i.Experimenter)
	return
}

func EncodeInstruction(i ofp13.Instruction) (data []byte, err error) {
	data = make([]byte, 0)
	switch i.(type) {
	case *ofp13.InstructionGotoTable:
		data, err = i.(*ofp13.InstructionGotoTable).PackBinary()
	case *InstructionWriteMetadata:
		data, err = i.(*InstructionWriteMetadata).PackBinary()
	case *ofp13.InstructionActions:
		data, err = i.(*ofp13.InstructionActions).PackBinary()
	case *InstructionMeter:
		data, err = i.(*InstructionMeter).PackBinary()
	case *InstructionExperimenter:
		data, err = i.(*InstructionExperimenter).PackBinary()
	}
	if err != nil {
		return
	}
	return
}

func DecodeInstruction(data []byte) (i Instruction, err error) {
	switch binary.BigEndian.Uint16(data[:]) {
	case OFPITGotoTable:
		i = new(ofp13.InstructionGotoTable)
	case OFPITWriteActions:
		i = new(ofp13.InstructionActions)
	case ofp13.OFPITApplyActions:
		i = new(ofp13.InstructionActions)
	case OFPITClearActions:
		i = new(ofp13.InstructionActions)
	case OFPITWriteMetadata:
		i = new(InstructionWriteMetadata)
	case OFPITMeter:
		i = new(InstructionMeter)
	case OFPITExperimenter:
		i = new(InstructionExperimenter)
	}
	buf := bytes.NewBuffer(data)
	is := make([]byte, i.Len())
	binary.Read(buf, binary.BigEndian, is)
	err = i.UnpackBinary(is)
	if err != nil {
		return
	}
	return
}
