package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/axgle/mahonia"
	"golang.org/x/net/html"
)

func init() {
	flag.Parse()
}

func main() {
	args := flag.Args()

	fetch(args[0])

	os.Exit(1)
}

func get_href(t html.Token) (ok bool, href string) {
	for _, a := range t.Attr {
		if a.Key == "href" {
			href = a.Val
			ok = true
		}
	}

	return
}

func fetch(url string) {
	resp, _ := http.Get(url)

	b := resp.Body
	defer b.Close()

	z := html.NewTokenizer(b)
	dec := mahonia.NewDecoder("gbk")

	for {
		tt := z.Next()
		switch {
		case tt == html.ErrorToken:
			return
		case tt == html.StartTagToken:
			t := z.Token()

			is_anchor := t.Data == "a"
			if !is_anchor {
				continue
			}

			ok, url := get_href(t)

			if !ok {
				continue
			}

			is_post := strings.Contains(url, "read.php?tid")

			if is_post {
				//fmt.Println(url)
				z.Next()
				link_name := z.Token().String()
				is_not_name := (strings.Contains(link_name, "<img") || strings.Contains(link_name, "201"))
				if !is_not_name {
					fmt.Println(dec.ConvertString(link_name))
				}
			}
		}
	}

}
