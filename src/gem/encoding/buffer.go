package encoding

import (
	"errors"
	"io"
	"sync"
)

// Buffer is similar to bytes.Buffer, but allows us to save and recall the read pointer,
// through the use of Try. Suitable for a long-living socket read buffer, as Trim allows
// us to discard and garbage collect data we've already dealt with.
type Buffer struct {
	s []byte

	i int
	m sync.Mutex
}

func NewBuffer() *Buffer {
	return &Buffer{
		s: make([]byte, 0),
		i: 0,
	}
}

func NewBufferBytes(s []byte) *Buffer {
	buffer := NewBuffer()
	buffer.s = append(buffer.s, s...)
	return buffer
}

// Trim discards all data before the current read pointer.
// blocks until we get a lock
func (b *Buffer) Trim() {
	b.m.Lock()
	defer b.m.Unlock()

	// perform a copy, so that the old array (and the discarded data) can be garbage collected
	oldSlice := b.s
	b.s = make([]byte, len(oldSlice) - b.i)
	copy(b.s, oldSlice)
	b.i = 0
}

// Try saves the current position, calls cb, and if cb returns an error, restores the previous position
// locks the buffer to trimming, to ensure we can always pop back to the original position
// since the trim mutex is locked until cb returns, deadlock can occur with incorrect usage
func (b *Buffer) Try(cb func (b *Buffer) error) error {
	b.m.Lock()
	defer b.m.Unlock()

	oldPtr := b.i
	err := cb(b)
	if err != nil {
		b.i = oldPtr
	}
	return err
}

func (b *Buffer) Read(p []byte) (n int, err error) {
	if b.i >= len(b.s) {
		// Buffer is empty
		if len(p) == 0 {
			return
		}
		return 0, io.EOF
	}
	n = copy(p, b.s[b.i:])
	b.i += n
	return n, nil
}

func (b *Buffer) ReadByte() (c byte, err error) {
	s := make([]byte, 1)
	n, err := b.Read(s)
	if err != nil {
		return 0, err
	}
	if n != 1 {
		return 0, io.EOF
	}
	return s[0], nil
}

func (b *Buffer) Write(p []byte) (n int, err error) {
	b.s = append(b.s, p...)
	return len(p), nil
}

func (b *Buffer) WriteByte(c byte) error {
	_, err := b.Write([]byte{c})
	return err
}

func (b *Buffer) Seek(offset int64, whence int) (int64, error) {
	var abs int64
	switch whence {
	case 0:
		abs = offset
	case 1:
		abs = int64(b.i) + offset
	case 2:
		abs = int64(len(b.s)) + offset
	default:
		return 0, errors.New("encoding.Buffer.Seek: invalid whence")
	}
	if abs < 0 {
		return 0, errors.New("encoding.Buffer.Seek: negative position")
	}
	if int(abs) > len(b.s) {
		return 0, errors.New("encoding.Buffer.Seek: out of bounds")
	}
	b.i = int(abs)
	return abs, nil
}