package buffer

import (
	"bytes"
)

type Buffer struct {
	bytes.Buffer
}

type BufferPool struct {
	Empty chan *bytes.Buffer
	Full  chan *bytes.Buffer
}

func NewBuffer(buf []byte) (b *Buffer) {
	b = new(Buffer)
	b.Buffer = *bytes.NewBuffer(buf)
	return
}

func (b *Buffer) Len() (l int) {
	return b.Buffer.Len()
}

func (b *Buffer) PackBinary() (data []byte, err error) {
	return b.Buffer.Bytes(), nil
}

func (b *Buffer) UnpackBinary(data []byte) error {
	b.Buffer.Reset()
	_, err := b.Buffer.Write(data)
	return err
}
