package main

import (
	"expvar"
	"flag"
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/JustinAzoff/http_flood/consts"
	"github.com/JustinAzoff/http_flood/randomreader"
	"github.com/JustinAzoff/http_flood/web"
)

type context map[string]interface{}

var connections = expvar.NewInt("connections")
var downloads = expvar.NewInt("downloads")
var uploads = expvar.NewInt("uploads")
var downloadMegs = expvar.NewInt("downloadMegs")
var uploadMegs = expvar.NewInt("uploadMegs")

func extractIP(addrPort string) string {
	lastColon := strings.LastIndex(addrPort, ":")
	return addrPort[0:lastColon]
}

func addConnection() {
	connections.Add(1)
}

func removeConnection(downmegs int64, upmegs int64) {
	connections.Add(-1)

	if downmegs != 0 {
		downloads.Add(1)
		downloadMegs.Add(downmegs)
	}
	if upmegs != 0 {
		uploads.Add(1)
		uploadMegs.Add(upmegs)
	}
}

var indexTemplate = template.Must(template.New("").Parse(string(web.Index_html())))

func handleIndex(w http.ResponseWriter, req *http.Request) {
	indexTemplate.Execute(w, context{"Connections": connections,
		"Downloads":    downloads,
		"Uploads":      uploads,
		"DownloadMegs": downloadMegs,
		"UploadMegs":   uploadMegs,
	})
}

func floodHelper(w http.ResponseWriter, req *http.Request, reader io.Reader) {
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

func handleFlood(w http.ResponseWriter, req *http.Request) {
	s := req.FormValue("s")
	if s != "" {
		handleTimeFlood(w, req)
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
	floodHelper(w, req, randomreader.LimitedRandomGen(m*consts.Megabyte))
}

func handleTimeFlood(w http.ResponseWriter, req *http.Request) {
	ss := req.FormValue("s")
	if ss == "" {
		ss = "10"
	}
	s, err := strconv.ParseUint(ss, 10, 0)
	if err != nil {
		s = 10
	}

	log.Printf("flood starting addr=%s seconds=%d\n", extractIP(req.RemoteAddr), s)
	floodHelper(w, req, randomreader.TimedRandomGen(s))
}

func handleUpload(w http.ResponseWriter, req *http.Request) {
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

func serverMain() {
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
	myHandler.HandleFunc("/", handleIndex)
	myHandler.HandleFunc("/flood", handleFlood)
	myHandler.HandleFunc("/upload", handleUpload)
	myHandler.HandleFunc("/Flashflood.swf", func(w http.ResponseWriter, r *http.Request) {
		w.Write(web.Flashflood_swf())
	})
	myHandler.HandleFunc("/debug/vars", expvarHandler)

	log.Printf("Listening on %s\n", *addr)
	log.Fatal(s.ListenAndServe())
}
