package main

import (
	"bytes"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/songjuncai1122/crawler/collect"
	"github.com/songjuncai1122/crawler/proxy"
	"time"
)

func main() {
	proxyURLs := []string{"http://127.0.0.1:7890", "http://127.0.0.1:7890"}
	p, err := proxy.RoundRobinProxySwithcher(proxyURLs...)
	if err != nil {
		fmt.Println("RoundRobin failed")
	}

	url := "https://google.com"
	var f collect.Fetcher = collect.BrowserFetch{
		Timeout: 30000 * time.Millisecond,
		Proxy:   p,
	}
	body, err := f.Get(url)
	if err != nil {
		fmt.Println("read content failed:%v", err)
		return
	}
	fmt.Println(string(body))

	// 加载HTML文档
	doc, err := goquery.NewDocumentFromReader((bytes.NewReader(body)))
	if err != nil {
		fmt.Println("read content ffailed:%v", err)
	}

	doc.Find("div.news_li h2 a[target=_blank]").Each(func(i int, s *goquery.Selection) {
		// 获取匹配元素的文本
		title := s.Text()
		fmt.Printf("Review %d: %s\n", i, title)
	})
}
