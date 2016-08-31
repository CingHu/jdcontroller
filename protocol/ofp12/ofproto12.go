package ofp12

import (
	"jd.com/jdcontroller/lib/buffer"
	"jd.com/jdcontroller/lib/packet/eth"
)

/* OpenFlow Switch Specification      */
/* Version 1.2 (Wire Protocol 0x03)   */
/* December 5, 2011                   */

const (
	Version = 3
)

/* A.1 OpenFlow Header */
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
	OFPTExperimenter

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
	OFPTGroupMod
	OFPTPortMod
	OFPTTableMod

	// Statistics messages.
	OFPTStatsRequest
	OFPTStatsReply

	// Barrier messages.
	OFPTBarrierRequest
	OFPTBarrierReply

	// Queue Configuration messages.
	OFPTQueueGetConfigRequest
	OFPTQueueGetConfigReply

	// Controller role change request messages.
	OFPTRoleRequest
	OFPTRoleReply
)

/* A.2 Common Structures */
// constant
const (
	OFPEthALen        = 6
	OFPMaxPortNameLen = 16
)

// A.2.1 Port Structures
type Port struct {
	PortNO uint32
	pad    [4]byte // Size 4
	HWAddr [OFPEthALen]byte
	pad2   [2]byte                 // Size 2
	Name   [OFPMaxPortNameLen]byte // Size 16

	Config uint32
	State  uint32

	Curr       uint32
	Advertised uint32
	Supported  uint32
	Peer       uint32

	CurrSpeed uint32
	MaxSpeed  uint32
}

// Port config
const (
	OFPPCPortDown = 1 << 0

	OFPPCNORecv     = 1 << 2
	OFPPCNOFWD      = 1 << 5
	OFPPCNOPacketIn = 1 << 6
)

// Port state
const (
	OFPPSLinkDown = 1 << 0
	OFPPSBlocked  = 1 << 1
	OFPPSLive     = 1 << 2
)

// Port numbering
const (
	OFPPMax = 0xffffff00

	OFPPInPort = 0xfffffff8
	OFPPTable  = 0xfffffff9

	OFPPNormal = 0xfffffffa
	OFPPFlood  = 0xfffffffb

	OFPPAll        = 0xfffffffc
	OFPPController = 0xfffffffd
	OFPPLocal      = 0xfffffffe
	OFPPAny        = 0xffffffff
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
	OFPPF40GBFD  = 1 << 7
	OFPPF100GBFD = 1 << 8
	OFPPF1TBFD   = 1 << 9
	OFPPFOther   = 1 << 10

	OFPPFCopper    = 1 << 11
	OFPPFFiber     = 1 << 12
	OFPPFAutoneg   = 1 << 13
	OFPPFPause     = 1 << 14
	OFPPFPauseAsym = 1 << 15
)

// A.2.2 Queue Structures
type PacketQueue struct {
	QueueID    uint32
	Port       uint32
	Length     uint16
	pad        [6]byte // Size 6
	Properties []QueueProp
}

// Queue properties
const (
	OFPQTMinRate      = 1
	OFPQTMaxRate      = 2
	OFPQTExperimenter = 0xffff
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

type QueuePropMaxRate struct {
	Header QueuePropHeader
	Rate   uint16
	pad    [6]byte // Size 6
}

type QueuePropExperimenter struct {
	Header       QueuePropHeader
	Experimenter uint32
	pad          [4]byte // Size 4
	Data         []uint8
}

// A.2.3 Flow Match Structures
// constant
const (
	OFPMTStandardLen = 88
)

// A.2.3.1 Flow Match Header
type Match struct {
	Type      uint16
	Length    uint16
	OXMFields [4]byte // Size 4
}

// Match type
const (
	OFPMTStandard = iota
	OFPMTOXM
)

// A.2.3.3 OXM classes
// OXM class
const (
	OFPXMCNXM0          = 0x0000
	OFPXMCNXM1          = 0x0001
	OFPXMCOpenFlowBasic = 0x8000
	OFPXMCExperimenter  = 0xffff
)

// A.2.3.7 Flow Match Fields
// oxm flow match field type
const (
	OFPXMTOFBInPort = iota
	OFPXMTFOFBInPhyPort
	OFPXMTFOFBMetadata
	OFPXMTFOFBEthDst
	OFPXMTFOFBEthSrc
	OFPXMTFOFBEthType
	OFPXMTFOFBVLANVID
	OFPXMTFOFBVLANPCP
	OFPXMTFOFBIPDSCP
	OFPXMTFOFBIPECN
	OFPXMTFOFBIPProto
	OFPXMTFOFBIPv4Src
	OFPXMTFOFBIPv4Dst
	OFPXMTFOFBTCPSrc
	OFPXMTFOFBTCPDst
	OFPXMTFOFBUDPSrc
	OFPXMTFOFBUDPDst
	OFPXMTFOFBSCTPSrc
	OFPXMTFOFBSCTPDst
	OFPXMTFOFBICMPv4Type
	OFPXMTFOFBICMPv4Code
	OFPXMTFOFBARPOP
	OFPXMTFOFBARPSpa
	OFPXMTFOFBARPTpa
	OFPXMTFOFBARPSha
	OFPXMTFOFBARPTha
	OFPXMTFOFBIPv6Src
	OFPXMTFOFBIPv6Dst
	OFPXMTFOFBIPv6FLabel
	OFPXMTFOFBICMPv6Type
	OFPXMTFOFBICMPv6Code
	OFPXMTFOFBIPv6NDTarget
	OFPXMTFOFBIPv6NDSll
	OFPXMTFOFBIPv6NDTll
	OFPXMTFOFBMPLSLabel
	OFPXMTFOFBMPLSTC
)

// VLAN id
const (
	OFPVIDPresent = 0x1000
	OFPVIDNone    = 0x0000
)

// A.2.3.8 Experimenter Flow Match Fields
type OXMExperimenterHeader struct {
	OXMHeader    uint32
	Experimenter uint32
}

// A.2.4 Flow Instruction Structures
// Instruction type
const (
	OFPITGotoTable     = 1
	OFPITWriteMetadata = 2
	OFPITWriteActions  = 3
	OFPITApplyActions  = 4
	OFPITClearActions  = 5
	OFPITExperimenter  = 0xffff
)

type Instruction interface {
	PackBinary() (data []byte, err error)
	UnpackBinary(data []byte) (err error)
	Len() int
}

type InstructionHeader struct {
	Type   uint16
	Length uint16
}

type InstructionGotoTable struct {
	Header  InstructionHeader
	TableID uint8
	pad     [3]byte // Size 3
}

type InstructionWriteMetadata struct {
	Header       InstructionHeader
	pad          [4]byte // Size 4
	Metadata     uint64
	MetadataMask uint64
}

type InstructionActions struct {
	Header  InstructionHeader
	pad     [4]byte // Size 4
	Actions []Action
}

// A.2.5 Action Structures
// Action type
const (
	OFPATOutput     = 0  // Output to switch port.
	OFPATCopyTTLOut = 11 // Set TCP/UDP destination port.
	OFPATCopyTTLIn  = 12 // Set TCP/UDP destination port.
	OFPATSetMPLSTTL = 15 // Set TCP/UDP destination port.
	OFPATDecMPLSTTL = 16 // Set TCP/UDP destination port.

	OFPATPushVLAN = 17 // Set TCP/UDP destination port.
	OFPATPopVLAN  = 18 // Set TCP/UDP destination port.
	OFPATPushMPLS = 19 // Set TCP/UDP destination port.
	OFPATPopMPLS  = 20 // Set TCP/UDP destination port.

	OFPATSetQueue = 21 // Set TCP/UDP destination port.
	OFPATGroup    = 22 // Set TCP/UDP destination port.
	OFPATSetNWTTL = 23 // Set TCP/UDP destination port.
	OFPATDecNWTTL = 24 // Set TCP/UDP destination port.
	OFPATSetField = 25 // Set TCP/UDP destination port.

	OFPATExperimenter = 0xffff
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
	Port   uint32
	MaxLen uint16
	pad    [6]byte // Size 6
}

// Controller max len
const (
	OFPCMLMax      = 0xffe5
	OFPCMLNoBuffer = 0xffff
)

type ActionGroup struct {
	Header  ActionHeader
	GroupID uint32
}

type ActionSetQueue struct {
	Header  ActionHeader
	QueueID uint32
}

type ActionMPLSTTL struct {
	Header  ActionHeader
	MPLSTTL uint8
	pad     [3]byte // Size 3
}

type ActionNWTTL struct {
	Header ActionHeader
	NWTTL  uint8
	pad    [3]byte // Size 3
}

type ActionPush struct {
	Header    ActionHeader
	Ethertype uint16
	pad       [2]byte // Size 2
}

type ActionPopMPLS struct {
	Header    ActionHeader
	Ethertype uint16
	pad       [2]byte // Size 2
}

type ActionSetField struct {
	Header ActionHeader
	Field  [4]byte // Size 4
}

type ActionExperimenterHeader struct {
	Header       ActionHeader
	Experimenter uint32
}

/* A.3 Controller to Switch Messages */
// A.3.1 Handshake
type SwitchFeatures struct {
	Header       Header
	DPID         uint64
	Buffers      uint32
	Tables       uint8
	pad          [3]byte // Size 3
	Capabilities uint32
	Reserved     uint32
	Ports        []Port
}

// Capabilities
const (
	OFPCFlowStats   = 1 << 0 // Flow statistics.
	OFPCTableStats  = 1 << 1 // Table statistics.
	OFPCPortStats   = 1 << 2 // Port statistics.
	OFPCGroupStats  = 1 << 3 // 802.1d spanning tree.
	OFPCIPReasm     = 1 << 5 // Can reassemble IP fragments.
	OFPCQueueStats  = 1 << 6 // Queue statistics.
	OFPCPortBlocked = 1 << 8 // Match IP address in ARP packets.
)

// A.3.2 Switch Configuration
// Switch config
type SwitchConfig struct {
	Header      Header
	Flags       uint16
	MissSendLen uint16
}

// Config flags
const (
	OFPCFragNormal             = 0
	OFPCFragDrop               = 1 << 0
	OFPCFragReasm              = 1 << 1
	OFPCFragMask               = 3
	OFPCInvalidTTLToController = 1 << 2
)

// A.3.3 Flow Table Configuration
// Table numbering
const (
	OFPTTMax = 0xfe
	OFPTTAll = 0xff
)

// Table mod
type TableMod struct {
	Header  Header
	TableID uint8
	pad     [3]byte // Size 3
	Config  uint32
}

// Table config
const (
	OFPTCTableMissController = 0
	OFPTCTableMissContinue   = 1 << 0
	OFPTCTableMissDrop       = 1 << 1
	OFTCTableMissMask        = 3
)

// A.3.4 Modify State Messages
// A.3.4.1 Modify Flow Entry Message
// Flow mod
type FlowMod struct {
	Header
	Cookie     uint64
	CookieMask uint64

	TableID      uint8
	Command      uint8
	IdleTimeout  uint16
	HardTimeout  uint16
	Priority     uint16
	BufferID     uint32
	OutPort      uint32
	OutGroup     uint32
	Flags        uint16
	pad          [2]byte // Size 2
	Match        Match
	Instructions []Instruction
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
	OFPFFResetCounts  = 1 << 2
)

// A.3.4.2 Modify Group Entry Message
// Group mod
type GroupMod struct {
	Header  Header
	Command uint16
	Type    uint8
	pad     [1]byte // Size 1
	GroupID uint32
	Buckets []Bucket
}

// Group mod command
const (
	OFPGCAdd = iota
	OFPGCModify
	OFPGCDelete
)

// Group type
const (
	OFPGTAll = iota
	OFPGTSelect
	OFPGTIndirect
	OFPGTFF
)

type Bucket struct {
	Length     uint16
	Weight     uint16
	WatchPort  uint32
	WatchGroup uint32
	pad        [4]byte // Size 4
	Actions    []Action
}

// A.3.4.3 Port Modification Message
// Port mod
type PortMod struct {
	Header    Header
	PortNO    uint32  //change uint16 to uint32
	pad       [4]byte // Size 4
	HWAddr    [OFPEthALen]byte
	pad2      [2]byte // Size 2
	Config    uint32
	Mask      uint32
	Advertise uint32
	pad3      [4]byte // Size 4
}

// A.3.5 Read State Messages

// Stats request
type StatsRequest struct {
	Header Header
	Type   uint16
	Flags  uint16
	pad    [4]byte // Size 4
	Body   buffer.Message
}

// Stats reply
type StatsReply struct {
	Header Header
	Type   uint16
	Flags  uint16
	pad    [4]byte // Size 4
	Body   buffer.Message
}

// Stats types
const (
	OFPSTDesc = iota
	OFPSTFlow
	OFPSTAggregate
	OFPSTTable
	OFPSTPort
	OFPSTQueue
	OFPSTGroup
	OFPSTGroupDesc
	OFPSTGroupFeatures
	OFPSTExperimenter = 0xffff
)

// A.3.5.1 Description Statistics
// constats
const (
	OFPDescStrLen      = 256
	OFPSerialNumLen    = 32
	OFPMaxTableNameLen = 32
)

// Desc stats
type DescStats struct {
	MfrDesc   [OFPDescStrLen]byte   // Size 256
	HWDesc    [OFPDescStrLen]byte   // Size 256
	SWDesc    [OFPDescStrLen]byte   // Size 256
	SerialNum [OFPSerialNumLen]byte // Size 32
	DPDesc    [OFPDescStrLen]byte   // Size 256
}

// A.3.5.2 Individual Flow Statistics
// Flow stats request
type FlowStatsRequest struct {
	TableID    uint8
	pad        [3]byte // Size 3
	OutPort    uint32
	OutGroup   uint32
	pad2       [4]byte // Size 4
	Cookie     uint64
	CookieMask uint64
	Match      Match
}

// Flow stats
type FlowStats struct {
	Length       uint16
	TableID      uint8
	pad          [1]byte // Size 1
	DurationSec  uint32
	DurationNSec uint32
	Priority     uint16
	IdleTimeout  uint16
	HardTimeout  uint16
	pad2         [6]byte // Size 6
	Cookie       uint64
	PacketCount  uint64
	ByteCount    uint64
	Match        Match
	Instructions []Instruction
}

// A.3.5.3 Aggregate Flow Statistics
// Aggregate stats request
type AggregateStatsRequest struct {
	TableID    uint8
	pad        [3]byte // Size 3
	OutPort    uint32
	OutGroup   uint32
	pad2       [4]byte // Size 4
	Cookie     uint64
	CookieMask uint64
	Match      Match
}

// Aggregate stats reply
type AggregateStatsReply struct {
	PacketCount uint64
	ByteCount   uint64
	FlowCount   uint32
	pad         [4]byte // Size 4
}

// A.3.5.4 Table Statistics
// Table stats
type TableStats struct {
	TableID        uint8
	pad            [7]uint8                 // Size 7
	Name           [OFPMaxTableNameLen]byte // Size 32
	Match          uint64
	Wildcards      uint64
	WriteActions   uint32
	ApplyActions   uint32
	WriteSetFields uint64
	ApplySetFields uint64
	MetadataMatch  uint64
	MetadataWrite  uint64
	Instructions   uint32
	Config         uint32
	MaxEntries     uint32
	ActiveCount    uint32
	LookupCount    uint64
	MatchedCount   uint64
}

// A.3.5.5 Port Statistics
// Port stats request
type PortStatsRequest struct {
	PortNO uint32
	pad    [4]byte // Size 4
}

// Port stats
type PortStats struct {
	PortNO     uint32
	pad        [4]byte // Size 4
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

// A.3.5.6 Queue Statistics
// Queue stats request
type QueueStatsRequest struct {
	PortNO  uint32
	QueueID uint32
}

// Queue stats
type QueueStats struct {
	PortNO    uint32
	QueueID   uint32
	TxBytes   uint64
	TxPackets uint64
	TxErrors  uint64
}

// A.3.5.7 Group Statistics
// Group stats request
type GroupStatsRequest struct {
	GroupID uint32
	pad     [4]byte // Size 4
}

// Group stats
type GroupStats struct {
	Length      uint16
	pad         [2]byte // Size 2
	GroupID     uint32
	RefCount    uint32
	pad2        [4]byte // Size 4
	PacketCount uint64
	ByteCount   uint64
	BucketStats []BucketCounter
}

type BucketCounter struct {
	PacketCount uint64
	ByteCount   uint64
}

// A.3.5.8 Group Description Statistics
type GroupDescStats struct {
	Length  uint16
	Type    uint8
	pad     [1]byte // Size 1
	GroupID uint32
	Buckets []Bucket
}

// A.3.5.9 Group Features Statistics
// Group features stats
type GroupFeaturesStats struct {
	Types        uint32
	Capabilities uint32
	MaxGroups    [4]uint32 // Size 4
	Actions      [4]uint32 // Size 4
}

// Group capabilities
const (
	OFPGFCSelectWeight   = 1 << 0
	OFPGFCSelectLiveness = 1 << 1
	OFPGFCChaining       = 1 << 2
	OFPGFCChainingChecks = 1 << 3
)

// A.3.5.10 Experimenter Statistics
// experimenter stats header
type ExperimenterStatsHeader struct {
	Experimenter uint32
	ExpType      uint32
}

// A.3.6 Queue Configuration Messages
type QueueGetConfigRequest struct {
	Header Header
	Port   uint32
	pad    [4]byte // Size 4
}

type QueueGetConfigReply struct {
	Header Header
	Port   uint32
	pad    [4]byte // Size 4
	Queues []PacketQueue
}

// A.3.7 Packet-Out Message
type PacketOut struct {
	Header     Header
	BufferID   uint32
	InPort     uint32
	ActionsLen uint16
	pad        [6]byte // Size 6
	Actions    []Action
	Data       buffer.Message
}

// A.3.9 Role Request Message
// Role request and reply message
type RoleRequest struct {
	Header       Header
	Role         uint32
	pad          [4]byte // Size 4
	GenerationID uint64
}

// Controller roles
const (
	OFPCRRoleNoChange = iota
	OFPCRRoleEqual
	OFPCRRoleMaster
	OFPCRRoleSlave
)

/* A.4 Asynchronous Messages */
// A.4.1 Packet-In Message
type PacketIn struct {
	Header   Header
	BufferID uint32
	TotalLen uint16
	Reason   uint8
	TableID  uint8
	Match    Match
	pad      [2]byte // Size 2
	//	Data     buffer.Message
	Data eth.Ethernet
}

// Packet-in reason
const (
	OFPRNoMatch = iota
	OFPRAction
	OFPRInvalidTTL
)

// A.4.2 Flow Removed Message
type FlowRemoved struct {
	Header   Header
	Cookie   uint64
	Priority uint16
	Reason   uint8
	TableID  uint8

	DurationSec  uint32
	DurationNSec uint32

	IdleTimeout uint16
	HardTimeout uint16
	PacketCount uint64
	ByteCount   uint64
	Match       Match
}

// Flow removed reason
const (
	OFPRRIdleTimeout = iota
	OFPRRHardTimeout
	OFPRRDelete
	OFPRRGroupDelete
)

// A.4.3 Port Status Message
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

// A.4.4 Error Message
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
	OFPETBadInstruction
	OFPETBadMatch
	OFPETFlowModFailed
	OFPETGroupModFailed
	OFPETPortModFailed
	OFPETTableModFailed
	OFPETQueueOPFailed
	OFPETSwitchConfigFailed
	OFPETRoleRequestFailed
	OFPETExperimenter = 0xffff
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
	OFPBRCBadExperimenter

	OFPBRCBadSubtype
	OFPBRCEperm
	OFPBRCBadLen
	OFPBRCBufferEmpty
	OFPBRCBufferUnknown
	OFPBRCBadTableID

	OFPBRCIsSlave
	OFPBRCBadPort
	OFPBRCBadPacket
)

// Bad action code
const (
	OFPBACBadType = iota
	OFPBACBadLen
	OFPBACBadExperimenter
	OFPBACBadExperimenterType
	OFPBACBadOutPort
	OFPBACBadArgument
	OFPBACEperm
	OFPBACTooMany
	OFPBACBadQueue
	OFPBACBadOutGroup
	OFPBACMatchInconsistent
	OFPBACUnsupportedOrder
	OFPBACBadTag
	OFPBACBadSetType
	OFPBACBadSetLen
	OFPBACBadSetArgument
)

// Bad instruction code
const (
	OFPBICUnknownInst = iota
	OFPBICUnsupInst
	OFPBICBadTableID
	OFPBICUnsupMetadata
	OFPBICUnsupMetadataMask
	OFPBICBadExperimenter
	OFPBICBadExpType
	OFPBICBadLen
	OFPBICEperm
)

// Bad Match code
const (
	OFPBMCBadType = iota
	OFPBMCBadLen
	OFPBMCBadTag
	OFPBMCBadDLAddrMask
	OFPBMCBadNWAddrMask
	OFPBMCBadWildcards
	OFPBMCBadField
	OFPBMCBadValue
	OFPBMCBadMask
	OFPBMCBadPrereq
	OFPBMCDupField
	OFPBMCEperm
)

// Flow mod failed code
const (
	OFPFMFCUnknown = iota
	OFPFMFCTableFull
	OFPFMFCBadTableID
	OFPFMFCOverlap
	OFPFMFCEperm
	OFPFMFCBadTimeout
	OFPFMFCBadCommand
	OFPFMFCBadFlags
)

// Group mod failed code
const (
	OFPGMFCGroupExists = iota
	OFPGMFCInvalidGroup
	OFPFMFCWeightUnsupported
	OFPFMFCOutOfGroups
	OFPFMFCOutOfBuckets
	OFPFMFCChainingUnsupported
	OFPFMFCWatchUnsupported
	OFPFMFCLoop
	OFPFMFCUnknownGroup
	OFPFMFChainedGroup
	OFPGMFCBadType
	OFPGMFCBadCommand
	OFPGMFCBadBucket
	OFPGMFCBadWatch
	OFPGMFCEperm
)

// Port mod failed code
const (
	OFPPMFCBadPort = iota
	OFPPMFCBadHWAddr
	OFPPMFCBadConfig
	OFPPMFCBadAdvertise
	OFPPMFCEper
)

// Table mod failed code
const (
	OFPTMFCBadTable = iota
	OFPTMFCBadConfig
	FPTMFCEperm
)

// Queue op failed code
const (
	OFPQOFCBadPort = iota
	OFPQOFCBadQueue
	OFPQOFCEperm
)

// Switch config failed code
const (
	OFPSCFCBadFlags = iota
	OFPSCFCBadLen
	OFPQCFCEperm
)

// Role request failed code
const (
	OFPRRFCStale = iota
	OFPRRFCUnsup
	OFPRRFCBadRole
)

// Error experimenter message
type ErrorExperimenterMessage struct {
	Header       Header
	Type         uint16
	ExpType      uint16
	Experimenter uint32
	Data         []uint8
}

/* A.5 Symmetric Messages */
// A.5.4 Experimenter
type ExperimenterHeader struct {
	Header       Header /*Type OFPT_VENDOR*/
	Experimenter uint32
	ExpType      uint32
}
