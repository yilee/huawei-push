# huawei-push
华为推送服务 Golang SDK

Production ready, full golang implementation of Huawei Push API (http://developer.huawei.com/push)

```Go
client := &HuaweiPushClient{
		clientID:     "clientID",
		clientSecret: "clientSecret",
	}
client.SingleSend(context.TODO(), NewSingleNotification("deviceToken", "message").SetRequestID("requestID").SetHighPriority())
```