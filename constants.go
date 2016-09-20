package huaweipush

const (
	accessTokenAPI      = "https://login.vmall.com/oauth2/token"
	baseAPI             = "https://api.vmall.com/rest.php"
	apiMethodPrefix     = "openpush.openapi."
	lbsSendURL          = "lbs_send"
	notificationSendURL = "notification_send"
	setUserTagURL       = "set_user_tag"
	queryAppTagsSendURL = "query_app_tags"
	deleteUserTagURL    = "delete_user_tag"
	queryUserTagSendURL = "query_user_tag"
	queryMsgResultURL   = "query_msg_result"
	getTokenByDateURL   = "get_token_by_date"
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
