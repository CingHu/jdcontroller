package ofpctl10

import (
	"jd.com/jdcontroller/lib/buffer"
	"jd.com/jdcontroller/protocol/ofp10"
)

type ConnectionUpReactor interface {
	ConnectionUp(dpid uint64)
}

type ConnectionDownReactor interface {
	ConnectionDown(dpid uint64, err error)
}

type HelloReactor interface {
	Hello(hello *ofp10.Header)
}

type ErrorReactor interface {
	Error(dpid uint64, err *ofp10.ErrorMsg)
}

type EchoRequestReactor interface {
	EchoRequest(dpid uint64)
}

type EchoReplyReactor interface {
	EchoReply(dpid uint64)
}

type VendorReactor interface {
	VendorHeader(dpid uint64, v *ofp10.VendorHeader)
}

type FeaturesRequestReactor interface {
	FeaturesRequest(features *ofp10.Header)
}

type FeaturesReplyReactor interface {
	FeaturesReply(dpid uint64, features *ofp10.SwitchFeatures)
}

type GetConfigRequestReactor interface {
	GetConfigRequest(config *ofp10.Header)
}

type GetConfigReplyReactor interface {
	GetConfigReply(dpid uint64, config *ofp10.SwitchConfig)
}

type SetConfigReactor interface {
	SetConfig(config *ofp10.SwitchConfig)
}

type PacketInReactor interface {
	PacketIn(dpid uint64, packet *ofp10.PacketIn)
}

type FlowRemovedReactor interface {
	FlowRemoved(dpid uint64, flow *ofp10.FlowRemoved)
}

type PortStatusReactor interface {
	PortStatus(dpid uint64, status *ofp10.PortStatus)
}

type PacketOutReactor interface {
	PacketOut(packet *ofp10.PacketOut)
}

type FlowModReactor interface {
	FlowMod(flowMod *ofp10.FlowMod)
}

type PortModReactor interface {
	PortMod(portMod *ofp10.PortMod)
}

type StatsRequestReactor interface {
	StatsRequest(req *ofp10.StatsRequest)
}

type StatsReplyReactor interface {
	StatsReply(dpid uint64, rep *ofp10.StatsReply)
}

type BarrierRequestReactor interface {
	BarrierRequest(req *ofp10.Header)
}

type BarrierReplyReactor interface {
	BarrierReply(dpid uint64, msg *ofp10.Header)
}

func ReactorParse(dpid uint64, app interface{}, msg buffer.Message) {
	switch t := msg.(type) {
	case *ofp10.Header:
		switch t.Type {
		case ofp10.OFPTHello:
			if actor, ok := app.(HelloReactor); ok {
				actor.Hello(t)
			}
		case ofp10.OFPTEchoRequest:
			if actor, ok := app.(EchoRequestReactor); ok {
				actor.EchoRequest(dpid)
			}
		case ofp10.OFPTEchoReply:
			if actor, ok := app.(EchoReplyReactor); ok {
				actor.EchoReply(dpid)
			}
		case ofp10.OFPTFeaturesRequest:
			if actor, ok := app.(FeaturesRequestReactor); ok {
				actor.FeaturesRequest(t)
			}
		case ofp10.OFPTGetConfigRequest:
			if actor, ok := app.(GetConfigRequestReactor); ok {
				actor.GetConfigRequest(t)
			}
		case ofp10.OFPTBarrierRequest:
			if actor, ok := app.(BarrierRequestReactor); ok {
				actor.BarrierRequest(t)
			}
		case ofp10.OFPTBarrierReply:
			if actor, ok := app.(BarrierReplyReactor); ok {
				actor.BarrierReply(dpid, t)
			}
		}
	case *ofp10.ErrorMsg:
		if actor, ok := app.(ErrorReactor); ok {
			actor.Error(dpid, t)
		}
	case *ofp10.VendorHeader:
		if actor, ok := app.(VendorReactor); ok {
			actor.VendorHeader(dpid, t)
		}
	case *ofp10.SwitchFeatures:
		if actor, ok := app.(FeaturesReplyReactor); ok {
			actor.FeaturesReply(dpid, t)
		}
	case *ofp10.SwitchConfig:
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
	case *ofp10.PacketIn:
		if actor, ok := app.(PacketInReactor); ok {
			actor.PacketIn(dpid, t)
		}
	case *ofp10.FlowRemoved:
		if actor, ok := app.(FlowRemovedReactor); ok {
			actor.FlowRemoved(dpid, t)
		}
	case *ofp10.PortStatus:
		if actor, ok := app.(PortStatusReactor); ok {
			actor.PortStatus(dpid, t)
		}
	case *ofp10.PacketOut:
		if actor, ok := app.(PacketOutReactor); ok {
			actor.PacketOut(t)
		}
	case *ofp10.FlowMod:
		if actor, ok := app.(FlowModReactor); ok {
			actor.FlowMod(t)
		}
	case *ofp10.PortMod:
		if actor, ok := app.(PortModReactor); ok {
			actor.PortMod(t)
		}
	case *ofp10.StatsRequest:
		if actor, ok := app.(StatsRequestReactor); ok {
			actor.StatsRequest(t)
		}
	case *ofp10.StatsReply:
		if actor, ok := app.(StatsReplyReactor); ok {
			actor.StatsReply(dpid, t)
		}
	}

}
