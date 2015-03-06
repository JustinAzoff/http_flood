package common

import (
	"crypto/rand"
	"io"
	"time"

	"github.com/JustinAzoff/http_flood/consts"
)

// A RandomGen is an io.Reader that returns random data
type RandomGen struct {
	buf []byte
}

// NewRandomGen creates a new RandomGen
func NewRandomGen() *RandomGen {
	randomBytes := make([]byte, consts.Blocksize)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return nil
	}

	return &RandomGen{
		buf: randomBytes,
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

// LimitedRandomGen returns a RandomGen of a specific Size
func LimitedRandomGen(n uint64) io.Reader {
	return io.LimitReader(NewRandomGen(), int64(n))
}

// NewTimedReader creates a new TimedReader
func NewTimedReader(r io.Reader, n uint64) io.Reader {
	start := time.Now()
	end := start.Add(time.Duration(n) * time.Second)
	return &TimedReader{r, end, 10, 0}
}

//TimedReader is similar to an io.LimitReader but limits based on time, not size
type TimedReader struct {
	R             io.Reader // underlying reader
	End           time.Time //When to stop reading
	CheckInterval int       //how often to check the time
	Calls         int       //how many calls have been made so far
}

func (l *TimedReader) Read(p []byte) (n int, err error) {
	l.Calls++
	if l.Calls == l.CheckInterval {
		l.Calls = 0
		if time.Now().After(l.End) {
			return 0, io.EOF
		}
	}
	n, err = l.R.Read(p)
	return
}

// TimedRandomGen returns a RandomGen for a specific time duration
func TimedRandomGen(s uint64) io.Reader {
	return NewTimedReader(NewRandomGen(), s)
}
