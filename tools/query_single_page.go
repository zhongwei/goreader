package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func init() {
	flag.Parse()
}

func main() {
	args := flag.Args()

	finding_urls := fetch(args[0])

	dstFile, _ := os.Create(args[1])
	defer dstFile.Close()
	for _, find_url := range finding_urls {
		full_url := args[0] + find_url
		fmt.Println(full_url)
		dstFile.WriteString(full_url + "\n")
	}
}

func fetch(url string) (urls []string) {
	doc, err := goquery.NewDocument(url)
	if err != nil {
		log.Fatal(err)
	}

	htmlcontent, err := doc.Html()
	fmt.Println(htmlcontent)
	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		current_url, exist := s.Attr("href")
		if exist {
			if strings.Contains(current_url, "galleries") {
				urls = append(urls, current_url)
			}
		}
	})

	return
}
