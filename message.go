package huaweipush

import "encoding/json"

type AndroidMessage struct {
	notificationTitle      string                   `json:"notification_title"`       // Notification bar上显示的标题
	notificationContent    string                   `json:"notification_content"`     // Notification bar上显示的内容
	notificationStatusIcon string                   `json:"notification_status_icon"` // 系统小图标名称, 该图标预置在客户端，在通知栏顶部展示
	doings                 int32                    `json:"doings"`                   // 1：直接打开应用, 2：通过自定义动作打开应用, 3：打开URL, 4：富媒体消息, 5：短信收件箱广告, 6：彩信收件箱广告
	url                    string                   `json:"doings"`                   // 链接 当doings的取值为3时，必须携带该字段
	intent                 string                   `json:"intent"`                   // 自定义打开应用动作 当doings的取值为2时，必须携带该字段
	pushType               int32                    `json:"push_type"`                // 推送范围, 1：指定用户，必须指定tokens字段; 2：所有人，无需指定tokens，tags，exclude_tags; 3：一群人，必须指定tags或者exclude_tags字段
	extra                  []map[string]interface{} `json:"extra"`                    // 用户自定义 dict
}

type IOSMessage struct {
}

type Notification struct {
	pushType       int32           `json:"push_type"`        // 推送范围, 1：指定用户，必须指定tokens字段; 2：所有人，无需指定tokens，tags，exclude_tags; 3：一群人，必须指定tags或者exclude_tags字段
	tokens         string          `json:"tokens"`           // 用户标识, 多个token用","分隔
	tags           string          `json:"tags"`             // 用户标签，目前仅对android用户生效
	exclude_tags   string          `json:"exclude_tags"`     // 需要剔除的用户的标签，目前仅对android用户生效
	android        *AndroidMessage `json:"android"`          // 给android设备发送消息时，必须填写该字段
	sendTime       string          `json:"send_time"`        // 消息生效时间。如果不携带该字段，则表示消息实时生效。实际使用时，该字段精确到分 消息发送时间戳，timestamp格式ISO 8601：2013-06-03T17:30:08+08:00
	expireTime     string          `json:"expire_time"`      // 消息过期删除时间, 格式同上
	deviceType     int32           `json:"device_type"`      // 目标设备类型, 1：android; 2：ios, 默认为android
	message        string          `json:"message"`          // 消息结构体 发送给非android设备的消息内容
	targetUserType int32           `json:"target_user_type"` // 1：IOS开发用户, 2：IOS生产用户
	allowPeriods   string          `json:"allow_periods"`    // 消息允许展示时间段，时间精确到半小时，24小时制，可以填写一个或者多个时间段, 时间段样例：[[09:30,12:00],[15:00,16:00]]，表示上午9点30到12点之间和下午3点到4点之间可
}

func NewAndroidMessage(notificationTitle, notificationContent string) *AndroidMessage {
	return &AndroidMessage{
		notificationTitle:   notificationTitle,
		notificationContent: notificationContent,
	}
}

func (a *AndroidMessage) AddExtra(k, v string) *AndroidMessage {
	extra := make(map[string]interface{})
	extra[k] = v
	a.extra = append(a.extra, extra)
	return a
}

func (a *AndroidMessage) String() string {
	bytes, err := json.Marshal(a)
	if err != nil {
		panic(err)
	}
	return string(bytes)
}

func NewNotification(pushType, deviceType int32) *Notification {
	return &Notification{
		pushType:       pushType,
		tokens:         "",
		tags:           "",
		exclude_tags:   "",
		android:        nil,
		sendTime:       "",
		expireTime:     "",
		deviceType:     deviceType,
		message:        "",
		targetUserType: 0,
		allowPeriods:   "",
	}
}

func (n *Notification) setTokens(tokens string) *Notification {
	n.tokens = tokens
	return n
}

func (n *Notification) setAndroid(android *AndroidMessage) *Notification {
	n.android = android
	return n
}
