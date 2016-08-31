package ofp11

import (
	"bytes"
	"encoding/binary"
	"errors"
	_ "fmt"
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
	binary.Read(buf, binary.BigEndian, &a.pad)
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

func NewActionVLANVID() (a *ActionVLANVID) {
	a = new(ActionVLANVID)
	a.Header.Type = uint16(OFPATSetVLANVID)
	a.Header.Length = uint16(a.Len())
	return
}

func (a *ActionVLANVID) Len() (l int) {
	l = 8
	return
}

func (a *ActionVLANVID) PackBinary() (data []byte, err error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	hs := make([]byte, 0)
	hs, err = a.Header.PackBinary()
	if err != nil {
		return
	}
	binary.Write(buf, binary.BigEndian, hs)
	binary.Write(buf, binary.BigEndian, a.VLANVID)
	binary.Write(buf, binary.BigEndian, a.pad)
	data = buf.Bytes()
	return
}

func (a *ActionVLANVID) UnpackBinary(data []byte) (err error) {
	if len(data) < a.Len() {
		return errors.New("The []byte the wrong size to Unpack an " +
			"ActiVLANVID message.")
	}
	buf := bytes.NewBuffer(data)
	hs := make([]byte, a.Header.Len())
	binary.Read(buf, binary.BigEndian, hs)
	err = a.Header.UnpackBinary(hs)
	if err != nil {
		return
	}
	binary.Read(buf, binary.BigEndian, &a.VLANVID)
	binary.Read(buf, binary.BigEndian, &a.pad)
	return
}

func NewActionVLANPCP() (a *ActionVLANPCP) {
	a = new(ActionVLANPCP)
	a.Header.Type = uint16(OFPATSetVLANPCP)
	a.Header.Length = uint16(a.Len())
	return
}

func (a *ActionVLANPCP) Len() (l int) {
	l = 8
	return
}

func (a *ActionVLANPCP) PackBinary() (data []byte, err error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	hs := make([]byte, 0)
	hs, err = a.Header.PackBinary()
	if err != nil {
		return
	}
	binary.Write(buf, binary.BigEndian, hs)
	binary.Write(buf, binary.BigEndian, a.VLANPCP)
	binary.Write(buf, binary.BigEndian, a.pad)
	data = buf.Bytes()
	return
}

func (a *ActionVLANPCP) UnpackBinary(data []byte) (err error) {
	if len(data) < a.Len() {
		return errors.New("The []byte the wrong size to Unpack an " +
			"ActionVLANPCP message.")
	}
	buf := bytes.NewBuffer(data)
	hs := make([]byte, a.Header.Len())
	binary.Read(buf, binary.BigEndian, hs)
	err = a.Header.UnpackBinary(hs)
	if err != nil {
		return
	}
	binary.Read(buf, binary.BigEndian, &a.VLANPCP)
	binary.Read(buf, binary.BigEndian, &a.pad)
	return
}

func NewActionMPLSLabel() (a *ActionMPLSLabel) {
	a = new(ActionMPLSLabel)
	a.Header.Type = uint16(OFPATSetMPLSLabel)
	a.Header.Length = uint16(a.Len())
	return
}

func (a *ActionMPLSLabel) Len() (l int) {
	l = 8
	return
}

func (a *ActionMPLSLabel) PackBinary() (data []byte, err error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	hs := make([]byte, 0)
	hs, err = a.Header.PackBinary()
	if err != nil {
		return
	}
	binary.Write(buf, binary.BigEndian, hs)
	binary.Write(buf, binary.BigEndian, a.MPLSLabel)
	data = buf.Bytes()
	return
}

func (a *ActionMPLSLabel) UnpackBinary(data []byte) (err error) {
	if len(data) < a.Len() {
		return errors.New("The []byte the wrong size to Unpack an " +
			"ActionMPLSLabel message.")
	}
	buf := bytes.NewBuffer(data)
	hs := make([]byte, a.Header.Len())
	binary.Read(buf, binary.BigEndian, hs)
	err = a.Header.UnpackBinary(hs)
	if err != nil {
		return
	}
	binary.Read(buf, binary.BigEndian, &a.MPLSLabel)
	return
}

func NewActionMPLSTC() (a *ActionMPLSTC) {
	a = new(ActionMPLSTC)
	a.Header.Type = uint16(OFPATSetMPLSTC)
	a.Header.Length = uint16(a.Len())
	return
}

func (a *ActionMPLSTC) Len() (l int) {
	l = 8
	return
}

func (a *ActionMPLSTC) PackBinary() (data []byte, err error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	hs := make([]byte, 0)
	hs, err = a.Header.PackBinary()
	if err != nil {
		return
	}
	binary.Write(buf, binary.BigEndian, hs)
	binary.Write(buf, binary.BigEndian, a.MPLSTC)
	binary.Write(buf, binary.BigEndian, a.pad)
	data = buf.Bytes()
	return
}

func (a *ActionMPLSTC) UnpackBinary(data []byte) (err error) {
	if len(data) < a.Len() {
		return errors.New("The []byte the wrong size to Unpack an " +
			"ActionMPLSTC message.")
	}
	buf := bytes.NewBuffer(data)
	hs := make([]byte, a.Header.Len())
	binary.Read(buf, binary.BigEndian, hs)
	err = a.Header.UnpackBinary(hs)
	if err != nil {
		return
	}
	binary.Read(buf, binary.BigEndian, &a.MPLSTC)
	binary.Read(buf, binary.BigEndian, &a.pad)
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

func NewActionDLAddr() (a *ActionDLAddr) {
	a = new(ActionDLAddr)
	a.Header.Length = uint16(a.Len())
	return
}

func NewActionDLSrc() (a *ActionDLAddr) {
	a = new(ActionDLAddr)
	a.Header.Type = uint16(OFPATSetDLSrc)
	a.Header.Length = uint16(a.Len())
	return
}

func NewActionDLDst() (a *ActionDLAddr) {
	a = new(ActionDLAddr)
	a.Header.Type = uint16(OFPATSetDLDst)
	a.Header.Length = uint16(a.Len())
	return
}

func (a *ActionDLAddr) Len() (l int) {
	l = 16
	return
}

func (a *ActionDLAddr) PackBinary() (data []byte, err error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	hs := make([]byte, 0)
	hs, err = a.Header.PackBinary()
	if err != nil {
		return
	}
	binary.Write(buf, binary.BigEndian, hs)
	binary.Write(buf, binary.BigEndian, a.DLAddr)
	binary.Write(buf, binary.BigEndian, a.pad)
	data = buf.Bytes()
	return
}

func (a *ActionDLAddr) UnpackBinary(data []byte) (err error) {
	if len(data) < a.Len() {
		return errors.New("The []byte the wrong size to Unpack an " +
			"ActionDLAddr message.")
	}
	buf := bytes.NewBuffer(data)
	hs := make([]byte, a.Header.Len())
	binary.Read(buf, binary.BigEndian, hs)
	err = a.Header.UnpackBinary(hs)
	if err != nil {
		return
	}
	binary.Read(buf, binary.BigEndian, &a.DLAddr)
	binary.Read(buf, binary.BigEndian, &a.pad)
	return
}

func NewActionNWAddr() (a *ActionNWAddr) {
	a = new(ActionNWAddr)
	a.Header.Length = uint16(a.Len())
	return
}

func NewActionNWSrc() (a *ActionNWAddr) {
	a = new(ActionNWAddr)
	a.Header.Type = uint16(OFPATSetNWSrc)
	a.Header.Length = uint16(a.Len())
	return
}

func NewActionNWDst() (a *ActionNWAddr) {
	a = new(ActionNWAddr)
	a.Header.Type = uint16(OFPATSetNWDst)
	a.Header.Length = uint16(a.Len())
	return
}

func (a *ActionNWAddr) Len() (l int) {
	l = 8
	return
}

func (a *ActionNWAddr) PackBinary() (data []byte, err error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	hs := make([]byte, 0)
	hs, err = a.Header.PackBinary()
	if err != nil {
		return
	}
	binary.Write(buf, binary.BigEndian, hs)
	binary.Write(buf, binary.BigEndian, a.NWAddr)
	data = buf.Bytes()
	return
}

func (a *ActionNWAddr) UnpackBinary(data []byte) (err error) {
	if len(data) < a.Len() {
		return errors.New("The []byte the wrong size to Unpack an " +
			"ActionNWAddr message.")
	}
	buf := bytes.NewBuffer(data)
	hs := make([]byte, a.Header.Len())
	binary.Read(buf, binary.BigEndian, hs)
	err = a.Header.UnpackBinary(hs)
	if err != nil {
		return
	}
	binary.Read(buf, binary.BigEndian, &a.NWAddr)
	return
}

func NewActionNWTOS() (a *ActionNWTOS) {
	a = new(ActionNWTOS)
	a.Header.Type = uint16(OFPATSetNWTos)
	a.Header.Length = uint16(a.Len())
	return
}

func (a *ActionNWTOS) Len() (l int) {
	l = 8
	return
}

func (a *ActionNWTOS) PackBinary() (data []byte, err error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	hs := make([]byte, 0)
	hs, err = a.Header.PackBinary()
	if err != nil {
		return
	}
	binary.Write(buf, binary.BigEndian, hs)
	binary.Write(buf, binary.BigEndian, a.NWTOS)
	binary.Write(buf, binary.BigEndian, a.pad)
	data = buf.Bytes()
	return
}

func (a *ActionNWTOS) UnpackBinary(data []byte) (err error) {
	if len(data) < a.Len() {
		return errors.New("The []byte the wrong size to Unpack an " +
			"ActionNWTOS message.")
	}
	buf := bytes.NewBuffer(data)
	hs := make([]byte, a.Header.Len())
	binary.Read(buf, binary.BigEndian, hs)
	err = a.Header.UnpackBinary(hs)
	if err != nil {
		return
	}
	binary.Read(buf, binary.BigEndian, &a.NWTOS)
	binary.Read(buf, binary.BigEndian, &a.pad)
	return
}

func NewActionNWECN() (a *ActionNWECN) {
	a = new(ActionNWECN)
	a.Header.Type = uint16(OFPATSetNWECN)
	a.Header.Length = uint16(a.Len())
	return
}

func (a *ActionNWECN) Len() (l int) {
	l = 8
	return
}

func (a *ActionNWECN) PackBinary() (data []byte, err error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	hs := make([]byte, 0)
	hs, err = a.Header.PackBinary()
	if err != nil {
		return
	}
	binary.Write(buf, binary.BigEndian, hs)
	binary.Write(buf, binary.BigEndian, a.NWECN)
	binary.Write(buf, binary.BigEndian, a.pad)
	data = buf.Bytes()
	return
}

func (a *ActionNWECN) UnpackBinary(data []byte) (err error) {
	if len(data) < a.Len() {
		return errors.New("The []byte the wrong size to Unpack an " +
			"ActionNWECN message.")
	}
	buf := bytes.NewBuffer(data)
	hs := make([]byte, a.Header.Len())
	binary.Read(buf, binary.BigEndian, hs)
	err = a.Header.UnpackBinary(hs)
	if err != nil {
		return
	}
	binary.Read(buf, binary.BigEndian, &a.NWECN)
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

func NewActionTPPort() (a *ActionTPPort) {
	a = new(ActionTPPort)
	a.Header.Length = uint16(a.Len())
	return
}

func NewActionTPSrc() (a *ActionTPPort) {
	a = new(ActionTPPort)
	a.Header.Type = uint16(OFPATSetTPSrc)
	a.Header.Length = uint16(a.Len())
	return
}

func NewActionTPDst() (a *ActionTPPort) {
	a = new(ActionTPPort)
	a.Header.Type = uint16(OFPATSetTPDst)
	a.Header.Length = uint16(a.Len())
	return
}

func (a *ActionTPPort) Len() (l int) {
	l = 8
	return
}

func (a *ActionTPPort) PackBinary() (data []byte, err error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	hs := make([]byte, 0)
	hs, err = a.Header.PackBinary()
	if err != nil {
		return
	}
	binary.Write(buf, binary.BigEndian, hs)
	binary.Write(buf, binary.BigEndian, a.TPPort)
	binary.Write(buf, binary.BigEndian, a.pad)
	data = buf.Bytes()
	return
}

func (a *ActionTPPort) UnpackBinary(data []byte) (err error) {
	if len(data) < a.Len() {
		return errors.New("The []byte the wrong size to Unpack an " +
			"ActionTPPort message.")
	}
	buf := bytes.NewBuffer(data)
	hs := make([]byte, a.Header.Len())
	binary.Read(buf, binary.BigEndian, hs)
	err = a.Header.UnpackBinary(hs)
	if err != nil {
		return
	}
	binary.Read(buf, binary.BigEndian, &a.TPPort)
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
	case *ActionTPPort:
		data, err = a.(*ActionTPPort).PackBinary()
	case *ActionOutput:
		data, err = a.(*ActionOutput).PackBinary()
	case *ActionDLAddr:
		data, err = a.(*ActionDLAddr).PackBinary()
	case *ActionGroup:
		data, err = a.(*ActionGroup).PackBinary()
	case *ActionMPLSLabel:
		data, err = a.(*ActionMPLSLabel).PackBinary()
	case *ActionMPLSTC:
		data, err = a.(*ActionMPLSTC).PackBinary()
	case *ActionMPLSTTL:
		data, err = a.(*ActionMPLSTTL).PackBinary()
	case *ActionNWAddr:
		data, err = a.(*ActionNWAddr).PackBinary()
	case *ActionNWECN:
		data, err = a.(*ActionNWECN).PackBinary()
	case *ActionNWTOS:
		data, err = a.(*ActionNWTOS).PackBinary()
	case *ActionPopMPLS:
		data, err = a.(*ActionPopMPLS).PackBinary()
	case *ActionPush:
		data, err = a.(*ActionPush).PackBinary()
	case *ActionSetQueue:
		data, err = a.(*ActionSetQueue).PackBinary()
	case *ActionVLANPCP:
		data, err = a.(*ActionVLANPCP).PackBinary()
	case *ActionVLANVID:
		data, err = a.(*ActionVLANVID).PackBinary()
	case *ActionNWTTL:
		data, err = a.(*ActionNWTTL).PackBinary()
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
	case OFPATSetVLANVID:
		a = new(ActionVLANVID)
	case OFPATSetVLANPCP:
		a = new(ActionVLANPCP)
	case OFPATSetDLSrc:
		a = new(ActionDLAddr)
	case OFPATSetDLDst:
		a = new(ActionDLAddr)
	case OFPATSetNWSrc:
		a = new(ActionNWAddr)
	case OFPATSetNWDst:
		a = new(ActionNWAddr)
	case OFPATSetNWTos:
		a = new(ActionNWTOS)
	case OFPATSetNWECN:
		a = new(ActionNWECN)
	case OFPATSetTPSrc:
		a = new(ActionTPPort)
	case OFPATSetTPDst:
		a = new(ActionTPPort)
	case OFPATCopyTTLIn:
		a = new(ActionHeader)
	case OFPATCopyTTLOut:
		a = new(ActionHeader)
	case OFPATSetMPLSLabel:
		a = new(ActionMPLSLabel)
	case OFPATSetMPLSTC:
		a = new(ActionMPLSTC)
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
