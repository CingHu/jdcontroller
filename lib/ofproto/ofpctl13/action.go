package ofpctl13

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"

	"jd.com/jdcontroller/protocol/ofp13"
)

type ActionCreator struct {
}

func NewActionHeader() (a *ofp13.ActionHeader) {
	a = new(ofp13.ActionHeader)
	return
}

func (a *ofp13.ActionHeader) Len() (l int) {
	l = 4
	return
}

func (a *ofp13.ActionHeader) PackBinary() (data []byte, err error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	binary.Write(buf, binary.BigEndian, a.Type)
	binary.Write(buf, binary.BigEndian, a.Length)
	data = buf.Bytes()
	return
}

func (a *ofp13.ActionHeader) UnpackBinary(data []byte) (err error) {
	buf := bytes.NewBuffer(data)
	binary.Read(buf, binary.BigEndian, &a.Type)
	binary.Read(buf, binary.BigEndian, &a.Length)
	return
}

func (a *ofp13.InstructionActions) AddOutputAction(p uint32) {
	a := new(ofp13.ActionOutput)
	a.Header.Type = uint16(ofp13.OFPATOutput)
	a.Header.Length = uint16(a.Len())
	a.Port = p
	a.MaxLen = 0xffff

	a.Actions = append(a.Actions, a)
	a.Header.Length += a.Header.Length
	fmt.Println("aaaaaaaaaaa", a.Actions)
	return
}

func (a *ofp13.ActionOutput) Len() (l int) {
	l = 16
	return
}

func (a *ofp13.ActionOutput) PackBinary() (data []byte, err error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	hs := make([]byte, 0)
	hs, err = a.Header.PackBinary()
	if err != nil {
		return
	}
	binary.Write(buf, binary.BigEndian, hs)
	binary.Write(buf, binary.BigEndian, a.Port)
	binary.Write(buf, binary.BigEndian, a.MaxLen)
	binary.Write(buf, binary.BigEndian, a.pad)
	data = buf.Bytes()
	return
}

func (a *ofp13.ActionOutput) UnpackBinary(data []byte) (err error) {
	if len(data) < a.Len() {
		return errors.New("The []byte the wrong size to Unpack an " +
			"ofp13.ActionOutput message.")
	}
	buf := bytes.NewBuffer(data)
	hs := make([]byte, a.Header.Len())
	binary.Read(buf, binary.BigEndian, hs)
	err = a.Header.UnpackBinary(hs)
	if err != nil {
		return
	}
	binary.Read(buf, binary.BigEndian, &a.Port)
	binary.Read(buf, binary.BigEndian, &a.MaxLen)
	return
}

func NewActionGroup() (a *ofp13.ActionGroup) {
	a = new(ActionGroup)
	a.Header.Type = uint16(ofp13.OFPATGroup)
	a.Header.Length = uint16(a.Len())
	return
}

func (a *ofp13.ActionGroup) Len() (l int) {
	l = 8
	return
}

func (a *ofp13.ActionGroup) PackBinary() (data []byte, err error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	hs := make([]byte, 0)
	hs, err = a.Header.PackBinary()
	if err != nil {
		return
	}
	binary.Write(buf, binary.BigEndian, hs)
	binary.Write(buf, binary.BigEndian, a.GroupID)
	data = buf.Bytes()
	return
}

func (a *ofp13.ActionGroup) UnpackBinary(data []byte) (err error) {
	if len(data) < a.Len() {
		return errors.New("The []byte the wrong size to Unpack an " +
			"ActionGroup message.")
	}
	buf := bytes.NewBuffer(data)
	hs := make([]byte, a.Header.Len())
	binary.Read(buf, binary.BigEndian, hs)
	err = a.Header.UnpackBinary(hs)
	if err != nil {
		return
	}
	binary.Read(buf, binary.BigEndian, &a.GroupID)
	return
}

func NewActionSetQueue() (a *ofp13.ActionSetQueue) {
	a = new(ActionSetQueue)
	a.Header.Type = uint16(ofp13.OFPATSetQueue)
	a.Header.Length = uint16(a.Len())
	return
}

func (a *ofp13.ActionSetQueue) Len() (l int) {
	l = 8
	return
}

func (a *ofp13.ActionSetQueue) PackBinary() (data []byte, err error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	hs := make([]byte, 0)
	hs, err = a.Header.PackBinary()
	if err != nil {
		return
	}
	binary.Write(buf, binary.BigEndian, hs)
	binary.Write(buf, binary.BigEndian, a.QueueID)
	data = buf.Bytes()
	return
}

func (a *ofp13.ActionSetQueue) UnpackBinary(data []byte) (err error) {
	if len(data) < a.Len() {
		return errors.New("The []byte the wrong size to Unpack an " +
			"ActionSetQueue message.")
	}
	buf := bytes.NewBuffer(data)
	hs := make([]byte, a.Header.Len())
	binary.Read(buf, binary.BigEndian, hs)
	err = a.Header.UnpackBinary(hs)
	if err != nil {
		return
	}
	binary.Read(buf, binary.BigEndian, &a.QueueID)
	return
}

func NewActionMPLSTTL() (a *ofp13.ActionMPLSTTL) {
	a = new(ofp13.ActionMPLSTTL)
	a.Header.Type = uint16(ofp13.OFPATSetMPLSTTL)
	a.Header.Length = uint16(a.Len())
	return
}

func (a *ofp13.ActionMPLSTTL) Len() (l int) {
	l = 8
	return
}

func (a *ofp13.ActionMPLSTTL) PackBinary() (data []byte, err error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	hs := make([]byte, 0)
	hs, err = a.Header.PackBinary()
	if err != nil {
		return
	}
	binary.Write(buf, binary.BigEndian, hs)
	binary.Write(buf, binary.BigEndian, a.MPLSTTL)
	binary.Write(buf, binary.BigEndian, a.pad)
	data = buf.Bytes()
	return
}

func (a *ofp13.ActionMPLSTTL) UnpackBinary(data []byte) (err error) {
	if len(data) < a.Len() {
		return errors.New("The []byte the wrong size to Unpack an " +
			"ActiMPLSTTL message.")
	}
	buf := bytes.NewBuffer(data)
	hs := make([]byte, a.Header.Len())
	binary.Read(buf, binary.BigEndian, hs)
	err = a.Header.UnpackBinary(hs)
	if err != nil {
		return
	}
	binary.Read(buf, binary.BigEndian, &a.MPLSTTL)
	binary.Read(buf, binary.BigEndian, &a.pad)
	return
}

func NewActionNWTTL() (a *ofp13.ActionNWTTL) {
	a = new(ActionNWTTL)
	a.Header.Type = uint16(ofp13.OFPATSetNWTTL)
	a.Header.Length = uint16(a.Len())
	return
}

func (a *ofp13.ActionNWTTL) Len() (l int) {
	l = 8
	return
}

func (a *ofp13.ActionNWTTL) PackBinary() (data []byte, err error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	hs := make([]byte, 0)
	hs, err = a.Header.PackBinary()
	if err != nil {
		return
	}
	binary.Write(buf, binary.BigEndian, hs)
	binary.Write(buf, binary.BigEndian, a.NWTTL)
	binary.Write(buf, binary.BigEndian, a.pad)
	data = buf.Bytes()
	return
}

func (a *ofp13.ActionNWTTL) UnpackBinary(data []byte) (err error) {
	if len(data) < a.Len() {
		return errors.New("The []byte the wrong size to Unpack an " +
			"ActionNWTTL message.")
	}
	buf := bytes.NewBuffer(data)
	hs := make([]byte, a.Header.Len())
	binary.Read(buf, binary.BigEndian, hs)
	err = a.Header.UnpackBinary(hs)
	if err != nil {
		return
	}
	binary.Read(buf, binary.BigEndian, &a.NWTTL)
	binary.Read(buf, binary.BigEndian, &a.pad)
	return
}

func NewActionPush() (a *ofp13.ActionPush) {
	a = new(ActionPush)
	a.Header.Length = uint16(a.Len())
	return
}

func NewActionPushMPLS() (a *ofp13.ActionPush) {
	a = new(ActionPush)
	a.Header.Type = uint16(ofp13.OFPATPushMPLS)
	a.Header.Length = uint16(a.Len())
	return
}

func NewActionPushVLAN() (a *ofp13.ActionPush) {
	a = new(ActionPush)
	a.Header.Type = uint16(ofp13.OFPATPushVLAN)
	a.Header.Length = uint16(a.Len())
	return
}

func NewActionPushPBB() (a *ofp13.ActionPush) {
	a = new(ActionPush)
	a.Header.Type = uint16(ofp13.OFPATPushPBB)
	a.Header.Length = uint16(a.Len())
	return
}

func (a *ofp13.ActionPush) Len() (l int) {
	l = 8
	return
}

func (a *ofp13.ActionPush) PackBinary() (data []byte, err error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	hs := make([]byte, 0)
	hs, err = a.Header.PackBinary()
	if err != nil {
		return
	}
	binary.Write(buf, binary.BigEndian, hs)
	binary.Write(buf, binary.BigEndian, a.Ethertype)
	binary.Write(buf, binary.BigEndian, a.pad)
	data = buf.Bytes()
	return
}

func (a *ofp13.ActionPush) UnpackBinary(data []byte) (err error) {
	if len(data) < a.Len() {
		return errors.New("The []byte the wrong size to Unpack an " +
			"ActionPush message.")
	}
	buf := bytes.NewBuffer(data)
	hs := make([]byte, a.Header.Len())
	binary.Read(buf, binary.BigEndian, hs)
	err = a.Header.UnpackBinary(hs)
	if err != nil {
		return
	}
	binary.Read(buf, binary.BigEndian, &a.Ethertype)
	binary.Read(buf, binary.BigEndian, &a.pad)
	return
}

func NewActionPopVLAN() (a *ofp13.ActionHeader) {
	a = new(ofp13.ActionHeader)
	a.Type = uint16(ofp13.OFPATPopVLAN)
	return
}

func NewActionPopPBB() (a *ofp13.ActionHeader) {
	a = new(ofp13.ActionHeader)
	a.Type = uint16(ofp13.OFPATPopPBB)
	return
}

func NewActionPopMPLS() (a *ofp13.ActionPopMPLS) {
	a = new(ActionPopMPLS)
	a.Header.Type = uint16(ofp13.OFPATPopMPLS)
	a.Header.Length = uint16(a.Len())
	return
}

func (a *ofp13.ActionPopMPLS) Len() (l int) {
	l = 8
	return
}

func (a *ofp13.ActionPopMPLS) PackBinary() (data []byte, err error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	hs := make([]byte, 0)
	hs, err = a.Header.PackBinary()
	if err != nil {
		return
	}
	binary.Write(buf, binary.BigEndian, hs)
	binary.Write(buf, binary.BigEndian, a.Ethertype)
	binary.Write(buf, binary.BigEndian, a.pad)
	data = buf.Bytes()
	return
}

func (a *ofp13.ActionPopMPLS) UnpackBinary(data []byte) (err error) {
	if len(data) < a.Len() {
		return errors.New("The []byte the wrong size to Unpack an " +
			"ActionPopMPLS message.")
	}
	buf := bytes.NewBuffer(data)
	hs := make([]byte, a.Header.Len())
	binary.Read(buf, binary.BigEndian, hs)
	err = a.Header.UnpackBinary(hs)
	if err != nil {
		return
	}
	binary.Read(buf, binary.BigEndian, &a.Ethertype)
	binary.Read(buf, binary.BigEndian, &a.pad)
	return
}

//func NewActionSetField() (a *ofp13.ActionSetField) {
func (ac ActionCreator)NewActionSetField() (a *ofp13.ActionSetField) {
	a = new(ActionSetField)
	a.Header.Type = uint16(ofp13.OFPATSetField)
	return
}

func (a *ofp13.ActionSetField) Len() (l int) {
	l = 4
	return
}

func (a *ofp13.ActionSetField) PackBinary() (data []byte, err error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	hs := make([]byte, 0)
	hs, err = a.Header.PackBinary()
	if err != nil {
		return
	}
	binary.Write(buf, binary.BigEndian, hs)
	binary.Write(buf, binary.BigEndian, a.Field)
	//binary.Write(buf, binary.BigEndian, a.pad)
	data = buf.Bytes()
	return
}

func (a *ofp13.ActionSetField) UnpackBinary(data []byte) (err error) {
	if len(data) < a.Len() {
		return errors.New("The []byte the wrong size to Unpack an " +
			"ActionPopMPLS message.")
	}
	buf := bytes.NewBuffer(data)
	hs := make([]byte, a.Header.Len())
	binary.Read(buf, binary.BigEndian, hs)
	err = a.Header.UnpackBinary(hs)
	if err != nil {
		return
	}
	binary.Read(buf, binary.BigEndian, &a.Field)
	return
}

func NewActionExperimenterHeader() (a *ofp13.ActionExperimenterHeader) {
	a = new(ActionExperimenterHeader)
	a.Header.Type = uint16(ofp13.OFPATExperimenter)
	a.Header.Length = uint16(a.Len())
	return
}

func (a *ofp13.ActionExperimenterHeader) Len() (l int) {
	l = 8
	return
}

func (a *ofp13.ActionExperimenterHeader) PackBinary() (data []byte, err error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	hs := make([]byte, 0)
	hs, err = a.Header.PackBinary()
	if err != nil {
		return
	}
	binary.Write(buf, binary.BigEndian, hs)
	binary.Write(buf, binary.BigEndian, a.Experimenter)
	data = buf.Bytes()
	return
}

func (a *ofp13.ActionExperimenterHeader) UnpackBinary(data []byte) (err error) {
	if len(data) < a.Len() {
		return errors.New("The []byte the wrong size to Unpack an " +
			"ActionExperimenterHeader message.")
	}
	buf := bytes.NewBuffer(data)
	hs := make([]byte, a.Header.Len())
	binary.Read(buf, binary.BigEndian, hs)
	err = a.Header.UnpackBinary(hs)
	if err != nil {
		return
	}
	binary.Read(buf, binary.BigEndian, &a.Experimenter)
	return
}

func EncodeAction(a Action) (data []byte, err error) {
	data = make([]byte, 0)
	switch a.(type) {
	case *ofp13.ActionOutput:
		data, err = a.(*ofp13.ActionOutput).PackBinary()
	case *ofp13.ActionGroup:
		data, err = a.(*ofp13.ActionGroup).PackBinary()
	case *ofp13.ActionSetQueue:
		data, err = a.(*ofp13.ActionSetQueue).PackBinary()
	case *ofp13.ActionMPLSTTL:
		data, err = a.(*ofp13.ActionMPLSTTL).PackBinary()
	case *ofp13.ActionNWTTL:
		data, err = a.(*ofp13.ActionNWTTL).PackBinary()
	case *ofp13.ActionPush:
		data, err = a.(*ofp13.ActionPush).PackBinary()
	case *ofp13.ActionPopMPLS:
		data, err = a.(*ofp13.ActionPopMPLS).PackBinary()
	case *ofp13.ActionSetField:
		data, err = a.(*ofp13.ActionSetField).PackBinary()
	case *ofp13.ActionHeader:
		data, err = a.(*ofp13.ActionHeader).PackBinary()
	case *ofp13.ActionExperimenterHeader:
		data, err = a.(*ofp13.ActionExperimenterHeader).PackBinary()
	}
	if err != nil {
		return
	}
	return
}

func DecodeAction(data []byte) (a Action, err error) {
	switch binary.BigEndian.Uint16(data[:]) {
	case ofp13.OFPATOutput:
		a = new(ofp13.ActionOutput)
	case ofp13.OFPATCopyTTLOut:
		a = new(ofp13.ActionHeader)
	case ofp13.OFPATCopyTTLIn:
		a = new(ofp13.ActionHeader)
	case ofp13.OFPATSetMPLSTTL:
		a = new(ActionMPLSTTL)
	case ofp13.OFPATDecMPLSTTL:
		a = new(ofp13.ActionMPLSTTL)
	case ofp13.OFPATPushVLAN:
		a = new(ofp13.ActionPush)
	case ofp13.OFPATPopVLAN:
		a = new(ofp13.ActionHeader)
	case ofp13.OFPATPushMPLS:
		a = new(ofp13.ActionPush)
	case ofp13.OFPATPopMPLS:
		a = new(ofp13.ActionPopMPLS)
	case ofp13.OFPATSetQueue:
		a = new(ofp13.ActionSetQueue)
	case ofp13.OFPATGroup:
		a = new(ofp13.ActionGroup)
	case ofp13.OFPATSetNWTTL:
		a = new(ofp13.ActionNWTTL)
	case ofp13.OFPATDecNWTTL:
		a = new(ofp13.ActionNWTTL)
	case ofp13.OFPATSetField:
		a = new(ofp13.ActionSetField)
	case ofp13.OFPATPushPBB:
		a = new(ofp13.ActionPush)
	case ofp13.OFPATPopPBB:
		a = new(ofp13.ActionHeader)
	case ofp13.OFPATExperimenter:
		a = new(ofp13.ActionExperimenterHeader)
	}
	buf := bytes.NewBuffer(data)
	as := make([]byte, a.Len())
	binary.Read(buf, binary.BigEndian, as)
	err = a.UnpackBinary(as)
	if err != nil {
		return
	}
	return
}
