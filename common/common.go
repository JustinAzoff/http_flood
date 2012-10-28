package common

import (
	"../consts"
	"crypto/rand"
	"io"
)

type LoopReader struct {
	pos  int
	buf  []byte
	size uint64
}

func NewLooper(megs int) *LoopReader {
	random_bytes := make([]byte, consts.Blocksize)
	_, err := rand.Read(random_bytes)
	if err != nil {
		return nil
	}

	return &LoopReader{
		buf:  random_bytes,
		size: uint64(megs) * consts.Megabyte,
	}
}

func (r *LoopReader) Read(p []byte) (n int, err error) {
	n = len(p)
	toread := n
	if n > len(r.buf) {
		toread = len(r.buf)
	}
	if uint64(toread) > r.size {
		toread = int(r.size)
	}
	r.size -= uint64(toread)

	copy(p[0:toread], r.buf[0:toread])
	if r.size == 0 {
		return toread, io.EOF
	}
	return toread, nil
}
