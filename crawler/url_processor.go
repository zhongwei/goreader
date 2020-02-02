package crawler

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
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

func GetDataFromJson(fileName, start, end string) (ids []string, err error) {
	var (
		content []byte
	)

	if content, err = ioutil.ReadFile(fileName); err != nil {
		fmt.Println(err)
		return
	}

	if err = json.Unmarshal(content, &ids); err != nil {
		fmt.Println(err)
		return
	}
	min, _ := strconv.Atoi(start)
	max, _ := strconv.Atoi(end)
	return ids[min-1 : max-1], err
}

func GenURLsFromData(ids []string, prefix, suffix string) []string {
	var (
		urls []string
	)

	for _, id := range ids {
		urls = append(urls, prefix+id+suffix)
	}

	return urls
}

func DownloadFiles(fileName, start, end, prefix, suffix string) {
	ids, _ := GetDataFromJson(fileName, start, end)
	for index, id := range ids {
		savedFileName := id + suffix
		url := prefix + savedFileName
		res, err := MockRequest(url)
		if err != nil {
			log.Fatal(err)
		}

		defer res.Body.Close()
		body, _ := ioutil.ReadAll(res.Body)
		ioutil.WriteFile(savedFileName, body, 0644)
		if index%5 == 0 {
			fmt.Printf("file: %v, fileName: %s\n", index, savedFileName)
		}
	}
}

func GenDetailFileURLS(fileName, start, end, prefix, suffix string) {
	ids, _ := GetDataFromJson(fileName, start, end)
	urls := GenURLsFromData(ids, prefix, suffix)
	txtURLs := []string{}
	for _, url := range urls {
		fmt.Printf("source: %s\n", url)
		res, err := MockRequest(url)
		if err != nil {
			log.Fatal(err)
		}

		defer res.Body.Close()

		if res.StatusCode != 200 {
			log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
		}

		txtURL := ProcessDetailPage(res.Body)
		fmt.Printf("url: %s\n", txtURL)

		txtURLs = append(txtURLs, txtURL)
	}
}

func ProcessDetailPage(r io.Reader) string {
	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("linkName: %v\n", doc)
	var url string
	doc.Find(".album_read").Each(func(i int, s *goquery.Selection) {
		fmt.Println("lllllllllllllllll")
		aNode := s.Find("a")
		linkName := ConvertToUTF8(aNode.Text())
		fmt.Printf("linkName: %s\n", linkName)
		if linkName == "TXT下载" {
			url, _ = aNode.Attr("href")
		}

	})

	return ConvertToUTF8(url)
}
