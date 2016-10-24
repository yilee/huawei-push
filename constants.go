package huaweipush

const (
	accessTokenAPI      = "https://login.vmall.com/oauth2/token"
	baseAPI             = "https://api.vmall.com/rest.php"
	singleSendURL       = "openpush.message.single_send"
	batchSendURL        = "openpush.message.batch_send"
	lbsSendURL          = "openpush.openapi.lbs_send"
	notificationSendURL = "openpush.openapi.notification_send"
	setUserTagURL       = "openpush.openapi.set_user_tag"
	queryAppTagsSendURL = "openpush.openapi.query_app_tags"
	deleteUserTagURL    = "openpush.openapi.delete_user_tag"
	queryUserTagSendURL = "openpush.openapi.query_user_tag"
	queryMsgResultURL   = "openpush.openapi.query_msg_result"
	getTokenByDateURL   = "openpush.openapi.get_token_by_date"
)

const (
	SessionTimeoutError = "session timeout"
	SessionInvalidError = "invalid session"
)
