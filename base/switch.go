package base

import (
	"reflect"
	"time"
	"fmt"
	"jd.com/jdcontroller/lib/buffer"

	"jd.com/jdcontroller/lib/packet/eth"
	"jd.com/jdcontroller/lib/ofproto/ofp10"
	"jd.com/jdcontroller/lib/ofproto/ofp11"
	"jd.com/jdcontroller/lib/ofproto/ofp12"
	"jd.com/jdcontroller/lib/ofproto/ofp13"
)

func NewNetwork() *Network {
	n := new(Network)
	n.Switches = make(map[string]*OFSwitch)
	return n
}

var network *Network

// Builds and populates a Switch struct then starts listening
// for OpenFlow messages on conn.
func NewSwitch(stream *buffer.MessageStream, msg OFSwitchFeatures, v uint8) {
	network.Lock()
	defer network.Unlock()
	immutable := reflect.ValueOf(msg).Elem()
	dpid := eth.NewHardwareAddr(immutable.FieldByName("DPID").Uint())

	s := new(OFSwitch)
	s.Version = v
	s.Stream = stream
	s.AppInstance = *new([]interface{})
	s.DPID = dpid.Int()
	s.Ports = make(map[uint32]OFPort)
	s.Links = make(map[string]*Link)
	s.Requests = make(map[uint32]chan buffer.Message)

	network.Switches[dpid.String()] = s
	go s.requestStatus()
	go s.receive()
}

func (s *OFSwitch) AddInstance(instance interface{}) {
	switch s.Version {
	case 1:
		if actor, ok := instance.(ofp10.ConnectionUpReactor); ok {
			actor.ConnectionUp(s.DPID)
		}
	case 2:
		if actor, ok := instance.(ofp11.ConnectionUpReactor); ok {
			actor.ConnectionUp(s.DPID)
		}
	case 3:
		if actor, ok := instance.(ofp12.ConnectionUpReactor); ok {
			actor.ConnectionUp(s.DPID)
		}
	case 4:
		if actor, ok := instance.(ofp13.ConnectionUpReactor); ok {
			actor.ConnectionUp(s.DPID)
		}
	}
	s.AppInstance = append(s.AppInstance, instance)
}

func (s *OFSwitch) SetPort(portNO uint32, port OFPort) {
	s.PortsMutex.Lock()
	defer s.PortsMutex.Unlock()
	s.Ports[portNO] = port
}

// Returns a pointer to the Switch mapped to dpid.
func Switch(dpid uint64) (*OFSwitch, bool) {
	d := eth.NewHardwareAddr(dpid)
	network.RLock()
	defer network.RUnlock()
	if s, ok := network.Switches[d.String()]; ok {
		return s, ok
	}
	return nil, false
}

// Returns a slice of *OFPSwitches for operations across all
// switches.
func Switches() (a []*OFSwitch) {
	network.RLock()
	defer network.RUnlock()
	a = make([]*OFSwitch, len(network.Switches))
	i := 0
	for _, s := range network.Switches {
		a[i] = s
		i++
	}
	return
}

// Disconnects Switch dpid.
func SwitchDisconnect(dpid uint64) {
	d := eth.NewHardwareAddr(dpid)
	network.Lock()
	defer network.Unlock()
	fmt.Printf("Closing connection with: %s", d.String())
	network.Switches[d.String()].Stream.Shutdown <- true
	delete(network.Switches, d.String())
}

// Returns a slice of all links connected to Switch s.
func (s *OFSwitch) LinkList() (a []Link) {
	s.LinksMutex.RLock()
	defer s.LinksMutex.RUnlock()
	a = make([]Link, 0)
	for _, v := range s.Links {
		a = append(a, *v)
	}
	return
}

// Returns the link between Switch s and the Switch dpid.
func (s *OFSwitch) Link(dpid uint64) (l Link, ok bool) {
	d := eth.NewHardwareAddr(dpid)
	s.LinksMutex.RLock()
	defer s.LinksMutex.RUnlock()
	if n, k := s.Links[d.String()]; k {
		l = *n
		ok = true
	}
	return
}

// Updates the link between s.DPID and l.DPID.
func (s *OFSwitch) setLink(dpid uint64, l *Link) {
	d := eth.NewHardwareAddr(dpid)
	ld := eth.NewHardwareAddr(l.DPID)
	s.LinksMutex.Lock()
	defer s.LinksMutex.Unlock()
	if _, ok := s.Links[ld.String()]; !ok {
		fmt.Println("Link discovered:", d.String(), l.Port, ld.String())
	}
	s.Links[ld.String()] = l
}

// Returns a slice of all the ports from Switch s.
func (s *OFSwitch) PortList() (a []OFPort) {
	s.PortsMutex.RLock()
	defer s.PortsMutex.RUnlock()
	a = make([]OFPort, len(s.Ports))
	i := 0
	for _, v := range s.Ports {
		a[i] = v
		i++
	}
	return
}

// Returns a pointer to the OFPPhyPort at port number from Switch s.
func (s *OFSwitch) Port(portNO uint32) (p OFPort, ok bool) {
	s.PortsMutex.RLock()
	defer s.PortsMutex.RUnlock()
	p, ok = s.Ports[portNO]
	return
}

// Sends an OpenFlow message to this Switch.
func (s *OFSwitch) Send(r buffer.Message) {
	s.Stream.Outbound <- r
}

// Receive loop for each Switch.
func (s *OFSwitch) receive() {
	for {
		select {
		case msg := <-s.Stream.Inbound:
			// New message has been received from message
			// stream.
			go s.distributeMessages(s.DPID, msg)
		case err := <-s.Stream.Error:
			// Message stream has been disconnected.
			for _, app := range s.AppInstance {
				switch s.Version {
				case 1:
					if actor, ok := app.(ofp10.ConnectionDownReactor); ok {
						actor.ConnectionDown(s.DPID, err)
					}
					return
				case 2:
					if actor, ok := app.(ofp11.ConnectionDownReactor); ok {
						actor.ConnectionDown(s.DPID, err)
					}
					return
				case 3:
					if actor, ok := app.(ofp12.ConnectionDownReactor); ok {
						actor.ConnectionDown(s.DPID, err)
					}
					return
				case 4:
					if actor, ok := app.(ofp13.ConnectionDownReactor); ok {
						actor.ConnectionDown(s.DPID, err)
					}
					return
				}
			}
		}
	}
}

func (s *OFSwitch) distributeMessages(dpid uint64, msg buffer.Message) {
	fmt.Println("Distribute msg.", msg)
	if dpid != s.DPID {
		fmt.Println("DPID Error.")
		return
	}
	for _, app := range s.AppInstance {
		fmt.Println("Switch version: ", s.Version)
		switch s.Version {
		case 1:
			ofp10.ReactorParse(dpid, app, msg)
		case 2:
			ofp11.ReactorParse(dpid, app, msg)
		case 3:
			ofp12.ReactorParse(dpid, app, msg)
		case 4:
			ofp13.ReactorParse(dpid, app, msg)
		}
	}
}

func (s *OFSwitch) requestStatus() {
		switch s.Version {
		//case 1:
		//	ofp10.requestStatus()
		//case 2:
		//	ofp11.requestStatus()
		//case 3:
		//	ofp12.requestStatus()
		case 4:
			for {
				s.Send(ofp13.NewMultipartFlowStatsRequest())
				time.Sleep(time.Second)
				//s.Send(ofp13.NewMultipartAggregateStatsRequest())
				//time.Sleep(time.Second)
				s.Send(ofp13.NewMultipartPortStatsRequest(0xffffffff))
				time.Sleep(time.Minute)//sleep for 1 minute
			}
		}
}
