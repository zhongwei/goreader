package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func merge(rootPath, outFileName string) {
	outFile, openErr := os.OpenFile(outFileName, os.O_CREATE|os.O_WRONLY, 0600)
	if openErr != nil {
		fmt.Printf("Can not open file %s", outFileName)
	}
	bWriter := bufio.NewWriter(outFile)
	filepath.Walk(rootPath, func(path string, info os.FileInfo, err error) error {
		fmt.Println("Processing:", path)
		//这里是文件过滤器，表示我仅仅处理txt文件
		if strings.HasSuffix(path, ".txt") {
			fp, fpOpenErr := os.Open(path)
			if fpOpenErr != nil {
				fmt.Printf("Can not open file %v", fpOpenErr)
				return fpOpenErr
			}
			bReader := bufio.NewReader(fp)
			for {
				buffer := make([]byte, 1024)
				readCount, readErr := bReader.Read(buffer)
				if readErr == io.EOF {
					break
				} else {
					bWriter.Write(buffer[:readCount])
				}
			}
		}
		return err
	})
	bWriter.Flush()
}

func init() {
	flag.Parse()
}

func main() {
	args := flag.Args()
	merge(args[0], args[1])
}
