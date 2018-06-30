package qiniu

import (
	"testing"

	"git.zhuzi.me/zzjz/qiniu_demo/app"
)

var (
	// config 读取你的配置文件信息
	accessKey = app.Config.Qiniu.AccessKey // your accessKey
	secretKey = app.Config.Qiniu.SecretKey // your secretKey
)

var KeyParams = &Key{
	AccessKey: accessKey,
	SecretKey: secretKey,
}

// 服务端表单直传 + 自定义回 JSON
func TestServerUpload(t *testing.T) {
	bucket := app.Config.Qiniu.Bucket // your bucket 非固定参数
	uploadToken, ret := ServerUpload(KeyParams, bucket)
	t.Logf("%+v", uploadToken)
	t.Logf("%+v", ret)
}

// 覆盖上传 Token 凭证
func TestGetOverUploadToken(t *testing.T) {
	keyToOverwrite := "sql.jpg"
	bucket := app.Config.Qiniu.Bucket // your bucket 非固定参数
	token := GetOverUploadToken(KeyParams, bucket, keyToOverwrite)
	t.Logf("%+v", token)
}

// 带回调业务服务器的凭证
func TestGetCallbackUploadToken(t *testing.T) {

	bucket := app.Config.Qiniu.Bucket                     // your bucket 非固定参数
	url := "http://api.example.com/qiniu/upload/callback" // 回调地址
	token := GetCallbackUploadToken(KeyParams, bucket, url)
	t.Logf("%+v", token)
}

// 客户端获取上传 Token 凭证
func TestGetUploadToken(t *testing.T) {
	bucket := app.Config.Qiniu.Bucket // your bucket 非固定参数
	uploadToken := GetUploadToken(KeyParams, bucket)
	t.Logf("%+v", uploadToken)
}
