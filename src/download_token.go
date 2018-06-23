package src

import (
	"time"

	"github.com/qiniu/api.v7/auth/qbox"
	"github.com/qiniu/api.v7/storage"
)

// 公有空间 可以直接拼接URL 下载即可
// 也可以调用SDK方法 publicAccessURL := storage.MakePublicURL(domain, key)来直接获取
// 对于 私有空间 下载文件 需要对这个链接进行私有授权签名才能下载。
func GetDownLoadToken(accessKey, secretKey, domain, key string) (privateAccessURL string) {
	mac := qbox.NewMac(accessKey, secretKey)

	deadline := time.Now().Add(time.Second * 3600).Unix() //1小时有效期
	privateAccessURL = storage.MakePrivateURL(mac, domain, key, deadline)
	return
}
