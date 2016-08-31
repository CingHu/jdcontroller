package ofp11

import (
	"jd.com/jdcontroller/lib/buffer"
)

type ConnectionUpReactor interface {
	ConnectionUp(dpid uint64)
}

type ConnectionDownReactor interface {
	ConnectionDown(dpid uint64, err error)
}

type HelloReactor interface {
	Hello(hello *Header)
}

type ErrorReactor interface {
	Error(dpid uint64, err *ErrorMsg)
}

type ExperimenterReactor interface {
	Experimenter(dpid uint64)
}

type EchoRequestReactor interface {
	EchoRequest(dpid uint64)
}

type EchoReplyReactor interface {
	EchoReply(dpid uint64)
}

type FeaturesRequestReactor interface {
	FeaturesRequest(features *Header)
}

type FeaturesReplyReactor interface {
	FeaturesReply(dpid uint64, features *SwitchFeatures)
}

type GetConfigRequestReactor interface {
	GetConfigRequest(config *Header)
}

type GetConfigReplyReactor interface {
	GetConfigReply(dpid uint64, config *SwitchConfig)
}

type QueueGetConfigRequestReactor interface {
	QueueGetConfigRequest(config *QueueGetConfigRequest)
}

type QueueGetConfigReplyReactor interface {
	QueueGetConfigReply(dpid uint64, config *QueueGetConfigReply)
}

type SetConfigReactor interface {
	SetConfig(config *SwitchConfig)
}

type PacketInReactor interface {
	PacketIn(dpid uint64, packet *PacketIn)
}

type FlowRemovedReactor interface {
	FlowRemoved(dpid uint64, flow *FlowRemoved)
}

type PortStatusReactor interface {
	PortStatus(dpid uint64, status *PortStatus)
}

type PacketOutReactor interface {
	PacketOut(packet *PacketOut)
}

type FlowModReactor interface {
	FlowMod(flowMod *FlowMod)
}

type GroupModReactor interface {
	GroupMod(groupMod *GroupMod)
}

type PortModReactor interface {
	PortMod(portMod *PortMod)
}

type TableModReactor interface {
	TableMod(tableMod *TableMod)
}

type StatsRequestReactor interface {
	StatsRequest(req *StatsRequest)
}

type StatsReplyReactor interface {
	StatsReply(dpid uint64, rep *StatsReply)
}

type BarrierRequestReactor interface {
	BarrierRequest(req *Header)
}

type BarrierReplyReactor interface {
	BarrierReply(dpid uint64, msg *Header)
}

func ReactorParse(dpid uint64, app interface{}, msg buffer.Message) {
	switch t := msg.(type) {
	case *Header:
		switch t.Type {
		case OFPTHello:
			if actor, ok := app.(HelloReactor); ok {
				actor.Hello(t)
			}
		case OFPTEchoRequest:
			if actor, ok := app.(EchoRequestReactor); ok {
				actor.EchoRequest(dpid)
			}
		case OFPTEchoReply:
			if actor, ok := app.(EchoReplyReactor); ok {
				actor.EchoReply(dpid)
			}
		case OFPTFeaturesRequest:
			if actor, ok := app.(FeaturesRequestReactor); ok {
				actor.FeaturesRequest(t)
			}
		case OFPTGetConfigRequest:
			if actor, ok := app.(GetConfigRequestReactor); ok {
				actor.GetConfigRequest(t)
			}
		case OFPTBarrierRequest:
			if actor, ok := app.(BarrierRequestReactor); ok {
				actor.BarrierRequest(t)
			}
		case OFPTBarrierReply:
			if actor, ok := app.(BarrierReplyReactor); ok {
				actor.BarrierReply(dpid, t)
			}
		}
	case *ErrorMsg:
		if actor, ok := app.(ErrorReactor); ok {
			actor.Error(dpid, t)
		}
	case *ExperimenterHeader:
		if actor, ok := app.(ExperimenterReactor); ok {
			actor.Experimenter(dpid)
		}
	case *SwitchFeatures:
		if actor, ok := app.(FeaturesReplyReactor); ok {
			actor.FeaturesReply(dpid, t)
		}
	case *SwitchConfig:
		switch t.Header.Type {
		case OFPTGetConfigReply:
			if actor, ok := app.(GetConfigReplyReactor); ok {
				actor.GetConfigReply(dpid, t)
			}
		case OFPTSetConfig:
			if actor, ok := app.(SetConfigReactor); ok {
				actor.SetConfig(t)
			}
		}
	case *PacketIn:
		if actor, ok := app.(PacketInReactor); ok {
			actor.PacketIn(dpid, t)
		}
	case *FlowRemoved:
		if actor, ok := app.(FlowRemovedReactor); ok {
			actor.FlowRemoved(dpid, t)
		}
	case *PortStatus:
		if actor, ok := app.(PortStatusReactor); ok {
			actor.PortStatus(dpid, t)
		}
	case *PacketOut:
		if actor, ok := app.(PacketOutReactor); ok {
			actor.PacketOut(t)
		}
	case *FlowMod:
		if actor, ok := app.(FlowModReactor); ok {
			actor.FlowMod(t)
		}
	case *GroupMod:
		if actor, ok := app.(GroupModReactor); ok {
			actor.GroupMod(t)
		}
	case *PortMod:
		if actor, ok := app.(PortModReactor); ok {
			actor.PortMod(t)
		}
	case *TableMod:
		if actor, ok := app.(TableModReactor); ok {
			actor.TableMod(t)
		}
	case *StatsRequest:
		if actor, ok := app.(StatsRequestReactor); ok {
			actor.StatsRequest(t)
		}
	case *StatsReply:
		if actor, ok := app.(StatsReplyReactor); ok {
			actor.StatsReply(dpid, t)
		}
	case *QueueGetConfigRequest:
		if actor, ok := app.(QueueGetConfigRequestReactor); ok {
			actor.QueueGetConfigRequest(t)
		}
	case *QueueGetConfigReply:
		if actor, ok := app.(QueueGetConfigReplyReactor); ok {
			actor.QueueGetConfigReply(dpid, t)
		}
	}

}
