package crawler

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/PuerkitoBio/goquery"
	"github.com/axgle/mahonia"
)

var (
	savePath string
)

type Article struct {
	ID       string
	Title    string
	Author   string
	Summary  string
	Category string
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
		res, err := MockRequest(url)
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
	articleBytes, _ := json.Marshal(&articles)
	// fmt.Printf("yaml: %v\n", articles)
	// fmt.Printf("json: %v\n", articleBytes)
	SavePage(strconv.Itoa(len(urls))+".tmp", articleBytes)
}

func ProcessArticle(r io.Reader) []Article {
	articles := make([]Article, 0)

	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		log.Fatal(err)
	}
	doc.Find(".bookinfo").Each(func(i int, s *goquery.Selection) {
		aNode := s.Find("h4 i a")
		title := ConvertToUTF8(aNode.Text())
		url, _ := aNode.Attr("href")
		id := ConvertToUTF8(ConvertToUTF8(url[9 : len(url)-5]))
		author := ConvertToUTF8(s.Find(".intro_line span").Text())
		summary := ConvertToUTF8(s.Find(".update span").Text())
		category := ConvertToUTF8(s.Find(".author").Text())

		var article = Article{id, title, author, summary, category}

		articles = append(articles, article)
	})
	return articles

}

func SavePage(name string, content []byte) {
	f, err := os.Create(savePath + string(os.PathSeparator) + name)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	f.WriteString(string(content))
}

func ConvertToString(src string, srcCode string, tagCode string) string {
	srcCoder := mahonia.NewDecoder(srcCode)
	srcResult := srcCoder.ConvertString(src)
	tagCoder := mahonia.NewDecoder(tagCode)
	_, cdata, _ := tagCoder.Translate([]byte(srcResult), true)
	result := string(cdata)
	return result
}

func ConvertToUTF8(src string) string {
	return ConvertToString(src, "gbk", "utf-8")
}

func MockRequest(url string) (resp *http.Response, err error) {
	client := &http.Client{}
	reqest, err := http.NewRequest("GET", url, nil)

	reqest.Header.Add("Connection", "keep-alive")
	reqest.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.88 Safari/537.36")
	reqest.Header.Add("Accept-Language", "zh-CN,zh;q=0.9")

	if err != nil {
		panic(err)
	}
	return client.Do(reqest)
}
