package common

import (
	"../consts"
	"crypto/rand"
	"io"
)

type RandomGen struct {
	buf []byte
}

func NewRandomGen() *RandomGen {
	random_bytes := make([]byte, consts.Blocksize)
	_, err := rand.Read(random_bytes)
	if err != nil {
		return nil
	}

	return &RandomGen{
		buf: random_bytes,
	}
}

func (r *RandomGen) Read(p []byte) (n int, err error) {
	n = len(p)
	toread := n
	if n > len(r.buf) {
		toread = len(r.buf)
	}

	copy(p[0:toread], r.buf[0:toread])
	return toread, nil
}

func LimitedRandomGen(n uint64) io.Reader {
    return io.LimitReader(NewRandomGen(), int64(n))
}
