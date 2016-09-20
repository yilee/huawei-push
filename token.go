package huaweipush

import (
	"encoding/json"
	"net/url"
	"time"
)

type HWToken struct {
	AccessToken      string `json:"access_token"`
	Expire           int64  `json:"expires_in"`
	Scope            string `json:"scope"`
	Error            int32  `json:"error"`
	ErrorDescription string `json:"error_description"`
}

var tokenInstance *HWToken

func init() {
	tokenInstance = &HWToken{
		AccessToken: "",
		Expire:      -1,
	}
}

func RequestAccess(clientID, clientSecret string) (*HWToken, error) {
	nowSeconds := time.Now().Unix()
	if tokenInstance.Expire > nowSeconds && tokenInstance.AccessToken != "" {
		return tokenInstance, nil
	}
	form := url.Values{}
	form.Add("client_id", clientID)
	form.Add("client_secret", clientSecret)
	form.Add("grant_type", "client_credentials")
	bytes, err := doPost(accessTokenAPI, form)
	if err != nil {
		return nil, err
	}
	var newToken HWToken
	err = json.Unmarshal(bytes, &newToken)
	if err != nil {
		return nil, err
	}
	tokenInstance = &newToken
	return tokenInstance, nil
}
