package huaweipush

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"strconv"
	"time"

	"golang.org/x/net/context"
)

type HuaweiPushClient struct {
	clientID, clientSecret string
}

func NewClient(clientID, clientSecret string) *HuaweiPushClient {
	return &HuaweiPushClient{
		clientID:     clientID,
		clientSecret: clientSecret,
	}
}

func (c *HuaweiPushClient) defaultParams(params url.Values) (url.Values, error) {
	accessToken, err := RequestAccess(c.clientID, c.clientSecret)
	if err != nil {
		return params, err
	}
	if accessToken.Error != 0 {
		return params, errors.New(accessToken.ErrorDescription)
	}
	params.Add("nsp_ts", strconv.FormatInt(time.Now().Unix(), 10))
	params.Add("nsp_fmt", "JSON")
	params.Add("access_token", accessToken.AccessToken)
	return params, nil
}

func (c *HuaweiPushClient) SingleSend(ctx context.Context, n *SingleNotification) (*PushResult, error) {
	params := n.Form()
	params, err := c.defaultParams(params)
	if err != nil {
		return nil, err
	}
	params.Add("nsp_svc", singleSendURL)
	bytes, err := doPost(ctx, baseAPI, params)
	if err != nil {
		return nil, err
	}
	var result PushResult
	err = json.Unmarshal(bytes, &result)
	if err != nil {
		return nil, err
	}
	if result.Error == SessionTimeoutError || result.Error == SessionInvalidError {
		fmt.Println("huawei token error", result)
		tokenInstance.AccessToken = ""
		return c.SingleSend(ctx, n)
	}
	return &result, nil
}

func (c *HuaweiPushClient) BatchSend(ctx context.Context, b *BatchNotification) (*PushResult, error) {
	params := b.Form()
	params, err := c.defaultParams(params)
	if err != nil {
		return nil, err
	}
	params.Add("nsp_svc", batchSendURL)
	bytes, err := doPost(ctx, baseAPI, params)
	if err != nil {
		return nil, err
	}
	var result PushResult
	err = json.Unmarshal(bytes, &result)
	if err != nil {
		return nil, err
	}
	if result.Error == SessionTimeoutError || result.Error == SessionInvalidError {
		fmt.Println("huawei token error", result)
		tokenInstance.AccessToken = ""
		return c.BatchSend(ctx, b)
	}
	return &result, nil
}

func (c *HuaweiPushClient) LBSSend(ctx context.Context, n *Notification, location string) (*Result, error) {
	params := n.Form()
	params, err := c.defaultParams(params)
	if err != nil {
		return nil, err
	}
	params.Add("location", location)
	params.Add("nsp_svc", lbsSendURL)
	bytes, err := doPost(ctx, baseAPI, params)
	if err != nil {
		return nil, err
	}
	var result Result
	err = json.Unmarshal(bytes, &result)
	if err != nil {
		return nil, err
	}
	if result.Error == SessionTimeoutError || result.Error == SessionInvalidError {
		fmt.Println("huawei token error", result)
		tokenInstance.AccessToken = ""
		return c.LBSSend(ctx, n, location)
	}
	return &result, nil
}

func (c *HuaweiPushClient) NotificationSend(ctx context.Context, n *Notification) (*NotificationSendResult, error) {
	params := n.Form()
	params, err := c.defaultParams(params)
	if err != nil {
		return nil, err
	}
	params.Add("nsp_svc", notificationSendURL)
	bytes, err := doPost(ctx, baseAPI, params)
	if err != nil {
		return nil, err
	}
	var result NotificationSendResult
	err = json.Unmarshal(bytes, &result)
	if err != nil {
		return nil, err
	}
	if result.Error == SessionTimeoutError || result.Error == SessionInvalidError {
		fmt.Println("huawei token error", result)
		tokenInstance.AccessToken = ""
		return c.NotificationSend(ctx, n)
	}
	return &result, nil
}

func (c *HuaweiPushClient) SetUserTag(ctx context.Context, token, tagKey, tagValue string) (*Result, error) {
	params := url.Values{}
	params, err := c.defaultParams(params)
	if err != nil {
		return nil, err
	}
	params.Add("token", token)
	params.Add("tag_key", tagKey)
	params.Add("tag_value", tagValue)
	params.Add("nsp_svc", setUserTagURL)
	bytes, err := doPost(ctx, baseAPI, params)
	if err != nil {
		return nil, err
	}
	var result Result
	err = json.Unmarshal(bytes, &result)
	if err != nil {
		return nil, err
	}
	if result.Error == SessionTimeoutError || result.Error == SessionInvalidError {
		fmt.Println("huawei token error", result)
		tokenInstance.AccessToken = ""
		return c.SetUserTag(ctx, token, tagKey, tagValue)
	}
	return &result, nil
}

func (c *HuaweiPushClient) QueryAppTags(ctx context.Context) (*TagsResult, error) {
	params := url.Values{}
	params, err := c.defaultParams(params)
	if err != nil {
		return nil, err
	}
	params.Add("nsp_svc", queryAppTagsSendURL)
	bytes, err := doPost(ctx, baseAPI, params)
	if err != nil {
		return nil, err
	}
	var result TagsResult
	err = json.Unmarshal(bytes, &result)
	if err != nil {
		return nil, err
	}
	if result.Error == SessionTimeoutError || result.Error == SessionInvalidError {
		fmt.Println("huawei token error", result)
		tokenInstance.AccessToken = ""
		return c.QueryAppTags(ctx)
	}
	return &result, nil
}

func (c *HuaweiPushClient) DeleteUserTag(ctx context.Context, token, tagKey string) (*Result, error) {
	params := url.Values{}
	params, err := c.defaultParams(params)
	if err != nil {
		return nil, err
	}
	params.Add("token", token)
	params.Add("tag_key", tagKey)
	params.Add("nsp_svc", deleteUserTagURL)
	bytes, err := doPost(ctx, baseAPI, params)
	if err != nil {
		return nil, err
	}
	var result Result
	err = json.Unmarshal(bytes, &result)
	if err != nil {
		return nil, err
	}
	if result.Error == SessionTimeoutError || result.Error == SessionInvalidError {
		fmt.Println("huawei token error", result)
		tokenInstance.AccessToken = ""
		return c.DeleteUserTag(ctx, token, tagKey)
	}
	return &result, nil
}

func (c *HuaweiPushClient) QueryUserTag(ctx context.Context, token string) (*TagsResult, error) {
	params := url.Values{}
	params, err := c.defaultParams(params)
	if err != nil {
		return nil, err
	}
	params.Add("token", token)
	params.Add("nsp_svc", queryUserTagSendURL)
	bytes, err := doPost(ctx, baseAPI, params)
	if err != nil {
		return nil, err
	}
	var result TagsResult
	err = json.Unmarshal(bytes, &result)
	if err != nil {
		return nil, err
	}
	if result.Error == SessionTimeoutError || result.Error == SessionInvalidError {
		fmt.Println("huawei token error", result)
		tokenInstance.AccessToken = ""
		return c.QueryUserTag(ctx, token)
	}
	return &result, nil
}

// 该接口仅能查询single_send和batch_send接口发送的消息
func (c *HuaweiPushClient) QueryMsgResult(ctx context.Context, requestID, token string) (*QueryMsgResult, error) {
	params := url.Values{}
	params, err := c.defaultParams(params)
	if err != nil {
		return nil, err
	}
	params.Add("requestID", requestID)
	if token != "" {
		params.Add("token", token)
	}
	params.Add("nsp_svc", queryMsgResultURL)
	bytes, err := doPost(ctx, baseAPI, params)
	if err != nil {
		return nil, err
	}
	var result QueryMsgResult
	err = json.Unmarshal(bytes, &result)
	if err != nil {
		return nil, err
	}
	if result.Error == SessionTimeoutError || result.Error == SessionInvalidError {
		fmt.Println("huawei token error", result)
		tokenInstance.AccessToken = ""
		return c.QueryMsgResult(ctx, requestID, token)
	}
	return &result, nil
}

func (c *HuaweiPushClient) GetTokenByDate(ctx context.Context, date string) (*GetTokenResult, error) {
	params := url.Values{}
	params, err := c.defaultParams(params)
	if err != nil {
		return nil, err
	}
	params.Add("date", date)
	params.Add("nsp_svc", getTokenByDateURL)
	bytes, err := doPost(ctx, baseAPI, params)
	if err != nil {
		return nil, err
	}
	var result GetTokenResult
	err = json.Unmarshal(bytes, &result)
	if err != nil {
		return nil, err
	}
	if result.Error == SessionTimeoutError || result.Error == SessionInvalidError {
		fmt.Println("huawei token error", result)
		tokenInstance.AccessToken = ""
		return c.GetTokenByDate(ctx, date)
	}
	return &result, nil
}
