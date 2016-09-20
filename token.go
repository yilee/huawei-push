package huaweipush

import (
	"encoding/json"
	"time"
)

type HWToken struct {
	accessToken string `json:"access_token"`
	expire      int64  `json:"expires_in"`
}

var tokenInstance *HWToken

func init() {
	tokenInstance = &HWToken{
		accessToken: "",
		expire:      -1,
	}
}

func requestAccess(clientID, clientSecret string) (string, error) {
	nowSeconds := time.Now().Second()
	if tokenInstance.expire > nowSeconds && tokenInstance.accessToken != "" {
		return tokenInstance.accessToken, nil
	}

	params := make(map[string]string)
	params["client_id"] = clientID
	params["client_secret"] = clientSecret
	params["grant_type"] = "client_credentials"
	bytes, err := doPost(accessTokenAPI, params)
	if err != nil {
		return "", err
	}
	var newToken HWToken
	err = json.Unmarshal(bytes, &newToken)
	if err != nil {
		return "", err
	}
	tokenInstance = &newToken
	return tokenInstance.accessToken, nil
}
