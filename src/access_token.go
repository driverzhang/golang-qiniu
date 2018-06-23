package src

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"git.zhuzi.me/zzjz/qiniu_demo/app"
)

// 返回数据结构体 规定字段
type AccessTokenRp struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    int64  `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
}

func GetAccessToken() (accessToken *AccessTokenRp, err error) {
	res, err := http.PostForm("https://acc.qbox.me/oauth2/token", // 官方接口地址
		url.Values{
			"grant_type": {"password"},              // 固定字段值 不可变
			"username":   {app.Config.Qiniu.Number}, // 填入你的七牛云 账号
			"password":   {app.Config.Qiniu.Pwd},    // 填入你的七牛云 密码
		})

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	accessToken = &AccessTokenRp{}
	err = json.Unmarshal(body, accessToken)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	return
}
