package ofpctl13

import (
	"fmt"
	"net"
	"strings"
	"strconv"

	"jd.com/jdcontroller/protocol/ofp13"
)

type MatchCreator struct {

}

func (match MatchCreator) NewInPortField()(m *ofp13.MatchField) {

	return
}

func(match MatchCreator) NewPhyPortField() (m *ofp13.MatchField){

	return
}

func(match MatchCreator) NewOfMetadataField() (m *ofp13.MatchField){

	return
}

func (m MatchCreator) NewEthDstField(EthDst string) (field *ofp13.MatchField) {
	field = new(ofp13.MatchField)
	field.Init(uint8(OFPXMTFOFBEthDst), uint8(OFPXMTFOFBEthLEN), uint8(OFPXMNOMASK))
	ethdstfield, _ := net.ParseMAC(EthDst)
	ethdst := []byte(ethdstfield)
	field.XMFields = append(field.XMFields, ethdst...)
	return
}

func (match MatchCreator) NewEthSrcField(EthSrc string) (field *ofp13.MatchField){
	field = new(ofp13.MatchField)
	field.Init(uint8(OFPXMTFOFBEthDst), uint8(OFPXMTFOFBEthLEN), uint8(OFPXMNOMASK))
	ethsrcfield, _ := net.ParseMAC(EthSrc)
	ethsrc := []byte(ethsrcfield)

	field.XMFields = append(field.XMFields, ethsrc...)

	return
}

func(match MatchCreator)NewEthTypeField(ethtype uint32) (field *ofp13.MatchField){
	field = new(ofp13.MatchField)
	field.Init(uint8(OFPXMTFOFBEthType), uint8(OFPXMTFOFBIPProtoLEN), uint8(OFPXMNOMASK))

	field.XMFields = append(field.XMFields, uint8(ethtype >> 8))
	field.XMFields = append(field.XMFields, uint8(ethtype & 0x00ff))

	return
}

func(match MatchCreator)NewVlanVidField() (field *ofp13.MatchField){

	return
}

func(match MatchCreator)NewVlanPcpField() (field *ofp13.MatchField){

	return
}

func(match MatchCreator)NewIpDscpField() (field *ofp13.MatchField){

	return
}

func(match MatchCreator)NewIpScnField() (field *ofp13.MatchField){

	return
}

func MaskToAddr(mask int) (maskaddr []uint8) {
	i := 31
	v := 0
	k := 0

	for j :=  0; j < mask; j++ {
		v = 1 << uint(i)
		k |= v
		i--
	}

	var n int
	for n = 24; n >= 0; n -= 8 {
		v = (k & (0xff << uint(n))) >> uint(n)
		maskaddr = append(maskaddr, uint8(v))
		fmt.Println(n)
	}

	return
}

/* cidr 192.168.1.1/24 */
func(match MatchCreator)NewIpSrcField(cidr string) (field *ofp13.MatchField){
	field = new(ofp13.MatchField)

	addr := strings.Split(cidr, "/")
	var mask int
	var ipaddr []uint8
	ip := strings.Split(addr[0], ".")
	for _, ipItem := range(ip) {
		ipItemNum, _ := strconv.Atoi(ipItem)
		ipaddr = append(ipaddr, uint8(ipItemNum))
	}

	if len(addr) > 1 {
		mask, _ = strconv.Atoi(string(addr[1]))
		if mask == 32 {
			mask = 0
		}
	}
	maskaddr := MaskToAddr(mask)

	if mask == 0 {
		field.Init(uint8(OFPXMTFOFBIPv4Src), uint8(OFPXMTFOFBIPv4LEN), uint8(OFPXMNOMASK))
		field.XMFields = append(field.XMFields, ipaddr...)
	} else {
		field.Init(uint8(OFPXMTFOFBIPv4Src), uint8(OFPXMTFOFBIPv4LEN) * 2, uint8(OFPXMHASMASK))
		field.XMFields = append(field.XMFields, ipaddr...)
		field.XMFields = append(field.XMFields, maskaddr...)
	}

	return
}

func(match MatchCreator)NewIpDstField(cidr string) (field *ofp13.MatchField){
	field = new(ofp13.MatchField)

	addr := strings.Split(cidr, "/")
	var mask int
	var ipaddr []uint8
	ip := strings.Split(addr[0], ".")
	for _, ipItem := range(ip) {
		ipItemNum, _ := strconv.Atoi(ipItem)
		ipaddr = append(ipaddr, uint8(ipItemNum))
	}

	if len(addr) > 1 {
		mask, _ = strconv.Atoi(string(addr[1]))
		if mask == 32 {
			mask = 0
		}
	}
	maskaddr := MaskToAddr(mask)

	if mask == 0 {
		field.Init(uint8(OFPXMTFOFBIPv4Dst), uint8(OFPXMTFOFBIPv4LEN), uint8(OFPXMNOMASK))
		field.XMFields = append(field.XMFields, ipaddr...)
	} else {
		field.Init(uint8(OFPXMTFOFBIPv4Dst), uint8(OFPXMTFOFBIPv4LEN) * 2, uint8(OFPXMHASMASK))
		field.XMFields = append(field.XMFields, ipaddr...)
		field.XMFields = append(field.XMFields, maskaddr...)
	}

	return
}

func(match MatchCreator)NewIpProtoField() (field *ofp13.MatchField){

	return
}

func(match MatchCreator)NewTpSrcField() (field *ofp13.MatchField){

	return
}

func(match MatchCreator)NewTpDstField() (field *ofp13.MatchField){

	return
}

func(match MatchCreator)NewIcmpv4TypeField() (field *ofp13.MatchField){

	return
}

func(match MatchCreator)NewIcmpv4CodeField() (field *ofp13.MatchField){

	return
}

func(match MatchCreator)NewArpOpField() (field *ofp13.MatchField){

	return
}

func(match MatchCreator)NewArpSpaField() (field *ofp13.MatchField){

	return
}

func(match MatchCreator)NewArpTpaField() (field *ofp13.MatchField){

	return
}

func(match MatchCreator)NewArpShaField() (field *ofp13.MatchField){

	return
}

func(match MatchCreator)NewArpThaField() (field *ofp13.MatchField) {

	return
}

func(match MatchCreator)NewIpv6SrcField() (field *ofp13.MatchField){

	return
}

func(match MatchCreator)NewIpv6DstField() (field *ofp13.MatchField){

	return
}

func(match MatchCreator)NewIpv6FlabelField() (field *ofp13.MatchField){

	return
}

func(match MatchCreator)NewIcmpv6TypeField() (field *ofp13.MatchField){

	return
}

func(match MatchCreator)NewIcmpv6CodeField() (field *ofp13.MatchField){

	return
}

func(match MatchCreator)NewIpv6NdTargetField() (field *ofp13.MatchField){

	return
}

func(match MatchCreator)NewIpv6NdSllField() (field *ofp13.MatchField){

	return
}

func(match MatchCreator)NewIpv6NdTllField() (field *ofp13.MatchField){

	return
}

func(match MatchCreator)NewMplsLableField() (field *ofp13.MatchField){

	return
}

func(match MatchCreator)NewMplsTcField() (field *ofp13.MatchField){

	return
}

func(match MatchCreator)NewMplsBosField() (field *ofp13.MatchField){

	return
}

func(match MatchCreator)NewPbbIsidField() (field *ofp13.MatchField){

	return
}

func(match MatchCreator)NewTunnelIdField() (field *ofp13.MatchField){

	return
}

func(match MatchCreator)NewIpv6ExthdrField() (field *ofp13.MatchField){

	return
}
