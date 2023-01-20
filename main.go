package main

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"text/template"
)

type myForm struct {
	Engines []string `form:"engines[]"`
}

var combinedUrl = "data:text/html, <!DOCTYPE html> <html>  <head>     <meta charset=\"utf-8\">     <title>搜索 %s</title> </head>  <body>     <h1>用多个搜索引擎搜索 <span style=\"color:red\"> \"%s\" </span></h1> </body> <script>  {{range .}}window.open('{{.}}'); {{end}} </script> </html>"

const (
	Google     = "google"
	Bing       = "bing"
	Baidu      = "baidu"
	Duckduckgo = "duckduckgo"
	Yahoo      = "yahoo"
	Sogou      = "sogou"
	SogouWx    = "sogouWx"
	SogouZh    = "sogouZh"
)

var searchEngineUrlMap = map[string]string{
	Google:     "https://www.google.com/search?q=%s",
	Bing:       "https://bing.com/search?q=%s",
	Baidu:      "https://www.baidu.com/s?wd=%s",
	Duckduckgo: "https://duckduckgo.com/?q=%s",
	Yahoo:      "https://search.yahoo.com/search?p=%s",
	Sogou:      "https://www.sogou.com/web?query=%s",
	SogouWx:    "https://weixin.sogou.com/weixin?query=%s&type=2",
	SogouZh:    "https://www.sogou.com/sogou?insite=zhihu.com&query=%s",
}

func main() {
	r := gin.Default()

	r.LoadHTMLGlob("views/*")
	r.GET("/", indexHandler)
	r.POST("/", formHandler)

	r.Run(":8080")
}

func indexHandler(c *gin.Context) {
	c.HTML(200, "form.html", nil)
}

func formHandler(c *gin.Context) {
	var form myForm
	c.Bind(&form)
	Create := func(name, t string) *template.Template {
		return template.Must(template.New(name).Parse(t))
	}
	t := Create("t",
		combinedUrl)
	buf := bytes.NewBufferString("")
	var urls []string
	for _, engine := range form.Engines {
		urls = append(urls, searchEngineUrlMap[engine])
	}
	t.Execute(buf, urls)
	c.HTML(200, "result.html", gin.H{
		"result": buf.String(),
	})
}
