package main

import (
	"bufio"
	"flag"
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

	urls := gen_urls(args[0])

	dstFile, _ := os.Create(args[1])
	defer dstFile.Close()

	for index, url := range urls {

		split_index := strings.Index(url, "|")

		current_url := url[split_index+1:]
		title := url[:split_index]

		content := fetch(current_url)

		no := gen_no_by_num(index)

		currentFile, _ := os.Create(args[2] + "/" + no + ".txt")
		defer currentFile.Close()
		currentFile.WriteString(content)

		dstFile.WriteString(no + "|" + title + "|" + current_url + "\n")
	}

}

func gen_no_by_num(i int) (no string) {
	no = strconv.Itoa(i)
	if i < 10 {
		no = "000" + no
	} else if i < 100 {
		no = "00" + no
	} else if i < 1000 {
		no = "0" + no
	}
	return
}

func gen_urls(path string) (urls []string) {
	fp, _ := os.Open(path)
	buf := bufio.NewReader(fp)
	for {
		line, err := buf.ReadString('\n')
		line = strings.TrimSpace(line)
		urls = append(urls, line)
		if err != nil {
			break
		}
	}

	return
}

func fetch(url string) (content string) {
	doc, err := goquery.NewDocument(url)
	if err != nil {
		log.Fatal(err)
	}

	content = doc.Find("td[id]").Text()

	return
}
