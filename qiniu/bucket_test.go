package qiniu

import (
	"testing"
)

func TestCreateBucket(t *testing.T) {
	data, err := CreateBucket("sp-test1")
	if err != nil {
		t.Log(data)
		t.Fatal(err)
	}
	t.Logf("%+v", data)
}
