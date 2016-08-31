package base

import (
	"jd.com/jdcontroller/lib/buffer"
	"sync"
	"time"
)

type OfpHandlerInstanceGenerator func() interface{}
type OfpHandlerInstanceList []OfpHandlerInstanceGenerator
type OfpHandlerInstanceMap map[int]OfpHandlerInstanceList

type Controller struct{}

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
	Stream        *buffer.MessageStream
	AppInstance   []interface{}
	DPID          uint64
	Ports         map[uint32]OFPort
	PortsMutex    sync.RWMutex
	Links         map[string]*Link
	LinksMutex    sync.RWMutex
	Requests      map[uint32]chan buffer.Message
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

type OFMultipart interface {
	PackBinary() (data []byte, err error)
	UnpackBinary(data []byte) (err error)
	Len() int
}

type OFPacketIn interface {
	PackBinary() (data []byte, err error)
	UnpackBinary(data []byte) (err error)
	Len() int
}
