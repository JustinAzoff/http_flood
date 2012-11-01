package main

import (
	"./common"
	"./consts"
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

type serverStatusStruct struct {
	Connections  int
	Downloads    int
	Uploads      int
	DownloadMegs uint64
	UploadMegs   uint64
}

var serverStatus = serverStatusStruct{}
var serverStatusChan = make(chan serverStatusStruct, 2)

func addConnection() {
	serverStatusChan <- serverStatusStruct{Connections: 1}
}

func removeConnection(downmegs uint64, upmegs uint64) {

	var downloads, uploads = 0, 0
	if downmegs != 0 {
		downloads = 1
	}
	if upmegs != 0 {
		uploads = 1
	}

	update := serverStatusStruct{
		Connections:  -1,
		Downloads:    downloads,
		Uploads:      uploads,
		DownloadMegs: downmegs,
		UploadMegs:   upmegs,
	}
	serverStatusChan <- update

}

func connectionTracker() {
	for {
		update := <-serverStatusChan
		serverStatus.Connections += update.Connections
		serverStatus.Downloads += update.Downloads
		serverStatus.Uploads += update.Uploads
		serverStatus.DownloadMegs += update.DownloadMegs
		serverStatus.UploadMegs += update.UploadMegs
	}
}

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
<p> Current connections: {{.Connections}} </p>
<p> Total Downloads: {{.Downloads}} </p>
<p> Total Uploads: {{.Uploads}} </p>
<p> Megabytes Downloaded: {{.DownloadMegs}} </p>
<p> Megabytes Uploaded: {{.UploadMegs}} </p>
</html>`))

func Hello(w http.ResponseWriter, req *http.Request) {
	indexTemplate.Execute(w, serverStatus)
}

func Flood(w http.ResponseWriter, req *http.Request) {
	addConnection()

	ms := req.FormValue("m")
	if ms == "" {
		ms = "1"
	}
	m, err := strconv.ParseUint(ms, 10, 0)
	if err != nil {
		m = 1
	}

	log.Printf("flood starting addr=%s megabytes=%d\n", req.RemoteAddr, m)
	start := time.Now()
	status := "finished"
	w.Header().Set("Content-length", strconv.FormatUint(m*consts.Megabyte, 10))
	written, err := io.Copy(w, common.LimitedRandomGen(m*consts.Megabyte))

	if err != nil {
		status = "aborted"
	}
	duration := time.Since(start)
	megabytes := float64(written) / consts.Megabyte
	mbs := megabytes / duration.Seconds()
	removeConnection(uint64(megabytes), 0)
	log.Printf("flood %s addr=%s duration=%s megabytes=%.1f speed=%.1fMB/s\n", status, req.RemoteAddr, duration, megabytes, mbs)
}

func Upload(w http.ResponseWriter, req *http.Request) {
	addConnection()

	log.Printf("upload starting addr=%s\n", req.RemoteAddr)
	start := time.Now()
	status := "finished"

	written, err := io.Copy(ioutil.Discard, req.Body)
	if err != nil {
		status = "aborted"
	}

	duration := time.Since(start)
	megabytes := float64(written) / consts.Megabyte
	removeConnection(0, uint64(megabytes))
	mbs := megabytes / duration.Seconds()
	message := fmt.Sprintf("upload %s addr=%s duration=%s megabytes=%.1f speed=%.1fMB/s\n", status, req.RemoteAddr, duration, megabytes, mbs)
	log.Print(message)
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
	log.Printf("Listening on %s\n", *addr)
	go connectionTracker()
	log.Fatal(s.ListenAndServe())
}
