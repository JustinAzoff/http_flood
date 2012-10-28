package main

import (
	"../consts"
	"../common"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func download(host string, megs int) {
	status := "finished"
	buffer := make([]byte, consts.Blocksize)
	url := fmt.Sprintf("http://%s/flood?m=%d", host, megs)

	var read uint64 = 0
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	for {
		n, err := io.ReadAtLeast(resp.Body, buffer, consts.Blocksize)
		if err != nil {
			if err != io.EOF {
				status = "aborted"
			}
			break
		}
		read += uint64(n)

	}
	duration := time.Since(start)
	megabytes := float64(read) / consts.Megabyte
	mbs := megabytes / duration.Seconds()

	fmt.Printf("download %s duration=%s megabytes=%.1f speed=%.1fMB/s\n", status, duration, megabytes, mbs)
}

func upload(host string, megs int) {
	url := fmt.Sprintf("http://%s/upload?m=%d", host, megs)

	resp, err := http.Post(url, "application/octet-stream", common.NewLooper(megs))
	if err != nil {
		log.Fatal(err)
	}

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
