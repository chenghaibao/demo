package main

import (
	"go.uber.org/zap"
	"time"
)

// 格式话输入日志
var Logger *zap.Logger

func main() {
	Logger = zap.NewExample()
	defer Logger.Sync()

	url := "http://example.org/api"
	Logger.Info("failed to fetch URL",
		zap.String("url", url),
		zap.Int("attempt", 3),
		zap.Duration("backoff", time.Second),
	)

	sugar := Logger.Sugar()
	sugar.Infow("failed to fetch URL",
		"url", url,
		"attempt", 6,
		"backoff", time.Second,
	)
	sugar.Infof("Failed to fetch URL: %s", url)
}