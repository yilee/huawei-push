# huawei-push
华为推送服务 Golang SDK

full golang implementation of Huawei Push API (http://developer.huawei.com/push)

```Go
defaultClient := &HuaweiPushClient{
		clientID:     "clientID",
		clientSecret: "clientSecret",
	}

result, err := defaultClient.SingleSend(NewSingleNotification("deviceToken", "message").SetRequestID("requestID").SetHighPriority())
if err != nil {
    t.Errorf("err=%v\n", err)
    return
}
t.Logf("result=%v\n", result)
```