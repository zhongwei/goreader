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

type article struct {
	title string
	url   string
}

func init() {
	flag.Parse()
}

func main() {
	args := flag.Args()

	urls := gen_urls(args[0], args[1], args[2])
	dstFile, _ := os.Create(args[3])
	defer dstFile.Close()

	for _, url := range urls {
		fmt.Println("url = " + url)
		articles := fetch(url)

		for _, article := range articles {
			fmt.Println(article.url)
			dstFile.WriteString(article.title + "|" + article.url + "\n")
		}

	}

}

func gen_urls(base, start, end string) (urls []string) {
	startNum, _ := strconv.Atoi(start)
	endNum, _ := strconv.Atoi(end)

	for i := startNum; i <= endNum; i++ {
		urls = append(urls, base+strconv.Itoa(i)+".html")
	}

	return
}

func fetch(url string) (articles []article) {
	last_slash := strings.LastIndex(url, "/")
	doc, err := goquery.NewDocument(url)
	if err != nil {
		log.Fatal(err)
	}

	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		current_url, exist := s.Attr("href")
		article_title := s.Text()
		article_url := ""
		if exist {
			if strings.Contains(current_url, "thread-") && len(article_title) > 0 {
				if strings.Contains(current_url, "http") {
					article_url = current_url
				} else {
					article_url = url[:last_slash+1] + current_url
				}
				this_article := article{}
				this_article = article{article_title, article_url}
				articles = append(articles, this_article)
			}
		}

	})

	return
}
