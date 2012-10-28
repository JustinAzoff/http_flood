package main

import (
	"../common"
	"../consts"
	"crypto/rand"
	"flag"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"runtime"
	"strconv"
	"time"
)

var connections = make(map[*http.Request]bool)

var indexTemplate = template.Must(template.New("").Parse(`
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
<p> Current connections: {{.}} </p>
</html>`))

func Hello(w http.ResponseWriter, req *http.Request) {
	indexTemplate.Execute(w, len(connections))
}

func Flood(w http.ResponseWriter, req *http.Request) {
	connections[req] = true
	defer delete(connections, req)

	random_bytes := make([]byte, consts.Blocksize)
	_, err := rand.Read(random_bytes)
	if err != nil {
		return
	}

	ms := req.FormValue("m")
	if ms == "" {
		ms = "1"
	}
	m, err := strconv.ParseUint(ms, 10, 0)
	if err != nil {
		m = 1
	}

	fmt.Printf("flood starting addr=%s megabytes=%d\n", req.RemoteAddr, m)
	start := time.Now()
	status := "finished"
	randReader := common.NewRandomGen(int(m))
	w.Header().Set("Content-length", strconv.FormatUint(m*consts.Megabyte, 10))
	written, err := io.Copy(w, randReader)

	if err != nil {
		status = "aborted"
	}
	duration := time.Since(start)
	megabytes := float64(written) / consts.Megabyte
	mbs := megabytes / duration.Seconds()
	fmt.Printf("flood %s addr=%s duration=%s megabytes=%.1f speed=%.1fMB/s\n", status, req.RemoteAddr, duration, megabytes, mbs)
}

func Upload(w http.ResponseWriter, req *http.Request) {
	connections[req] = true
	defer delete(connections, req)

	buf := make([]byte, consts.Blocksize)
	fmt.Printf("upload starting addr=%s\n", req.RemoteAddr)
	start := time.Now()
	status := "finished"
	var written uint64 = 0

	for {
		nr, er := req.Body.Read(buf)
		written += uint64(nr)
		if er != nil {
			break
		}
	}
	duration := time.Since(start)
	megabytes := float64(written) / consts.Megabyte
	mbs := megabytes / duration.Seconds()
	message := fmt.Sprintf("upload %s addr=%s duration=%s megabytes=%.1f speed=%.1fMB/s\n", status, req.RemoteAddr, duration, megabytes, mbs)
	fmt.Print(message)
	io.WriteString(w, message)
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	var port int
	flag.IntVar(&port, "port", 7070, "Port to bind to")
	flag.Parse()
	myHandler := http.NewServeMux()
	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", port),
		Handler:        myHandler,
		ReadTimeout:    60 * time.Second,
		WriteTimeout:   60 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	myHandler.HandleFunc("/", Hello)
	myHandler.HandleFunc("/flood", Flood)
	myHandler.HandleFunc("/upload", Upload)
	fmt.Printf("Listening on port %d\n", port)
	s.ListenAndServe()
}
