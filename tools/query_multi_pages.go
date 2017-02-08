package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func init() {
	flag.Parse()
}

func main() {
	args := flag.Args()

	urls := gen_urls(args[0], args[1], args[2])

	var finding_urls []string

	for _, url := range urls {
		fmt.Println("url = " + url)
		finding_urls = fetch(url)
	}

	dstFile, _ := os.Create(args[3])
	defer dstFile.Close()
	for _, find_url := range finding_urls {
		fmt.Println(find_url)
		dstFile.WriteString(find_url + "\n")
	}
}

func gen_urls(base, start, end string) (urls []string) {
	startNum, _ := strconv.Atoi(start)
	endNum, _ := strconv.Atoi(end)

	for i := startNum; i <= endNum; i++ {
		urls = append(urls, base+strconv.Itoa(i))
	}

	return
}

func fetch(url string) (urls []string) {

	doc, err := goquery.NewDocument(url)
	if err != nil {
		log.Fatal(err)
	}
	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		current_url, exist := s.Attr("href")
		if exist {
			if strings.Contains(current_url, "http") {
				urls = append(urls, current_url[15:])
			}
		}
	})

	return
}
