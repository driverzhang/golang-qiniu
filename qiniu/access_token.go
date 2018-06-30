package qiniu

import (
	"fmt"
	"io"
	"net/http"

	"github.com/qiniu/api.v7/auth/qbox"
)

type AccessToken struct {
	AccessKey string    `json:"access_key"` // 七牛公钥
	SecretKey string    `json:"secret_key"` // 七牛秘钥
	Url       string    `json:"url"`        // 需要调用接口的 URL
	Body      io.Reader `json:"body"`       // 需要调用接口的 body 中的数据 没有就传入 nil
}

func GetAccessToken(a *AccessToken) (access_token string, err error) {
	request, err := http.NewRequest("POST", a.Url, a.Body)
	mac := qbox.NewMac(a.AccessKey, a.SecretKey)
	access_token, err = mac.SignRequest(request)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	return
}

func CreateAuthorization(a *AccessToken) (tokenUrl string) {
	accessToken, err := GetAccessToken(a)
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}
	tokenUrl = fmt.Sprintf("%s %s", "QBox", accessToken)
	return
}
