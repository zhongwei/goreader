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

		desc := gen_desc(url)

		no := gen_no_by_num(seed_num)
		seed_num = seed_num + 1

		os.Mkdir(no, os.ModePerm)

		files_urls := fetch(url)
		for _, file_url := range files_urls {
			err := download_file(file_url, no, gen_filename(file_url))
			if err != nil {
				fmt.Println(no + "|" + file_url)
			}
		}

		dstFile.WriteString(no + "|" + desc + "\n")
	}
}

func gen_filename(url string) (filename string) {
	filename_index := strings.LastIndex(url, "/")
	filename = url[filename_index+1:]
	return
}

func download_file(url, path, filename string) (err error) {
	res, err := http.Get(url)
	if err == nil {
		defer res.Body.Close()
	}

	if err != nil {
		err = errors.New("File downloaded error.")
		return
	}

	file, _ := os.Create(path + "/" + filename)
	io.Copy(file, res.Body)
	contentlength := res.ContentLength
	fileinfo, _ := file.Stat()
	if fileinfo.Size() != contentlength {
		err = errors.New("File was not downloaded completely.")
	}
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

func gen_desc(url string) (desc string) {
	desc_index := strings.Index(url, "-")
	desc_with_slash := url[desc_index:]
	desc_without_slash := desc_with_slash[:len(desc_with_slash)-1]
	desc = strings.Replace(desc_without_slash, "-", " ", -1)
	desc = strings.TrimSpace(desc)
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

func fetch(url string) (urls []string) {

	doc, err := goquery.NewDocument(url)
	if err != nil {
		log.Fatal(err)
	}
	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		current_url, exist := s.Attr("href")
		if exist {
			if strings.Contains(current_url, "images") {
				urls = append(urls, current_url)
			}
		}
	})

	return urls
}
