package huaweipush

type PushResult struct {
	resultCode string `json:"resultcode"`
	message    string `json:"message"`
	requestID  string `json:"requestID"`
}
