package cmd

import (
	"fmt"

	"goreader/fiber"

	"github.com/spf13/cobra"
)

// fiberCmd represents the fiber command
var fiberCmd = &cobra.Command{
	Use:   "fiber",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("fiber called")
		fiber.Start()
	},
}

func init() {
	rootCmd.AddCommand(fiberCmd)
}
