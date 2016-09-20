package huaweipush

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

func NewAndroidMessage(notificationTitle, notificationContent string) *AndroidMessage {
	return &AndroidMessage{
		notificationTitle:   notificationTitle,
		notificationContent: notificationContent,
	}
}

func (a *AndroidMessage) addExtra(k, v string) *AndroidMessage {
	extra := make(map[string]interface{})
	extra[k] = v
	a.extra = append(a.extra, extra)
	return a
}
