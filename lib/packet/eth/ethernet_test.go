package eth

import (
	"bytes"
	"encoding/hex"
	"net"
	"strings"
	"testing"
)

func TestEthPackBinary(t *testing.T) {
	b := "   0a b0 0c 0d e0 0f " + // HWDst
		"00 00 00 00 00 ff " + // HWSrc
		"08 00 " // Ethertype
	b = strings.Replace(b, " ", "", -1)

	e := New()
	e.HWDst, _ = net.ParseMAC("0a:b0:0c:0d:e0:0f")
	e.HWSrc, _ = net.ParseMAC("00:00:00:00:00:ff")
	data, _ := e.PackBinary()
	d := hex.EncodeToString(data)
	if (len(b) != len(d)) || (b != d) {
		t.Log("Exp:", b)
		t.Log("Rec:", d)
		t.Errorf("Received length of %d, expected %d", len(d), len(b))
	}
}

func TestEthUnpackBinary(t *testing.T) {
	b := "   00 " + // Delim
		"0a b0 0c 0d e0 0f " + // HWDst
		"00 00 00 00 00 ff " + // HWSrc
		"88 00 " // Ethertype
	b = strings.Replace(b, " ", "", -1)
	byte, _ := hex.DecodeString(b)

	a := New() // Ensure type is set correctly
	a.UnpackBinary(byte)

	dst, _ := net.ParseMAC("0a:b0:0c:0d:e0:0f")
	src, _ := net.ParseMAC("00:00:00:00:00:ff")

	if int(a.Len()) != (len(byte) - 1) {
		t.Errorf("Got length of %d, expected %d.", a.Len(), (len(byte) - 1))
	} else if a.Ethertype != 0x8800 {
		t.Errorf("Got type %d, expected %d.", a.Ethertype, 0x0880)
	} else if bytes.Compare(net.HardwareAddr(a.HWDst), net.HardwareAddr(dst)) != 0 {
		t.Log("Exp:", dst)
		t.Log("Rec:", a.HWDst)
		t.Errorf("Received length of %d, expected %d", len(a.HWDst), len(dst))
	} else if bytes.Compare(net.HardwareAddr(a.HWSrc), net.HardwareAddr(src)) != 0 {
		t.Log("Exp:", src)
		t.Log("Rec:", a.HWSrc)
		t.Errorf("Received length of %d, expected %d", len(a.HWSrc), len(src))
	}
}
