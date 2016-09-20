package huaweipush

import (
	"fmt"
	"testing"
)

func init() {
	Init("your appSecret")
}

func TestHuaweiPushClient_SingleSend(t *testing.T) {
	result, err := defaultClient.SingleSend(NewSingleNotification("deviceToken", "message", "requestID"))
	fmt.Println("result", result)
	fmt.Println("err", err)
}
