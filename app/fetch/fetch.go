package fetch

import (
	"time"

	"github.com/go-resty/resty/v2"
)

var Client *resty.Client

func init() {
	Client = resty.New().
		SetTimeout(10 * time.Second).     // 设置请求超时时间
		SetRetryCount(3).                 // 设置重试次数
		SetRetryWaitTime(5 * time.Second) // 设置重试间隔
}
