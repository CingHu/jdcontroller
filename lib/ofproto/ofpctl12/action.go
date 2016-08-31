package ofp12

import (
	"bytes"
	"encoding/binary"
	"errors"
)

func NewActionHeader() (a *ActionHeader) {
	a = new(ActionHeader)
	return
}

func (a *ActionHeader) Len() (l int) {
	l = 4
	return
}

func (a *ActionHeader) PackBinary() (data []byte, err error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	binary.Write(buf, binary.BigEndian, a.Type)
	binary.Write(buf, binary.BigEndian, a.Length)
	data = buf.Bytes()
	return
}

func (a *ActionHeader) UnpackBinary(data []byte) (err error) {
	buf := bytes.NewBuffer(data)
	binary.Read(buf, binary.BigEndian, &a.Type)
	binary.Read(buf, binary.BigEndian, &a.Length)
	return
}

func NewActionOutput(p uint32) (a *ActionOutput) {
	a = new(ActionOutput)
	a.Header.Type = uint16(OFPATOutput)
	a.Header.Length = uint16(a.Len())
	a.Port = p
	return
}

func (a *ActionOutput) Len() (l int) {
	l = 16
	return
}

func (a *ActionOutput) PackBinary() (data []byte, err error) {
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

func (a *ActionOutput) UnpackBinary(data []byte) (err error) {
	if len(data) < a.Len() {
		return errors.New("The []byte the wrong size to Unpack an " +
			"ActionOutput message.")
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

func NewActionGroup() (a *ActionGroup) {
	a = new(ActionGroup)
	a.Header.Type = uint16(OFPATGroup)
	a.Header.Length = uint16(a.Len())
	return
}

func (a *ActionGroup) Len() (l int) {
	l = 8
	return
}

func (a *ActionGroup) PackBinary() (data []byte, err error) {
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

func (a *ActionGroup) UnpackBinary(data []byte) (err error) {
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

func NewActionSetQueue() (a *ActionSetQueue) {
	a = new(ActionSetQueue)
	a.Header.Type = uint16(OFPATSetQueue)
	a.Header.Length = uint16(a.Len())
	return
}

func (a *ActionSetQueue) Len() (l int) {
	l = 8
	return
}

func (a *ActionSetQueue) PackBinary() (data []byte, err error) {
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

func (a *ActionSetQueue) UnpackBinary(data []byte) (err error) {
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

func NewActionMPLSTTL() (a *ActionMPLSTTL) {
	a = new(ActionMPLSTTL)
	a.Header.Type = uint16(OFPATSetMPLSTTL)
	a.Header.Length = uint16(a.Len())
	return
}

func (a *ActionMPLSTTL) Len() (l int) {
	l = 8
	return
}

func (a *ActionMPLSTTL) PackBinary() (data []byte, err error) {
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

func (a *ActionMPLSTTL) UnpackBinary(data []byte) (err error) {
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

func NewActionNWTTL() (a *ActionNWTTL) {
	a = new(ActionNWTTL)
	a.Header.Type = uint16(OFPATSetNWTTL)
	a.Header.Length = uint16(a.Len())
	return
}

func (a *ActionNWTTL) Len() (l int) {
	l = 8
	return
}

func (a *ActionNWTTL) PackBinary() (data []byte, err error) {
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

func (a *ActionNWTTL) UnpackBinary(data []byte) (err error) {
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

func NewActionPush() (a *ActionPush) {
	a = new(ActionPush)
	a.Header.Length = uint16(a.Len())
	return
}

func NewActionPushMPLS() (a *ActionPush) {
	a = new(ActionPush)
	a.Header.Type = uint16(OFPATPushMPLS)
	a.Header.Length = uint16(a.Len())
	return
}

func NewActionPushVLAN() (a *ActionPush) {
	a = new(ActionPush)
	a.Header.Type = uint16(OFPATPushVLAN)
	a.Header.Length = uint16(a.Len())
	return
}

func (a *ActionPush) Len() (l int) {
	l = 8
	return
}

func (a *ActionPush) PackBinary() (data []byte, err error) {
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

func (a *ActionPush) UnpackBinary(data []byte) (err error) {
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

func NewActionPopMPLS() (a *ActionPopMPLS) {
	a = new(ActionPopMPLS)
	a.Header.Type = uint16(OFPATPopMPLS)
	a.Header.Length = uint16(a.Len())
	return
}

func NewActionPopVLAN() (a *ActionHeader) {
	a = new(ActionHeader)
	a.Type = uint16(OFPATPopVLAN)
	return
}

func (a *ActionPopMPLS) Len() (l int) {
	l = 8
	return
}

func (a *ActionPopMPLS) PackBinary() (data []byte, err error) {
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

func (a *ActionPopMPLS) UnpackBinary(data []byte) (err error) {
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

func NewActionSetField() (a *ActionSetField) {
	a = new(ActionSetField)
	a.Header.Type = uint16(OFPATSetField)
	return
}

func (a *ActionSetField) Len() (l int) {
	l = 8
	return
}

func (a *ActionSetField) PackBinary() (data []byte, err error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	hs := make([]byte, 0)
	hs, err = a.Header.PackBinary()
	if err != nil {
		return
	}
	binary.Write(buf, binary.BigEndian, hs)
	binary.Write(buf, binary.BigEndian, a.Field)
	data = buf.Bytes()
	return
}

func (a *ActionSetField) UnpackBinary(data []byte) (err error) {
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

func NewActionExperimenterHeader() (a *ActionExperimenterHeader) {
	a = new(ActionExperimenterHeader)
	a.Header.Type = uint16(OFPATExperimenter)
	a.Header.Length = uint16(a.Len())
	return
}

func (a *ActionExperimenterHeader) Len() (l int) {
	l = 8
	return
}

func (a *ActionExperimenterHeader) PackBinary() (data []byte, err error) {
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

func (a *ActionExperimenterHeader) UnpackBinary(data []byte) (err error) {
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
	case *ActionOutput:
		data, err = a.(*ActionOutput).PackBinary()
	case *ActionGroup:
		data, err = a.(*ActionGroup).PackBinary()
	case *ActionSetQueue:
		data, err = a.(*ActionSetQueue).PackBinary()
	case *ActionMPLSTTL:
		data, err = a.(*ActionMPLSTTL).PackBinary()
	case *ActionNWTTL:
		data, err = a.(*ActionNWTTL).PackBinary()
	case *ActionPush:
		data, err = a.(*ActionPush).PackBinary()
	case *ActionPopMPLS:
		data, err = a.(*ActionPopMPLS).PackBinary()
	case *ActionSetField:
		data, err = a.(*ActionSetField).PackBinary()
	case *ActionExperimenterHeader:
		data, err = a.(*ActionExperimenterHeader).PackBinary()
	}
	if err != nil {
		return
	}
	return
}

func DecodeAction(data []byte) (a Action, err error) {
	switch binary.BigEndian.Uint16(data[:]) {
	case OFPATOutput:
		a = new(ActionOutput)
	case OFPATCopyTTLOut:
		a = new(ActionHeader)
	case OFPATCopyTTLIn:
		a = new(ActionHeader)
	case OFPATSetMPLSTTL:
		a = new(ActionMPLSTTL)
	case OFPATDecMPLSTTL:
		a = new(ActionMPLSTTL)
	case OFPATPushVLAN:
		a = new(ActionPush)
	case OFPATPopVLAN:
		a = new(ActionHeader)
	case OFPATPushMPLS:
		a = new(ActionPush)
	case OFPATPopMPLS:
		a = new(ActionPopMPLS)
	case OFPATSetQueue:
		a = new(ActionSetQueue)
	case OFPATGroup:
		a = new(ActionGroup)
	case OFPATSetNWTTL:
		a = new(ActionNWTTL)
	case OFPATDecNWTTL:
		a = new(ActionNWTTL)
	case OFPATSetField:
		a = new(ActionSetField)
	case OFPATExperimenter:
		a = new(ActionExperimenterHeader)
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
