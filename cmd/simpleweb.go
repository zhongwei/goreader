package cmd

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/spf13/cobra"
)

// simplewebCmd represents the simpleweb command
var simplewebCmd = &cobra.Command{
	Use:   "simpleweb",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("simpleweb called")
		start()
	},
}

func init() {
	rootCmd.AddCommand(simplewebCmd)
}

type webHandler1 struct {
}

func (webHandler1) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("web1"))
}

type webHandler2 struct {
}

func (webHandler2) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("web2"))
}

func start() {
	c := make(chan os.Signal)
	go func() {
		http.ListenAndServe(":9091", webHandler1{})
	}()

	go func() {
		http.ListenAndServe(":9092", webHandler2{})
	}()

	signal.Notify(c, os.Interrupt)
	s := <-c
	log.Println(s)
}
