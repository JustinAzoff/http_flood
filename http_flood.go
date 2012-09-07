package main
import (
    "net/http"
    "fmt"
    "io"
    "time"
    "crypto/rand"
    "runtime"
    "strconv"
)

const blocksize = 8*1024
const MEGABYTE = 1024*1024


func Hello(w http.ResponseWriter, req *http.Request) { 
    io.WriteString(w, `
<html>
<head><title>HTTP Flood Server </title></head>
<body>
<h1>HTTP Flood Server</h1>
<ul>
 <li><a href="/flood?m=1">1 megabyte file</a></li>
 <li><a href="/flood?m=10">10 megabyte file</a></li>
 <li><a href="/flood?m=100">100 megabyte file</a></li>
 <li><a href="/flood?m=1024">1 gigabyte file</a></li>
 <li><a href="/flood?m=10240">10 gigabyte file</a></li>
</ul>
</body>
</html>`)
} 

func Flood(w http.ResponseWriter, req *http.Request) { 
    random_bytes := make([]byte, blocksize)
    _, err := rand.Read(random_bytes)
    if err != nil {
        return
    }

    ms := req.FormValue("m")
    if ms == "" {
        ms = "1"
    }
    m, err := strconv.ParseUint(ms, 10, 0)
    if err != nil  {
        m = 1
    }
        
    fmt.Printf("flood starting addr=%s\n", req.RemoteAddr)
    start := time.Now()
    status := "finished"
    var written uint64 = 0

    fmt.Println(uint64(m)*MEGABYTE)
    w.Header().Set("Content-length", strconv.FormatUint(m*MEGABYTE, 10))
    for ; written < m*MEGABYTE; written += blocksize{
        _, err := w.Write(random_bytes)
        if err != nil {
            status = "aborted"
            break;
        }
        runtime.Gosched()
    }
    duration := time.Since(start)
    megabytes := float64(written)/MEGABYTE
    mbs := megabytes/duration.Seconds()
    fmt.Printf("flood %s addr=%s duration=%s megabytes=%.1f speed=%.1fMB/s\n", status, req.RemoteAddr, duration, megabytes, mbs)
} 

func main() { 
    myHandler := http.NewServeMux()
    s := &http.Server{
        Addr:           ":7070",
        Handler:        myHandler,
        ReadTimeout:    10 * time.Second,
        WriteTimeout:   60 * time.Second,
        MaxHeaderBytes: 1 << 20,
    }
    myHandler.HandleFunc("/", Hello) 
    myHandler.HandleFunc("/flood", Flood) 
    fmt.Println("Listening on port 7070")
    s.ListenAndServe()
} 
