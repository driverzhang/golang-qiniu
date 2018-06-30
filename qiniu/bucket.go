package qiniu

import (
	"errors"
	"fmt"
	"net/http"

	"git.zhuzi.me/zzjz/qiniu_demo/app"
	"go.zhuzi.me/go/util/encoder"
)

func CreateBucket(bucket string) (status int, err error) {
	client := &http.Client{}
	bucket = encoder.Base64EncodeString(bucket) // base64 加密
	fmt.Println(bucket)

	bucketUrl := fmt.Sprintf("%s%s%s", "https://rs.qiniu.com/mkbucketv2/", bucket, "/region/z0")

	key := &AccessToken{
		AccessKey: app.Config.Qiniu.AccessKey,
		SecretKey: app.Config.Qiniu.SecretKey,
		Url:       bucketUrl,
		Body:      nil,
	}

	tokenUrl := CreateAuthorization(key)

	request, err := http.NewRequest("POST", bucketUrl, nil)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	request.Header.Set("Authorization", tokenUrl)

	response, err := client.Do(request)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	status = response.StatusCode
	if status == 614 {
		return status, errors.New("目标资源已存在") // 请求的 Bucket 已经存在,但没有设置 Region 参数
	}

	if status == 400 {
		return status, errors.New("invalid arguments") // 创建的 Bucket 不符合命名规范
	}

	if status == 401 {
		return status, errors.New("bad token") // 发起创建 Bucket 请求的时候，没有传入认证信息
	}

	if status == 630 {
		return status, errors.New("too many buckets") // 创建 Bucket 的时候超过最大创建数（默认 20 个）
	}

	return
}
