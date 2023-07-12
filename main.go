package main

import (
	"github.com/songjuncai1122/crawler/collect"
	"github.com/songjuncai1122/crawler/log"
	"github.com/songjuncai1122/crawler/proxy"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"time"
)

func main() {
	plugin, c := log.NewFilePlugin("./log.txt", zapcore.InfoLevel)
	defer c.Close()
	logger := log.NewLogger(plugin)
	logger.Info("log init end")

	proxyURLs := []string{"http://127.0.0.1:7890", "http://127.0.0.1:7890"}
	p, err := proxy.RoundRobinProxySwithcher(proxyURLs...)
	if err != nil {
		logger.Error("Roundboinproxy failed")
	}

	url := "https://google.com"
	var f collect.Fetcher = collect.BrowserFetch{
		Timeout: 3000 * time.Millisecond,
		Proxy:   p,
	}
	body, err := f.Get(url)
	if err != nil {
		logger.Error("read content filed",
			zap.Error(err),
		)
		return
	}
	logger.Info("get content", zap.Int("len", len(body)))
}
