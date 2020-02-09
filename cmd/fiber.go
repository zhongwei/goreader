package cmd

import (
	"fmt"

	"goreader/fiber"

	"github.com/spf13/cobra"
)

var (
	fiberPath string
	fiberPort string
)

// fiberCmd represents the fiber command
var fiberCmd = &cobra.Command{
	Use:   "fiber",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("fiber called")
		fiber.Start(fiberPath, fiberPort)
	},
}

func init() {
	rootCmd.AddCommand(fiberCmd)
	fiberCmd.Flags().StringVarP(&fiberPath, "root", "r", "", "Doc root path of web server.")
	fiberCmd.Flags().StringVarP(&fiberPort, "port", "p", "", "Port number of web server.")
}
