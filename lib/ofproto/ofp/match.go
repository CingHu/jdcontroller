package ofp

import (
	"fmt"
	"net"
	"strings"
	"strconv"
	"jd.com/jdcontroller/lib/packet/eth"
	"jd.com/jdcontroller/lib/ofproto/ofp13"
)

/*Match.Type,Match.Length and TLV header's bytes is always 8*/
const OXMHEADERLEN = 8

type PacketField interface {
	NewEthDstField(EthDst string,  EthDstMask string) (field *ofp13.MatchField)
	NewEthSrcField(EthSrc string, EthSrcMask string) (field *ofp13.MatchField)
	NewIpSrcField(cidr string) (field *ofp13.MatchField)
	NewIpDstField(cidr string) (field *ofp13.MatchField)
}

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

func SetMacField(XMFields []uint8, mac []byte) ([]uint8){
	for i := 0; i < 6; i++ {
		XMFields = append(XMFields, mac[i])
	}
	return XMFields
}

func (m MatchCreator) NewEthDstField(EthDst string,  EthDstMask string) (field *ofp13.MatchField) {

	/*mac地址参数转换*/
	field = new(ofp13.MatchField)
	ethdstfield, _ := net.ParseMAC(EthDst)
	ethdstmaskfield, _ := net.ParseMAC(EthDstMask)
	ethdst := []byte(ethdstfield)
	ethdstmask := []byte(ethdstmaskfield)

	// oxm_class (用两个uint8组成ofp13.OFPXMCOpenFlowBasic（0x8000）)
	field.XMFields = append(field.XMFields, 1<<7)
	field.XMFields = append(field.XMFields, 0)
	if !eth.Ethaddriszero(ethdstmask) {
		if eth.Ethmaskisexact(ethdstmask) {
			// oxm_field(7bit) + oxm_hasmask(1bit)
			field.XMFields = append(field.XMFields, ofp13.OFPXMTFOFBEthDst<<1 | 0)
			// oxm_length
			field.XMFields = append(field.XMFields, 6)
			field.XMFields = SetMacField(field.XMFields, ethdst)
		} else {
			// oxm_field(7bit) + oxm_hasmask(1bit)
			field.XMFields = append(field.XMFields, ofp13.OFPXMTFOFBEthDst<<1 | 1)
			// oxm_length
			field.XMFields = append(field.XMFields, 6 * 2)
			field.XMFields = SetMacField(field.XMFields, ethdst)
			field.XMFields = SetMacField(field.XMFields, ethdstmask)
		}
	}

	return
}

func (match MatchCreator) NewEthSrcField(EthSrc string, EthSrcMask string) (field *ofp13.MatchField){

	/*mac地址参数转换*/
	field = new(ofp13.MatchField)
	ethsrcfield, _ := net.ParseMAC(EthSrc)
	ethsrcmaskfield, _ := net.ParseMAC(EthSrcMask)
	ethsrc := []byte(ethsrcfield)
	ethsrcmask := []byte(ethsrcmaskfield)

	// oxm_class (用两个uint8组成ofp13.OFPXMCOpenFlowBasic（0x8000）)
	field.XMFields = append(field.XMFields, 1<<7)
	field.XMFields = append(field.XMFields, 0)
	if !eth.Ethaddriszero(ethsrcmask) {
		if eth.Ethmaskisexact(ethsrcmask) {
			// oxm_field(7bit) + oxm_hasmask(1bit)
			field.XMFields = append(field.XMFields, ofp13.OFPXMTFOFBEthDst<<1 | 0)
			// oxm_length
			field.XMFields = append(field.XMFields, 6)
			field.XMFields = SetMacField(field.XMFields, ethsrc)
		} else {
			// oxm_field(7bit) + oxm_hasmask(1bit)
			field.XMFields = append(field.XMFields, ofp13.OFPXMTFOFBEthDst<<1 | 1)
			// oxm_length
			field.XMFields = append(field.XMFields, 6 * 2)
			field.XMFields = SetMacField(field.XMFields, ethsrc)
			field.XMFields = SetMacField(field.XMFields, ethsrcmask)
		}
	}

	return
}

func(match MatchCreator)NewEthTypeField() (field *ofp13.MatchField){

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
	// oxm_class (用两个uint8组成ofp13.OFPXMCOpenFlowBasic（0x8000）)
	field.XMFields = append(field.XMFields, 1<<7)
	field.XMFields = append(field.XMFields, 0)

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
		// oxm_field(7bit) + oxm_hasmask(1bit)
		field.XMFields = append(field.XMFields, ofp13.OFPXMTFOFBIPv4Src<<1 | 0)
		// oxm_length
		field.XMFields = append(field.XMFields, 4)
		field.XMFields = append(field.XMFields, ipaddr...)
	} else {
		// oxm_field(7bit) + oxm_hasmask(1bit)
		field.XMFields = append(field.XMFields, ofp13.OFPXMTFOFBIPv4Src<<1 | 1)
		// oxm_length
		field.XMFields = append(field.XMFields, 4 * 2)
		field.XMFields = append(field.XMFields, ipaddr...)
		field.XMFields = append(field.XMFields, maskaddr...)
	}

	return
}

func(match MatchCreator)NewIpDstField(cidr string) (field *ofp13.MatchField){
	field = new(ofp13.MatchField)
	// oxm_class (用两个uint8组成ofp13.OFPXMCOpenFlowBasic（0x8000）)
	field.XMFields = append(field.XMFields, 1<<7)
	field.XMFields = append(field.XMFields, 0)

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
		// oxm_field(7bit) + oxm_hasmask(1bit)
		field.XMFields = append(field.XMFields, ofp13.OFPXMTFOFBIPv4Dst<<1 | 0)
		// oxm_length
		field.XMFields = append(field.XMFields, 4)
		field.XMFields = append(field.XMFields, ipaddr...)
	} else {
		// oxm_field(7bit) + oxm_hasmask(1bit)
		field.XMFields = append(field.XMFields, ofp13.OFPXMTFOFBIPv4Dst<<1 | 1)
		// oxm_length
		field.XMFields = append(field.XMFields, 4 * 2)
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
