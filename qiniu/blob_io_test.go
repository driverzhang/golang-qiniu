package qiniu

import "testing"

func TestGetBlobIoFlow(t *testing.T) {
	params := &BlobIoFlowParams{
		Bucket:    "zhuzi-test",
		StartTime: "20170218140000",
		EndTime:   "20180629140000",
	}
	data, status, err := GetBlobIoFlow(params, "flow") // hits
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", data)
	t.Logf("%+v", status)
}
