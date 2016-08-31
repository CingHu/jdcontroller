package dhcpserver

import (
	"fmt"
	"net"
	"time"
//	"math/rand"

	"jd.com/jdcontroller/lib/packet/eth"
	"jd.com/jdcontroller/lib/packet/ipv4"
	"jd.com/jdcontroller/lib/packet/udp"
)

type lease struct {
	nic    string    // Client's CHAddr
	expiry time.Time // When the lease expires
}

type DHCPHandler struct {
	ip            net.IP        // Server IP to use
	options       Options  // Options to send to DHCP Clients
	start         net.IP        // Start of IP range to distribute
	leaseRange    int           // Number of IPs to distribute (starting from start)
	leaseDuration time.Duration // Lease period
	leases        map[int]lease // Map to keep track of leases
}

type Handler interface {
	ServeDHCP(req Packet, msgType MessageType, options Options) Packet
}

func NewDHCPHandler(name string) (* DHCPHandler, error) {
	serverIP, err := GetIpByName(name)
	if err != nil {
		return nil, err
	}

	DhcpHandler := &DHCPHandler{
		ip:            serverIP,
		leaseDuration: 2 * time.Hour,
		//start:         net.IP{192, 168, 100, 10},
		start:         stringToIp("192.168.100.70"),
		//leaseRange:    50,
		//leases:        make(map[int]lease, 30),
		options: Options{
			//OptionSubnetMask:       []byte{255, 255, 255, 0},
			//OptionRouter:           []byte{192, 168, 100, 1},                   // Presuming Server is also your router
			OptionSubnetMask:       stringToIp("255.255.255.0"),
			OptionRouter:           stringToIp("192.168.100.1"),                   // Presuming Server is also your router
			OptionDomainNameServer: []byte(net.IP{114, 114, 114, 114}), // Presuming Server is also your DNS server
		},
	}
	return DhcpHandler, err
}

func UpdateDHCPHandler() {
	return
}

func (h *DHCPHandler) ServeDHCP(p Packet, msgType MessageType, options Options) (d Packet) {
	switch msgType {
	case Request:
		return ReplyPacket(p, ACK, h.ip, h.start, h.leaseDuration,
			h.options.SelectOrderOrAll(options[OptionParameterRequestList]))
	case Discover:
		/*free, nic := -1, p.CHAddr().String()
		for i, v := range h.leases { // Find previous lease
			if v.nic == nic {
				free = i
				goto reply
			}
		}
		if free = h.freeLease(); free == -1 {
			return
		}
		reply:
		return ReplyPacket(p, Offer, h.ip, IPAdd(h.start, free), h.leaseDuration,
			h.options.SelectOrderOrAll(options[OptionParameterRequestList]))*/
		return ReplyPacket(p, Offer, h.ip, h.start, h.leaseDuration,
			h.options.SelectOrderOrAll(options[OptionParameterRequestList]))

	/*case Request:
		if server, ok := options[OptionServerIdentifier]; ok && !net.IP(server).Equal(h.ip) {
			return nil // Message not for this dhcp server
		}
		reqIP := net.IP(options[OptionRequestedIPAddress])
		if reqIP == nil {
			reqIP = net.IP(p.CIAddr())
		}

		if len(reqIP) == 4 && !reqIP.Equal(net.IPv4zero) {
			if leaseNum := IPRange(h.start, reqIP) - 1; leaseNum >= 0 && leaseNum < h.leaseRange {
				if l, exists := h.leases[leaseNum]; !exists || l.nic == p.CHAddr().String() {
					h.leases[leaseNum] = lease{nic: p.CHAddr().String(), expiry: time.Now().Add(h.leaseDuration)}
					return ReplyPacket(p, ACK, h.ip, net.IP(options[OptionRequestedIPAddress]), h.leaseDuration,
						h.options.SelectOrderOrAll(options[OptionParameterRequestList]))
				}
			}
			return ReplyPacket(p, ACK, h.ip, net.IP(options[OptionRequestedIPAddress]), h.leaseDuration,
						h.options.SelectOrderOrAll(options[OptionParameterRequestList]))
		}
		return ReplyPacket(p, NAK, h.ip, nil, 0, nil) */

	case Release, Decline:
		nic := p.CHAddr().String()
		for i, v := range h.leases {
			if v.nic == nic {
				delete(h.leases, i)
				break
			}
		}
	}
	return nil
}

/*func (h *DHCPHandler) freeLease() int {
	now := time.Now()
	b := rand.Intn(h.leaseRange) // Try random first
	for _, v := range [][]int{[]int{b, h.leaseRange}, []int{0, b}} {
		for i := v[0]; i < v[1]; i++ {
			if l, ok := h.leases[i]; !ok || l.expiry.Before(now) {
				return i
			}
		}
	}
	return -1
}*/

// Serve takes a ServeConn (such as a net.PacketConn) that it uses for both
// reading and writing DHCP packets. Every packet is passed to the handler,
// which processes it and optionally return a response packet for writing back
// to the network.
//
// To capture limited broadcast packets (sent to 255.255.255.255), you must
// listen on a socket bound to IP_ADDRANY (0.0.0.0). This means that broadcast
// packets sent to any interface on the system may be delivered to this
// socket.  See: https://code.google.com/p/go/issues/detail?id=7106
//
// Additionally, response packets may not return to the same
// interface that the request was received from.  Writing a custom ServeConn,
// or using ServeIf() can provide a workaround to this problem.
func (h *DHCPHandler) dhcpReply(buffer []byte, buf_len int) (res []byte){
	if buf_len < 240 { // Packet too small to be DHCP
		return
	}

	req := Packet(buffer[:buf_len])
	if req.HLen() > 16 { // Invalid size
		return
	}
	fmt.Printf("client mac:=%s\n", req.CHAddr())
	options := req.ParseOptions()
	fmt.Println(options)
	var reqType MessageType
	if t := options[OptionDHCPMessageType]; len(t) != 1 {
		return
	} else {
		reqType = MessageType(t[0])
		if reqType < Discover || reqType > Inform {
			return
		}
	}
	fmt.Printf("request message type=%d\n", reqType)
	if res = h.ServeDHCP(req, reqType, options); res != nil {
		fmt.Println(res)
	}
	return
}

func (h* DHCPHandler)transEthernet(oldFrame eth.Ethernet, newFrame *eth.Ethernet) error {
	newFrame.Ethertype = oldFrame.Ethertype
	newFrame.HWDst = oldFrame.HWSrc
	//newFrame.HWSrc = {1,2,3,4,5,6}
	localInterface, err := net.InterfaceByName("br0")
	if err != nil {
		return err
	}
	fmt.Println(newFrame.HWSrc, oldFrame.HWDst)
	for i:=0; i<6; i++ {
		newFrame.HWSrc[i] = localInterface.HardwareAddr[i]
	}
	return nil
}

func (h* DHCPHandler)transIpv4(oldIpv4Msg *ipv4.IPv4, newIpv4Msg *ipv4.IPv4) {
	newIpv4Msg.Version = oldIpv4Msg.Version
	newIpv4Msg.DSCP = oldIpv4Msg.DSCP //0
	newIpv4Msg.ECN = oldIpv4Msg.ECN //0
	newIpv4Msg.Id = oldIpv4Msg.Id //random
	newIpv4Msg.Flags = oldIpv4Msg.Flags //0
	newIpv4Msg.TTL = oldIpv4Msg.TTL //64
	newIpv4Msg.Protocol = oldIpv4Msg.Protocol //udp
	newIpv4Msg.NWSrc = h.ip
	if net.ParseIP(oldIpv4Msg.NWSrc.String()).Equal(net.IPv4zero) {
		newIpv4Msg.NWDst = net.IPv4bcast
	} else {
		newIpv4Msg.NWDst = oldIpv4Msg.NWSrc
	}

	newIpv4Msg.FragmentOffset = 0
	newIpv4Msg.IHL = 20 //ipv4 header length
	newIpv4Msg.Length = uint16(newIpv4Msg.Len()) //total length
	newIpv4Msg.Checksum = 0
	newIpv4Msg.CheckSum()
}

func (h* DHCPHandler)transUDP(oldUdpMsg *udp.UDP, newUdpMsg *udp.UDP) {
	newUdpMsg.PortSrc = oldUdpMsg.PortDst
	newUdpMsg.PortDst = oldUdpMsg.PortSrc
	newUdpMsg.Length = uint16(newUdpMsg.Len())
	newUdpMsg.Checksum = 0
}

func DhcpServer(frame eth.Ethernet) (*eth.Ethernet, error) {
	DhcpHandler, err := NewDHCPHandler("br0")
	if err != nil {
		return nil, err
	}

	outFrame := eth.New()
	outIpv4Msg := ipv4.New()
	outUdpMsg := udp.New()

	ipv4Msg := frame.Data.(*ipv4.IPv4)
	udpMsg := ipv4Msg.Data.(*udp.UDP)
	dhcpBuffer := DhcpHandler.dhcpReply(udpMsg.Data, len(udpMsg.Data))

	outUdpMsg.Data = dhcpBuffer
	DhcpHandler.transUDP(udpMsg, outUdpMsg)

	outIpv4Msg.Data = outUdpMsg
	DhcpHandler.transIpv4(ipv4Msg, outIpv4Msg)

	outFrame.Data = outIpv4Msg
	DhcpHandler.transEthernet(frame, outFrame)

	fmt.Println("data ", outFrame)

	return outFrame, err
}
