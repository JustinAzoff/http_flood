package main

import (
	"github.com/JustinAzoff/http_flood/commands"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	commands.Execute()
}
