package app

import (
	"fmt"
	_"bytes"
	"strings"
	"strconv"
	"jd.com/jdcontroller/base"
	"jd.com/jdcontroller/lib/packet/eth"
	"jd.com/jdcontroller/lib/ofproto/ofp13"
)

func (b *XeniumdInstance)FlowRemoved(dpid uint64, flow *ofp13.FlowRemoved) {
	if (flow.Reason != ofp13.OFPRRDelete) || (flow.Match.Length == 4) {
		return
	}

	var dstmac [6]byte
	match := flow.Match

	dpidhw := eth.NewHardwareAddr(dpid)
	dpidStr := strings.Replace(dpidhw.String(), ":", "", -1)
	dpidStr = "0000" + dpidStr
	if dpidStr == "" {
		fmt.Println(">> Meet unknow dpid: " + dpidStr)
		return
	}
	fmt.Println(">> FlowRemoved: %v", flow)
	for i := 0; i < 6; i++ {
		dstmac[i] = match.OXMFields[i + 4]
	}
//	dmacstr := eth.NewHardwareAddr(dstmac).String()

	return
}

func DeleteOvsFlow(dpid, dmac string) bool {
	flowMod := ofp13.NewFlowMod(uint16(ofp13.FLOWDEFPRIORITY))
	flowMod.Command = ofp13.OFPFCDeleteStrict
	flowMod.TableID = ofp13.OFPTTAll
	flowMod.BufferID = ofp13.OFPNOBUFFER
	flowMod.OutPort = ofp13.OFPPAny
	flowMod.OutGroup = ofp13.OFPGAny

	var mc ofp13.MatchCreator
	ethdst := mc.NewEthDstField(dmac)
	flowMod.AddMatch(ethdst)

	switchindex, _ := strconv.ParseUint(dpid, 16, 64)
	fmt.Println(">> DeleteOvsFlow, %x", switchindex)
	sw, ok := base.Switch(switchindex)
	if ok {
		fmt.Println(">> DeleteOvsFlow ok", dpid, dmac)
		sw.Send(flowMod)
		return true
	}
	return false
}
