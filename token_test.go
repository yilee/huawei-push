package huaweipush

import "testing"

func TestRequestAccess(t *testing.T) {
	token, err := RequestAccess("clientID", "clientSecret")
	if err != nil {
		t.Errorf("err=%v\n", err)
		return
	}
	if token.Error != 0 {
		t.Errorf("token=%v\n", token)
		return
	}
	t.Logf("token=%#v\n", token)
}
