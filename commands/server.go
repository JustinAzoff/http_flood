package commands

import (
	"github.com/JustinAzoff/http_flood/floodlib"
	"github.com/spf13/cobra"
)

var bindAddr string

var cmdServer = &cobra.Command{
	Use:   "server [args]",
	Short: "Run server",
	Long:  "Run server",
	Run: func(cmd *cobra.Command, args []string) {
		floodlib.RunServer(bindAddr)
	},
}

func init() {
	cmdServer.Flags().StringVarP(&bindAddr, "bind", "b", ":7070", "Address to bind to")
}
