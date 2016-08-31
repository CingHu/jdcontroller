package ofp10

import (
	//"jd.com/jdcontroller/lib/buffer"
	"jd.com/jdcontroller/lib/packet/eth"
)

/* OpenFlow Switch Specification      */
/* Version 1.0.0 (Wire Protocol 0x01) */
/* December 31, 2009                  */

const (
	Version = 1
)

/* 5.1 OpenFlow Header */
type Header struct {
	Version uint8
	Type    uint8
	Length  uint16
	XID     uint32
}

// Type
const (
	// Immutable messages.
	OFPTHello = iota
	OFPTError
	OFPTEchoRequest
	OFPTEchoReply
	OFPTVendor

	// Switch configuration messages.
	OFPTFeaturesRequest
	OFPTFeaturesReply
	OFPTGetConfigRequest
	OFPTGetConfigReply
	OFPTSetConfig

	// Asynchronous messages.
	OFPTPacketIn
	OFPTFlowRemoved
	OFPTPortStatus

	// Controller command messages.
	OFPTPacketOut
	OFPTFlowMod
	OFPTPortMod

	// Statistics messages.
	OFPTStatsRequest
	OFPTStatsReply

	// Barrier messages.
	OFPTBarrierRequest
	OFPTBarrierReply

	// Queue Configuration messages.
	OFPTQueueGetConfigRequest
	OFPTQueueGetConfigReply
)

/* 5.2 Common Structures */
// constant
const (
	OFPEthALen        = 6
	OFPMaxPortNameLen = 16
)

// 5.2.1 Port Structures
type Port struct {
	PortNO uint16
	HWAddr [OFPEthALen]byte
	Name   [OFPMaxPortNameLen]byte // Size 16

	Config uint32
	State  uint32

	Curr       uint32
	Advertised uint32
	Supported  uint32
	Peer       uint32
}

// Port config
const (
	OFPPCPortDown = 1 << 0

	OFPPCNOSTP  = 1 << 1
	OFPPCNORecv = 1 << 2

	OFPPCNOSTPRecv  = 1 << 3
	OFPPCNOFlood    = 1 << 4
	OFPPCNOFWD      = 1 << 5
	OFPPCNOPacketIn = 1 << 6
)

// Port state
const (
	OFPPSLinkDown = 1 << 0

	OFPPSSTPListen  = 0 << 8 /* Not learning or relaying frames. */
	OFPPSSTPLearn   = 1 << 8 /* Learning but not relaying frames. */
	OFPPSSTPForward = 2 << 8 /* Learning and relaying frames. */
	OFPPSSTPBlock   = 3 << 8 /* Not part of spanning tree. */
	OFPPSSTPMask    = 3 << 8 /* Bit mask for OFPPS_STP_* values. */
)

// Port numbering
const (
	OFPPMax = 0Xff00

	OFPPInPort = 0xfff8
	OFPPTable  = 0xfff9

	OFPPNormal = 0xfffa
	OFPPFlood  = 0xfffb

	OFPPAll        = 0xfffc
	OFPPController = 0xfffd
	OFPPLocal      = 0xfffe
	OFPPNone       = 0xffff
)

// Port features
const (
	OFPPF10MBHD  = 1 << 0
	OFPPF10MBFD  = 1 << 1
	OFPPF100MBHD = 1 << 2
	OFPPF100MBFD = 1 << 3
	OFPPF1GBHD   = 1 << 4
	OFPPF1GBFD   = 1 << 5
	OFPPF10GBFD  = 1 << 6

	OFPPFCopper    = 1 << 7
	OFPPFFiber     = 1 << 8
	OFPPFAutoneg   = 1 << 9
	OFPPFPause     = 1 << 10
	OFPPFPauseAsym = 1 << 11
)

// 5.2.2 Queue Structures
type PacketQueue struct {
	QueueID    uint32
	Length     uint16
	pad        [2]byte // Size 2
	Properties []QueueProp
}

// Queue properties
const (
	OFPQTNone = iota
	OFPQTMinRate
)

type QueueProp interface {
	PackBinary() (data []byte, err error)
	UnpackBinary(data []byte) (err error)
	Len() int
}

type QueuePropHeader struct {
	Property uint16
	Length   uint16
	pad      [4]byte // Size 4
}

type QueuePropMinRate struct {
	Header QueuePropHeader
	Rate   uint16
	pad    [6]byte // Size 6
}

// 5.2.3 Flow Match Structures
type Match struct {
	Wildcards uint32           /* Wildcard fields. */
	InPort    uint16           /* Input switch port. */
	DLSrc     [OFPEthALen]byte //[ETH_ALEN]uint8 /* Ethernet source address. */
	DLDst     [OFPEthALen]byte //[ETH_ALEN]uint8 /* Ethernet destination address. */
	DLVLAN    uint16           /* Input VLAN id. */
	DLVLANPCP uint8            /* Input VLAN priority. */
	pad       [1]byte          /* Align to 64-bits Size 1 */
	DLType    uint16           /* Ethernet frame type. */
	NWTos     uint8            /* IP ToS (actually DSCP field, 6 bits). */
	NWProto   uint8            /* IP protocol or lower 8 bits of ARP opcode. */
	pad2      [2]byte          /* Align to 64-bits Size 2 */
	NWSrc     [4]byte          /* IP source address. */
	NWDst     [4]byte          /* IP destination address. */
	TPSrc     uint16           /* TCP/UDP source port. */
	TPDst     uint16           /* TCP/UDP destination port. */
}

// Flow wildcards
const (
	OFPFWInPort  = 1 << 0
	OFPFWDLVLAN  = 1 << 1
	OFPFWDLSrc   = 1 << 2
	OFPFWDLDst   = 1 << 3
	OFPFWDLType  = 1 << 4
	OFPFWNWProto = 1 << 5
	OFPFWTPSrc   = 1 << 6
	OFPFWTPDst   = 1 << 7

	OFPFWNWSrcShift = 8
	OFPFWNWSrcBits  = 6
	OFPFWNWSrcMask  = ((1 << OFPFWNWSrcBits) - 1) << OFPFWNWSrcShift
	OFPFWNWSrcAll   = 32 << OFPFWNWSrcShift

	OFPFWNWDstShift = 14
	OFPFWNWDstBits  = 6
	OFPFWNWDstMask  = ((1 << OFPFWNWDstBits) - 1) << OFPFWNWDstShift
	OFPFWNWDstAll   = 32 << OFPFWNWDstShift

	OFPFWDLVLANPCP = 1 << 20
	OFPFWNWTos     = 1 << 21

	OFPFWAll = ((1 << 22) - 1)
)

// 5.2.4 Flow Action Structures
// Action type
const (
	OFPATOutput     = iota // Output to switch port.
	OFPATSetVLANVID        // Set the 802.1q VLAN id.
	OFPATSetVLANPCP        // Set the 802.1q priority.
	OFPATStripVLAN         // Strip the 802.1q header.
	OFPATSetDLSrc          // Set ethernet source address.
	OFPATSetDLDst          // Set ethernet destination address.
	OFPATSetNWSrc          // Set IP source address.
	OFPATSetNWDst          // Set IP destination address.
	OFPATSetNWTos          // Set IP ToS (DSCP field, 6bits).
	OFPATSetTPSrc          // Set TCP/UDP source port.
	OFPATSetTPDst          // Set TCP/UDP destination port.
	OFPATEnqueue           // Output to queue.
	OFPATVendor     = 0xffff
)

type Action interface {
	PackBinary() (data []byte, err error)
	UnpackBinary(data []byte) (err error)
	Len() int
}

type ActionHeader struct {
	Type   uint16
	Length uint16
}

type ActionOutput struct {
	Header ActionHeader
	Port   uint16
	MaxLen uint16
}

type ActionEnqueue struct {
	Header  ActionHeader
	Port    uint16
	pad     [6]byte // Size 6
	QueueID uint32
}

type ActionVLANVID struct {
	Header  ActionHeader
	VLANVID uint16
	pad     [2]byte // Size 2
}

type ActionVLANPCP struct {
	Header  ActionHeader
	VLANPCP uint8
	pad     [3]byte // Size 3
}

type ActionDLAddr struct {
	Header ActionHeader
	DLAddr [OFPEthALen]byte
	pad    [6]byte // Size 6
}

type ActionNWAddr struct {
	Header ActionHeader
	NWAddr [4]byte
}

type ActionNWTOS struct {
	Header ActionHeader
	NWTOS  uint8
	pad    [3]byte // Size 3
}

type ActionTPPort struct {
	Header ActionHeader
	TPPort uint16
	pad    [2]byte // Size 2
}

type ActionVendorHeader struct {
	Header ActionHeader
	Vendor uint32
}

/* 5.3 Controller to Switch Messages */
// 5.3.1 Handshake
type SwitchFeatures struct {
	Header       Header
	DPID         uint64
	Buffers      uint32
	Tables       uint8
	pad          [3]byte // Size 3
	Capabilities uint32
	Actions      uint32
	Ports        []Port
}

// Capabilities
const (
	OFPCFlowStats  = 1 << 0 // Flow statistics.
	OFPCTableStats = 1 << 1 // Table statistics.
	OFPCPortStats  = 1 << 2 // Port statistics.
	OFPCSTP        = 1 << 3 // 802.1d spanning tree.
	OFPCReserved   = 1 << 4 // Reserved, must not be set.
	OFPCIPReasm    = 1 << 5 // Can reassemble IP fragments.
	OFPCQueueStats = 1 << 6 // Queue statistics.
	OFPCARPMatchIP = 1 << 7 // Match IP address in ARP packets.
)

// 5.3.2 Switch Configuration
// Switch config
type SwitchConfig struct {
	Header      Header
	Flags       uint16
	MissSendLen uint16
}

// Config flags
const (
	OFPCFragNormal = iota
	OFPCFragDrop
	OFPCFragReasm
	OFPCFragMask
)

// 5.3.3 Modify State Messages
// Flow mod
type FlowMod struct {
	Header      Header
	Match       Match
	Cookie      uint64
	Command     uint16
	IdleTimeout uint16
	HardTimeout uint16
	Priority    uint16
	BufferID    uint32
	OutPort     uint16
	Flags       uint16
	Actions     []Action
}

// Flow mod command
const (
	OFPFCAdd          = iota // New flow.
	OFPFCModify              // Modify all matching flow.
	OFPFCModifyStrict        // Modify entry strictly matching wildcards.
	OFPFCDelete              // Delete all matching flow.
	OFPFCDeleteStrict        // Strictly match wildcards and priority.
)

// Flow mod flags
const (
	OFPFFSendFlowRem  = 1 << 0
	OFPFFCheckOverlap = 1 << 1
	OFPFFEmerg        = 1 << 2
)

// Port mod
type PortMod struct {
	Header    Header
	PortNO    uint16
	HWAddr    [OFPEthALen]uint8
	Config    uint32
	Mask      uint32
	Advertise uint32
	pad       [4]byte // Size 4
}

// 5.3.4 Queue Configuration Messages
type QueueGetConfigRequest struct {
	Header Header
	Port   uint16
	pad    [2]byte // Size 2
}

type QueueGetConfigReply struct {
	Header Header
	Port   uint16
	pad    [6]byte // Size 6
	Queues []PacketQueue
}

// 5.3.5 Read State Messages
// constats
const (
	OFPDescStrLen      = 256
	OFPSerialNumLen    = 32
	OFPMaxTableNameLen = 32
)

// Stats request
type StatsRequest struct {
	Header Header
	Type   uint16
	Flags  uint16
	//Body   buffer.Message
	Body   interface{}
}

// Stats reply
type StatsReply struct {
	Header Header
	Type   uint16
	Flags  uint16
	//Body   buffer.Message
	Body   interface{}
}

// Stats types
const (
	OFPSTDesc = iota
	OFPSTFlow
	OFPSTAggregate
	OFPSTTable
	OFPSTPort
	OFPSTQueue
	OFPSTVendor = 0xffff
)

// Desc stats
type DescStats struct {
	MfrDesc   [OFPDescStrLen]byte   // Size 256
	HWDesc    [OFPDescStrLen]byte   // Size 256
	SWDesc    [OFPDescStrLen]byte   // Size 256
	SerialNum [OFPSerialNumLen]byte // Size 32
	DPDesc    [OFPDescStrLen]byte   // Size 256
}

// Flow stats request
type FlowStatsRequest struct {
	Match   Match
	TableID uint8
	pad     [1]byte // Size 1
	OutPort uint16
}

// Flow stats
type FlowStats struct {
	Length       uint16
	TableID      uint8
	pad          [1]byte // Size 1
	Match        Match
	DurationSec  uint32
	DurationNSec uint32
	Priority     uint16
	IdleTimeout  uint16
	HardTimeout  uint16
	pad2         [6]byte // Size 6
	Cookie       uint64
	PacketCount  uint64
	ByteCount    uint64
	Actions      []Action
}

// Aggregate stats request
type AggregateStatsRequest struct {
	Match   Match
	TableID uint8
	pad     [1]byte // Size 1
	OutPort uint16
}

// Aggregate stats reply
type AggregateStatsReply struct {
	PacketCount uint64
	ByteCount   uint64
	FlowCount   uint32
	pad         [4]byte // Size 4
}

// Table stats
type TableStats struct {
	TableID      uint8
	pad          [3]byte                  // Size 3
	Name         [OFPMaxTableNameLen]byte // Size 32
	Wildcards    uint32
	MaxEntries   uint32
	ActiveCount  uint32
	LookupCount  uint64
	MatchedCount uint64
}

// Port stats request
type PortStatsRequest struct {
	PortNO uint16
	pad    [6]byte // Size 6
}

// Port stats
type PortStats struct {
	PortNO     uint16
	pad        [6]byte // Size 6
	RxPackets  uint64
	TxPackets  uint64
	RxBytes    uint64
	TxBytes    uint64
	RxDropped  uint64
	TxDropped  uint64
	RxErrors   uint64
	TxErrors   uint64
	RxFrameErr uint64
	RxOverErr  uint64
	RxCRCErr   uint64
	Collisions uint64
}

// Queue stats request
type QueueStatsRequest struct {
	PortNO  uint16
	pad     [2]byte // Size 2
	QueueID uint32
}

// Queue stats
type QueueStats struct {
	PortNO    uint16
	pad       [2]byte // Size 2
	QueueID   uint32
	TxBytes   uint64
	TxPackets uint64
	TxErrors  uint64
}

// 5.3.6 Send Packet Message
type PacketOut struct {
	Header     Header
	BufferID   uint32
	InPort     uint16
	ActionsLen uint16
	Actions    []Action
	Data        interface{}
}

/* 5.4 Asynchronous Messages */
// 5.4.1 Packet-In Message
type PacketIn struct {
	Header   Header
	BufferID uint32
	TotalLen uint16
	InPort   uint16
	Reason   uint8
	pad      [1]byte // Size 1
	//	Data     buffer.Message
	Data eth.Ethernet
}

// Packet-in reason
const (
	OFPRNoMatch = iota
	OFPRAction
)

// 5.4.2 Flow Removed Message
type FlowRemoved struct {
	Header   Header
	Match    Match
	Cookie   uint64
	Priority uint16
	Reason   uint8
	pad      [1]byte // Size 1

	DurationSec  uint32
	DurationNSec uint32

	IdleTimeout uint16
	pad2        [2]byte // Size 2
	PacketCount uint64
	ByteCount   uint64
}

// Removed reason
const (
	OFPRRIdleTimeout = iota
	OFPRRHardTimeout
	OFPRRDelete
)

// 5.4.3 Port Status Message
type PortStatus struct {
	Header Header
	Reason uint8
	pad    [7]uint8 // Size 7
	Desc   Port
}

// Port reason
const (
	OFPPRAdd = iota
	OFPPRDelete
	OFPPRModify
)

// 5.4.4 Error Message
type ErrorMsg struct {
	Header Header
	Type   uint16
	Code   uint16
	Data   []uint8
}

// Error type
const (
	OFPETHelloFailed = iota
	OFPETBadRequest
	OFPETBadAction
	OFPETFlowModFailed
	OFPETPortModFailed
	OFPETQueueOPFailed
)

// Hello failed code
const (
	OFPHFCIncompatible = iota
	OFPHFCEperm
)

// Bad request code
const (
	OFPBRCBadVersion = iota
	OFPBRCBadType
	OFPBRCBadStat
	OFPBRCBadVendor

	OFPBRCBadSubtype
	OFPBRCEperm
	OFPBRCBadLen
	OFPBRCBufferEmpty
	OFPBRCBufferUnknown
)

// Bad action code
const (
	OFPBACBadType = iota
	OFPBACBadLen
	OFPBACBadVendor
	OFPBACBadVendorType
	OFPBACBadOutPort
	OFPBACBadArgument
	OFPBACEperm
	OFPBACTooMany
	OFPBACBadQueue
)

// Flow mod failed code
const (
	OFPFMFCAllTablesFull = iota
	OFPFMFCOverlap
	OFPFMFCEperm
	OFPFMFCBadEmergTimeout
	OFPFMFCBadCommand
	OFPFMFCUnsupported
)

// Port mod failed code
const (
	OFPPMFCBadPort = iota
	OFPPMFCBadHWAddr
)

// Queue op failed code
const (
	OFPQOFCBadPort = iota
	OFPQOFCBadQueue
	OFPQOFCEperm
)

/* 5.5 Symmetric Messages */
type VendorHeader struct {
	Header Header /*Type OFPT_VENDOR*/
	Vendor uint32
}
