package huaweipush

import (
	"testing"
	"time"

	"golang.org/x/net/context"
)

func init() {
	Init("clientID", "clientSecret")
}

var defaultClient *HuaweiPushClient

func Init(clientID, clientSecret string) {
	defaultClient = &HuaweiPushClient{
		clientID:     clientID,
		clientSecret: clientSecret,
	}
}

func TestHuaweiPushClient_SingleSend(t *testing.T) {
	result, err := defaultClient.SingleSend(context.TODO(), NewSingleNotification("deviceToken", "hi baby").SetRequestID("requestID"))
	if err != nil {
		t.Errorf("err=%v\n", err)
		return
	}
	result, err = defaultClient.SingleSend(context.TODO(), NewSingleNotification("deviceToken", "hi baby").SetRequestID("requestID"))
	if err != nil {
		t.Errorf("err=%v\n", err)
		return
	}
	t.Logf("result=%#v\n", result)
}

func TestHuaweiPushClient_NotificationSend(t *testing.T) {
	result, err := defaultClient.NotificationSend(
		context.TODO(),
		NewNotification(1, 1).
			AddTokens("tokens1").
			AddTokens("tokens2").
			SetAndroid(
				NewAndroidMessage("hi", "baby")))

	if err != nil {
		t.Errorf("err=%v\n", err)
		return
	}
	t.Logf("result=%#v\n", result)
}

func TestHuaweiPushClient_NotificationSendWithTimeout(t *testing.T) {
	ctx, _ := context.WithTimeout(context.Background(), time.Second)
	time.Sleep(time.Second)
	result, err := defaultClient.NotificationSend(
		ctx,
		NewNotification(1, 1).
			AddTokens("tokens1").
			AddTokens("tokens2").
			SetAndroid(
				NewAndroidMessage("hi", "baby")))

	if err == nil {
		t.Errorf("err=%v\n", err)
		return
	} else {
		t.Logf("result=%v,err=%v\n", result, err)
	}
}

func TestHuaweiPushClient_NotificationSendWithCancel(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	result, err := defaultClient.NotificationSend(
		ctx,
		NewNotification(1, 1).
			AddTokens("tokens1").
			AddTokens("tokens2").
			SetAndroid(
				NewAndroidMessage("hi", "baby")))

	if err == nil {
		t.Errorf("err=%v\n", err)
		return
	} else {
		t.Logf("result=%v,err=%v\n", result, err)
	}
}

func TestHuaweiPushClient_SetUserTag(t *testing.T) {
	result, err := defaultClient.SetUserTag(context.TODO(), "token1", "地区", "上海")
	if err != nil {
		t.Errorf("TestHuaweiPushClient_SetUserTag err=%v\n", err)
		return
	}
	t.Logf("result=%#v\n", result)
}

func TestHuaweiPushClient_QueryAppTags(t *testing.T) {
	result, err := defaultClient.QueryAppTags(context.TODO())
	if err != nil {
		t.Errorf("TestHuaweiPushClient_QueryAppTags err=%v\n", err)
		return
	}
	t.Logf("result=%#v\n", result)
}

func TestHuaweiPushClient_DeleteUserTag(t *testing.T) {
	result, err := defaultClient.DeleteUserTag(context.TODO(), "token1", "地区")
	if err != nil {
		t.Errorf("TestHuaweiPushClient_DeleteUserTag err=%v\n", err)
		return
	}
	t.Logf("result=%#v\n", result)
}

func TestHuaweiPushClient_QueryUserTag(t *testing.T) {
	result, err := defaultClient.QueryUserTag(context.TODO(), "token1")
	if err != nil {
		t.Errorf("TestHuaweiPushClient_QueryUserTag err=%v\n", err)
		return
	}
	t.Logf("result=%#v\n", result)
}

func TestHuaweiPushClient_QueryMsgResult(t *testing.T) {
	result, err := defaultClient.QueryMsgResult(context.TODO(), "1474696910794906661310000000", "token1")
	if err != nil {
		t.Errorf("TestHuaweiPushClient_QueryUserTag err=%v\n", err)
		return
	}
	t.Logf("result=%#v\n", result)
}

func TestHuaweiPushClient_GetTokenByDate(t *testing.T) {
	result, err := defaultClient.GetTokenByDate(context.TODO(), "2016-09-23")
	if err != nil {
		t.Errorf("TestHuaweiPushClient_GetTokenByDate err=%v\n", err)
		return
	}
	t.Logf("result=%#v\n", result)
}
