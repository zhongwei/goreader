package netdata

import (
	"log"
	"testing"
)

func TestGet(t *testing.T) {
	log.Print(Get("http://www.baidu.com"))
}
