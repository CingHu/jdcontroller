package ofpctl10

import (
	"errors"
	"log"

	"jd.com/jdcontroller/lib/buffer"
	"jd.com/jdcontroller/protocol/ofp10"
)

func Parse(bytes []byte) (msg buffer.Message, err error) {
	switch bytes[1] {
	case ofp10.OFPTHello:
		log.Println("[IN] Hello message.")
		msg = new(ofp10.Header)
		msg.UnpackBinary(bytes)
	case ofp10.OFPTError:
		msg = new(ofp10.ErrorMsg)
		log.Println("[IN] Error Message.")
		msg.UnpackBinary(bytes)
	case ofp10.OFPTEchoRequest:
		log.Println("[IN] Echo Request Message.")
		msg = new(ofp10.Header)
		msg.UnpackBinary(bytes)
	case ofp10.OFPTEchoReply:
		log.Println("[IN] Echo Reply Message.")
		msg = new(ofp10.Header)
		msg.UnpackBinary(bytes)
	case ofp10.OFPTVendor:
		log.Println("[IN] Vendor Message.")
		msg = new(ofp10.VendorHeader)
		msg.UnpackBinary(bytes)
	case ofp10.OFPTFeaturesRequest:
		log.Println("[IN] Features Request Message.")
		msg = NewFeaturesRequest()
		msg.UnpackBinary(bytes)
	case ofp10.OFPTFeaturesReply:
		log.Println("[IN] Features Reply Message.")
		msg = NewFeaturesReply()
		msg.UnpackBinary(bytes)
	case ofp10.OFPTGetConfigRequest:
		log.Println("[IN] Get Config Request Message.")
		msg = new(ofp10.Header)
		msg.UnpackBinary(bytes)
	case ofp10.OFPTGetConfigReply:
		log.Println("[IN] Get Config Reply Message.")
		msg = new(ofp10.SwitchConfig)
		msg.UnpackBinary(bytes)
	case ofp10.OFPTSetConfig:
		log.Println("[IN] Set Config Message.")
		msg = NewSetConfig()
		msg.UnpackBinary(bytes)
	case ofp10.OFPTPacketIn:
		log.Println("[IN] Packet In Message.")
		msg = new(ofp10.PacketIn)
		msg.UnpackBinary(bytes)
	case ofp10.OFPTFlowRemoved:
		log.Println("[IN] Flow Removed Message.")
		msg = NewFlowRemoved()
		msg.UnpackBinary(bytes)
	case ofp10.OFPTPortStatus:
		log.Println("[IN] Port Status Message.")
		msg = new(ofp10.PortStatus)
		msg.UnpackBinary(bytes)
	case ofp10.OFPTPacketOut:
		log.Println("[IN] Packet Out Message.")
		break
	case ofp10.OFPTFlowMod:
		log.Println("[IN] Flow Mod Message.")
		msg = NewFlowMod()
		msg.UnpackBinary(bytes)
	case ofp10.OFPTPortMod:
		log.Println("[IN] Port Mod Message.")
		break
	case ofp10.OFPTStatsRequest:
		log.Println("[IN] Stats Request Message.")
		msg = new(ofp10.StatsRequest)
		msg.UnpackBinary(bytes)
	case ofp10.OFPTStatsReply:
		log.Println("[IN] Stats Reply Message.")
		msg = new(ofp10.StatsReply)
		msg.UnpackBinary(bytes)
	case ofp10.OFPTBarrierRequest:
		log.Println("[IN] Barrier Request Message.")
		msg = new(ofp10.Header)
		msg.UnpackBinary(bytes)
	case ofp10.OFPTBarrierReply:
		log.Println("[IN] Barrier Reply Message.")
		msg = new(ofp10.Header)
		msg.UnpackBinary(bytes)
	case ofp10.OFPTQueueGetConfigRequest:
		log.Println("[IN] Queue Get Config Request Message.")
		break
	case ofp10.OFPTQueueGetConfigReply:
		log.Println("[IN] Queue Get Config Reply Message.")
		break
	default:
		err = errors.New("[IN] An unknown v1.0 packet type was received. Parse function will discard data.")
	}
	return
}
