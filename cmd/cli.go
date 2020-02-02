package cmd

import (
	"fmt"
	"goreader/crawler"

	"github.com/spf13/cobra"
)

var (
	baseURL    string
	paramStart string
	paramEnd   string
	urlSuffix  string
	path       string
)

// cliCmd represents the cli command
var cliCmd = &cobra.Command{
	Use:   "cli",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("crawl page %s, start: %s, end: %s, urlSuffix: %s, savePath: %s\n", baseURL, paramStart, paramEnd, urlSuffix, path)
		crawler.SetSavePath(path)
		// urls := crawler.GenURLs(baseURL, paramStart, paramEnd, urlSuffix)
		// crawler.ProcessURLs(urls)
		crawler.DownloadFiles(baseURL, paramStart, paramEnd, urlSuffix, path)
	},
}

func init() {
	rootCmd.AddCommand(cliCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// cliCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// cliCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	cliCmd.Flags().StringVarP(&baseURL, "url", "u", "", "URL of page.")
	cliCmd.Flags().StringVarP(&paramStart, "start", "s", "", "Start value.")
	cliCmd.Flags().StringVarP(&paramEnd, "end", "e", "", "End value.")
	cliCmd.Flags().StringVarP(&urlSuffix, "suffix", "x", "", "Suffix of url.")
	cliCmd.Flags().StringVarP(&path, "path", "p", "", "Save path of page.")
}
