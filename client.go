package huaweipush

import (
	"encoding/json"
	"errors"
	"net/url"
	"strconv"
	"time"
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

func (c *HuaweiPushClient) SingleSend(n *SingleNotification) (*PushResult, error) {
	params := n.Form()
	params, err := c.defaultParams(params)
	if err != nil {
		return nil, err
	}
	params.Add("nsp_svc", singleSendURL)
	bytes, err := doPost(baseAPI, params)
	if err != nil {
		return nil, err
	}
	var result PushResult
	err = json.Unmarshal(bytes, &result)
	if err != nil {
		return nil, err
	}
	if result.Error != "" {
		return nil, errors.New(result.Error)
	}
	return &result, nil
}

func (c *HuaweiPushClient) BatchSend(b *BatchNotification) (*PushResult, error) {
	params := b.Form()
	params, err := c.defaultParams(params)
	if err != nil {
		return nil, err
	}
	params.Add("nsp_svc", batchSendURL)
	bytes, err := doPost(baseAPI, params)
	if err != nil {
		return nil, err
	}
	var result PushResult
	err = json.Unmarshal(bytes, &result)
	if err != nil {
		return nil, err
	}
	if result.Error != "" {
		return nil, errors.New(result.Error)
	}
	return &result, nil
}

func (c *HuaweiPushClient) LBSSend(n *Notification, location string) (*Result, error) {
	params := n.Form()
	params, err := c.defaultParams(params)
	if err != nil {
		return nil, err
	}
	params.Add("location", location)
	params.Add("nsp_svc", lbsSendURL)
	bytes, err := doPost(baseAPI, params)
	if err != nil {
		return nil, err
	}
	var result Result
	err = json.Unmarshal(bytes, &result)
	if err != nil {
		return nil, err
	}
	if result.Error != "" {
		return nil, errors.New(result.Error)
	}
	return &result, nil
}

func (c *HuaweiPushClient) NotificationSend(n *Notification) (*NotificationSendResult, error) {
	params := n.Form()
	params, err := c.defaultParams(params)
	if err != nil {
		return nil, err
	}
	params.Add("nsp_svc", notificationSendURL)
	bytes, err := doPost(baseAPI, params)
	if err != nil {
		return nil, err
	}
	s2, err := strconv.Unquote(string(bytes))
	if err != nil {
		return nil, err
	}
	var result NotificationSendResult
	err = json.Unmarshal([]byte(s2), &result)
	if err != nil {
		return nil, err
	}
	if result.Error != "" {
		return nil, errors.New(result.Error)
	}
	if result.ResultCode != 0 {
		return nil, errors.New(result.ResultDesc)
	}
	return &result, nil
}

func (c *HuaweiPushClient) SetUserTag(token, tagKey, tagValue string) (*Result, error) {
	params := url.Values{}
	params, err := c.defaultParams(params)
	if err != nil {
		return nil, err
	}
	params.Add("tag_key", tagKey)
	params.Add("tag_value", tagValue)
	params.Add("nsp_svc", setUserTagURL)
	bytes, err := doPost(baseAPI, params)
	if err != nil {
		return nil, err
	}
	s2, err := strconv.Unquote(string(bytes))
	if err != nil {
		return nil, err
	}
	var result Result
	err = json.Unmarshal([]byte(s2), &result)
	if err != nil {
		return nil, err
	}
	if result.Error != "" {
		return nil, errors.New(result.Error)
	}
	if result.ResultCode != "0" {
		return nil, errors.New(result.ResultDesc)
	}
	return &result, nil
}

func (c *HuaweiPushClient) QueryAppTags() (*TagsResult, error) {
	params := url.Values{}
	params, err := c.defaultParams(params)
	if err != nil {
		return nil, err
	}
	params.Add("nsp_svc", queryAppTagsSendURL)
	bytes, err := doPost(baseAPI, params)
	if err != nil {
		return nil, err
	}
	s2, err := strconv.Unquote(string(bytes))
	if err != nil {
		return nil, err
	}
	var result TagsResult
	err = json.Unmarshal([]byte(s2), &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *HuaweiPushClient) DeleteUserTag(token, tagKey string) (*Result, error) {
	params := url.Values{}
	params, err := c.defaultParams(params)
	if err != nil {
		return nil, err
	}
	params.Add("tag_key", tagKey)
	params.Add("nsp_svc", deleteUserTagURL)
	bytes, err := doPost(baseAPI, params)
	if err != nil {
		return nil, err
	}
	s2, err := strconv.Unquote(string(bytes))
	if err != nil {
		return nil, err
	}
	var result Result
	err = json.Unmarshal([]byte(s2), &result)
	if err != nil {
		return nil, err
	}
	if result.Error != "" {
		return nil, errors.New(result.Error)
	}
	if result.ResultCode != "0" {
		return nil, errors.New(result.ResultDesc)
	}
	return &result, nil
}

func (c *HuaweiPushClient) QueryUserTag(token string) (*TagsResult, error) {
	params := url.Values{}
	params, err := c.defaultParams(params)
	if err != nil {
		return nil, err
	}
	params.Add("nsp_svc", queryUserTagSendURL)
	bytes, err := doPost(baseAPI, params)
	if err != nil {
		return nil, err
	}
	s2, err := strconv.Unquote(string(bytes))
	if err != nil {
		return nil, err
	}
	var result TagsResult
	err = json.Unmarshal([]byte(s2), &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// 该接口仅能查询single_send和batch_send接口发送的消息
func (c *HuaweiPushClient) QueryMsgResult(requestID, token string) (*QueryMsgResult, error) {
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
	bytes, err := doPost(baseAPI, params)
	if err != nil {
		return nil, err
	}
	s2, err := strconv.Unquote(string(bytes))
	if err != nil {
		return nil, err
	}
	var result QueryMsgResult
	err = json.Unmarshal([]byte(s2), &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *HuaweiPushClient) GetTokenByDate(date string) (*GetTokenResult, error) {
	params := url.Values{}
	params, err := c.defaultParams(params)
	if err != nil {
		return nil, err
	}
	params.Add("date", date)
	params.Add("nsp_svc", getTokenByDateURL)
	bytes, err := doPost(baseAPI, params)
	if err != nil {
		return nil, err
	}
	var result GetTokenResult
	err = json.Unmarshal(bytes, &result)
	if err != nil {
		return nil, err
	}
	if result.Error != "" {
		return nil, errors.New(result.Error)
	}
	return &result, nil
}
