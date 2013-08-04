package main

import (
	"./common"
    "log"
	"./consts"
	"fmt"
	"io"
	"io/ioutil"
    "time"
)

func test_helper(reader io.Reader) {
    start := time.Now()
	written, err := io.Copy(ioutil.Discard, reader)
    if err != nil {
        log.Fatal(err)
    }
    duration := time.Since(start)
    megabytes := float64(written) / consts.Megabyte;
    mbs := float64(megabytes) / duration.Seconds()
    fmt.Printf("%f Megabytes in %f seconds. %.1f MB/s\n", megabytes, duration.Seconds(), mbs)
}

func size_limited(megabytes uint64) {
	sl := common.LimitedRandomGen(megabytes * consts.Megabyte)
    test_helper(sl)
}

func time_limited(seconds uint64) {
	tl := common.TimedRandomGen(seconds)
    test_helper(tl)
}

func main() {
    size_limited(4000)
    time_limited(4)
}
