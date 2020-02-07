package web

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"golang.org/x/text/encoding/simplifiedchinese"
)

func Start(rootPath, port string) {
	r := gin.Default()
	r.Static("/s", rootPath)
	r.GET("/t/:fileName", func(c *gin.Context) {
		fileName := c.Param("fileName")
		fullPath := rootPath + "/t/" + fileName
		fmt.Println(fullPath)
		b, _ := ioutil.ReadFile(fullPath)
		utf8Str, _ := ConvertGBK2Str(string(b))
		s := strings.Replace(utf8Str, "\r\n\r\n", "\r\n", -1)
		s = strings.Replace(s, "\r\n\r\n", "\r\n", -1)
		c.String(http.StatusOK, s)
	})
	r.Run(port) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func ConvertGBK2Str(gbkStr string) (string, error) {
	return simplifiedchinese.GBK.NewDecoder().String(gbkStr)
}
