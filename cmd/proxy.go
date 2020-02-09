package cmd

import (
	"fmt"

	"goreader/proxy"

	"github.com/spf13/cobra"
)

// proxyCmd represents the proxy command
var proxyCmd = &cobra.Command{
	Use:   "proxy",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("proxy called")
		proxy.Start()
	},
}

func init() {
	rootCmd.AddCommand(proxyCmd)
}
