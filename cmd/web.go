package cmd

import (
	"fmt"

	"goreader/web"

	"github.com/spf13/cobra"
)

var (
	rootPath string
	port     string
)

// webCmd represents the web command
var webCmd = &cobra.Command{
	Use:   "web",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("web called rootpath: %s, port: %s", rootPath, port)
		web.Start(rootPath, port)
	},
}

func init() {
	rootCmd.AddCommand(webCmd)
	webCmd.Flags().StringVarP(&rootPath, "root", "r", "", "Doc root path of web server.")
	webCmd.Flags().StringVarP(&port, "port", "t", "", "Port number of web server.")
}
