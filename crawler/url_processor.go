package crawler

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/PuerkitoBio/goquery"
)

var (
	savePath string
)

type Article struct {
	page     int
	title    string
	url      string
	author   string
	summary  string
	words    int32
	category string
}

func SetSavePath(path string) {
	savePath = path
}

func GenURLs(baseURL, start, end, suffix string) []string {

	min, _ := strconv.Atoi(start)
	max, _ := strconv.Atoi(end)

	var urls []string = make([]string, 0)

	for i := min; i <= max; i++ {
		fullURL := baseURL + strconv.Itoa(i) + suffix
		urls = append(urls, fullURL)
	}

	return urls
}

func ProcessURLs(urls []string) {
	for index, url := range urls {
		response, _ := http.Get(url)
		defer response.Body.Close()
		ProcessArticle(response.Body)
		body, _ := ioutil.ReadAll(response.Body)
		bodystr := string(body)
		SavePage(strconv.Itoa(index+1)+".htm", bodystr)
	}
}

func ProcessArticle(r io.Reader) {
	articles := make([]Article, 0)

	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		log.Fatal(err)
	}
	doc.Find(".bookinfo").Each(func(i int, s *goquery.Selection) {
		aNode := s.Find("h4 i a")
		title := aNode.Text()
		url, _ := aNode.Attr("href")
		author := s.Find(".intro_line span").Text()
		summary := s.Find(".update span").Text()
		category := s.Find(".author").Text()

		var article = Article{title: title, url: url, author: author, summary: summary, category: category}

		articles = append(articles, article)
	})
	fmt.Printf("articles: %v \n", articles)
	return

}

func SavePage(name, content string) {
	f, err := os.Create(savePath + string(os.PathSeparator) + name)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	f.WriteString(content)
}
