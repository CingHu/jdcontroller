package buffer

import (
	"fmt"
	"bytes"
	"net"
	"encoding/binary"
	"time"
	"errors"
	"jd.com/jdcontroller/lib/ofproto/ofp"
)

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

type LinkDiscovery struct {
	SrcDPID uint64
	Nsec    int64 /* Number of nanoseconds elapsed since Jan 1, 1970. */
}

func NewLinkDiscovery() (d *LinkDiscovery) {
	d = new(LinkDiscovery)
	d.Nsec = time.Now().UnixNano()
	return
}

func (d *LinkDiscovery) Len() (l int) {
	return 22
}

func (d *LinkDiscovery) PackBinary() (data []byte, err error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	binary.Write(buf, binary.BigEndian, d.SrcDPID)
	binary.Write(buf, binary.BigEndian, d.Nsec)
	data = buf.Bytes()
	return
}

func (d *LinkDiscovery) UnpackBinary(data []byte) (err error) {
	buf := bytes.NewBuffer(data)
	binary.Read(buf, binary.BigEndian, &d.SrcDPID)
	binary.Read(buf, binary.BigEndian, &d.Nsec)
	return
}

func (m *MessageStream) GetAddr() net.Addr {
	return m.Conn.RemoteAddr()
}

// Listen for a Shutdown signal or Outbound messages.
func (m *MessageStream) Out() {
	for {
		select {
		case <-m.Shutdown:
			fmt.Println("Closing OpenFlow message stream.")
			m.Conn.Close()
			return
		case msg := <-m.Outbound:
			// Forward outbound messages to conn
			fmt.Println("msg: ", msg)
			fmt.Println("msg type: %T", msg)
			data, _ := msg.PackBinary()
			fmt.Println("msg data: ", data)
			if _, err := m.Conn.Write(data); err != nil {
				fmt.Println("OutboundError:", err)
				m.Error <- err
				m.Shutdown <- true
			}
		}
	}
}

func (m *MessageStream) In() {
	msg := 0
	hdr := 0
	hdrBuf := make([]byte, 4)

	tmp := make([]byte, 2048)
	buf := <-m.Pool.Empty
	for {
		n, err := m.Conn.Read(tmp)
		if err != nil {
			if err.Error() == "EOF" {
				m.Error <- errors.New("Connection is closed.")
				m.Shutdown <- true
				return
			}
			fmt.Println("InboundError", err)
			m.Error <- err
			m.Shutdown <- true
			return
		}

		for i := 0; i < n; i++ {
			if hdr < 4 {
				hdrBuf[hdr] = tmp[i]
				buf.WriteByte(tmp[i])
				hdr += 1
				if hdr >= 4 {
					msg = int(binary.BigEndian.Uint16(hdrBuf[2:])) - 4
				}
				continue
			}
			if msg > 0 {
				buf.WriteByte(tmp[i])
				msg = msg - 1
				if msg == 0 {
					hdr = 0
					m.Pool.Full <- buf
					buf = <-m.Pool.Empty
				}
				continue
			}
		}
	}
}

func (m *MessageStream) Parse() {
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
