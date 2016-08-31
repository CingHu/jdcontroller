package dhcp

import (
	"fmt"
	"jd.com/jdcontroller/lib/packet/eth"
	"jd.com/jdcontroller/lib/packet/ipv4"
	"jd.com/jdcontroller/lib/packet/udp"
//	"bytes"
//	"encoding/binary"
//	"errors"
//	"io"
//	"math/rand"
//	"net"
)

func IfDhcpPkt (frame eth.Ethernet) (bool) {
	if eth.GetMacAddr(frame.HWDst) ==  "ff:ff:ff:ff:ff:ff" && frame.Ethertype == eth.IPv4Msg{
		ipv4msg := frame.Data.(*ipv4.IPv4)

		if ipv4msg.Protocol == ipv4.UDP {
			fmt.Println("ipv4data:", ipv4msg)
			udpmsg := ipv4msg.Data.(*udp.UDP)

			if udpmsg.PortDst == udp.DHCP {
				fmt.Println("udpdata:", udpmsg.Data)
				return true
			}
		}
	}
	return false
}

//
//const (
//	DHCPMsgBootReq byte = iota
//	DHCPMsgBootRes
//)
//
//type DHCPOperation uint8
//
//const (
//	DHCPMsgUnspec DHCPOperation = iota
//	DHCPMsgDiscover
//	DHCPMsgOffer
//	DHCPMsgRequest
//	DHCPMsgDecline
//	DHCPMsgAck
//	DHCPMsgNak
//	DHCPMsgRelease
//	DHCPMsgInform
//)
//
//var dhcpMagic uint32 = 0x63825363
//
//type DHCP struct {
//	Operation    DHCPOperation
//	HardwareType uint8
//	HardwareLen  uint8
//	HardwareOpts uint8
//	Xid          uint32
//	Secs         uint16
//	Flags        uint16
//	ClientIP     [4]byte
//	YourIP       [4]byte
//	ServerIP     [4]byte
//	GatewayIP    [4]byte
//	ClientHWAddr [6]byte
//	ServerName   [64]byte
//	File         [128]byte
//	Options      []DHCPOption
//}
//
//const (
//	DHCPOptRequestIP     byte = iota + 50 // 0x32, 4, net.IP
//	DHCPOptLeaseTime                      // 0x33, 4, uint32
//	DHCPOptExtOpts                        // 0x34, 1, 1/2/3
//	DHCPOptMessageType                    // 0x35, 1, 1-7
//	DHCPOptServerID                       // 0x36, 4, net.IP
//	DHCPOptParamsRequest                  // 0x37, n, []byte
//	DHCPOptMessage                        // 0x38, n, string
//	DHCPOptMaxDHCPSize                    // 0x39, 2, uint16
//	DHCPOptT1                             // 0x3a, 4, uint32
//	DHCPOptT2                             // 0x3b, 4, uint32
//	DHCPOptClassID                        // 0x3c, n, []byte
//	DHCPOptClientID                       // 0x3d, n >=  2, []byte
//
//)
//
//const (
//	DHCPHWEthernet byte = 0x01
//)
//
//const (
//	DHCPFlagBroadcast uint16 = 0x80
//
////	FLAG_BROADCAST_MASK uint16 = (1 << FLAG_BROADCAST)
//)
//
//func New(xid uint32, op DHCPOperation, hwtype byte) (d *DHCP, err error) {
//	if xid == 0 {
//		xid = rand.Uint32()
//	}
//	switch hwtype {
//	case DHCPHWEthernet:
//		break
//	default:
//		err = errors.New("Bad HardwareType")
//		return
//	}
//	d := &DHCP{
//		Operation:    op,
//		HardwareType: hwtype,
//		Xid:          xid,
//	}
//	return
//}
//
//func (d *DHCP) Len() (l int) {
//	n := 0
//	n += uint16(240)
//	optend := false
//	for _, opt := range d.Options {
//		n += opt.Len()
//		if opt.OptionType() == DHCPOptEnd {
//			optend = true
//		}
//	}
//	if !optend {
//		n += 1
//	}
//	return
//}
//
//func (d *DHCP) PackBinary() (data []byte, err error) {
//	buf := bytes.NewBuffer(make([]byte, 0))
//	binary.Write(buf, binary.BigEndian, d.Operation)
//	binary.Write(buf, binary.BigEndian, d.HardwareType)
//	binary.Write(buf, binary.BigEndian, d.HardwareLen)
//	binary.Write(buf, binary.BigEndian, d.HardwareOpts)
//	binary.Write(buf, binary.BigEndian, d.Xid)
//	binary.Write(buf, binary.BigEndian, d.Secs)
//	binary.Write(buf, binary.BigEndian, d.Flags)
//	binary.Write(buf, binary.BigEndian, d.ClientIP)
//	binary.Write(buf, binary.BigEndian, d.YourIP)
//	binary.Write(buf, binary.BigEndian, d.ServerIP)
//	binary.Write(buf, binary.BigEndian, d.GatewayIP)
//	binary.Write(buf, binary.BigEndian, d.ClientHWAddr)
//	binary.Write(buf, binary.BigEndian, d.ServerName)
//	binary.Write(buf, binary.BigEndian, d.File)
////	binary.Write(buf, binary.BigEndian, dhcpMagic)
//
//	for _, opt := range d.Options {
//		m, err := DHCPWriteOption(buf, opt)
//		n += m
//		if err != nil {
//			return n, err
//		}
//		if opt.OptionType() == DHCPOptEnd {
//			optend = true
//		}
//	}
//	if !optend {
//		m, err := DHCPWriteOption(buf, DHCPNewOption(DHCPOptEnd, nil))
//		n += m
//		if err != nil {
//			return n, err
//		}
//	}
//	if n, err = buf.Read(b); n == 0 {
//		return
//	}
//	return n, nil
//}
//
//func (d *DHCP) UnpackBinary(data []byte) (err error) {
//	if len(data) < 240 {
//		err = errors.New("ErrTruncated")
//		return
//	}
//	buf := bytes.NewBuffer(data)
//
//	binary.Read(buf, binary.BigEndian, &d.Operation)
//	binary.Read(buf, binary.BigEndian, &d.HardwareType)
//	binary.Read(buf, binary.BigEndian, &d.HardwareLen)
//	binary.Read(buf, binary.BigEndian, &d.HardwareOpts)
//	binary.Read(buf, binary.BigEndian, &d.Xid)
//	binary.Read(buf, binary.BigEndian, &d.Secs)
//	binary.Read(buf, binary.BigEndian, &d.Flags)
//	binary.Read(buf, binary.BigEndian, &d.ClientIP)
//	binary.Read(buf, binary.BigEndian, &d.YourIP)
//	binary.Read(buf, binary.BigEndian, &d.ServerIP)
//	binary.Read(buf, binary.BigEndian, &d.GatewayIP)
//	binary.Read(buf, binary.BigEndian, &d.ClientHWAddr)
//	binary.Read(buf, binary.BigEndian, &d.ServerName)
//	binary.Read(buf, binary.BigEndian, &d.File)
////	var magic uint32
////	binary.Read(buf, binary.BigEndian, &magic)
////	if magic != dhcpMagic {
////		return n, errors.New("Bad DHCP header")
////	}
//
//	n := d.Len()
//	for n < len(data) {
//		bs := make([]byte, 0)
//
//	}
//	optlen := buf.Len()
//	opts := make([]byte, optlen)
//	if err = binary.Read(buf, binary.BigEndian, &opts); err != nil {
//		return
//	}
//	n += optlen
//
//	if d.Options, err = DHCPParseOptions(opts); err != nil {
//		return
//	}
//
//	return
//}
//
//// Standard options (RFC1533)
//const (
//	DHCPOptPad                   byte = iota
//	DHCPOptSubnetMask                 // 0x01, 4, net.IP
//	DHCPOptTimeOffset                 // 0x02, 4, int32 (signed seconds from UTC)
//	DHCPOptDefaultGateway             // 0x03, n*4, [n]net.IP
//	DHCPOptTimeServer                 // 0x04, n*4, [n]net.IP
//	DHCPOptNameServer                 // 0x05, n*4, [n]net.IP
//	DHCPOptDomainNameServers          // 0x06, n*4, [n]net.IP
//	DHCPOptLogServer                  // 0x07, n*4, [n]net.IP
//	DHCPOptCookieServer               // 0x08, n*4, [n]net.IP
//	DHCPOptLPRServer                  // 0x09, n*4, [n]net.IP
//	DHCPOptImpressServer              // 0x0a, n*4, [n]net.IP
//	DHCPOptRlserver                   // 0x0b, n*4, [n]net.IP
//	DHCPOptHostName                   // 0x0c, n, string
//	DHCPOptBootfileSize               // 0x0d, 2, uint16
//	DHCPOptMeritDumpFile              // 0x0e, >1, string
//	DHCPOptDomainName                 // 0x0f, n, string
//	DHCPOptSwapServer                 // 0x10, n*4, [n]net.IP
//	DHCPOptRootPath                   // 0x11, n, string
//	DHCPOptExtensionsPath             // 0x12, n, string
//	DHCPOptIPForwarding               // 0x13, 1, bool
//	DHCPOptSourceRouting              // 0x14, 1, bool
//	DHCPOptPolicyFilter               // 0x15, 8*n, [n]{net.IP/net.IP}
//	DHCPOptDGramMTU                   // 0x16, 2, uint16
//	DHCPOptDefaultTTL                 // 0x17, 1, byte
//	DHCPOptPathMTUAgingTimeout        // 0x18, 4, uint32
//	DHCPOptPathPlateuTableOption      // 0x19, 2*n, []uint16
//	DHCPOptInterfaceMTU               //0x1a, 2, uint16
//	DHCPOptAllSubsLocal               // 0x1b, 1, bool
//	DHCPOptBroadcastAddr              // 0x1c, 4, net.IP
//	DHCPOptMaskDiscovery              // 0x1d, 1, bool
//	DHCPOptMaskSupplier               // 0x1e, 1, bool
//	DHCPOptRouterDiscovery            // 0x1f, 1, bool
//	DHCPOptRouterSolicitAddr          // 0x20, 4, net.IP
//	DHCPOptStaticRoute                // 0x21, n*8, [n]{net.IP/net.IP} -- note the 2nd is router not mask
//	DHCPOptArpTrailers                // 0x22, 1, bool
//	DHCPOptArpTimeout                 // 0x23, 4, uint32
//	DHCPOptEthernetEncap              // 0x24, 1, bool
//	DHCPOptTcpTTL                     // 0x25,1, byte
//	DHCPOptTcpKeepaliveInt            // 0x26,4, uint32
//	DHCPOptTcpKeepaliveGarbage        // 0x27,1, bool
//	DHCPOptNisDomain                  // 0x28,n, string
//	DHCPOptNisServers                 // 0x29,4*n,  [n]net.IP
//	DHCPOptNtpServers                 // 0x2a, 4*n, [n]net.IP
//	DHCPOptVendorOpt                  // 0x2b, n, [n]byte // may be encapsulated.
//	DHCPOptNetbiosIPNS                // 0x2c, 4*n, [n]net.IP
//	DHCPOptNetbiosDDS                 // 0x2d, 4*n, [n]net.IP
//	DHCPOptNetbiosNodeType            // 0x2e, 1, magic byte
//	DHCPOptNetbiosScope               // 0x2f, n, string
//	DHCPOptXFontServer                // 0x30, n, string
//	DHCPOptXDisplayManager            // 0x31, n, string
//
//	DHCPOptSipServers byte = 0x78 // 0x78!, n, url
//	DHCPOptEnd        byte = 0xff
//)
//
//// I'm amazed that this is syntactically valid.
//// cool though.
//var DHCPOptionTypeStrings = [256]string{
//	DHCPOptPad:                   "(padding)",
//	DHCPOptSubnetMask:            "SubnetMask",
//	DHCPOptTimeOffset:            "TimeOffset",
//	DHCPOptDefaultGateway:        "DefaultGateway",
//	DHCPOptTimeServer:            "rfc868", // old time server protocol, stringified to dissuade confusion w. NTP
//	DHCPOptNameServer:            "ien116", // obscure nameserver protocol, stringified to dissuade confusion w. DNS
//	DHCPOptDomainNameServers:     "DNS",
//	DHCPOptLogServer:             "mitLCS", // MIT LCS server protocol, yada yada w. Syslog
//	DHCPOptCookieServer:          "OPT_COOKIE_SERVER",
//	DHCPOptLPRServer:             "OPT_LPR_SERVER",
//	DHCPOptImpressServer:         "OPT_IMPRESS_SERVER",
//	DHCPOptRlserver:              "OPT_RLSERVER",
//	DHCPOptHostName:              "Hostname",
//	DHCPOptBootfileSize:          "BootfileSize",
//	DHCPOptMeritDumpFile:         "OPT_MERIT_DUMP_FILE",
//	DHCPOptDomainName:            "DomainName",
//	DHCPOptSwapServer:            "OPT_SWAP_SERVER",
//	DHCPOptRootPath:              "RootPath",
//	DHCPOptExtensionsPath:        "OPT_EXTENSIONS_PATH",
//	DHCPOptIPForwarding:          "OPT_IP_FORWARDING",
//	DHCPOptSourceRouting:         "OPT_SOURCE_ROUTING",
//	DHCPOptPolicyFilter:          "OPT_POLICY_FILTER",
//	DHCPOptDGramMTU:              "OPT_DGRAM_MTU",
//	DHCPOptDefaultTTL:            "OPT_DEFAULT_TTL",
//	DHCPOptPathMTUAgingTimeout:   "OPT_PATH_MTU_AGING_TIMEOUT",
//	DHCPOptPathPlateuTableOption: "OPT_PATH_PLATEU_TABLE_OPTION",
//	DHCPOptInterfaceMTU:          "OPT_INTERFACE_MTU",
//	DHCPOptAllSubsLocal:          "OPT_ALL_SUBS_LOCAL",
//	DHCPOptBroadcastAddr:         "OPT_BROADCAST_ADDR",
//	DHCPOptMaskDiscovery:         "OPT_MASK_DISCOVERY",
//	DHCPOptMaskSupplier:          "OPT_MASK_SUPPLIER",
//	DHCPOptRouterDiscovery:       "OPT_ROUTER_DISCOVERY",
//	DHCPOptRouterSolicitAddr:     "OPT_ROUTER_SOLICIT_ADDR",
//	DHCPOptStaticRoute:           "OPT_STATIC_ROUTE",
//	DHCPOptArpTrailers:           "OPT_ARP_TRAILERS",
//	DHCPOptArpTimeout:            "OPT_ARP_TIMEOUT",
//	DHCPOptEthernetEncap:         "OPT_ETHERNET_ENCAP",
//	DHCPOptTcpTTL:                "OPT_TCP_TTL",
//	DHCPOptTcpKeepaliveInt:       "OPT_TCP_KEEPALIVE_INT",
//	DHCPOptTcpKeepaliveGarbage:   "OPT_TCP_KEEPALIVE_GARBAGE",
//	DHCPOptNisDomain:             "OPT_NIS_DOMAIN",
//	DHCPOptNisServers:            "OPT_NIS_SERVERS",
//	DHCPOptNtpServers:            "OPT_NTP_SERVERS",
//	DHCPOptVendorOpt:             "OPT_VENDOR_OPT",
//	DHCPOptNetbiosIPNS:           "OPT_NETBIOS_IPNS",
//	DHCPOptNetbiosDDS:            "OPT_NETBIOS_DDS",
//	DHCPOptNetbiosNodeType:       "OPT_NETBIOS_NODE_TYPE",
//	DHCPOptNetbiosScope:          "OPT_NETBIOS_SCOPE",
//	DHCPOptXFontServer:           "OPT_X_FONT_SERVER",
//	DHCPOptXDisplayManager:       "OPT_X_DISPLAY_MANAGER",
//	DHCPOptEnd:                   "(end)",
//	DHCPOptSipServers:            "SipServers",
//	DHCPOptRequestIP:             "RequestIP",
//	DHCPOptLeaseTime:             "LeaseTime",
//	DHCPOptExtOpts:               "ExtOpts",
//	DHCPOptMessageType:           "MessageType",
//	DHCPOptServerID:              "ServerID",
//	DHCPOptParamsRequest:         "ParamsRequest",
//	DHCPOptMessage:               "Message",
//	DHCPOptMaxDHCPSize:           "MaxDHCPSize",
//	DHCPOptT1:                    "Timer1",
//	DHCPOptT2:                    "Timer2",
//	DHCPOptClassID:               "ClassID",
//	DHCPOptClientID:              "ClientID",
//}
//
//type DHCPOption interface {
//	OptionType() byte
//	Bytes() []byte
//	Len() uint16
//}
//
//// Write an option to an io.Writer, including tag  & length
//// (if length is appropriate to the tag type).
//// Utilizes the PackOption as the underlying serializer.
//func DHCPWriteOption(w io.Writer, a DHCPOption) (n int, err error) {
//	out, err := DHCPPackOption(a)
//	if err == nil {
//		n, err = w.Write(out)
//	}
//	return
//}
//
//type dhcpoption struct {
//	tag  byte
//	data []byte
//}
//
//// A more json.Pack like version of WriteOption.
//func DHCPPackOption(o DHCPOption) (out []byte, err error) {
//	switch o.OptionType() {
//	case DHCPOptPad, DHCPOptEnd:
//		out = []byte{o.OptionType()}
//	default:
//		dlen := len(o.Bytes())
//		if dlen > 253 {
//			err = errors.New("Data too long to Pack")
//		} else {
//			out = make([]byte, dlen+2)
//			out[0], out[1] = o.OptionType(), byte(dlen)
//			copy(out[2:], o.Bytes())
//		}
//	}
//	return
//}
//
//func (self dhcpoption) Len() uint16      { return uint16(len(self.data) + 2) }
//func (self dhcpoption) Bytes() []byte    { return self.data }
//func (self dhcpoption) OptionType() byte { return self.tag }
//
//func DHCPNewOption(tag byte, data []byte) DHCPOption {
//	return &dhcpoption{tag: tag, data: data}
//}
//
//// NB: We don't validate that you have /any/ IP's in the option here,
//// simply that if you do that they're valid. Most DHCP options are only
//// valid with 1(+|) values
//func DHCPIP4sOption(tag byte, ips []net.IP) (opt DHCPOption, err error) {
//	var out []byte = make([]byte, 4*len(ips))
//	for i := range ips {
//		ip := ips[i].To4()
//		if ip == nil {
//			err = errors.New("ip is not a valid IPv4 address")
//		} else {
//			copy(out[i*4:], []byte(ip))
//		}
//		if err != nil {
//			break
//		}
//	}
//	opt = DHCPNewOption(tag, out)
//	return
//}
//
//// NB: We don't validate that you have /any/ IP's in the option here,
//// simply that if you do that they're valid. Most DHCP options are only
//// valid with 1(+|) values
//func DHCPIP4Option(tag byte, ips net.IP) (opt DHCPOption, err error) {
//	ips = ips.To4()
//	if ips == nil {
//		err = errors.New("ip is not a valid IPv4 address")
//		return
//	}
//	opt = DHCPNewOption(tag, []byte(ips))
//	return
//}
//
//// NB: I'm not checking tag : min length here!
//func DHCPStringOption(tag byte, s string) (opt DHCPOption, err error) {
//	opt = &dhcpoption{tag: tag, data: bytes.NewBufferString(s).Bytes()}
//	return
//}
//
//func DHCPParseOptions(in []byte) (opts []DHCPOption, err error) {
//	pos := 0
//	for pos < len(in) && err == nil {
//		var tag = in[pos]
//		pos++
//		switch tag {
//		case DHCPOptPad:
//			opts = append(opts, DHCPNewOption(tag, []byte{}))
//		case DHCPOptEnd:
//			return
//		default:
//			if len(in)-pos >= 1 {
//				_len := in[pos]
//				pos++
//				opts = append(opts, DHCPNewOption(tag, in[pos:pos+int(_len)]))
//				pos += int(_len)
//			}
//		}
//	}
//	return
//}
//
//func NewDHCPDiscover(xid uint32, hwAddr net.HardwareAddr) (d *DHCP, err error) {
//	if d, err = New(xid, DHCPMsgDiscover, DHCPHWEthernet); err != nil {
//		return
//	}
//	d.HardwareLen = uint8(len(hwAddr))
//	d.ClientHWAddr = hwAddr
//	d.Options = append(d.Options, DHCPNewOption(byte(53), []byte{byte(DHCPMsgDiscover)}))
//	d.Options = append(d.Options, DHCPNewOption(DHCPOptClientID, []byte(hwAddr)))
//	return
//}
//
//func NewDHCPOffer(xid uint32, hwAddr net.HardwareAddr) (d *DHCP, err error) {
//	if d, err = New(xid, DHCPMsgOffer, DHCPHWEthernet); err != nil {
//		return
//	}
//	d.HardwareLen = uint8(len(hwAddr))
//	d.ClientHWAddr = hwAddr
//	d.Options = append(d.Options, DHCPNewOption(byte(53), []byte{byte(DHCPMsgOffer)}))
//	return
//}
//
//func NewDHCPRequest(xid uint32, hwAddr net.HardwareAddr) (d *DHCP, err error) {
//	if d, err = New(xid, DHCPMsgRequest, DHCPHWEthernet); err != nil {
//		return
//	}
//	d.HardwareLen = uint8(len(hwAddr))
//	d.ClientHWAddr = hwAddr
//	d.Options = append(d.Options, DHCPNewOption(byte(53), []byte{byte(DHCPMsgRequest)}))
//	return
//}
//
//func NewDHCPAck(xid uint32, hwAddr net.HardwareAddr) (d *DHCP, err error) {
//	if d, err = New(xid, DHCPMsgAck, DHCPHWEthernet); err != nil {
//		return
//	}
//	d.HardwareLen = uint8(len(hwAddr))
//	d.ClientHWAddr = hwAddr
//	d.Options = append(d.Options, DHCPNewOption(byte(53), []byte{byte(DHCPMsgAck)}))
//	return
//}
//
//func NewDHCPNak(xid uint32, hwAddr net.HardwareAddr) (d *DHCP, err error) {
//	if d, err = New(xid, DHCPMsgNak, DHCPHWEthernet); err != nil {
//		return
//	}
//	d.HardwareLen = uint8(len(hwAddr))
//	d.ClientHWAddr = hwAddr
//	d.Options = append(d.Options, DHCPNewOption(byte(53), []byte{byte(DHCPMsgNak)}))
//	return
//}
