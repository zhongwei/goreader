package processor

import (
	"fmt"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

func CollectLinks(url string) {
	resp, _ := http.Get(url)
	body := resp.Body
	defer body.Close()

	tokenizer := html.NewTokenizer(body)

	for {
		tokenType := tokenizer.Next()

		switch {
		case tokenType == html.ErrorToken:
			return
		case tokenType == html.StartTagToken:
			token := tokenizer.Token()

			isAnchor := token.Data == "a"
			if isAnchor {
				_, currentURL := getHref(token)

				hasProto := strings.Index(currentURL, "http") == 0
				if hasProto {
					fmt.Println(currentURL)
				}
			}
		}
	}
}

func getHref(t html.Token) (ok bool, href string) {
	for _, a := range t.Attr {
		if a.Key == "href" {
			href = a.Val
			ok = true
		}
	}

	return
}
