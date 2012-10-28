package main

import (
	"../common"
	"../consts"
	"crypto/rand"
	"flag"
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"runtime"
	"strconv"
	"time"
)

var connections = make(map[*http.Request]bool)
var random_bytes = make([]byte, consts.Blocksize)

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

func Init() {
	_, err := rand.Read(random_bytes)
	if err != nil {
		log.Fatal(err)
	}
}

func Flood(w http.ResponseWriter, req *http.Request) {
	connections[req] = true
	defer delete(connections, req)

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
	randReader := common.NewRandomGen(m)
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

	fmt.Printf("upload starting addr=%s\n", req.RemoteAddr)
	start := time.Now()
	status := "finished"

	written, err := io.Copy(ioutil.Discard, req.Body)
	if err != nil {
		status = "aborted"
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
	addr := flag.String("addr", ":7070", "Address to bind to")
	flag.Parse()
	myHandler := http.NewServeMux()
	s := &http.Server{
		Addr:           *addr,
		Handler:        myHandler,
		ReadTimeout:    60 * time.Second,
		WriteTimeout:   60 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	myHandler.HandleFunc("/", Hello)
	myHandler.HandleFunc("/flood", Flood)
	myHandler.HandleFunc("/upload", Upload)
	Init()
	fmt.Printf("Listening on %s\n", *addr)
	log.Fatal(s.ListenAndServe())
}
