package qiniu

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"git.zhuzi.me/zzjz/qiniu_demo/app"
)

type GetSpaceParams struct {
	Bucket    string `json:"bucket"`
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
}

// 定义返回参数格式
type SpaceRp struct {
	Times []int `json:"times"`
	Datas []int `json:"datas"`
}

// 该接口可以获取标准存储的当前存储量。监控统计可能会延迟 1 天。
func GetSpace(p *GetSpaceParams) (data *SpaceRp, status int, err error) {
	client := &http.Client{}
	spaceUrl := fmt.Sprintf("%s%s%s%s%s%s%s",
		"https://api.qiniu.com/v6/space?bucket=", p.Bucket,
		"&region=z0&begin=", p.StartTime,
		"&end=", p.EndTime,
		"&g=day")

	request, err := http.NewRequest("POST", spaceUrl, nil)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	key := &AccessToken{
		AccessKey: app.Config.Qiniu.AccessKey,
		SecretKey: app.Config.Qiniu.SecretKey,
		Url:       spaceUrl,
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
	data = &SpaceRp{}
	err = json.Unmarshal(body, data)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	return
}
