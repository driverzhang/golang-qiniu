package qiniu

import "testing"

func TestGetRsPut(t *testing.T) {
	params := &RsPutParams{
		Bucket:    "zhuzi-test",
		StartTime: "20170218140000",
		EndTime:   "20180629140000",
	}
	data, status, err := GetRsPut(params)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", data)
	t.Logf("%+v", status)
}
