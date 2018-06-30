package qiniu

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"git.zhuzi.me/zzjz/qiniu_demo/app"
)

type RsPutParams struct {
	Bucket    string `json:"bucket"`
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
}

type RsPutRp struct {
	Time   string `json:"time"`
	Values Hits   `json:"values"`
}

// 统计 获取 PUT 请求次数
// 该接口可以获取 PUT 请求次数。监控统计延迟大概 5 分钟。
func GetRsPut(p *RsPutParams) (data *RsPutRp, status int, err error) {
	client := &http.Client{}
	RsPutUrl := fmt.Sprintf("%s%s%s%s%s%s%s",
		"https://api.qiniu.com/v6/rs_put?bucket=", p.Bucket,
		"&region=z0&begin=", p.StartTime,
		"&end=", p.EndTime,
		"&g=5min&select=hits&$src=origin")

	request, err := http.NewRequest("POST", RsPutUrl, nil)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	key := &AccessToken{
		AccessKey: app.Config.Qiniu.AccessKey,
		SecretKey: app.Config.Qiniu.SecretKey,
		Url:       RsPutUrl,
		Body:      nil,
	}

	tokenUrl := CreateAuthorization(key)
	request.Header.Set("Authorization", tokenUrl)
	response, err := client.Do(request)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer response.Body.Close()
	status = response.StatusCode
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	data = &RsPutRp{}
	err = json.Unmarshal(body, data)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	return
}
