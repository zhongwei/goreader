package netdata

import (
    "net/http"
    "io/ioutil"
)

func Get(url string) string {
    response, _ := http.Get(url)
    defer response.Body.Close()
    body, _ := ioutil.ReadAll(response.Body)
    bodystr := string(body)
    return bodystr
}
