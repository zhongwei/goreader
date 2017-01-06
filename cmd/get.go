package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/zhongwei/goreader/netdata"
)

var (
    url string
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
        siteContent := netdata.Get(url)
		fmt.Println("get called url " + url)
        fmt.Println(siteContent)
	},
}

func init() {
	RootCmd.AddCommand(getCmd)
    getCmd.Flags().StringVarP(&url, "url", "u", "", "Get site url address.")
	getCmd.PersistentFlags().String("foo", "", "A help for foo")
	getCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
