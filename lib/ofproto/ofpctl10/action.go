package ofpctl10

import (
	"bytes"
	"encoding/binary"
	"errors"

	"jd.com/jdcontroller/protocol/ofp10"
)

func NewActionHeader() (a *ofp10.ActionHeader) {
	a = new(ofp10.ActionHeader)
	a.Type = 0
	a.Length = 4
	return
}

func (a *ofp10.ActionHeader) Len() (l int) {
	l = 4
	return
}

func (a *ofp10.ActionHeader) PackBinary() (data []byte, err error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	binary.Write(buf, binary.BigEndian, a.Type)
	binary.Write(buf, binary.BigEndian, a.Length)
	data = buf.Bytes()
	return
}

func (a *ofp10.ActionHeader) UnpackBinary(data []byte) (err error) {
	if len(data) != 4 {
		return errors.New("The []byte the wrong size to Unpack an " +
			"ofp10.ActionHeader message.")
	}
	buf := bytes.NewBuffer(data)
	binary.Read(buf, binary.BigEndian, &a.Type)
	binary.Read(buf, binary.BigEndian, &a.Length)
	return
}

// Returns a new ofp10.Action Output message which sends packets out
// port number.
func NewActionOutput(p uint16) (a *ofp10.ActionOutput) {
	a = new(ofp10.ActionOutput)
	a.Header.Type = uint16(OFPATOutput)
	a.Port = p
	return
}

func (a *ofp10.ActionOutput) Len() (l int) {
	l = 8
	return
}

func (a *ofp10.ActionOutput) PackBinary() (data []byte, err error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	hs := make([]byte, 0)
	hs, err = a.Header.PackBinary()
	if err != nil {
		return
	}
	binary.Write(buf, binary.BigEndian, hs)
	binary.Write(buf, binary.BigEndian, a.Port)
	binary.Write(buf, binary.BigEndian, a.MaxLen)
	data = buf.Bytes()
	return
}

func (a *ofp10.ActionOutput) UnpackBinary(data []byte) (err error) {
	if len(data) < int(a.Len()) {
		return errors.New("The []byte the wrong size to Unpack an " +
			"ofp10.ActionOutput message.")
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

func NewActionEnqueue() (a *ActionEnqueue) {
	a = new(ofp10.ActionEnqueue)
	a.Header.Type = uint16(OFPATEnqueue)
	return
}

func (a *ofp10.ActionEnqueue) Len() (l int) {
	l = 16
	return
}

func (a *ofp10.ActionEnqueue) PackBinary() (data []byte, err error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	hs := make([]byte, 0)
	hs, err = a.Header.PackBinary()
	if err != nil {
		return
	}
	binary.Write(buf, binary.BigEndian, hs)
	binary.Write(buf, binary.BigEndian, a.Port)
	binary.Write(buf, binary.BigEndian, a.pad)
	binary.Write(buf, binary.BigEndian, a.QueueID)
	data = buf.Bytes()
	return
}

func (a *ofp10.ActionEnqueue) UnpackBinary(data []byte) (err error) {
	if len(data) != int(a.Len()) {
		return errors.New("The []byte the wrong size to Unpack an " +
			"ActionEnqueue message.")
	}
	buf := bytes.NewBuffer(data)
	hs := make([]byte, a.Header.Len())
	binary.Read(buf, binary.BigEndian, hs)
	err = a.Header.UnpackBinary(hs)
	if err != nil {
		return
	}
	binary.Read(buf, binary.BigEndian, &a.Port)
	binary.Read(buf, binary.BigEndian, &a.pad)
	binary.Read(buf, binary.BigEndian, &a.QueueID)
	return
}

// Sets a VLAN ID on tagged packets. VLAN ID may be added to
// untagged packets on some switches.
func NewActionVLANVID() (a *ofp10.ActionVLANVID) {
	a = new(ofp10.ActionVLANVID)
	a.Header.Type = uint16(OFPATSetVLANVID)
	return
}

func (a *ofp10.ActionVLANVID) Len() (l int) {
	l = 8
	return
}

func (a *ofp10.ActionVLANVID) PackBinary() (data []byte, err error) {
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

func (a *ofp10.ActionVLANVID) UnpackBinary(data []byte) (err error) {
	if len(data) != int(a.Len()) {
		return errors.New("The []byte the wrong size to Unpack an " +
			"ActionVLANVID message.")
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

// Modifies PCP on VLAN tagged packets.
func NewActionVLANPCP() (a *ofp10.ActionVLANPCP) {
	a = new(ofp10.ActionVLANPCP)
	a.Header.Type = uint16(OFPATSetVLANPCP)
	return
}

func (a *ofp10.ActionVLANPCP) Len() (l int) {
	l = 8
	return
}

func (a *ofp10.ActionVLANPCP) PackBinary() (data []byte, err error) {
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

func (a *ofp10.ActionVLANPCP) UnpackBinary(data []byte) (err error) {
	if len(data) != int(a.Len()) {
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

// ofp10.Action to strip VLAN IDs from tagged packets.
func NewActionStripVLAN() (a *ofp10.ActionHeader) {
	a = new(ofp10.ActionHeader)
	a.Type = OFPATStripVLAN
	return
}

// Sets the source MAC adddress to dlAddr
func NewActionDLSrc() (a *ofp10.ActionDLAddr) {
	a = new(ofp10.ActionDLAddr)
	a.Header.Type = uint16(OFPATSetDLSrc)
	return
}

// Sets the destination MAC adddress to dlAddr
func NewActionDLDst() (a *ofp10.ActionDLAddr) {
	a = new(ofp10.ActionDLAddr)
	a.Header.Type = uint16(OFPATSetDLDst)
	return
}

func (a *ofp10.ActionDLAddr) Len() (l int) {
	l = 16
	return
}

func (a *ofp10.ActionDLAddr) PackBinary() (data []byte, err error) {
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

func (a *ofp10.ActionDLAddr) UnpackBinary(data []byte) (err error) {
	if len(data) != int(a.Len()) {
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

// Sets the source IP adddress to nwAddr
func NewActionNWSrc() (a *ofp10.ActionNWAddr) {
	a = new(ofp10.ActionNWAddr)
	a.Header.Type = uint16(OFPATSetNWSrc)
	return
}

// Sets the destination IP adddress to nwAddr
func NewActionNWDst() (a *ofp10.ActionNWAddr) {
	a = new(ofp10.ActionNWAddr)
	a.Header.Type = uint16(OFPATSetNWDst)
	return
}

func (a *ofp10.ActionNWAddr) Len() (l int) {
	l = 8
	return
}

func (a *ofp10.ActionNWAddr) PackBinary() (data []byte, err error) {
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

func (a *ofp10.ActionNWAddr) UnpackBinary(data []byte) (err error) {
	if len(data) != int(a.Len()) {
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
	binary.Read(buf, binary.BigEndian, &a.NWAddr)
	return
}

// Set ToS field in IP packets.
func NewActionNWTos() (a *ofp10.ActionNWTOS) {
	a = new(ofp10.ActionNWTOS)
	a.Header.Type = uint16(OFPATSetNWTos)
	return
}

func (a *ofp10.ActionNWTOS) Len() (l int) {
	l = 8
	return
}

func (a *ofp10.ActionNWTOS) PackBinary() (data []byte, err error) {
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

func (a *ofp10.ActionNWTOS) UnpackBinary(data []byte) (err error) {
	if len(data) != int(a.Len()) {
		return errors.New("The []byte the wrong size to Unpack an " +
			"ofp10.ActionDLAddr message.")
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

// Returns an action that sets the transport layer source port.
func NewActionTPSrc() (a *ofp10.ActionTPPort) {
	a = new(ofp10.ActionTPPort)
	a.Header.Type = uint16(OFPATSetTPSrc)
	return
}

// Returns an action that sets the transport layer destination
// port.
func NewActionTPDst() (a *ofp10.ActionTPPort) {
	a = new(ofp10.ActionTPPort)
	a.Header.Type = uint16(OFPATSetTPDst)
	return
}

func (a *ofp10.ActionTPPort) Len() (l int) {
	l = 8
	return
}

func (a *ofp10.ActionTPPort) PackBinary() (data []byte, err error) {
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

func (a *ofp10.ActionTPPort) UnpackBinary(data []byte) (err error) {
	if len(data) != int(a.Len()) {
		return errors.New("The []byte the wrong size to Unpack an " +
			"ofp10.ActionNWTOS message.")
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

func NewActionVendor() (a *ofp10.ActionVendorHeader) {
	a = new(ofp10.ActionVendorHeader)
	a.Header.Type = uint16(OFPATVendor)
	return
}

func (a *ofp10.ActionVendorHeader) Len() (l int) {
	l = 8
	return
}

func (a *ofp10.ActionVendorHeader) PackBinary() (data []byte, err error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	hs := make([]byte, 0)
	hs, err = a.Header.PackBinary()
	if err != nil {
		return
	}
	binary.Write(buf, binary.BigEndian, hs)
	binary.Write(buf, binary.BigEndian, a.Vendor)
	data = buf.Bytes()
	return
	return
}

func (a *ofp10.ActionVendorHeader) UnpackBinary(data []byte) (err error) {
	if len(data) != int(a.Len()) {
		return errors.New("The []byte the wrong size to Unpack an " +
			"ofp10.ActionVendor message.")
	}
	buf := bytes.NewBuffer(data)
	hs := make([]byte, a.Header.Len())
	binary.Read(buf, binary.BigEndian, hs)
	err = a.Header.UnpackBinary(hs)
	if err != nil {
		return
	}
	binary.Read(buf, binary.BigEndian, &a.Vendor)
	return
}

func EncodeAction(a ofp10.Action) (data []byte, err error) {
	data = make([]byte, 0)
	switch a.(type) {
	case *ofp10.ActionOutput:
		data, err = a.(*ofp10.ActionOutput).PackBinary()
	case *ofp10.ActionEnqueue:
		data, err = a.(*ofp10.ActionEnqueue).PackBinary()
	case *ofp10.ActionVLANVID:
		data, err = a.(*ofp10.ActionVLANVID).PackBinary()
	case *ofp10.ActionVLANPCP:
		data, err = a.(*ofp10.ActionVLANPCP).PackBinary()
	case *ofp10.ActionDLAddr:
		data, err = a.(*ofp10.ActionDLAddr).PackBinary()
	case *ofp10.ActionNWAddr:
		data, err = a.(*ofp10.ActionNWAddr).PackBinary()
	case *ofp10.ActionNWTOS:
		data, err = a.(*ofp10.ActionNWTOS).PackBinary()
	case *ofp10.ActionTPPort:
		data, err = a.(*ofp10.ActionTPPort).PackBinary()
	}
	if err != nil {
		return
	}
	return
}

func DecodeAction(data []byte) (a ofp10.Action, err error) {
	buf := bytes.NewBuffer(data)
	switch binary.BigEndian.Uint16(data[:]) {
	case ofp10.OFPATOutput:
		a = new(ofp10.ActionOutput)
		r := a.(*ofp10.ActionOutput)
		as := make([]byte, r.Len())
		binary.Read(buf, binary.BigEndian, as)
		err = r.UnpackBinary(as)
		if err != nil {
			return
		}
	case ofp10.OFPATSetVLANVID:
		a = new(ofp10.ActionVLANVID)
		r := a.(*ofp10.ActionVLANVID)
		as := make([]byte, r.Len())
		binary.Read(buf, binary.BigEndian, as)
		err = r.UnpackBinary(as)
		if err != nil {
			return
		}
	case ofp10.OFPATSetVLANPCP:
		a = new(ofp10.ActionVLANPCP)
		r := a.(*ofp10.ActionVLANPCP)
		as := make([]byte, r.Len())
		binary.Read(buf, binary.BigEndian, as)
		err = r.UnpackBinary(as)
		if err != nil {
			return
		}
	case ofp10.OFPATStripVLAN:
		a = new(ofp10.ActionHeader)
		r := a.(*ofp10.ActionHeader)
		as := make([]byte, r.Len())
		binary.Read(buf, binary.BigEndian, as)
		err = r.UnpackBinary(as)
		if err != nil {
			return
		}
	case ofp10.OFPATSetDLSrc:
		a = new(ofp10.ActionDLAddr)
		r := a.(*ofp10.ActionDLAddr)
		as := make([]byte, r.Len())
		binary.Read(buf, binary.BigEndian, as)
		err = r.UnpackBinary(as)
		if err != nil {
			return
		}
	case ofp10.OFPATSetDLDst:
		a = new(ofp10.ActionDLAddr)
		r := a.(*ofp10.ActionDLAddr)
		as := make([]byte, r.Len())
		binary.Read(buf, binary.BigEndian, as)
		err = r.UnpackBinary(as)
		if err != nil {
			return
		}
	case ofp10.OFPATSetNWSrc:
		a = new(ofp10.ActionNWAddr)
		r := a.(*ofp10.ActionNWAddr)
		as := make([]byte, r.Len())
		binary.Read(buf, binary.BigEndian, as)
		err = r.UnpackBinary(as)
		if err != nil {
			return
		}
	case ofp10.OFPATSetNWDst:
		a = new(ofp10.ActionNWAddr)
		r := a.(*ofp10.ActionNWAddr)
		as := make([]byte, r.Len())
		binary.Read(buf, binary.BigEndian, as)
		err = r.UnpackBinary(as)
		if err != nil {
			return
		}
	case ofp10.OFPATSetNWTos:
		a = new(ofp10.ActionNWTOS)
		r := a.(*ofp10.ActionNWTOS)
		as := make([]byte, r.Len())
		binary.Read(buf, binary.BigEndian, as)
		err = r.UnpackBinary(as)
		if err != nil {
			return
		}
	case ofp10.OFPATSetTPSrc:
		a = new(ofp10.ActionTPPort)
		r := a.(*ofp10.ActionTPPort)
		as := make([]byte, r.Len())
		binary.Read(buf, binary.BigEndian, as)
		err = r.UnpackBinary(as)
		if err != nil {
			return
		}
	case ofp10.OFPATSetTPDst:
		a = new(ofp10.ActionTPPort)
		r := a.(*ofp10.ActionTPPort)
		as := make([]byte, r.Len())
		binary.Read(buf, binary.BigEndian, as)
		err = r.UnpackBinary(as)
		if err != nil {
			return
		}
	case ofp10.OFPATEnqueue:
		a = new(ofp10.ActionEnqueue)
		r := a.(*ofp10.ActionEnqueue)
		as := make([]byte, r.Len())
		binary.Read(buf, binary.BigEndian, as)
		err = r.UnpackBinary(as)
		if err != nil {
			return
		}
	case ofp10.OFPATVendor:
		a = new(ofp10.ActionVendorHeader)
		r := a.(*ofp10.ActionVendorHeader)
		as := make([]byte, r.Len())
		binary.Read(buf, binary.BigEndian, as)
		err = r.UnpackBinary(as)
		if err != nil {
			return
		}
	}
	return
}
