package qiniu

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"git.zhuzi.me/zzjz/qiniu_demo/app"
)

type BlobIoFlowParams struct {
	Bucket    string `json:"bucket"`
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
}

type Hits struct {
	Hits int `json:"hits"`
	Flow int `json:"flow"`
}

type BlobIoFlowRp struct {
	Time   string `json:"time"`
	Values Hits   `json:"values"`
}

// 统计 flow外网流出流量大小  hits GET 请求次数
// 该接口可以获取外网流出流量统计和 GET 请求次数。监控统计延迟大概 5 分钟。
func GetBlobIoFlow(p *BlobIoFlowParams, selectKey string) (data *BlobIoFlowRp, status int, err error) {
	client := &http.Client{}
	blobIoUrl := fmt.Sprintf("%s%s%s%s%s%s%s%s%s",
		"https://api.qiniu.com/v6/space?bucket=", p.Bucket,
		"&region=z0&begin=", p.StartTime,
		"&end=", p.EndTime,
		"&g=5min&select=", selectKey,
		"&$src=origin")

	request, err := http.NewRequest("POST", blobIoUrl, nil)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	key := &AccessToken{
		AccessKey: app.Config.Qiniu.AccessKey,
		SecretKey: app.Config.Qiniu.SecretKey,
		Url:       blobIoUrl,
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
	data = &BlobIoFlowRp{}
	err = json.Unmarshal(body, data)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	return
}
