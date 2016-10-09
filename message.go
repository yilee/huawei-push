package huaweipush

import "encoding/json"

type AndroidMessage struct {
	NotificationTitle      string                   `json:"notification_title"`                 // Notification bar上显示的标题
	NotificationContent    string                   `json:"notification_content"`               // Notification bar上显示的内容
	NotificationStatusIcon string                   `json:"notification_status_icon,omitempty"` // 系统小图标名称, 该图标预置在客户端，在通知栏顶部展示
	Doings                 int32                    `json:"doings"`                             // 1：直接打开应用, 2：通过自定义动作打开应用, 3：打开URL, 4：富媒体消息, 5：短信收件箱广告, 6：彩信收件箱广告
	URL                    string                   `json:"url,omitempty"`                      // 链接 当doings的取值为3时，必须携带该字段
	Intent                 string                   `json:"intent,omitempty"`                   // 自定义打开应用动作 当doings的取值为2时，必须携带该字段
	Extras                 []map[string]interface{} `json:"extras,omitempty"`                   // 用户自定义 dict
}

type IOSMessage struct {
	aps    map[string]interface{} `json:"aps"`
	doings int32                  `json:"doings"`        // 1：直接打开应用, 3：打开URL
	url    string                 `json:"url,omitempty"` // 链接 当doings的取值为3时，必须携带该字段

}

func NewAndroidMessage(notificationTitle, notificationContent string) *AndroidMessage {
	return &AndroidMessage{
		NotificationTitle:   notificationTitle,
		NotificationContent: notificationContent,
		Doings:              1,
		Extras:              make([]map[string]interface{}, 1),
	}
}

func (a *AndroidMessage) SetNotificationStatusIcon(notificationStatusIcon string) *AndroidMessage {
	a.NotificationStatusIcon = notificationStatusIcon
	return a
}

func (a *AndroidMessage) SetDoings(doings int32) *AndroidMessage {
	a.Doings = doings
	return a
}

func (a *AndroidMessage) SetURL(url string) *AndroidMessage {
	a.URL = url
	return a
}

func (a *AndroidMessage) SetIntent(intent string) *AndroidMessage {
	a.Intent = intent
	return a
}

func (a *AndroidMessage) AddExtra(k, v string) *AndroidMessage {
	if a.Extras[0] == nil {
		a.Extras[0] = make(map[string]interface{})
	}
	a.Extras[0][k] = v
	return a
}

func (a *AndroidMessage) String() string {
	bytes, err := json.Marshal(a)
	if err != nil {
		panic(err)
	}
	return string(bytes)
}

func NewIOSMessage(aps map[string]interface{}, doings int32, url string) *IOSMessage {
	return &IOSMessage{
		aps:    aps,
		doings: doings,
		url:    url,
	}
}

func (i *IOSMessage) String() string {
	bytes, err := json.Marshal(i)
	if err != nil {
		panic(err)
	}
	return string(bytes)
}
