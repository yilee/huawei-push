package huaweipush

import "time"

type Notification struct {
	tokens      string
	appMethod   string
	nspTS       int64
	nspSVC      string
	nspFmt      string
	accessToken string
	android     string
}

func NewNotification(token, access_token, android string) *Notification {
	return &Notification{
		tokens:      token,
		nspTS:       time.Now().Second(),
		nspSVC:      apiMethodPrefix + notificationSendURL,
		nspFmt:      "JSON",
		accessToken: access_token,
		android:     android,
	}
}

func (n *Notification) Send() {

}
