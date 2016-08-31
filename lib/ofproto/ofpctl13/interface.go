package ofpctl13

import (
	"jd.com/jdcontroller/lib/buffer"
	"jd.com/jdcontroller/protocol/ofp13"
)

type ConnectionUpReactor interface {
	ConnectionUp(dpid uint64)
}

type ConnectionDownReactor interface {
	ConnectionDown(dpid uint64, err error)
}

type HelloReactor interface {
	Hello(hello *ofp13.Header)
}

type ErrorReactor interface {
	Error(dpid uint64, err *ofp13.ErrorMsg)
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
	FeaturesRequest(features *ofp13.Header)
}

type FeaturesReplyReactor interface {
	FeaturesReply(dpid uint64, features *ofp13.SwitchFeatures)
}

type GetConfigRequestReactor interface {
	GetConfigRequest(config *ofp13.Header)
}

type GetConfigReplyReactor interface {
	GetConfigReply(dpid uint64, config *ofp13.SwitchConfig)
}

type QueueGetConfigRequestReactor interface {
	QueueGetConfigRequest(config *ofp13.QueueGetConfigRequest)
}

type QueueGetConfigReplyReactor interface {
	QueueGetConfigReply(dpid uint64, config *ofp13.QueueGetConfigReply)
}

type RoleRequestReactor interface {
	RoleRequest(req *ofp13.RoleRequest)
}

type RoleReplyReactor interface {
	RoleReply(dpid uint64, req *ofp13.RoleRequest)
}

type SetConfigReactor interface {
	SetConfig(config *ofp13.SwitchConfig)
}

type PacketInReactor interface {
	PacketIn(dpid uint64, packet *ofp13.PacketIn)
}

type FlowRemovedReactor interface {
	FlowRemoved(dpid uint64, flow *ofp13.FlowRemoved)
}

type PortStatusReactor interface {
	PortStatus(dpid uint64, status *ofp13.PortStatus)
}

type PacketOutReactor interface {
	PacketOut(packet *ofp13.PacketOut)
}

type FlowModReactor interface {
	FlowMod(flowMod *ofp13.FlowMod)
}

type GroupModReactor interface {
	GroupMod(groupMod *ofp13.GroupMod)
}

type PortModReactor interface {
	PortMod(portMod *ofp13.PortMod)
}

type TableModReactor interface {
	TableMod(tableMod *ofp13.TableMod)
}

type BarrierRequestReactor interface {
	BarrierRequest(req *ofp13.Header)
}

type BarrierReplyReactor interface {
	BarrierReply(dpid uint64, msg *ofp13.Header)
}

type MultipartRequestReactor interface {
	MultipartRequest(req *ofp13.MultipartRequest)
}

type MultipartReplyReactor interface {
	MultipartReply(dpid uint64, req *ofp13.MultipartReply)
}

type SetAsyncReactor interface {
	SetAsync(config *ofp13.AsyncConfig)
}

type GetAsyncRequestReactor interface {
	GetAsyncRequest(req *ofp13.Header)
}

type GetAsyncReplyReactor interface {
	GetAsyncReply(dpid uint64, config *ofp13.AsyncConfig)
}

type MeterModReactor interface {
	MeterMod(meterMod *ofp13.MeterMod)
}

func ReactorParse(dpid uint64, app interface{}, msg buffer.Message) {
	switch t := msg.(type) {
	case *ofp13.Header:
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
		case OFPTGetAsyncRequest:
			if actor, ok := app.(BarrierReplyReactor); ok {
				actor.BarrierReply(dpid, t)
			}
		}
	case *ofp13.ErrorMsg:
		if actor, ok := app.(ErrorReactor); ok {
			actor.Error(dpid, t)
		}
	case *ofp13.ExperimenterHeader:
		if actor, ok := app.(ExperimenterReactor); ok {
			actor.Experimenter(dpid)
		}
	case *ofp13.SwitchFeatures:
		if actor, ok := app.(FeaturesReplyReactor); ok {
			actor.FeaturesReply(dpid, t)
		}
	case *ofp13.SwitchConfig:
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
	case *ofp13.PacketIn:
		if actor, ok := app.(PacketInReactor); ok {
			actor.PacketIn(dpid, t)
		}
	case *ofp13.FlowRemoved:
		if actor, ok := app.(FlowRemovedReactor); ok {
			actor.FlowRemoved(dpid, t)
		}
	case *ofp13.PortStatus:
		if actor, ok := app.(PortStatusReactor); ok {
			actor.PortStatus(dpid, t)
		}
	case *ofp13.PacketOut:
		if actor, ok := app.(PacketOutReactor); ok {
			actor.PacketOut(t)
		}
	case *ofp13.FlowMod:
		if actor, ok := app.(FlowModReactor); ok {
			actor.FlowMod(t)
		}
	case *ofp13.GroupMod:
		if actor, ok := app.(GroupModReactor); ok {
			actor.GroupMod(t)
		}
	case *ofp13.PortMod:
		if actor, ok := app.(PortModReactor); ok {
			actor.PortMod(t)
		}
	case *ofp13.TableMod:
		if actor, ok := app.(TableModReactor); ok {
			actor.TableMod(t)
		}
	case *ofp13.MultipartRequest:
		if actor, ok := app.(MultipartRequestReactor); ok {
			actor.MultipartRequest(t)
		}
	case *ofp13.MultipartReply:
		if actor, ok := app.(MultipartReplyReactor); ok {
			actor.MultipartReply(dpid, t)
		}
	case *ofp13.QueueGetConfigRequest:
		if actor, ok := app.(QueueGetConfigRequestReactor); ok {
			actor.QueueGetConfigRequest(t)
		}
	case *ofp13.QueueGetConfigReply:
		if actor, ok := app.(QueueGetConfigReplyReactor); ok {
			actor.QueueGetConfigReply(dpid, t)
		}
	case *ofp13.RoleRequest:
		switch t.Header.Type {
		case OFPTRoleRequest:
			if actor, ok := app.(RoleRequestReactor); ok {
				actor.RoleRequest(t)
			}
		case OFPTRoleReply:
			if actor, ok := app.(RoleReplyReactor); ok {
				actor.RoleReply(dpid, t)
			}
		}
	case *ofp13.AsyncConfig:
		switch t.Header.Type {
		case OFPTSetAsync:
			if actor, ok := app.(SetAsyncReactor); ok {
				actor.SetAsync(t)
			}
		case OFPTGetAsyncReply:
			if actor, ok := app.(GetAsyncReplyReactor); ok {
				actor.GetAsyncReply(dpid, t)
			}
		}
	case *ofp13.MeterMod:
		if actor, ok := app.(MeterModReactor); ok {
			actor.MeterMod(t)
		}
	}
}
