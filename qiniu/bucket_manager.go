package qiniu

import (
	"fmt"

	"github.com/qiniu/api.v7/auth/qbox"
	"github.com/qiniu/api.v7/storage"
)

// 获取资源管理 bucketManager 对象
func BucketManager(accessKey, secretKey string) (bucketManager *storage.BucketManager) {
	mac := qbox.NewMac(accessKey, secretKey)
	cfg := storage.Config{
		// 是否使用https域名进行资源管理
		UseHTTPS: false,
	}
	// 指定空间所在的区域，如果不指定将自动探测
	// 如果没有特殊需求，默认不需要指定
	cfg.Zone = &storage.ZoneHuadong

	bucketManager = storage.NewBucketManager(mac, &cfg)
	return
}

// 获取文件信息...
func GetFileInfo(accessKey, secretKey string) (fileInfo storage.FileInfo) {
	bucket := "zhuzi-test"
	key := "sql.jpg"

	fileInfo, sErr := BucketManager(accessKey, secretKey).Stat(bucket, key)
	if sErr != nil {
		fmt.Println(sErr)
		return
	}
	fmt.Println(fileInfo.String())
	//可以解析文件的PutTime
	fmt.Println(storage.ParsePutTime(fileInfo.PutTime))
	return
}
