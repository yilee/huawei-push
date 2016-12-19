package huaweipush

import (
	"encoding/json"
	"net/url"
	"strconv"
	"strings"
	"time"
)

type SingleNotification struct {
	DeviceToken string // 32 字节长度，由系统分配的合法TMID
	Message     string // 发送到设备上的消息，最长为4096 字节
	Priority    int32  // 0：高优先级; 1：普通优先级, 缺省值为1
	CacheMode   int32  // 消息是否需要缓存 0：不缓存; 1：缓存, 缺省值为0
	MsgType     int32  // 标识消息类型（缓存机制），由调用端赋值，取值范围（1~100）。当TMID+msgType的值一样时，仅缓存最新的一条消息
	RequestID   string // 如果开发者填写了该字段，则需要保证该字段唯一
	ExpireTime  string // 消息过期删除时间 unix时间戳,格式：2013-08-29 19:55, 如果不填写，默认超时时间为当前时间后48小时
}

type BatchNotification struct {
	DeviceTokenList []string // Device token列表，最多填1000个
	Message         string   // 发送到设备上的消息，最长为4096 字节
	CacheMode       int32    // 消息是否需要缓存 0：不缓存; 1：缓存, 缺省值为0
	MsgType         int32    // 标识消息类型（缓存机制），由调用端赋值，取值范围（1~100）。当TMID+msgType的值一样时，仅缓存最新的一条消息
	ExpireTime      string   // 消息过期删除时间 unix时间戳,格式：2013-08-29 19:55, 如果不填写，默认超时时间为当前时间后48小时
}

func NewSingleNotification(deviceToken, message string) *SingleNotification {
	return &SingleNotification{
		DeviceToken: deviceToken,
		Message:     message,
		Priority:    1,
		CacheMode:   0,
		MsgType:     0,
		RequestID:   "",
		ExpireTime:  "",
	}
}

func (s *SingleNotification) SetMsgType(msgType int32) *SingleNotification {
	s.MsgType = msgType
	return s
}

func (s *SingleNotification) SetRequestID(requestID string) *SingleNotification {
	s.RequestID = requestID
	return s
}

func (s *SingleNotification) SetHighPriority() *SingleNotification {
	s.Priority = 0
	return s
}

func (s *SingleNotification) SetExpireTime(expireTime string) *SingleNotification {
	s.ExpireTime = expireTime
	return s
}

func (s *SingleNotification) SetCacheMode(cacheMode int32) *SingleNotification {
	s.CacheMode = cacheMode
	return s
}

func (s *SingleNotification) Form() url.Values {
	m := url.Values{}
	m.Add("deviceToken", s.DeviceToken)
	m.Add("message", s.Message)
	m.Add("priority", strconv.FormatInt(int64(s.Priority), 10))
	m.Add("cacheMode", strconv.FormatInt(int64(s.CacheMode), 10))
	if s.MsgType > 0 {
		m.Add("msgType", strconv.FormatInt(int64(s.MsgType), 10))
	}
	if s.RequestID != "" {
		m.Add("requestID", s.RequestID)
	}
	if s.ExpireTime != "" {
		m.Add("expireTime", s.ExpireTime)
	}
	return m
}

func NewBatchNotification(deviceTokenList []string, message string) *BatchNotification {
	return &BatchNotification{
		DeviceTokenList: deviceTokenList,
		Message:         message,
		CacheMode:       0,
		MsgType:         0,
		ExpireTime:      "",
	}
}

func (b *BatchNotification) Form() url.Values {
	m := url.Values{}
	m.Add("deviceTokenList", strings.Join(b.DeviceTokenList, ","))
	m.Add("message", b.Message)
	m.Add("cacheMode", strconv.FormatInt(int64(b.CacheMode), 10))
	m.Add("msgType", strconv.FormatInt(int64(b.MsgType), 10))
	m.Add("expireTime", b.ExpireTime)
	return m
}

// 通知栏消息
type Notification struct {
	pushType       int32           `json:"push_type"`        // 推送范围, 1：指定用户，必须指定tokens字段; 2：所有人，无需指定tokens，tags，exclude_tags; 3：一群人，必须指定tags或者exclude_tags字段
	tokens         []string        `json:"tokens"`           // 用户标识, 多个token用","分隔
	tags           []string        `json:"tags"`             // 用户标签，目前仅对android用户生效
	excludeTags    []string        `json:"exclude_tags"`     // 需要剔除的用户的标签，目前仅对android用户生效
	android        *AndroidMessage `json:"android"`          // 给android设备发送消息时，必须填写该字段
	sendTime       string          `json:"send_time"`        // 消息生效时间。如果不携带该字段，则表示消息实时生效。实际使用时，该字段精确到分 消息发送时间戳，timestamp格式ISO 8601：2013-06-03T17:30:08+08:00
	expireTime     string          `json:"expire_time"`      // 消息过期删除时间, 格式同上
	deviceType     int32           `json:"device_type"`      // 目标设备类型, 1：android; 2：ios, 默认为android
	message        *IOSMessage     `json:"message"`          // 消息结构体 发送给非android设备的消息内容
	targetUserType int32           `json:"target_user_type"` // 1：IOS开发用户, 2：IOS生产用户
	allowPeriods   string          `json:"allow_periods"`    // 消息允许展示时间段，时间精确到半小时，24小时制，可以填写一个或者多个时间段, 时间段样例：[[09:30,12:00],[15:00,16:00]]，表示上午9点30到12点之间和下午3点到4点之间可
}

func NewNotification(pushType, deviceType int32) *Notification {
	return &Notification{
		pushType:       pushType,
		tokens:         nil,
		tags:           nil,
		excludeTags:    nil,
		android:        nil,
		sendTime:       "",
		expireTime:     "",
		deviceType:     deviceType,
		message:        nil,
		targetUserType: 0,
		allowPeriods:   "",
	}
}

func (n *Notification) AddTokens(tokens ...string) *Notification {
	n.tokens = append(n.tokens, tokens...)
	return n
}

func (n *Notification) ClearTokens() *Notification {
	n.tokens = nil
	return n
}

func (n *Notification) SetAndroid(android *AndroidMessage) *Notification {
	n.android = android
	return n
}

func (n *Notification) SetTimeToLive(timeToLive int64) *Notification {
	expireTimeStr := time.Now().Add(time.Millisecond * time.Duration(timeToLive)).Format(time.RFC3339)
	n.expireTime = expireTimeStr
	return n
}

func (n *Notification) SetMessage(message *IOSMessage) *Notification {
	n.message = message
	return n
}

func (n *Notification) Form() url.Values {
	m := url.Values{}
	m.Add("push_type", strconv.FormatInt(int64(n.pushType), 10))
	m.Add("tokens", strings.Join(n.tokens, ","))
	if len(n.tags) > 0 {
		m.Add("tags", strings.Join(n.tags, ","))
	}
	if len(n.excludeTags) > 0 {
		m.Add("exclude_tags", strings.Join(n.excludeTags, ","))
	}
	if n.android != nil {
		m.Add("android", n.android.String())
	}
	if n.sendTime != "" {
		m.Add("send_time", n.sendTime)
	}
	if n.expireTime != "" {
		m.Add("expireTime", n.expireTime)
	}
	m.Add("device_type", strconv.FormatInt(int64(n.deviceType), 10))
	if n.message != nil {
		m.Add("message", n.message.String())
	}
	if n.targetUserType != 0 {
		m.Add("target_user_type", strconv.FormatInt(int64(n.targetUserType), 10))
	}
	if n.allowPeriods != "" {
		m.Add("allow_periods", n.allowPeriods)
	}
	return m
}

func (n *Notification) String() string {
	bytes, err := json.Marshal(n)
	if err != nil {
		panic(err)
	}
	return string(bytes)
}
