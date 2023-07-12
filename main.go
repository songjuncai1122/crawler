package main

import (
	"bytes"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/songjuncai1122/crawler/collect"
	"time"
)

func main() {
	url := "https://book.douban.com/subject/1007305/"
	var f collect.Fetcher = collect.BrowserFetch{
		Timeout: 3000 * time.Millisecond,
	}
	body, err := f.Get(url)
	if err != nil {
		fmt.Println("read content failed:%v", err)
		return
	}
	//fmt.Println(string(body))

	// 加载HTML文档
	doc, err := goquery.NewDocumentFromReader((bytes.NewReader(body)))
	if err != nil {
		fmt.Println("read content ffailed:%v", err)
	}

	doc.Find("div.small_cardcontent__BTALp h2").Each(func(i int, s *goquery.Selection) {
		// 获取匹配元素的文本
		title := s.Text()
		fmt.Printf("Review %d: %s\n", i, title)
	})
}
