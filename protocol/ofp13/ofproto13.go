package ofp13

import (
	"fmt"
	"jd.com/jdcontroller/lib/buffer"
	"jd.com/jdcontroller/lib/packet/eth"
)

/* OpenFlow Switch Specification      */
/* Version 1.3.4 (Wire Protocol 0x04) */
/* March 27, 2014                     */

const FLOWDEFPRIORITY = 1000
const OFPNOBUFFER = 0Xffffffff

const (
	Version = 4
)

/* 7.1.1 OpenFlow Header */
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

	// Multipart messages.
	OFPTMultipartRequest
	OFPTMultipartReply

	// Barrier messages.
	OFPTBarrierRequest
	OFPTBarrierReply

	// Queue Configuration messages.
	OFPTQueueGetConfigRequest
	OFPTQueueGetConfigReply

	// Controller role change request messages.
	OFPTRoleRequest
	OFPTRoleReply

	// Asynchronous mesage configuration.
	OFPTGetAsyncRequest
	OFPTGetAsyncReply
	OFPTSetAsync

	// Meters and rate limiters configuration messages.
	OFPTMeterMod
)

/* 7.2 Common Structures */
// 7.2.1 Port Structures
// constant
const (
	OFPEthALen        = 6
	OFPMaxPortNameLen = 16
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

type Port struct {
	PortNO uint32
	pad    [4]uint8 // Size 4
	HWAddr [6]byte
	pad2   [2]uint8                // Size 2
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

// 7.2.2 Queue Structures
type PacketQueue struct {
	QueueID    uint32
	Port       uint32
	Length     uint16
	pad        [6]uint8 // Size 6
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
	pad      [4]uint8 // Size 4
}

type QueuePropMinRate struct {
	Header QueuePropHeader
	Rate   uint16
	pad    [6]uint8 // Size 6
}

type QueuePropMaxRate struct {
	Header QueuePropHeader
	Rate   uint16
	pad    [6]uint8 // Size 6
}

type QueuePropExperimenter struct {
	Header       QueuePropHeader
	Experimenter uint32
	pad          [4]uint8 // Size 4
	Data         []uint8
}

// 7.2.3 Flow Match Structures
// 7.2.3.1 Flow Match Header
type Match struct {
	Type      uint16
	Length    uint16 	/* Length of ofp_match (excluding padding) */
	OXMFields []uint8
	pad       []uint8
}

const MATCHHEADERLEN = 4
// Match type
const (
	OFPMTStandard = iota
	OFPMTOXM
)

// 7.2.3.3 OXM classes
// OXM class
const (
	OFPXMCNXM0          = 0x0000
	OFPXMCNXM1          = 0x0001
	OFPXMCOpenFlowBasic = 0x8000
	OFPXMCExperimenter  = 0xffff
)

// 7.2.3.7 Flow Match Fields
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
	OFPXMTFOFBMPLSBOS
	OFPXMTFOFBPBBISID
	OFPXMTFOFBTunnelID
	OFPXMTFOFBIPv6ExtHdr
)

const (
	NXMNXTUNIPV4SRC = 31
	NXMNXTUNIPV4DST = 32
)

const (
	OFPXMHEADERLEN = 4
	OFPXMTFOFBEthLEN = 6
	OFPXMTFOFBIPProtoLEN = 2
	OFPXMTFOFBIPv4LEN = 4
	NXMNXTUNIPV4LEN = 4
)

const (
	OFPXMNOMASK = 0
	OFPXMHASMASK = 1
)

// 7.2.3.8 Header Match Fields
// VLAN id
const (
	OFPVIDPresent = 0x1000
	OFPVIDNone    = 0x0000
)

// ipv6 exthdr flags
const (
	OFPIEHNoNext = 1 << 0
	OFPIEHESP    = 1 << 1
	OFPIEHAuth   = 1 << 2
	OFPIEHDest   = 1 << 3
	OFPIEHFrag   = 1 << 4
	OFPIEHRouter = 1 << 5
	OFPIEHHop    = 1 << 6
	OFPIEHUnrep  = 1 << 7
	OFPIEHUnseq  = 1 << 8
)

// 7.2.3.10 Experimenter Flow Match Fields
type OXMExperimenterHeader struct {
	OXMHeader    uint32
	Experimenter uint32
}

// 7.2.4 Flow Instruction Structures
// Instruction type
const (
	OFPITGotoTable     = 1
	OFPITWriteMetadata = 2
	OFPITWriteActions  = 3
	OFPITApplyActions  = 4
	OFPITClearActions  = 5
	OFPITMeter         = 6
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
	pad     [3]uint8 // Size 3
}

type InstructionWriteMetadata struct {
	Header       InstructionHeader
	pad          [4]uint8 // Size 4
	Metadata     uint64
	MetadataMask uint64
}

type InstructionActions struct {
	Header  InstructionHeader
	pad     [4]uint8 // Size 4
	Actions []Action
}

type InstructionMeter struct {
	Header  InstructionHeader
	MeterID uint32
}

type InstructionExperimenter struct {
	Header       InstructionHeader
	Experimenter uint32
}

// 7.2.5 Action Structures
// Action type
const (
	OFPATOutput     = 0  /* Output to switch port. */
	OFPATCopyTTLOut = 11 /* Copy TTL "outwards" -- from next-to-outermostto outermost */
	OFPATCopyTTLIn  = 12 /* Copy TTL "inwards" -- from outermost tonext-to-outermost */
	OFPATSetMPLSTTL = 15 /* MPLS TTL */
	OFPATDecMPLSTTL = 16 /* Decrement MPLS TTL */

	OFPATPushVLAN = 17 	 /*Push a new VLAN tag */
	OFPATPopVLAN  = 18 	 /*Pop the outer VLAN tag */
	OFPATPushMPLS = 19 	 /*Push a new MPLS tag */
	OFPATPopMPLS  = 20 	 /*Pop the outer MPLS tag */

	OFPATSetQueue = 21 	 /*Set queue id when outputting to a port */
	OFPATGroup    = 22	 /*Apply group. */
	OFPATSetNWTTL = 23 	 /*IP TTL. */
	OFPATDecNWTTL = 24 	 /*Decrement IP TTL. */
	OFPATSetField = 25 	 /*Set a header field using OXM TLV format. */
	OFPATPushPBB  = 26 	 /*Push a new PBB service tag (I-TAG) */
	OFPATPopPBB   = 27 	 /*Pop the outer PBB service tag (I-TAG) */

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
	pad    [6]uint8 // Size 6
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
	pad     [3]uint8 // Size 3
}

type ActionNWTTL struct {
	Header ActionHeader
	NWTTL  uint8
	pad    [3]uint8 // Size 3
}

type ActionPush struct {
	Header    ActionHeader
	Ethertype uint16
	pad       [2]uint8 // Size 2
}

type ActionPopMPLS struct {
	Header    ActionHeader
	Ethertype uint16
	pad       [2]uint8 // Size 2
}

type ActionSetField struct {
	Header ActionHeader
	Field  []uint8
}

type ActionExperimenterHeader struct {
	Header       ActionHeader
	Experimenter uint32
}

/* 7.3 Controller to Switch Messages */
// 7.3.1 Handshake
type SwitchFeatures struct {
	Header       Header
	DPID         uint64
	Buffers      uint32
	Tables       uint8
	AuxiliaryID  uint8
	pad          [2]uint8 // Size 2
	Capabilities uint32
	Reserved     uint32
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

// 7.3.2 Switch Configuration
// Switch config
type SwitchConfig struct {
	Header      Header
	Flags       uint16
	MissSendLen uint16
}

// Config flags
const (
	OFPCFragNormal = 0
	OFPCFragDrop   = 1 << 0
	OFPCFragReasm  = 1 << 1
	OFPCFragMask   = 3
)

// 7.3.3 Flow Table Configuration
// Table numbering
const (
	OFPTTMax = 0xfe
	OFPTTAll = 0xff
)

// Table mod
type TableMod struct {
	Header  Header
	TableID uint8
	pad     [3]uint8 // Size 3
	Config  uint32
}

// Table config
const (
	OFTCTableMissMask = 3
)

// 7.3.4 Modify State Messages
// 7.3.4.1 Modify Flow Entry Message
// Flow mod
type FlowMod struct {
	Header     Header
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
	pad          [2]uint8 // Size 2
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
	OFPFFSendFlowRem    = 1 << 0
	OFPFFCheckOverlap   = 1 << 1
	OFPFFResetCounts    = 1 << 2
	OFPFFNoPacketCounts = 1 << 3
	OFPFFNoByteCounts   = 1 << 4
)

// 7.3.4.2 Modify Group Entry Message
// Group mod
type GroupMod struct {
	Header
	Command uint16
	Type    uint8
	pad     [1]uint8 // Size 1
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

// Group numbering
const (
	OFPGMax = 0xffffff00
	OFPGAll = 0xfffffffc
	OFPGAny = 0xffffffff
)

type Bucket struct {
	Length     uint16
	Weight     uint16
	WatchPort  uint32
	WatchGroup uint32
	pad        [4]uint8 // Size 4
	Actions    []Action
}

// 7.3.4.3 Port Modification Message
// Port mod
type PortMod struct {
	Header    Header
	PortNO    uint32
	pad       [4]uint8 // Size 4
	HWAddr    [OFPEthALen]byte
	pad2      [2]uint8 // Size 2
	Config    uint32
	Mask      uint32
	Advertise uint32
	pad3      [4]uint8 // Size 4
}

// 7.3.4.4 Meter Modification Message
// Meter mod
type MeterMod struct {
	Header  Header
	Command uint16
	Flags   uint16
	MeterID uint32
	Bands   []MeterBand
}

// Meter numbering
const (
	OFPMMax        = 0xffff0000
	OFPMSlowPath   = 0xfffffffd
	OFPMController = 0xfffffffe
	OFPMAll        = 0xffffffff
)

// Meter commands
const (
	OFPMCAdd = iota
	OFPMCModify
	OFPMCDelete
)

// Meter flags
const (
	OFPMFKbps  = 1 << 0
	OFPMFPktps = 1 << 1
	OFPMFBurst = 1 << 2
	OFPMFStats = 1 << 3
)

type MeterBand interface {
	PackBinary() (data []byte, err error)
	UnpackBinary(data []byte) (err error)
	Len() int
}

// Meter band header
type MeterBandHeader struct {
	Type      uint16
	Length    uint16
	Rate      uint32
	BurstSize uint32
}

// Meter band types
const (
	OFPMBTDrop         = 1
	OFPMBTDscpRemark   = 2
	OFPMBTExperimenter = 0xffff
)

// Meter band drop
type MeterBandDrop struct {
	Header MeterBandHeader
	pad    [4]uint8 // Size 4
}

// Meter band dscp remark
type MeterBandDscpRemark struct {
	Header    MeterBandHeader
	PrecLevel uint8
	pad       [3]uint8 // Size 3
}

// Meter band experimenter
type MeterBandExperimenter struct {
	Header       MeterBandHeader
	Experimenter uint32
}

// 7.3.5 Multipart Messages
type MultipartRequest struct {
	Header Header
	Type   uint16
	Flags  uint16
	pad    [4]uint8 // Size 4
	Body   []buffer.Message
}

type MultipartReply struct {
	Header Header
	Type   uint16
	Flags  uint16
	pad    [4]byte // Size 4
	//	Body   []uint8
	Body []buffer.Message
}

// Multipart request flags
const (
	OFPMPFRequestMore = 1 << 0
)

// Multipart reply flags
const (
	OFPMPFReplyMore = 1 << 0
)

// Multipart type
const (
	OFPMPDesc = iota
	OFPMPFlow
	OFPMPAggregate
	OFPMPTable
	OFPMPPortStats
	OFPMPQueue
	OFPMPGroup
	OFPMPGroupDesc
	OFPMPGroupFeatures
	OFPMPMeter
	OFPMPMeterConfig
	OFPMPMeterFeatures
	OFPMPTableFeatures
	OFPMPPortDesc
	OFPMPExperimenter = 0xffff
)

// 7.3.5.1 Description
// constats
const (
	OFPDescStrLen      = 256
	OFPSerialNumLen    = 32
	OFPMaxTableNameLen = 32
)

type Desc struct {
	MfrDesc   [OFPDescStrLen]byte   // Size 256
	HWDesc    [OFPDescStrLen]byte   // Size 256
	SWDesc    [OFPDescStrLen]byte   // Size 256
	SerialNum [OFPSerialNumLen]byte // Size 32
	DPDesc    [OFPDescStrLen]byte   // SizeOFPDescStrLen 256
}

// 7.3.5.2 Individual Flow Statistics
// Flow stats request
type FlowStatsRequest struct {
	TableID    uint8
	pad        [3]uint8 // Size 3
	OutPort    uint32
	OutGroup   uint32
	pad2       [4]uint8 // Size 4
	Cookie     uint64
	CookieMask uint64
	Match      Match
}

// Flow stats
type FlowStats struct {
	Length       uint16
	TableID      uint8
	pad          [1]uint8 // Size 1
	DurationSec  uint32
	DurationNSec uint32
	Priority     uint16
	IdleTimeout  uint16
	HardTimeout  uint16
	Flags        uint16
	pad2         [4]uint8 // Size 4
	Cookie       uint64
	PacketCount  uint64
	ByteCount    uint64
	Match        Match
	Instructions []Instruction
}

// 7.3.5.3 Aggregate Flow Statistics
// Aggregate stats request
type AggregateStatsRequest struct {
	TableID    uint8
	pad        [3]uint8 // Size 3
	OutPort    uint32
	OutGroup   uint32
	pad2       [4]uint8 // Size 4
	Cookie     uint64
	CookieMask uint64
	Match      Match
}

// Aggregate stats reply
type AggregateStatsReply struct {
	PacketCount uint64
	ByteCount   uint64
	FlowCount   uint32
	pad         [4]uint8 // Size 4
}

// 7.3.5.4 Table Statistics
// Table stats
type TableStats struct {
	TableID      uint8
	pad          [3]uint8 // Size 3
	ActiveCount  uint32
	LookupCount  uint64
	MatchedCount uint64
}

// 7.3.5.5.1 Table Features request and reply
type TableFeatures struct {
	Length        uint16
	TableID       uint8
	pad           [5]uint8                 // Size 5
	Name          [OFPMaxTableNameLen]byte // Size 32
	MetadataMatch uint64
	MetadataWrite uint64
	Config        uint32
	MaxEntries    uint32
	Properties    []TableFeatureProp
}

// 7.3.5.5.2 Table Features properties
// Table feature prop type
const (
	OFPTFPTInstructions = iota
	OFPTFPTInstructionsMiss
	OFPTFPTNextTables
	OFPTFPTNextTablesMiss
	OFPTFPTWriteActions
	OFPTFPTWriteActionsMiss
	OFPTFPTApplyActions
	OFPTFPTApplyActionsMiss
	OFPTFPTMatch
	OFPTFPTWildcards
	OFPTFPTWriteSetField
	OFPTFPTWriteSetFieldMiss
	OFPTFPTApplySetField
	OFPTFPTApplySetFieldMiss
	OFPTFPTExperimenter     = 0xfffe
	OFPTFPTExperimenterMiss = 0xffff
)

// Table feature prop header
type TableFeatureProp interface {
	PackBinary() (data []byte, err error)
	UnpackBinary(data []byte) (err error)
	Len() int
}

type TableFeaturePropHeader struct {
	Type   uint16
	Length uint16
}

// Table features prop instructions
type TableFeaturePropInstructions struct {
	Header       TableFeaturePropHeader
	Instructions []Instruction
}

// Table feature prop next tables
type TableFeaturePropNextTables struct {
	Header    TableFeaturePropHeader
	NextTable []uint8 // Size 1
}

// Table feature prop actions
type TableFeaturePropActions struct {
	Header  TableFeaturePropHeader
	Actions []Action
}

// Table feature prop oxm
type TableFeaturePropOXM struct {
	Header TableFeaturePropHeader
	OXM    []uint32
}

// Table feature prop experimenter
type TableFeaturePropExperimenter struct {
	Header       TableFeaturePropHeader
	Experimenter uint32
	ExpType      uint32
	Data         []uint32
}

// 7.3.5.6 Port Statistics
// Port stats request
type PortStatsRequest struct {
	PortNO uint32
	pad    [4]uint8 // Size 4
}

// Port stats
type PortStats struct {
	PortNO       uint32
	pad          [4]uint8 // Size 4
	RxPackets    uint64
	TxPackets    uint64
	RxBytes      uint64
	TxBytes      uint64
	RxDropped    uint64
	TxDropped    uint64
	RxErrors     uint64
	TxErrors     uint64
	RxFrameErr   uint64
	RxOverErr    uint64
	RxCRCErr     uint64
	Collisions   uint64
	DurationSec  uint32
	DurationNSec uint32
}

// 7.3.5.8 Queue Statistics
// Queue stats request
type QueueStatsRequest struct {
	PortNO  uint32
	QueueID uint32
}

// Queue stats
type QueueStats struct {
	PortNO       uint32
	QueueID      uint32
	TxBytes      uint64
	TxPackets    uint64
	TxErrors     uint64
	DurationSec  uint32
	DurationNSec uint32
}

// 7.3.5.9 Group Statistics
// Group stats request
type GroupStatsRequest struct {
	GroupID uint32
	pad     [4]uint8 // Size 4
}

// Group stats
type GroupStats struct {
	Length       uint16
	pad          [2]uint8 // Size 2
	GroupID      uint32
	RefCount     uint32
	pad2         [4]uint8 // Size 4
	PacketCount  uint64
	ByteCount    uint64
	DurationSec  uint32
	DurationNSec uint32
	BucketStats  []BucketCounter
}

type BucketCounter struct {
	PacketCount uint64
	ByteCount   uint64
}

// 7.3.5.10 Group Description
type GroupDescStats struct {
	Length  uint16
	Type    uint8
	pad     [1]uint8 // Size 1
	GroupID uint32
	Buckets []Bucket
}

// 7.3.5.11 Group Features
// Group features
type GroupFeatures struct {
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

// 7.3.5.12 Meter Statistics
type MeterMultipartRequest struct {
	MeterID uint32
	pad     [4]uint8 // Size 4
}

// Meter stats
type MeterStats struct {
	MeterID       uint32
	Length        uint16
	pad           [6]uint8 // Size 6
	FlowCount     uint32
	PacketInCount uint64
	ByteInCount   uint64
	DurationSec   uint32
	DurationNSec  uint32
	BandStats     []MeterBandStats
}

// Meter band stats
type MeterBandStats struct {
	PacketBandCount uint64
	ByteBandCount   uint64
}

// 7.3.5.13 Meter Configuration
// Meter config
type MeterConfig struct {
	Length  uint16
	Flags   uint16
	MeterID uint32
	Bands   []MeterBand
}

// 7.3.5.14 Meter Features
// Meater features
type MeterFeatures struct {
	MaxMeter     uint32
	BandTypes    uint32
	Capabilities uint32
	MaxBands     uint8
	MaxColor     uint8
	pad          [2]uint8 // Size 2
}

// 7.3.5.15 Experimenter Multipart
// Experimenter multipart header
type ExperimenterMultipartHeader struct {
	Experimenter uint32
	ExpType      uint32
}

// 7.3.6 Queue Configuration Messages
type QueueGetConfigRequest struct {
	Header Header
	Port   uint32
	pad    [4]uint8 // Size 4
}

type QueueGetConfigReply struct {
	Header Header
	Port   uint32
	pad    [4]uint8 // Size 4
	Queues []PacketQueue
}

// A.3.7 Packet-Out Message
type PacketOut struct {
	Header     Header
	BufferID   uint32
	InPort     uint32
	ActionsLen uint16
	pad        [6]uint8 // Size 6
	Actions    []Action
	//Data       buffer.Message
	Data 	   eth.Ethernet
}

// 7.3.9 Role Request Message
// Role request and reply message
type RoleRequest struct {
	Header       Header
	Role         uint32
	pad          [4]uint8 // Size 4
	GenerationID uint64
}

// Controller roles
const (
	OFPCRRoleNoChange = iota
	OFPCRRoleEqual
	OFPCRRoleMaster
	OFPCRRoleSlave
)

// 7.3.10 Set Asynchronous Configuration Message

type AsyncConfig struct {
	Header          Header
	PacketInMask    [2]uint32 // Size 2
	PortStatusMask  [2]uint32 // Size 2
	FlowRemovedMask [2]uint32 // Size 2
}

/* 7.4 Asynchronous Messages */
// 7.4.1 Packet-In Message
type PacketIn struct {
	Header   Header
	BufferID uint32
	TotalLen uint16
	Reason   uint8
	TableID  uint8
	Cookie   uint64
	Match    Match
	pad      [2]byte // Size 2
	//Data     buffer.Message
	Data 	eth.Ethernet
}

// Packet-in reason
const (
	OFPRNoMatch = iota
	OFPRAction
	OFPRInvalidTTL
)

// 7.4.2 Flow Removed Message
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

// 7.4.3 Port Status Message
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

// 7.4.4 Error Message
type ErrorMsg struct {
	Header Header
	Type   uint16
	Code   uint16
	Data   []uint8
}


type MultipartInfo struct {
	OfpDesc Desc
	FlowStats FlowStats
	AggregateStat AggregateStatsReply
	TableStats TableStats
	TableFeatures TableFeatures
	PortStats map[uint32]*PortStats
	PortDesc []Port
	QueueStats QueueStats
	GroupStats GroupStats
	GroupDescStats GroupDescStats
	GroupFeatures GroupFeatures
	MeterStats MeterStats
	MeterFeatures MeterFeatures
	MeterConfig MeterConfig
	ExperimenterHeader
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
	OFPETMeterModFailed
	OFPETTableFeaturesFailed
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
	OFPBRCBadMultipart
	OFPBRCBadExperimenter

	OFPBRCBadExpType
	OFPBRCEperm
	OFPBRCBadLen
	OFPBRCBufferEmpty
	OFPBRCBufferUnknown
	OFPBRCBadTableID

	OFPBRCIsSlave
	OFPBRCBadPort
	OFPBRCBadPacket
	OFPBRCMultipartBufferOverflow
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
	OFPBICBadExperimenterType
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
	OFPPMFCEperm
)

// Table mod failed code
const (
	OFPTMFCBadTable = iota
	OFPTMFCBadConfig
	OFPTMFCEperm
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

// Meter mod failed code
const (
	OFPMMFCUnknown = iota
	OFPMMFCMeterExists
	OFPMMFCInvalidMeter
	OFPMMFCUnknownMeter
	OFPMMFCBadCommand
	OFPMMFCBadFlags
	OFPMMFCBadRate
	OFPMMFCBadBurst
	OFPMMFCBadBand
	OFPMMFCBadBandValue
	OFPMMFCOutOfMeters
	OFPMMFCOutOfBands
)

// Table features failed code
const (
	OFPTFFCBadTable = iota
	OFPTFFCBadMetadata
	OFPTFFCBadType
	OFPTFFCBadLen
	OFPTFFCBadArgument
	OFPTFFCEperm
)

// Error experimenter message
type ErrorExperimenterMessage struct {
	Header       Header
	Type         uint16
	ExpType      uint16
	Experimenter uint32
	Data         []uint8
}

/* 7.5 Symmetric Messages */
// 7.5.1 Hello
type Hello struct {
	Header   Header
	Elements []HelloElem
}

// Hello elements types
const (
	OFPHETVersionBitmap = 1
)

type HelloElem interface{}

type HelloElemHeader struct {
	Type   uint16
	Length uint16
}

type HelloElemVersionBitmap struct {
	Header  HelloElemHeader
	Bitmaps []uint32
}

// 7.5.4 Experimenter
type ExperimenterHeader struct {
	Header       Header /*Type OFPT_VENDOR*/
	Experimenter uint32
	ExpType      uint32
}

type MatchField struct{
	XMFields []uint8;
}

type ActionField MatchField


func (field *MatchField)Init(ofpxm uint8, length uint8, hasmask uint8) {
	// oxm_class (用两个uint8组成ofp13.OFPXMCOpenFlowBasic（0x8000）)
	field.XMFields = append(field.XMFields, 1<<7)
	field.XMFields = append(field.XMFields, 0)
	// oxm_field(7bit) + oxm_hasmask(1bit)
	field.XMFields = append(field.XMFields, (ofpxm << 1) | hasmask) //no mask
	// oxm_length
	field.XMFields = append(field.XMFields, length)
	return
}

func (match MatchField) GetLen() (len uint16){
	len = uint16(match.XMFields[3]) & uint16(0xff)
	fmt.Println("GetLen", len, match.XMFields)
	return
}
