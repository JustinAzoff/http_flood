package main

import (
	"expvar"
	"flag"
	"fmt"
	"github.com/JustinAzoff/http_flood/common"
	"github.com/JustinAzoff/http_flood/consts"
	"github.com/JustinAzoff/http_flood/web"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"runtime"
	"strconv"
	"strings"
	"time"
)

type Context map[string]interface{}

var connections = expvar.NewInt("connections")
var downloads = expvar.NewInt("downloads")
var uploads = expvar.NewInt("uploads")
var download_megs = expvar.NewInt("download_megs")
var upload_megs = expvar.NewInt("upload_megs")

func extractIP(s string) string {
	return strings.Split(s, ":")[0]
}

func addConnection() {
	connections.Add(1)
}

func removeConnection(downmegs int64, upmegs int64) {
	connections.Add(-1)

	if downmegs != 0 {
		downloads.Add(1)
		download_megs.Add(downmegs)
	}
	if upmegs != 0 {
		uploads.Add(1)
		upload_megs.Add(upmegs)
	}
}

var indexTemplate = template.Must(template.New("").Parse(string(web.Index_html())))

func Hello(w http.ResponseWriter, req *http.Request) {
	indexTemplate.Execute(w, Context{"Connections": connections,
		"Downloads":    downloads,
		"Uploads":      uploads,
		"DownloadMegs": download_megs,
		"UploadMegs":   upload_megs,
	})
}

func FloodHelper(w http.ResponseWriter, req *http.Request, reader io.Reader) {
	addConnection()
	start := time.Now()
	status := "finished"
	written, err := io.Copy(w, reader)

	if err != nil {
		status = "aborted"
	}
	duration := time.Since(start)
	megabytes := float64(written) / consts.Megabyte
	mbs := megabytes / duration.Seconds()
	removeConnection(int64(megabytes), 0)
	log.Printf("flood %s addr=%s duration=%s megabytes=%.1f speed=%.1fMB/s\n", status, extractIP(req.RemoteAddr), duration, megabytes, mbs)
}

func Flood(w http.ResponseWriter, req *http.Request) {
	s := req.FormValue("s")
	if s != "" {
		TimeFlood(w, req)
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
	w.Header().Set("Content-length", strconv.FormatUint(m*consts.Megabyte, 10))
	log.Printf("flood starting addr=%s megabytes=%d\n", extractIP(req.RemoteAddr), m)
	FloodHelper(w, req, common.LimitedRandomGen(m*consts.Megabyte))
}

func TimeFlood(w http.ResponseWriter, req *http.Request) {
	ss := req.FormValue("s")
	if ss == "" {
		ss = "10"
	}
	s, err := strconv.ParseUint(ss, 10, 0)
	if err != nil {
		s = 10
	}

	log.Printf("flood starting addr=%s seconds=%d\n", extractIP(req.RemoteAddr), s)
	FloodHelper(w, req, common.TimedRandomGen(s))
}

func Upload(w http.ResponseWriter, req *http.Request) {
	addConnection()

	log.Printf("upload starting addr=%s\n", extractIP(req.RemoteAddr))
	start := time.Now()
	status := "finished"

	written, err := io.Copy(ioutil.Discard, req.Body)
	if err != nil {
		status = "aborted"
	}

	duration := time.Since(start)
	megabytes := float64(written) / consts.Megabyte
	removeConnection(0, int64(megabytes))
	mbs := megabytes / duration.Seconds()
	message := fmt.Sprintf("upload %s addr=%s duration=%s megabytes=%.1f speed=%.1fMB/s\n", status, extractIP(req.RemoteAddr), duration, megabytes, mbs)
	log.Print(message)
	io.WriteString(w, message)
}

func expvarHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	fmt.Fprintf(w, "{\n")
	first := true
	expvar.Do(func(kv expvar.KeyValue) {
		if !first {
			fmt.Fprintf(w, ",\n")
		}
		first = false
		fmt.Fprintf(w, "%q: %s", kv.Key, kv.Value)
	})
	fmt.Fprintf(w, "\n}\n")
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
	myHandler.HandleFunc("/Flashflood.swf", func(w http.ResponseWriter, r *http.Request) {
		w.Write(web.Flashflood_swf())
	})
	myHandler.HandleFunc("/debug/vars", expvarHandler)

	log.Printf("Listening on %s\n", *addr)
	log.Fatal(s.ListenAndServe())
}
