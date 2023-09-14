package internal

import (
	"bytes"
	"fmt"
	"io"
)

type BufferWriteSeeker struct {
	buf bytes.Buffer
	off int64
}

func (b *BufferWriteSeeker) Write(p []byte) (n int, err error) {
	n, err = b.buf.Write(p)
	b.off += int64(n)
	return
}

func (b *BufferWriteSeeker) Seek(offset int64, whence int) (int64, error) {
	switch whence {
	case io.SeekStart:
		if offset < 0 || offset > int64(b.buf.Len()) {
			return 0, io.EOF
		}
		b.off = offset
	case io.SeekCurrent:
		newOffset := b.off + offset
		if newOffset < 0 || newOffset > int64(b.buf.Len()) {
			return 0, io.EOF
		}
		b.off = newOffset
	case io.SeekEnd:
		newOffset := int64(b.buf.Len()) + offset
		if newOffset < 0 || newOffset > int64(b.buf.Len()) {
			return 0, io.EOF
		}
		b.off = newOffset
	default:
		return 0, fmt.Errorf("invalid whence")
	}
	return b.off, nil
}

func (b *BufferWriteSeeker) Bytes() []byte {
	return b.buf.Bytes()
}
