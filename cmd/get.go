package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/zhongwei/goreader/netdata"
	"github.com/zhongwei/goreader/store"
)

var (
    url string
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "A brief description of your command",
	Long: ` and usage of using your command. For example: Cobra is a CLI library for Go that empockly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		    log.Println("get content from : " + url)
        siteContent := netdata.Get(url)
        store.Save(url, siteContent)
	},
}

func init() {
	RootCmd.AddCommand(getCmd)
  getCmd.PersistentFlags().String("foo", "", "A help for foo")
  getCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
  getCmd.Flags().StringVarP(&url, "url", "u", "", "Get site url address.")
}
