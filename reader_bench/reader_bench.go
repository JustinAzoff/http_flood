package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"time"

	"github.com/JustinAzoff/http_flood/consts"
	"github.com/JustinAzoff/http_flood/randomreader"
)

func testHelper(reader io.Reader) {
	start := time.Now()
	written, err := io.Copy(ioutil.Discard, reader)
	if err != nil {
		log.Fatal(err)
	}
	duration := time.Since(start)
	megabytes := float64(written) / consts.Megabyte
	mbs := float64(megabytes) / duration.Seconds()
	fmt.Printf("%f Megabytes in %f seconds. %.1f MB/s\n", megabytes, duration.Seconds(), mbs)
}

func sizeLimited(megabytes uint64) {
	sl := randomreader.LimitedRandomGen(megabytes * consts.Megabyte)
	testHelper(sl)
}

func timeLimited(seconds uint64) {
	tl := randomreader.TimedRandomGen(seconds)
	testHelper(tl)
}

func main() {
	sizeLimited(200000)
	timeLimited(4)
}
