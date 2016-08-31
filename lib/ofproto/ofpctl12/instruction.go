package ofp12

import (
	"bytes"
	"encoding/binary"
)

func NewInstructionHeader() (i *InstructionHeader) {
	i = new(InstructionHeader)
	return
}

func (i *InstructionHeader) Len() (l int) {
	l = 4
	return
}

func (i *InstructionHeader) PackBinary() (data []byte, err error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	binary.Write(buf, binary.BigEndian, i.Type)
	binary.Write(buf, binary.BigEndian, i.Length)
	data = buf.Bytes()
	return
}

func (i *InstructionHeader) UnpackBinary(data []byte) (err error) {
	buf := bytes.NewBuffer(data)
	binary.Read(buf, binary.BigEndian, &i.Type)
	binary.Read(buf, binary.BigEndian, &i.Length)
	return
}

func NewInstructionGotoTable() (i *InstructionGotoTable) {
	i = new(InstructionGotoTable)
	i.Header.Type = uint16(OFPITGotoTable)
	i.Header.Length = uint16(i.Len())
	return
}

func (i *InstructionGotoTable) Len() (l int) {
	l = 8
	return
}

func (i *InstructionGotoTable) PackBinary() (data []byte, err error) {
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

func (i *InstructionGotoTable) UnpackBinary(data []byte) (err error) {
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

func NewInstructionActions() (i *InstructionActions) {
	i = new(InstructionActions)
	return
}

func (i *InstructionActions) Len() (l int) {
	l = 8
	return
}

func (i *InstructionActions) PackBinary() (data []byte, err error) {
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

func (i *InstructionActions) UnpackBinary(data []byte) (err error) {
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

func EncodeInstruction(i Instruction) (data []byte, err error) {
	data = make([]byte, 0)
	switch i.(type) {
	case *InstructionGotoTable:
		data, err = i.(*InstructionGotoTable).PackBinary()
	case *InstructionWriteMetadata:
		data, err = i.(*InstructionWriteMetadata).PackBinary()
	case *InstructionActions:
		data, err = i.(*InstructionActions).PackBinary()
	}
	if err != nil {
		return
	}
	return
}

func DecodeInstruction(data []byte) (i Instruction, err error) {
	switch binary.BigEndian.Uint16(data[:]) {
	case OFPITGotoTable:
		i = new(InstructionGotoTable)
	case OFPITWriteActions:
		i = new(InstructionActions)
	case OFPITApplyActions:
		i = new(InstructionActions)
	case OFPITClearActions:
		i = new(InstructionActions)
	case OFPITWriteMetadata:
		i = new(InstructionWriteMetadata)
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
