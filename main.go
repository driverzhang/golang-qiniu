package main

import (
	"fmt"

	"git.zhuzi.me/zzjz/qiniu_demo/app"
	"git.zhuzi.me/zzjz/qiniu_demo/src"
)

var (
	// config 读取你的配置文件信息
	accessKey = app.Config.Qiniu.AccessKey
	secretKey = app.Config.Qiniu.SecretKey
	bucket    = app.Config.Qiniu.Bucket // 空间名不一定每次都固定，可以让前端 post请求传入参数获取
)

var keyParams = &src.Key{
	AccessKey: accessKey,
	SecretKey: secretKey,
}

func main() {
	uploadToken := src.GetUploadToken(keyParams, bucket)
	fmt.Println(uploadToken)
}
