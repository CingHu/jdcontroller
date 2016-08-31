package ofpctl13

import (
	"net"
	"strings"
	"strconv"

	"jd.com/jdcontroller/protocol/ofp13"
)

func (setfield *ofp13.ActionSetField) SetFieldInit(ofpoxmclass uint32, fieldxm uint8, length uint8) {
	setfield.Field = append(setfield.Field, uint8(ofpoxmclass >> 8))
	setfield.Field = append(setfield.Field, uint8(ofpoxmclass & 0xff))

	// oxm_field(7bit) + oxm_hasmask(1bit)
	setfield.Field = append(setfield.Field, (fieldxm << 1) | 0) //no mask
	// oxm_length
	setfield.Field = append(setfield.Field, length)

	return
}

func (i *ofp13.InstructionActions) SetInPortField() {

}

func(i *ofp13.InstructionActions) SetPhyPortField() {

}

func(i *ofp13.InstructionActions) SetOfMetadataField() {

}

func (i *ofp13.InstructionActions) AddSetEthSrcField(EthSrc string) {
	a := i.Newofp13.ActionSetField()
	a.SetFieldInit(uint32(ofp13.OFPXMCOpenFlowBasic), uint8(ofp13.OFPXMTFOFBEthSrc), uint8(ofp13.OFPXMTFOFBEthLEN))
	ethsrcfield, _ := net.ParseMAC(EthSrc)
	ethsrc := []byte(ethsrcfield)
	a.Field = append(a.Field, ethsrc...)
	a.Header.Length += (ofp13.OFPXMHEADERLEN + ofp13.OFPXMTFOFBEthLEN)

	i.AddSetField(a)
	return
}

func(i *ofp13.InstructionActions) AddSetEthDstField(EthDst string) {
	a := i.Newofp13.ActionSetField()
	a.SetFieldInit(uint32(ofp13.OFPXMCOpenFlowBasic), uint8(ofp13.OFPXMTFOFBEthDst), uint8(ofp13.OFPXMTFOFBEthLEN))
	ethdstfield, _ := net.ParseMAC(EthDst)
	ethdst := []byte(ethdstfield)
	a.Field = append(a.Field, ethdst...)
	a.Header.Length += (ofp13.OFPXMHEADERLEN + ofp13.OFPXMTFOFBEthLEN)

	i.AddSetField(a)
	return
}

func(i *ofp13.InstructionActions)SetEthTypeField(){

	return
}

func(i *ofp13.InstructionActions)SetVlanVidField() {

	return
}

func(i *ofp13.InstructionActions)SetVlanPcpField() {

	return
}

func(i *ofp13.InstructionActions)SetIpDscpField() {

	return
}

func(i *ofp13.InstructionActions)SetIpScnField() {

	return
}

func(i *ofp13.InstructionActions)AddSetIpSrcField(ipaddr string) {
	a := i.Newofp13.ActionSetField()
	var ipv4src []uint8
	a.SetFieldInit(uint32(ofp13.OFPXMCOpenFlowBasic), uint8(ofp13.OFPXMTFOFBIPv4Src), uint8(ofp13.OFPXMTFOFBIPv4LEN))

	addr := strings.Split(ipaddr, "/")
	ip := strings.Split(addr[0], ".")
	for _, ipItem := range(ip) {
		ipItemNum, _ := strconv.Atoi(ipItem)
		ipv4src = append(ipv4src, uint8(ipItemNum))
	}
	a.Header.Length += (ofp13.OFPXMHEADERLEN + ofp13.OFPXMTFOFBIPv4LEN)
	a.Field = append(a.Field, ipv4src...)
	i.AddSetField(a)
	return
}

func(i *ofp13.InstructionActions)AddSetIpDstField(ipaddr string) {
	a := i.Newofp13.ActionSetField()
	var ipv4dst []uint8
	a.SetFieldInit(uint32(ofp13.OFPXMCOpenFlowBasic), uint8(ofp13.OFPXMTFOFBIPv4Dst), uint8(ofp13.OFPXMTFOFBIPv4LEN))

	addr := strings.Split(ipaddr, "/")
	ip := strings.Split(addr[0], ".")
	for _, ipItem := range(ip) {
		ipItemNum, _ := strconv.Atoi(ipItem)
		ipv4dst = append(ipv4dst, uint8(ipItemNum))
	}
	a.Header.Length += (ofp13.OFPXMHEADERLEN + ofp13.OFPXMTFOFBIPv4LEN)
	a.Field = append(a.Field, ipv4dst...)
	i.AddSetField(a)
	return
}

func(i *ofp13.InstructionActions)SetIpProtoField() {

	return
}

func(i *ofp13.InstructionActions)SetTpSrcField() {

	return
}

func(i *ofp13.InstructionActions)SetTpDstField() {

	return
}

func(i *ofp13.InstructionActions)SetIcmpv4TypeField() {

	return
}

func(i *ofp13.InstructionActions)SetIcmpv4CodeField() {

	return
}

func(i *ofp13.InstructionActions)SetArpOpField() {

	return
}

func(i *ofp13.InstructionActions)SetArpSpaField() {

	return
}

func(i *ofp13.InstructionActions)SetArpTpaField() {

	return
}

func(i *ofp13.InstructionActions)SetArpShaField() {

	return
}

func(i *ofp13.InstructionActions)SetArpThaField()  {

	return
}

func(i *ofp13.InstructionActions)SetIpv6SrcField() {

	return
}

func(i *ofp13.InstructionActions) SetIpv6DstField() {

	return
}

func(i *ofp13.InstructionActions) SetIpv6FlabelField() {

	return
}

func(i *ofp13.InstructionActions) SetIcmpv6TypeField() {

	return
}

func(i *ofp13.InstructionActions) SetIcmpv6CodeField() {

	return
}

func(i *ofp13.InstructionActions) SetIpv6NdTargetField() {

	return
}

func(i *ofp13.InstructionActions) SetIpv6NdSllField() {

	return
}

func(i *ofp13.InstructionActions) SetIpv6NdTllField() {

	return
}

func(i *ofp13.InstructionActions) SetMplsLableField() {

	return
}

func(i *ofp13.InstructionActions) SetMplsTcField() {

	return
}

func(i *ofp13.InstructionActions) SetMplsBosField() {

	return
}

func(i *ofp13.InstructionActions) SetPbbIsidField() {

	return
}

func(i *ofp13.InstructionActions) SetTunnelIdField() {

	return
}

func(i *ofp13.InstructionActions) SetIpv6ExthdrField() {

	return
}

func(i *ofp13.InstructionActions) AddSetTunnelIpSrcField(ipaddr string) {
	a := i.Newofp13.ActionSetField()
	var ipv4src []uint8
	a.SetFieldInit(uint32(ofp13.OFPXMCNXM1), uint8(NXMNXTUNIPV4SRC), uint8(ofp13.NXMNXTUNIPV4LEN))

	addr := strings.Split(ipaddr, "/")
	ip := strings.Split(addr[0], ".")
	for _, ipItem := range(ip) {
		ipItemNum, _ := strconv.Atoi(ipItem)
		ipv4src = append(ipv4src, uint8(ipItemNum))
	}
	a.Header.Length += (ofp13.OFPXMHEADERLEN + ofp13.NXMNXTUNIPV4LEN)
	a.Field = append(a.Field, ipv4src...)
	i.AddSetField(a)
	return
}

func(i *ofp13.InstructionActions) AddSetTunnelIpDstField(ipaddr string) {
	a := i.Newofp13.ActionSetField()
	var ipv4dst []uint8
	a.SetFieldInit(uint32(ofp13.OFPXMCNXM1), uint8(NXMNXTUNIPV4DST), uint8(ofp13.NXMNXTUNIPV4LEN))

	addr := strings.Split(ipaddr, "/")
	ip := strings.Split(addr[0], ".")
	for _, ipItem := range(ip) {
		ipItemNum, _ := strconv.Atoi(ipItem)
		ipv4dst = append(ipv4dst, uint8(ipItemNum))
	}
	a.Header.Length += (ofp13.OFPXMHEADERLEN + ofp13.NXMNXTUNIPV4LEN)
	a.Field = append(a.Field, ipv4dst...)
	i.AddSetField(a)
	return
}

func(i *ofp13.InstructionActions) NewActionSetField() (a *ofp13.ActionSetField) {
	a = new(ofp13.ActionSetField)
	a.Header.Type = uint16(OFPATSetField)
	a.Header.Length = 4
	return
}

func(i *ofp13.InstructionActions) AddSetField(setfieldact *ofp13.ActionSetField) {
	i.Header.Length += setfieldact.Header.Length
	//padLen := 8 - int(i.Header.Length) % 8
	padLen := ((setfieldact.Header.Length + 7) / 8 * 8) - setfieldact.Header.Length
	i.Header.Length += uint16(padLen)
	setfieldact.Header.Length += uint16(padLen)

	for i := 0; i < int(padLen); i++ {
		setfieldact.Field = append(setfieldact.Field, uint8(0))
	}

	i.Actions = append(i.Actions, setfieldact)

	return
}
