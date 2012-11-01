package main

import (
	"../common"
	"../consts"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

func download(host string, megs uint64) {
	status := "finished"
	url := fmt.Sprintf("http://%s/flood?m=%d", host, megs)

	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	read, err := io.Copy(ioutil.Discard, resp.Body)
	if err != nil {
		status = "aborted"
	}
	duration := time.Since(start)
	megabytes := float64(read) / consts.Megabyte
	mbs := megabytes / duration.Seconds()

	fmt.Printf("download %s duration=%s megabytes=%.1f speed=%.1fMB/s\n", status, duration, megabytes, mbs)
}

func upload(host string, megs uint64) {
	url := fmt.Sprintf("http://%s/upload?m=%d", host, megs)

	resp, err := http.Post(url, "application/octet-stream", common.LimitedRandomGen(megs*consts.Megabyte))
	if err != nil {
		log.Fatal(err)
	}

	io.Copy(os.Stdout, resp.Body)
}

func notify(c chan bool, f func()) {
	f()
	c <- true
}

func main() {
	megabytes := flag.Uint64("megs", 1024, "megabytes to download")
	host := flag.String("host", "localhost:7070", "Host to connect to")
	fd := flag.Bool("full", false, "Download and upload at the same time")
	flag.Parse()

	if *fd {
		c := make(chan bool, 2)
		go notify(c, func() { download(*host, *megabytes) })
		go notify(c, func() { upload(*host, *megabytes) })
		<-c
		<-c
	} else {
		download(*host, *megabytes)
		upload(*host, *megabytes)
	}

}
