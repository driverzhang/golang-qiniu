package qiniu

import "testing"

func TestGetSpace(t *testing.T) {
	params := &GetSpaceParams{
		Bucket:    "zhuzi-test",
		StartTime: "20180218140000",
		EndTime:   "20180629140000",
	}
	data, status, err := GetSpace(params)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", data)
	t.Logf("%+v", status)
}
