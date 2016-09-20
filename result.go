package huaweipush

type PushResult struct {
	resultCode string `json:"resultcode"`
	message    string `json:"message"`
	requestID  string `json:"requestID"`
}

type NotificationResult struct {
	resultCode string `json:"result_code"` // 0：成功
	requestID  string `json:"request_id"`  // 由服务器生成，方便用户问题追查与定位
	resultDesc string `json:"result_desc"` // 失败原因
	Error      string `json:"error"`
}
