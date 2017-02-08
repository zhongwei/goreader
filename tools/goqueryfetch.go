package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/axgle/mahonia"
)

type article struct {
	title string
	url   string
	count int
	page  int
}

var articles []article

func init() {
	flag.Parse()
}

func main() {
	args := flag.Args()

	urls := gen_urls(args[0], args[1], args[2])

	for _, url := range urls {
		fmt.Println("url = " + url)
		fetch(url)
		time.Sleep(time.Second * 0)
	}

	dstFile, _ := os.Create(args[3])
	defer dstFile.Close()
	for _, current_article := range articles {
		aticle_address := current_article.url + "|" + current_article.title
		dstFile.WriteString(aticle_address + "\n")
	}

	os.Exit(1)
}

func gen_urls(base, start, end string) (urls []string) {
	startNum, _ := strconv.Atoi(start)
	endNum, _ := strconv.Atoi(end)

	for i := startNum; i <= endNum; i++ {
		urls = append(urls, base+strconv.Itoa(i))
	}

	return
}

func fetch(url string) {
	dec := mahonia.NewDecoder("gbk")

	last_slash := strings.LastIndex(url, "/")

	doc, err := goquery.NewDocument(url)
	htmlcontent, err := doc.Html()
	fmt.Println("-----------------------------------------------------------")
	fmt.Println(dec.ConvertString(htmlcontent))
	fmt.Println("-----------------------------------------------------------")
	if err != nil {
		log.Fatal(err)
	}
	doc.Find(".subject_t").Each(func(i int, s *goquery.Selection) {
		current_url, exist := s.Attr("href")
		this_article := article{}
		if exist {
			article_title := dec.ConvertString(s.Text())
			article_url := url[:last_slash+1] + current_url
			this_article = article{article_title, article_url, 0, 0}
		}
		articles = append(articles, this_article)
	})
}
