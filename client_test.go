package huaweipush

import (
	"fmt"
	"testing"
)

func init() {
	Init("clientID", "clientSecret")
}

func TestHuaweiPushClient_SingleSend(t *testing.T) {
	result, err := defaultClient.SingleSend(NewSingleNotification("deviceToken", "message", "requestID"))
	fmt.Println("result", result)
	fmt.Println("err", err)
}

func TestHuaweiPushClient_NotificationSend(t *testing.T) {
	result, err := defaultClient.NotificationSend(
		NewNotification(1, 1).
			addTokens("tokens1").
			addTokens("tokens2").
			setAndroid(
			NewAndroidMessage("notificationTitle", "notificationContent")))

	if err != nil {
		t.Errorf("err=%v\n", err)
		return
	}
	t.Logf("result=%v\n", result)
}
