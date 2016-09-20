package huaweipush

type PushResult struct {
	ResultCode string `json:"resultcode"`
	Message    string `json:"message"`
	RequestID  string `json:"requestID"`
	Error      string `json:"error"`
}

type Result struct {
	ResultCode string `json:"result_code"` // 0：成功
	RequestID  string `json:"request_id"`  // 由服务器生成，方便用户问题追查与定位
	ResultDesc string `json:"result_desc"` // 失败原因
	Error      string `json:"error"`
}

type TagsResult struct {
	RequestID string `json:"request_id"` // 由服务器生成，方便用户问题追查与定位
	Tags      string `json:"tags"`       // 标签列表, 样例：{"tags":[{"location":["ShangHai","GuangZhou"]},{"age":["20","30"]}]}
}

type MsgResult struct {
	RequestID string `json:"request_id"` // 由服务器生成，方便用户问题追查与定位
	Result    []struct {
		Token  string `json:"token"`  // 用户标识
		Status int32  `json:"status"` // 消息状态, 0：成功送达, 1：待发送 （没送到，但又没过期、没被覆盖的消息，还在等待补发的）, 2：被覆盖, 3：过期丢弃
	} `json:"result"`
}

type GetTokenResult struct {
	RequestID     string `json:"request_id"`
	ResultCode    string `json:"result_code"`
	ResultDesc    string `json:"result_desc"`
	TokenFileURL  string `json:"tokenFile_url"`
	UnzipPassword string `json:"unzip_password"`
}
