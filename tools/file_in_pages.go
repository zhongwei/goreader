package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
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
	seed_num, _ := strconv.Atoi(args[1])

	dstFile, _ := os.Create(args[2])
	defer dstFile.Close()
	for _, url := range urls {

		no := gen_no_by_num(seed_num)
		seed_num = seed_num + 1

		os.Mkdir(no, os.ModePerm)

		title, name, desc, files_urls := fetch(url)
		for _, file_url := range files_urls {
			err := download_file(file_url, no)
			if err != nil {
				fmt.Println(no + "|" + file_url)
			}
		}

		dstFile.WriteString(no + "|" + title + "|" + name + "|" + desc + "\n")
	}
}

func gen_filename(url string) (filename string) {
	filename_index := strings.LastIndex(url, "/")
	filename = url[filename_index+1:]
	return
}

func download_file(url, path string) (err error) {
	img_url := img_url_from_page_url(url)
	res, err := http.Get(img_url)
	if err == nil {
		defer res.Body.Close()
	}
	file, _ := os.Create(path + "/" + gen_filename(img_url))
	io.Copy(file, res.Body)
	contentlength := res.ContentLength
	fileinfo, _ := file.Stat()
	if fileinfo.Size() != contentlength {
		err = errors.New("File was not downloaded completely.")
	}
	return
}

func img_url_from_page_url(url string) (img_url string) {
	doc, err := goquery.NewDocument(url)
	if err != nil {
		log.Fatal(err)
	}
	doc.Find("#picture").Each(func(i int, s *goquery.Selection) {
		img_url, _ = s.Attr("src")
	})
	return
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

func fetch(url string) (title, name, desc string, urls []string) {
	domain_slash := strings.LastIndex(url, "galleries")
	doc, err := goquery.NewDocument(url)
	if err != nil {
		log.Fatal(err)
	}
	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		current_url, exist := s.Attr("href")
		if exist {
			if strings.Contains(current_url, "galleries") {
				urls = append(urls, url[:domain_slash-1]+current_url)
			}
		}
	})

	doc.Find("h1").Each(func(i int, s *goquery.Selection) {
		title = s.Text()
	})

	doc.Find("h2 a").Each(func(i int, s *goquery.Selection) {
		name = s.Text()
	})

	doc.Find("#description").Each(func(i int, s *goquery.Selection) {
		desc = s.Text()
	})

	return
}
