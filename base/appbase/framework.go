package appbase

import (
	"bytes"
	"net"
	"sync"
	"time"
)

type ApplicationInstanceGenerator func() interface{}
type ApplicationInstanceList []ApplicationInstanceGenerator
type ApplicationInstanceMap map[int]ApplicationInstanceList

//type PigeonInstance interface {
//	ConnectionUp(dpid uint64)
//	ConnectionDown(dpid uint64)
//	EchoRequest(dpid uint64)
//	EchoReply(dpid uint64)
//	linkDiscoveryLoop(dpid uint64)
//}

type Controller struct{}

type Buffer struct {
	bytes.Buffer
}

type BufferPool struct {
	Empty chan *bytes.Buffer
	Full  chan *bytes.Buffer
}

type Message interface {
	//encoding.BinaryPacker
	PackBinary() (data []byte, err error)
	//encoding.BinaryUnpacker
	UnpackBinary(data []byte) (err error)
	//len
	Len() int
}

type MessageStream struct {
	Conn *net.TCPConn
	Pool *BufferPool
	// OpenFlow Version
	Version uint8
	// Channel on which to publish connection errors
	Error chan error
	// Channel on which to publish inbound messages
	Inbound chan Message
	// Channel on which to receive outbound messages
	Outbound chan Message
	// Channel on which to receive a shutdown command
	Shutdown chan bool
}

type Link struct {
	DPID      uint64
	Port      uint16
	Latency   time.Duration
	Bandwidth int
}
type Network struct {
	sync.RWMutex
	Switches map[string]*OFSwitch
}

type OFSwitch struct {
	Version       uint8
	Stream        *MessageStream
	AppInstance   []interface{}
	DPID          uint64
	Ports         map[uint32]OFPort
	PortsMutex    sync.RWMutex
	Links         map[string]*Link
	LinksMutex    sync.RWMutex
	Requests      map[uint32]chan Message
	RequestsMutex sync.RWMutex
}

type OFPort interface {
	PackBinary() (data []byte, err error)
	UnpackBinary(data []byte) (err error)
	Len() int
}

type OFSwitchFeatures interface {
	PackBinary() (data []byte, err error)
	UnpackBinary(data []byte) (err error)
	Len() int
}

type OFPacketIn interface {
	PackBinary() (data []byte, err error)
	UnpackBinary(data []byte) (err error)
	Len() int
}
