package io

import "bytes"

type Buffer struct {
	*bytes.Buffer
}

func (b *Buffer) Close() error {
	b.Reset()
	return nil
}

func NewBuffer() *Buffer {
	return &Buffer{
		Buffer: &bytes.Buffer{},
	}
}
