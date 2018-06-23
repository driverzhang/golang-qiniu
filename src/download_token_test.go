package src

import (
	"testing"
)

func TestGetDownLoadToken(t *testing.T) {
	domain := "https://image.example.com" // 你设置的外域 链接地址
	key := "sql.jpg"                      // 你要下载的文件名称
	url := GetDownLoadToken(accessKey, secretKey, domain, key)
	t.Logf("%+v", url)
}
