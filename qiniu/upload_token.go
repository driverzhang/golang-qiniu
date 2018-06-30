package qiniu

import (
	"context"
	"fmt"

	"github.com/qiniu/api.v7/auth/qbox"
	"github.com/qiniu/api.v7/storage"
)

type Key struct {
	AccessKey string `json:"access_key"`
	SecretKey string `json:"secret_key"`
}

// 自定义返回值结构体
type MyPutRet struct {
	Key    string
	Hash   string
	Fsize  int
	Bucket string
}

// 服务端表单直传 + 自定义回 JSON
func ServerUpload(k *Key, bucket string) (token string, ret MyPutRet) {
	key := "sql.jpg"                                                       // 自定义上传文件名称 可以说是时间+string.后缀的形式
	localFile := "/Users/zhangxiaoxin/Documents/1281518437357_.pic_hd.jpg" // 填入你本地图片的绝对地址，你也可以把图片放入项目文件中
	putExtra := storage.PutExtra{}                                         // 可选配置 自定义返回字段
	// 上传文件自定义返回值结构体
	putPolicy := storage.PutPolicy{
		Scope:      bucket,
		ReturnBody: `{"key":"$(key)","hash":"$(etag)","fsize":$(fsize),"bucket":"$(bucket)"}`,
	}
	mac := qbox.NewMac(k.AccessKey, k.SecretKey)
	upToken := putPolicy.UploadToken(mac)
	fmt.Println(upToken)

	cfg := storage.Config{}
	//华东机房
	cfg.Zone = &storage.ZoneHuadong
	// 是否使用https域名
	cfg.UseHTTPS = false
	// 上传是否使用CDN上传加速
	cfg.UseCdnDomains = false

	// 构建表单上传的对象
	formUploader := storage.NewFormUploader(&cfg)
	ret = MyPutRet{} // 你的自定义返回值的结构体

	err := formUploader.PutFile(context.Background(), &ret, upToken, key, localFile, &putExtra)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(ret)

	return
}

// 客户端获取上传 Token 凭证
func GetUploadToken(k *Key, bucket string) (upToken string) {

	// 简单上传
	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}
	mac := qbox.NewMac(k.AccessKey, k.SecretKey)
	upToken = putPolicy.UploadToken(mac)
	fmt.Println(upToken)

	return
}

// 覆盖上传 Token 凭证
func GetOverUploadToken(k *Key, bucket, keyToOverwrite string) (overUpToken string) {
	// 需要覆盖的文件名
	keyToOverwrite = "sql.jpg" // 这个文件名称同时也是客户端之前上传代码中指定的文件名，两者必须一致
	putPolicy := storage.PutPolicy{
		Scope: fmt.Sprintf("%s:%s", bucket, keyToOverwrite),
	}
	mac := qbox.NewMac(k.AccessKey, k.SecretKey)
	overUpToken = putPolicy.UploadToken(mac)
	fmt.Println(overUpToken)

	return
}

// 带回调业务服务器的凭证
// 场景使用： 客服端获取到token后，上传给七牛后，七牛会产生一个POST回调到你设置的回调地址并返回JSON数据
func GetCallbackUploadToken(k *Key, bucket, url string) (upToken string) {
	putPolicy := storage.PutPolicy{
		Scope:            bucket,
		CallbackURL:      url,                                                                       // 填入你的回调地址路径接口
		CallbackBody:     `{"key":"$(key)","hash":"$(etag)","fsize":$(fsize),"bucket":"$(bucket)"}`, // CallbackBody 的设置会在客服端的文件上传到七牛之后，触发七牛回调下内容給业务服务器
		CallbackBodyType: "application/json",
	}
	mac := qbox.NewMac(k.AccessKey, k.SecretKey)
	upToken = putPolicy.UploadToken(mac)
	return
}
