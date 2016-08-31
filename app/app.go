package app

import (
	"fmt"
	"jd.com/jdcontroller/lib/ofproto/ofp13"
	"jd.com/jdcontroller/lib/packet/eth"
	_"jd.com/jdcontroller/lib/packet/lldp"
	"jd.com/jdcontroller/lib/packet/dhcp"
	"jd.com/jdcontroller/app/dhcpserver"
)

type SdnApplication struct {
	Name string
	Version uint8
	Handler func(uint64, interface{})(bool)
}

type AppList []*SdnApplication
type AppHandler func(uint64, interface{})(bool)
type AppHandlerList []AppHandler
type AppHandlerMap map[uint8]AppHandlerList

var Apps AppList
var AppHandlers AppHandlerMap

func L2ForwardApp(dpid uint64, msg interface{}) bool {
	pkt := msg.(*ofp13.PacketIn)
	// Ignore link discovery packet types.
	if pkt.Data.Ethertype == 0xa0f1 || pkt.Data.Ethertype == 0x88cc {
		return true
	}
	inport := pkt.GetInport()
	frame := pkt.Data
	fmt.Println("HWSrc:", eth.GetMacAddr(frame.HWSrc))
	fmt.Println("HWDst:", eth.GetMacAddr(frame.HWDst))
	fmt.Println("L2ForwardApp", dpid, inport)

	sendJingYangFlow(dpid, pkt)
	return true
}

func DhcpServerApp(dpid uint64, msg interface{}) bool {
	pkt := msg.(*ofp13.PacketIn)
	frame := pkt.Data

	fmt.Println("DhcpServerApp", dpid, pkt)
	if dhcp.IfDhcpPkt(frame) {
		pktoutdata, err := dhcpserver.DhcpServer(frame)
		if err != nil {
			return true
		}
		sendDhcpPacketOut(dpid, pkt, pktoutdata)
	}

	return true
}

func LLDPApp(dpid uint64, msg interface{}) bool {
	pkt := msg.(*ofp13.PacketIn)
	if pkt.Data.Ethertype == 0xa0f1 && pkt.Data.Ethertype == 0x88cc {
		inport := pkt.GetInport()
		frame := pkt.Data
		fmt.Println("HWSrc:", eth.GetMacAddr(frame.HWSrc))
		fmt.Println("HWDst:", eth.GetMacAddr(frame.HWDst))
		fmt.Println("LLDPApp", dpid, inport)
		return true
	}

	fmt.Println("LLDPApp Not Match")
	return true
}

func RegisterApplication(app *SdnApplication) {
	Apps = append(Apps, app)
	AppHandlers[app.Version] = append(AppHandlers[app.Version], app.Handler)
	return
}

func AppInit(name string, Version uint8, handler AppHandler) (*SdnApplication){
	app := new(SdnApplication)
	app.Name = name
	app.Version = Version
	app.Handler = handler
	return app
}

func NewApp(name string, fn AppHandler) {
	if AppHandlers == nil {
		AppHandlers = make(AppHandlerMap)
	}

	app := AppInit(name, uint8(ofp13.Version), fn)
	RegisterApplication(app)
	return
}

