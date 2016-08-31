package ofp12

import (
	"errors"
	"log"

	"jd.com/jdcontroller/lib/buffer"
	"jd.com/jdcontroller/protocol"
)

func Parse(data []byte) (msg buffer.Message, err error) {
	switch data[1] {
	case protocol.OFPTHello:
		log.Println("[IN] Hello message.")
		msg = new(Header)
		msg.UnpackBinary(data)
	case protocol.OFPTError:
		msg = new(ErrorMsg)
		log.Println("[IN] Error Message.")
		msg.UnpackBinary(data)
	case protocol.OFPTEchoRequest:
		log.Println("[IN] Echo Request Message.")
		msg = new(Header)
		msg.UnpackBinary(data)
	case protocol.OFPTEchoReply:
		log.Println("[IN] Echo Reply Message.")
		msg = new(Header)
		msg.UnpackBinary(data)
	case protocol.OFPTExperimenter:
		log.Println("[IN] Experimenter Message.")
		msg = new(ExperimenterHeader)
		msg.UnpackBinary(data)
	case protocol.OFPTFeaturesRequest:
		log.Println("[IN] Features Request Message.")
		msg = new(Header)
		msg.UnpackBinary(data)
	case protocol.OFPTFeaturesReply:
		log.Println("[IN] Features Reply Message.")
		msg = new(SwitchFeatures)
		msg.UnpackBinary(data)
	case protocol.OFPTGetConfigRequest:
		log.Println("[IN] Get Config Request Message.")
		msg = new(Header)
		msg.UnpackBinary(data)
	case protocol.OFPTGetConfigReply:
		log.Println("[IN] Get Config Reply Message.")
		msg = new(SwitchConfig)
		msg.UnpackBinary(data)
	case protocol.OFPTSetConfig:
		log.Println("[IN] Set Config Message.")
		msg = new(SwitchConfig)
		msg.UnpackBinary(data)
	case protocol.OFPTPacketIn:
		log.Println("[IN] Packet In Message.")
		msg = new(PacketIn)
		msg.UnpackBinary(data)
	case protocol.OFPTFlowRemoved:
		log.Println("[IN] Flow Removed Message.")
		msg = new(FlowRemoved)
		msg.UnpackBinary(data)
	case protocol.OFPTPortStatus:
		log.Println("[IN] Port Status Message.")
		msg = new(PortStatus)
		msg.UnpackBinary(data)
	case protocol.OFPTPacketOut:
		log.Println("[IN] Packet Out Message.")
		msg = new(PacketOut)
		msg.UnpackBinary(data)
	case protocol.OFPTFlowMod:
		log.Println("[IN] Flow Mod Message.")
		msg = new(FlowMod)
		msg.UnpackBinary(data)
	case protocol.OFPTGroupMod:
		log.Println("[IN] Group Mod Message.")
		msg = new(GroupMod)
		msg.UnpackBinary(data)
	case protocol.OFPTPortMod:
		log.Println("[IN] Port Mod Message.")
		msg = new(PortMod)
		msg.UnpackBinary(data)
	case protocol.OFPTTableMod:
		log.Println("[IN] Table Mod Message.")
		msg = new(TableMod)
		msg.UnpackBinary(data)
	case protocol.OFPTStatsRequest:
		log.Println("[IN] Stats Request Message.")
		msg = new(StatsRequest)
		msg.UnpackBinary(data)
	case protocol.OFPTStatsReply:
		log.Println("[IN] Stats Reply Message.")
		msg = new(StatsReply)
		msg.UnpackBinary(data)
	case protocol.OFPTBarrierRequest:
		log.Println("[IN] Barrier Request Message.")
		msg = new(Header)
		msg.UnpackBinary(data)
	case protocol.OFPTBarrierReply:
		log.Println("[IN] Barrier Reply Message.")
		msg = new(Header)
		msg.UnpackBinary(data)
	case protocol.OFPTQueueGetConfigRequest:
		log.Println("[IN] Queue Get Config Request Message.")
		msg = new(QueueGetConfigRequest)
		msg.UnpackBinary(data)
	case protocol.OFPTQueueGetConfigReply:
		log.Println("[IN] Queue Get Config Reply Message.")
		msg = new(QueueGetConfigReply)
		msg.UnpackBinary(data)
	default:
		err = errors.New("[IN] An unknown v1.1 packet type was received. Parse function will discard data.")
	}
	return
}
