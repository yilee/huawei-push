package huaweipush

import (
	"encoding/json"
	"fmt"
)

type HuaweiPushClient struct {
	appSecret string
}

var defaultClient *HuaweiPushClient

func Init(appSecret string) {
	defaultClient = &HuaweiPushClient{
		appSecret: appSecret,
	}
}

func (c *HuaweiPushClient) SingleSend(n *SingleNotification) (*PushResult, error) {
	params := n.Map()
	fmt.Printf("params=%#v\n", params)
	bytes, err := doPost(baseAPI+apiMethodPrefix+singleSendURL, params)
	if err != nil {
		fmt.Errorf("post err:%v\n", err)
		return nil, err
	}
	var result PushResult
	err = json.Unmarshal(bytes, &result)
	if err != nil {
		fmt.Errorf("unmarshal err:%v\n", err)
		return nil, err
	}
	return &result, nil
}
