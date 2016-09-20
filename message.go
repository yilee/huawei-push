package huaweipush

type Message struct {
	title       string                 `json:"title"`        // Notification bar上显示的标题
	content     string                 `json:"content"`      // Notification bar上显示的内容
	statusIcon  string                 `json:"status_icon"`  // 系统小图标名称, 该图标预置在客户端，在通知栏顶部展示
	pushType    int32                  `json:"push_type"`    // 推送范围, 1：指定用户，必须指定tokens字段; 2：所有人，无需指定tokens，tags，exclude_tags; 3：一群人，必须指定tags或者exclude_tags字段
	tags        map[string]interface{} `json:"tags"`         // 用户标签，目前仅对android用户生效
	excludeTags map[string]interface{} `json:"exclude_tags"` // 需要剔除的用户的标签，目前仅对android用户生效
	sendTime    string                 `json:"send_time"`    // 消息生效时间。如果不携带该字段，则表示消息实时生效。实际使用时，该字段精确到分, timestamp格式ISO 8601：2013-06-03T17:30:08+08:00
	expireTime  string                 `json:"expire_time"`  // 消息过期删除时间, 格式同上
	deviceType  int32                  `json:"device_type"`  // 目标设备类型, 1：android; 2：ios
	doings      int32                  `json:"doings"`       // 1：直接打开应用, 2：通过自定义动作打开应用, 3：打开URL, 4：富媒体消息, 5：短信收件箱广告, 6：彩信收件箱广告
	extra       map[string]string      `json:"extra"`        // 用户自定义 dict
}
