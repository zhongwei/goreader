package crawler

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/PuerkitoBio/goquery"
	"gopkg.in/yaml.v2"
)

var (
	savePath string
)

type Article struct {
	title    string
	id       string
	author   string
	summary  string
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
	articles := []Article{}
	for index, url := range urls {
		res, err := http.Get(url)
		if err != nil {
			log.Fatal(err)
		}

		defer res.Body.Close()

		if res.StatusCode != 200 {
			log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
		}

		thisPageArticles := ProcessArticle(res.Body)
		articles = append(articles, thisPageArticles...)

		if index%5 == 0 {
			fmt.Printf("page: %v\n", index)
		}
	}
	articleYaml, _ := yaml.Marshal(&articles)
	SavePage(strconv.Itoa(len(urls))+".tmp", string(articleYaml))
}

func ProcessArticle(r io.Reader) []Article {
	articles := make([]Article, 0)

	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		log.Fatal(err)
	}
	doc.Find(".bookinfo").Each(func(i int, s *goquery.Selection) {
		aNode := s.Find("h4 i a")
		title := aNode.Text()
		url, _ := aNode.Attr("href")
		id := url[9 : len(url)-5]
		author := s.Find(".intro_line span").Text()
		summary := s.Find(".update span").Text()
		category := s.Find(".author").Text()

		var article = Article{title, id, author, summary, category}

		articles = append(articles, article)
	})
	return articles

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
