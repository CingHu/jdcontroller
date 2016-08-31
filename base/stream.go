package base

import (
	"bytes"
//	"encoding/binary"
//	"errors"
	"net"
	"jd.com/jdcontroller/lib/buffer"
//	"jd.com/jdcontroller/lib/ofproto/ofp"
)

func NewBufferPool() (b *buffer.BufferPool) {
	b = new(buffer.BufferPool)
	b.Empty = make(chan *bytes.Buffer, 50)
	b.Full = make(chan *bytes.Buffer, 50)

	for i := 0; i < 50; i++ {
		b.Empty <- bytes.NewBuffer(make([]byte, 0, 2048))
	}
	return
}

// Returns a pointer to a new MessageStream. Used to parse
// OpenFlow messages from conn.
func NewMessageStream(conn *net.TCPConn) (m *buffer.MessageStream) {
	m = &buffer.MessageStream{
		conn,
		NewBufferPool(),
		0,
		make(chan error, 1),   // Error
		make(chan buffer.Message, 1), // Inbound
		make(chan buffer.Message, 1), // Outbound
		make(chan bool, 1),    // Shutdown
	}

	go m.Out()
	go m.In()

	for i := 0; i < 25; i++ {
		go m.Parse()
	}
	return
}

