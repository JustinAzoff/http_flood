package commands

import (
	"github.com/JustinAzoff/http_flood/floodlib"
	"github.com/spf13/cobra"
)

var host string
var seconds uint64
var megabytes uint64
var full_duplex bool

var cmdClient = &cobra.Command{
	Use:   "client [args]",
	Short: "Run client",
	Long:  "Run client",
	Run: func(cmd *cobra.Command, args []string) {
		floodlib.RunClient(host, seconds, megabytes, full_duplex)
	},
}

func init() {
	cmdClient.Flags().StringVar(&host, "host", "localhost:7070", "Host to connect to")
	cmdClient.Flags().Uint64VarP(&seconds, "seconds", "s", 10, "seconds to download")
	cmdClient.Flags().Uint64VarP(&megabytes, "megs", "m", 0, "seconds to download")
	cmdClient.Flags().BoolVarP(&full_duplex, "full-duplex", "f", false, "Full duplex test: download and upload at the same time")
}
