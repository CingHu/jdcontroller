package base

import (
	"fmt"
	"net"
	"time"

	"jd.com/jdcontroller/lib/ofproto/ofp"
	"jd.com/jdcontroller/lib/ofproto/ofpctl10"
	"jd.com/jdcontroller/lib/ofproto/ofpctl11"
	"jd.com/jdcontroller/lib/ofproto/ofpctl12"
	"jd.com/jdcontroller/lib/ofproto/ofpctl13"
	"jd.com/jdcontroller/lib/buffer"
)

var OfpHandlers OfpHandlerInstanceMap

func NewController(v int) (c *Controller) {
	c = new(Controller)
	OfpHandlers = make(OfpHandlerInstanceMap)
	network = NewNetwork()

	ofpctl13.NewMultipartInfo()
	return
}

func (c *Controller) RegisterOfpHandler(aig OfpHandlerInstanceGenerator, v int) {
	OfpHandlers[v] = append(OfpHandlers[v], aig)
}

func (c *Controller) Listen(port string) {
	addr, _ := net.ResolveTCPAddr("tcp4", port)
	sock, err := net.ListenTCP("tcp4", addr)
	if err != nil {
		fmt.Println(err)
	}
	defer sock.Close()

	fmt.Println("Listening for connections on", addr)
	for {
		conn, err := sock.AcceptTCP()
		if err != nil {
			fmt.Println(err)
		}
		go c.handleConnection(conn)
	}
}

func (c *Controller) handleConnection(conn *net.TCPConn) {
	stream := NewMessageStream(conn)

	fmt.Println("Current connection: ", stream.GetAddr().String())
	fmt.Println("Create Hello message.")
	h := ofpctl13.NewHello()
	fmt.Println("[OUT] Hello Message.")
	stream.Outbound <- h

	for {
		select {
		case msg := <-stream.Inbound:
			fmt.Println("Get new message.")
			switch msg.(type) {
			case *ofp10.Header:
				fmt.Println("Get openflow 1.0 message.")
				m := msg.(*ofp10.Header)
				stream.Version = m.Version
				stream.Outbound <- ofp10.NewFeaturesRequest()
			case *ofp11.Header:
				fmt.Println("Get openflow 1.1 message.")
				m := msg.(*ofp11.Header)
				stream.Version = m.Version
				stream.Outbound <- ofp11.NewFeaturesRequest()
			case *ofp12.Header:
				fmt.Println("Get openflow 1.2 message.")
				m := msg.(*ofp12.Header)
				stream.Version = m.Version
				stream.Outbound <- ofp12.NewFeaturesRequest()
			case *ofp13.Header:
				fmt.Println("Get openflow 1.3 message.")
				m := msg.(*ofp13.Header)
				stream.Version = m.Version
				stream.Outbound <- ofp13.NewFeaturesRequest()
				fmt.Println("Send features request.")
			case *ofp10.SwitchFeatures:
				fmt.Println("Switch Register of ofp10.")
				m := msg.(*ofp10.SwitchFeatures)
				NewSwitch(stream, m, m.Header.Version)
				for _, newInstance := range OfpHandlers[int(m.Header.Version)] {
					if sw, ok := Switch(m.DPID); ok {
						i := newInstance()
						sw.AddInstance(i)
					}
				}
				return
			case *ofp11.SwitchFeatures:
				fmt.Println("Switch Register of ofp11.")
				m := msg.(*ofp11.SwitchFeatures)
				NewSwitch(stream, m, m.Header.Version)
				for _, newInstance := range OfpHandlers[int(m.Header.Version)] {
					if sw, ok := Switch(m.DPID); ok {
						i := newInstance()
						sw.AddInstance(i)
					}
				}
				return
			case *ofp12.SwitchFeatures:
				fmt.Println("Switch Register of ofp12.")
				m := msg.(*ofp12.SwitchFeatures)
				NewSwitch(stream, m, m.Header.Version)
				for _, newInstance := range OfpHandlers[int(m.Header.Version)] {
					if sw, ok := Switch(m.DPID); ok {
						i := newInstance()
						sw.AddInstance(i)
					}
				}
				return
			case *ofp13.SwitchFeatures:
				fmt.Println("Switch Register of ofp13.")
				m := msg.(*ofp13.SwitchFeatures)
				NewSwitch(stream, m, m.Header.Version)
				for _, newInstance := range OfpHandlers[int(m.Header.Version)] {
					if sw, ok := Switch(m.DPID); ok {
						i := newInstance()
						sw.AddInstance(i)
					}
				}
				//stream.Outbound <- ofp13.NewMultipartPortDescRequest()
				stream.Outbound <- ofp13.NewSwitchConfig()
				return
			case *ofp10.ErrorMsg:
				fmt.Println("Default Error openflow 1.0 message.")
				fmt.Println(msg)
				stream.Shutdown <- true
			case *ofp11.ErrorMsg:
				fmt.Println("Default Error openflow 1.1 message.")
				fmt.Println(msg)
				stream.Shutdown <- true
			case *ofp12.ErrorMsg:
				fmt.Println("Default Error openflow 1.2 message.")
				fmt.Println(msg)
				stream.Shutdown <- true
			case *ofp13.ErrorMsg:
				fmt.Println("Default Error openflow 1.3 message.")
				fmt.Println(msg)
				stream.Shutdown <- true
			}
		case err := <-stream.Error:
			fmt.Println(err)
			return
		case <-time.After(time.Second * 5):
			fmt.Println("Connection timed out.")
			return
		}
	}
}

func (m *buffer.MessageStream) Parse() {
	for {
		b := <-m.Pool.Full
		msg, err := ofp.Parse(b.Bytes())
		// Log all message parsing errors.
		if err != nil {
			fmt.Println("Message stream parse error:", err)
		}

		m.Inbound <- msg
		b.Reset()
		m.Pool.Empty <- b
	}
}
