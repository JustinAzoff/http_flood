package main

import (
    "flag"
    "fmt"
    "io"
    "log"
    "net/http"
    "time"
)

const blocksize = 8 * 1024
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

    fmt.Printf("flood %s duration=%s megabytes=%.1f speed=%.1fMB/s\n", status, duration, megabytes, mbs)
}

func main() {
    var megabytes int
    var host string
    flag.IntVar(&megabytes, "megs", 1024, "megabytes to download")
    flag.StringVar(&host, "host", "localhost:7070", "Host to connect to")
    flag.Parse()

    download(host, megabytes)
}
