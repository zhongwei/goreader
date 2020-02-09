package proxy

import (
	"io/ioutil"
	"log"
	"net/http"
)

type ProxyServer struct {
}

func (ProxyServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			w.WriteHeader(500)
			log.Println(err)
		}
	}()

	if r.URL.Path == "/a" {
		newReq, _ := http.NewRequest(r.Method, "http://localhost:9091", r.Body)
		newResponse, _ := http.DefaultClient.Do(newReq)
		defer newResponse.Body.Close()
		resContent, _ := ioutil.ReadAll(newResponse.Body)
		w.Write(resContent)
		return
	}
	w.Write([]byte("default index"))
}

func Start() {
	http.ListenAndServe(":8080", ProxyServer{})
}
