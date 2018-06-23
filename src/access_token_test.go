package src

import "testing"

func TestGetAccessToken(t *testing.T) {
	data, err := GetAccessToken()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", data)
}
