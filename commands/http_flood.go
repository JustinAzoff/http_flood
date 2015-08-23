package commands

import (
	"github.com/spf13/cobra"
)

var HttpFloodCmd = &cobra.Command{
	Use:   "http_flood",
	Short: "http_flood sends or receives data over http. fast.",
	Long:  "http_flood sends or receives data over http. fast.",
	Run:   nil,
}

func Execute() {
	HttpFloodCmd.AddCommand(cmdClient)
	HttpFloodCmd.AddCommand(cmdServer)
	HttpFloodCmd.Execute()
}
