package app

import (
	"go.zhuzi.me/go/config"
	"go.zhuzi.me/go/log"
)

// 使用者根据自己需要修改这个结构体
var Config struct {
	// LogDebug开启后会使用颜色
	LogDebug bool   `yaml:"log_debug"`
	HttpAddr string `yaml:"http_addr"`
	Qiniu    struct {
		AccessKey string `yaml:"access_key"`
		SecretKey string `yaml:"secret_key"`
		Bucket    string `yaml:"bucket" `
		Number    string `yaml:"number"`
		Pwd       string `yaml:"pwd"`
	} `yaml:"qiniu"`

	QiniuURL struct {
		AccessToken string `yaml:"access_token"`
	}
}

// 初始化
// - config
func init() {
	log.SetDebug(Config.LogDebug)
	config.Init(&Config)
}
