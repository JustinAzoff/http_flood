package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/JustinAzoff/http_flood/common"
	"github.com/JustinAzoff/http_flood/consts"
)

func download(host string, megs uint64, seconds uint64) {
	status := "finished"
	url := fmt.Sprintf("http://%s/flood?m=%d", host, megs)
	if seconds != 0 {
		url = fmt.Sprintf("http://%s/flood?s=%d", host, seconds)
	}

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

	log.Printf("download %s duration=%s megabytes=%.1f speed=%.1fMB/s\n", status, duration, megabytes, mbs)
}

func upload(host string, megs uint64, seconds uint64) {
	url := fmt.Sprintf("http://%s/upload", host)
	var reader io.Reader
	if seconds != 0 {
		reader = common.TimedRandomGen(seconds)
	} else {
		reader = common.LimitedRandomGen(megs * consts.Megabyte)
	}

	resp, err := http.Post(url, "application/octet-stream", reader)
	if err != nil {
		log.Fatal(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%s", body)
}

func notify(c chan bool, f func()) {
	f()
	c <- true
}

func main() {
	seconds := flag.Uint64("seconds", 10, "seconds to download")
	megabytes := flag.Uint64("megs", 0, "megabytes to download")
	host := flag.String("host", "localhost:7070", "Host to connect to")
	fd := flag.Bool("full", false, "Download and upload at the same time")
	flag.Parse()

	if *megabytes != 0 {
		*seconds = 0
	}

	if *fd {
		c := make(chan bool, 2)
		go notify(c, func() { download(*host, *megabytes, *seconds) })
		go notify(c, func() { upload(*host, *megabytes, *seconds) })
		<-c
		<-c
	} else {
		download(*host, *megabytes, *seconds)
		upload(*host, *megabytes, *seconds)
	}

}
