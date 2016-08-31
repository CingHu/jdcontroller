package buffer

import (
	"errors"
	"log"

	"jd.com/jdcontroller/protocol/ofp13"
)

func ofp13MsgParse(data []byte) (msg Message, err error) {
	switch data[1] {
	case ofp13.OFPTHello:
		log.Println("[IN] Hello message.")
		msg = new(ofp13.Header)
		msg.UnpackBinary(data)
	case ofp13.OFPTError:
		msg = new(ofp13.ErrorMsg)
		log.Println("[IN] Error Message.")
		msg.UnpackBinary(data)
	case ofp13.OFPTEchoRequest:
		log.Println("[IN] Echo Request Message.")
		msg = new(ofp13.Header)
		msg.UnpackBinary(data)
	case ofp13.OFPTEchoReply:
		log.Println("[IN] Echo Reply Message.")
		msg = new(ofp13.Header)
		msg.UnpackBinary(data)
	case ofp13.OFPTExperimenter:
		log.Println("[IN] Experimenter Message.")
		msg = new(ofp13.ExperimenterHeader)
		msg.UnpackBinary(data)
	case ofp13.OFPTFeaturesRequest:
		log.Println("[IN] Features Request Message.")
		msg = new(ofp13.Header)
		msg.UnpackBinary(data)
	case ofp13.OFPTFeaturesReply:
		log.Println("[IN] Features Reply Message.")
		msg = new(ofp13.SwitchFeatures)
		msg.UnpackBinary(data)
	case ofp13.OFPTGetConfigRequest:
		log.Println("[IN] Get Config Request Message.")
		msg = new(ofp13.Header)
		msg.UnpackBinary(data)
	case ofp13.OFPTGetConfigReply:
		log.Println("[IN] Get Config Reply Message.")
		msg = new(ofp13.SwitchConfig)
		msg.UnpackBinary(data)
	case ofp13.OFPTSetConfig:
		log.Println("[IN] Set Config Message.")
		msg = new(ofp13.SwitchConfig)
		msg.UnpackBinary(data)
	case ofp13.OFPTPacketIn:
		log.Println("[IN] Packet In Message.")
		msg = new(ofp13.PacketIn)
		msg.UnpackBinary(data)
	case ofp13.OFPTFlowRemoved:
		log.Println("[IN] Flow Removed Message.")
		msg = new(ofp13.FlowRemoved)
		msg.UnpackBinary(data)
	case ofp13.OFPTPortStatus:
		log.Println("[IN] Port Status Message.")
		msg = new(ofp13.PortStatus)
		msg.UnpackBinary(data)
	case ofp13.OFPTPacketOut:
		log.Println("[IN] Packet Out Message.")
		msg = new(ofp13.PacketOut)
		msg.UnpackBinary(data)
	case ofp13.OFPTFlowMod:
		log.Println("[IN] Flow Mod Message.")
		msg = new(ofp13.FlowMod)
		msg.UnpackBinary(data)
	case ofp13.OFPTGroupMod:
		log.Println("[IN] Group Mod Message.")
		msg = new(ofp13.GroupMod)
		msg.UnpackBinary(data)
	case ofp13.OFPTPortMod:
		log.Println("[IN] Port Mod Message.")
		msg = new(ofp13.PortMod)
		msg.UnpackBinary(data)
	case ofp13.OFPTTableMod:
		log.Println("[IN] Table Mod Message.")
		msg = new(ofp13.TableMod)
		msg.UnpackBinary(data)
	case ofp13.OFPTMultipartRequest:
		log.Println("[IN] Multipart Request Message.")
		msg = new(ofp13.MultipartRequest)
		msg.UnpackBinary(data)
	case ofp13.OFPTMultipartReply:
		log.Println("[IN] Multipart Reply Message.")
		msg = new(ofp13.MultipartReply)
		msg.UnpackBinary(data)
	case ofp13.OFPTBarrierRequest:
		log.Println("[IN] Barrier Request Message.")
		msg = new(ofp13.Header)
		msg.UnpackBinary(data)
	case ofp13.OFPTBarrierReply:
		log.Println("[IN] Barrier Reply Message.")
		msg = new(ofp13.Header)
		msg.UnpackBinary(data)
	case ofp13.OFPTQueueGetConfigRequest:
		log.Println("[IN] Queue Get Config Request Message.")
		msg = new(ofp13.QueueGetConfigRequest)
		msg.UnpackBinary(data)
	case ofp13.OFPTQueueGetConfigReply:
		log.Println("[IN] Queue Get Config Reply Message.")
		msg = new(ofp13.QueueGetConfigReply)
		msg.UnpackBinary(data)
	case ofp13.OFPTRoleRequest:
		log.Println("[IN] Role request Message.")
		msg = new(ofp13.RoleRequest)
		msg.UnpackBinary(data)
	case ofp13.OFPTRoleReply:
		log.Println("[IN] Role reply Message.")
		msg = new(ofp13.RoleRequest)
		msg.UnpackBinary(data)
	case ofp13.OFPTGetAsyncRequest:
		log.Println("[IN] Get Async request Message.")
		msg = new(ofp13.Header)
		msg.UnpackBinary(data)
	case ofp13.OFPTGetAsyncReply:
		log.Println("[IN] Get Async reply Message.")
		msg = new(ofp13.AsyncConfig)
		msg.UnpackBinary(data)
	case ofp13.OFPTSetAsync:
		log.Println("[IN] Set Async Message.")
		msg = new(ofp13.AsyncConfig)
		msg.UnpackBinary(data)
	case ofp13.OFPTMeterMod:
		log.Println("[IN] Meter mod Message.")
		msg = new(ofp13.MeterMod)
		msg.UnpackBinary(data)
	default:
		err = errors.New("[IN] An unknown v1.3 packet type was received. Parse function will discard data.")
	}
	return
}
