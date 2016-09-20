package huaweipush

const (
	accessTokenAPI      = "https://login.vmall.com/oauth2/token"
	baseAPI             = "https://api.vmall.com/rest.php"
	apiMethodPrefix     = "openpush."
	singleSendURL       = "message.single_send"
	batchSendURL        = "message.batch_send"
	lbsSendURL          = "openapi.lbs_send"
	notificationSendURL = "openapi.notification_send"
	setUserTagURL       = "openapi.set_user_tag"
	queryAppTagsSendURL = "openapi.query_app_tags"
	deleteUserTagURL    = "openapi.delete_user_tag"
	queryUserTagSendURL = "openapi.query_user_tag"
	queryMsgResultURL   = "openapi.query_msg_result"
	getTokenByDateURL   = "openapi.get_token_by_date"
)

var (
	SystemErrorCodeMap = map[int64]string{
		0: "成功",
		1: "一个未知的错误发生",
		2: "服务临时不可用",
		3: "未知的方法",
		4: "应用已达到设定的请求上限",
	}

	ServcieErrorCodeMap = map[int64]string{
		20001: "tmID is null",
		20002: "tmID is too many",
		20003: "Message is null",
		20004: "Data too long, please send data less than 1024 bytes",
	}
)
