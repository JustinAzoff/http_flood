package main

import (
	"fmt"
	"os"
	"runtime"
)

func usage() {
	bin := os.Args[0]
	fmt.Printf("Usage:\n  %s client ...\n  %s server ...\n", bin, bin)
	os.Exit(1)
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	if len(os.Args) < 2 {
		usage()
	}
	mode := os.Args[1]
	os.Args = os.Args[1:]

	switch mode {
	case "client":
		clientMain()
	case "server":
		serverMain()
	default:
		usage()
	}
}
