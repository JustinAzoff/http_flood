package floodlib

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/JustinAzoff/http_flood/consts"
	"github.com/JustinAzoff/http_flood/randomreader"
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
		reader = randomreader.TimedRandomGen(seconds)
	} else {
		reader = randomreader.LimitedRandomGen(megs * consts.Megabyte)
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

func RunClient(host string, seconds, megabytes uint64, full_duplex bool) {
	if megabytes != 0 {
		seconds = 0
	}

	if full_duplex {
		c := make(chan bool, 2)
		go notify(c, func() { download(host, megabytes, seconds) })
		go notify(c, func() { upload(host, megabytes, seconds) })
		<-c
		<-c
	} else {
		download(host, megabytes, seconds)
		upload(host, megabytes, seconds)
	}

}
