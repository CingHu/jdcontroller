package app

import (
	"fmt"
	"sync"

	"jd.com/jdcontroller/base"
	"jd.com/jdcontroller/lib/ofproto/ofp13"
	"jd.com/jdcontroller/lib/packet/eth"
	"jd.com/jdcontroller/lib/packet/ipv4"
)

const (
	BRD_MAC = "ff:ff:ff:ff:ff:ff"
)


// Structure to track hosts that we discover.
type Host struct {
	mac  eth.HardwareAddr
	port uint32
}

// A thread safe map to store our hosts.
type HostMap struct {
	hosts map[string]Host
	sync.RWMutex
}

func NewHostMap() *HostMap {
	h := new(HostMap)
	h.hosts = make(map[string]Host)
	return h
}

// Returns the host associated with mac.
func (m *HostMap) Host(mac eth.HardwareAddr) (h Host, ok bool) {
	m.RLock()
	defer m.RUnlock()
	h, ok = m.hosts[mac.String()]
	return
}

// Records the host mac address and the port where mac was discovered.
func (m *HostMap) SetHost(mac eth.HardwareAddr, port uint32) {
	m.Lock()
	defer m.Unlock()
	m.hosts[mac.String()] = Host{mac, port}
}

var hostMap HostMap

// Returns a new instance that implements one of the many
// interfaces found in ofp/ofp13/interface.go. One
// XeniumdInstance will be created for each switch that connects
// to the network.
func NewXeniumdInstance() interface{} {
	hostMap = *NewHostMap()
	return &XeniumdInstance{&hostMap}
}

// Acts as a simple learning switch.
type XeniumdInstance struct {
	*HostMap
}

func (b *XeniumdInstance) MultipartReply(dpid uint64, pkt *ofp13.MultipartReply) {
	if _, ok := base.Switch(dpid); ok {
		ofp13.ParseMultipartReply(dpid, pkt)
	}
}

func (b *XeniumdInstance) EchoRequest(dpid uint64) {
	d := eth.NewHardwareAddr(dpid)
	fmt.Println("Switch Request:", d.String())
	if sw, ok := base.Switch(dpid); ok {
		pkt := ofp13.NewEchoReply()
		sw.Send(pkt)
	}
}

func (b *XeniumdInstance) ConnectionDown(dpid uint64, err error) {
	d := eth.NewHardwareAddr(dpid)
	fmt.Println("ConnectionDown dpid and error is:", d.String(), err)
	base.SwitchDisconnect(dpid)
	return
}

func (b *XeniumdInstance) ConnectionUp(dpid uint64) {
	SetDefaultFlow(dpid)
	//SetArpProxyFlow(dpid)
}

func (b *XeniumdInstance) PacketIn(dpid uint64, pkt *ofp13.PacketIn) {
	for k, app := range AppHandlers[ofp13.Version] {
		fmt.Println("PacketIn for range AppHandlers", k + 1)
		next := app(dpid, pkt)
		if !next {
			break
		}
	}

	return
}

func getSwitchL3FlowPolicy(pkt *ofp13.PacketIn) bool {
	ethPkt := pkt.Data
	if ethPkt.Ethertype != eth.IPv4Msg {
		return true
	}
	var ethPktData interface {} = ethPkt.Data
	ipv4Pkt := ethPktData.(*ipv4.IPv4)
	fmt.Println("^^^IPv4Msg: %s --> %s", ipv4Pkt.NWSrc, ipv4Pkt.NWDst)
	return true
}

func SetDefaultFlow(dpid uint64) {
	sw, found := base.Switch(dpid)
	if !found {
		fmt.Println("Can't find switch:", dpid)
		return
	}

	var priority uint16 = 0
	flow := ofp13.NewFlowMod(priority)
	flow.BufferID = 0xffffffff
	flow.OutPort = 0xffffffff
	flow.OutGroup = 0xffffffff

	/* Add Match */

	/*Add Insturction actions*/
	insActs := ofp13.NewInstructionActions()
	insActs.AddOutputAction(uint32(ofp13.OFPPController))
	flow.AddInsAction(insActs)

	fmt.Println("SetDefaultFlow:", flow.Header)
	sw.Send(flow)
}

func SetArpProxyFlow(dpid uint64) {
	sw, found := base.Switch(dpid)
	if !found {
		fmt.Println("Can't find switch:", dpid)
		return
	}

	var priority uint16 = 1
	flow := ofp13.NewFlowMod(priority)
	flow.BufferID = 0xffffffff

	/* Add Match */
	var mc ofp13.MatchCreator
	ethdst := mc.NewEthDstField("ff:ff:ff:ff:ff:ff")
	flow.AddMatch(ethdst)
	ethtype := mc.NewEthTypeField(uint32(0x0806))
	flow.AddMatch(ethtype)

	/*Add Insturction actions*/
	insActs := ofp13.NewInstructionActions()
	insActs.AddOutputAction(uint32(ofp13.OFPPLocal))
	flow.AddInsAction(insActs)

	fmt.Println("SetArpProxyFlow:", flow.Header)
	sw.Send(flow)
}

func sendTestFlow(dpid uint64, pkt *ofp13.PacketIn) {
	sw, found := base.Switch(dpid)
	if !found {
		fmt.Println("Can't find switch:", dpid)
		return
	}

	var priority uint16 = 2000
	flow := ofp13.NewFlowMod(priority)
	flow.Header.XID = 0
	//flow.Header.XID = 100

	if pkt != nil {
		//flow.BufferID = 0
		flow.BufferID = pkt.BufferID
		//fmt.Println("buffer id", pkt.BufferID)
	}

	/* Add Match */
	var mc ofp13.MatchCreator
	ethdst := mc.NewEthDstField("00:11:11:11:11:11")
	flow.AddMatch(ethdst)
	fmt.Println("sendTestFlowmatch:", flow.Match)

	/*Add Insturction actions*/
	insActs := ofp13.NewInstructionActions()
	insActs.AddOutputAction(uint32(888))
	flow.AddInsAction(insActs)
	sw.Send(flow)
}

func sendTestFlow2(dpid uint64, pkt *ofp13.PacketIn) {
	sw, found := base.Switch(dpid)
	if !found {
		fmt.Println("Can't find switch:", dpid)
		return
	}

	var priority uint16 = 2000
	flow := ofp13.NewFlowMod(priority)
	flow.Header.XID = 0
	//flow.Header.XID = 100

	if pkt != nil {
		//flow.BufferID = 0
		flow.BufferID = pkt.BufferID
		//fmt.Println("buffer id", pkt.BufferID)
	}

	/* Add Match */
	var mc ofp13.MatchCreator
	ethdst := mc.NewEthDstField("00:22:22:22:22:22")
	flow.AddMatch(ethdst)
	ethtype := mc.NewEthTypeField(uint32(0x0800))
	flow.AddMatch(ethtype)
	ipsrc := mc.NewIpSrcField("10.1.1.123/24")
	flow.AddMatch(ipsrc)

	/*Add Insturction actions*/
	insActs := ofp13.NewInstructionActions()
	insActs.AddSetTunnelIpDstField("10.1.1.1")
	insActs.AddSetEthDstField("00:11:11:11:11:11")
	insActs.AddSetIpDstField("10.2.2.2")
	insActs.AddOutputAction(uint32(888))
	flow.AddInsAction(insActs)

	sw.Send(flow)
}

func sendPacketOut(dpid uint64, pkt *ofp13.PacketIn) {
	sw, found := base.Switch(dpid)
	if !found {
		fmt.Println("Can't find switch:", dpid)
		return
	}

	dstmac := pkt.Data.HWDst
	eth.NewHardwareAddr(dstmac).String()
	dstmacstr := eth.NewHardwareAddr(dstmac).String()

	ethPkt := pkt.Data
	packetout := ofp13.NewPacketOut()
	packetout.InPort = pkt.GetInport()

	if dstmacstr == "00:50:56:26:8f:6b" {
		packetout.AddSetEthDstField("00:00:00:11:11:11")
		packetout.AddOutputAction(uint32(1))
		fmt.Println("sendPacketOut output:1")
	} else if dstmacstr == "00:50:56:2e:08:4c" {
		packetout.AddSetEthDstField("00:00:00:11:11:11")
		packetout.AddSetTunnelIpDstField("10.2.2.2")
		packetout.AddOutputAction(uint32(2))
		fmt.Println("sendPacketOut output:2")
	} else {
		return
	}

	packetout.BufferID = pkt.BufferID

	if packetout.BufferID == 0xffffffff {
		datalen := uint16(ethPkt.Len()) + ethPkt.Data.(*ipv4.IPv4).Length
		packetout.Header.Length += datalen
		packetout.Data = ethPkt
	}

	sw.Send(packetout)
}

func sendDhcpPacketOut(dpid uint64, pkt *ofp13.PacketIn, frame *eth.Ethernet) {
	sw, found := base.Switch(dpid)
	if !found {
		fmt.Println("Can't find switch:", dpid)
		return
	}

	ethPkt := *frame
	packetout := ofp13.NewPacketOut()
	packetout.BufferID = 0xffffffff
	packetout.InPort = pkt.GetInport()
	packetout.AddOutputAction(uint32(ofp13.OFPPInPort))
	packetout.Data = ethPkt
	datalen := uint16(ethPkt.Len()) + ethPkt.Data.(*ipv4.IPv4).Length
	packetout.Header.Length += datalen

	sw.Send(packetout)
}

func sendJingYangFlow(dpid uint64, pkt *ofp13.PacketIn) {
	sw, found := base.Switch(dpid)
	if !found {
		fmt.Println("Can't find switch:", dpid)
		return
	}

	var priority uint16 = 2000
	flow := ofp13.NewFlowMod(priority)

	if pkt != nil {
		flow.BufferID = pkt.BufferID
		fmt.Println("sendJingYangFlow buffer id", pkt.BufferID)
	}

	dstmac := pkt.Data.HWDst
	eth.NewHardwareAddr(dstmac).String()
	dstmacstr := eth.NewHardwareAddr(dstmac).String()

	/* Add Match */
	var mc ofp13.MatchCreator
	ethdst := mc.NewEthDstField(dstmacstr)
	flow.AddMatch(ethdst)
	fmt.Println("sendTestFlowmatch:", pkt.BufferID, flow.Match)

	/*Add Insturction actions*/
	insActs := ofp13.NewInstructionActions()

	if dstmacstr == "00:50:56:26:8f:6b" {
		insActs.AddOutputAction(uint32(1))
	} else if dstmacstr == "00:50:56:2e:08:4c" {
		insActs.AddOutputAction(uint32(2))
	} else {
		return
	}

	fmt.Println("send jingyang's flow", dstmacstr)
	flow.AddInsAction(insActs)

	sw.Send(flow)
}
