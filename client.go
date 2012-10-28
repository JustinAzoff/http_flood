package main

import (
	"crypto/rand"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

const blocksize = 32  * 1024
const MEGABYTE = 1024 * 1024

func download(host string, megs int) {
	status := "finished"
	buffer := make([]byte, blocksize)
	url := fmt.Sprintf("http://%s/flood?m=%d", host, megs)

	var read uint64 = 0
	start := time.Now()
	resp, err := http.Get(url)
	defer resp.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	for {
		n, err := io.ReadAtLeast(resp.Body, buffer, blocksize)
		if err != nil {
			if err != io.EOF {
				status = "aborted"
			}
			break
		}
		read += uint64(n)

	}
	duration := time.Since(start)
	megabytes := float64(read) / MEGABYTE
	mbs := megabytes / duration.Seconds()

	fmt.Printf("download %s duration=%s megabytes=%.1f speed=%.1fMB/s\n", status, duration, megabytes, mbs)
}

type LoopReader struct {
	pos  int
	buf  []byte
	size uint64
}

func NewLooper(megs int) *LoopReader {
	random_bytes := make([]byte, blocksize)
	_, err := rand.Read(random_bytes)
	if err != nil {
		return nil
	}

	return &LoopReader{
		buf:  random_bytes,
		size: uint64(megs) * MEGABYTE,
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

func upload(host string, megs int) {
	url := fmt.Sprintf("http://%s/upload?m=%d", host, megs)

	resp, _ := http.Post(url, "application/octet-stream", NewLooper(megs))

	io.Copy(os.Stdout, resp.Body)
}

func main() {
	var megabytes int
	var host string
	flag.IntVar(&megabytes, "megs", 1024, "megabytes to download")
	flag.StringVar(&host, "host", "localhost:7070", "Host to connect to")
	flag.Parse()

	download(host, megabytes)
	upload(host, megabytes)

}
