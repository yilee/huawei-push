package huaweipush

import (
	"encoding/json"
	"fmt"
	"net/url"
	"time"

	"golang.org/x/net/context"
)

type HWToken struct {
	AccessToken      string `json:"access_token"`
	ExpireIn         int64  `json:"expires_in"` // expires_in秒后token过期
	ExpireAt         int64  `json:"expires_at"`
	Scope            string `json:"scope"`
	Error            int32  `json:"error"`
	ErrorDescription string `json:"error_description"`
}

var tokenInstance *HWToken

func init() {
	tokenInstance = &HWToken{
		AccessToken: "",
		ExpireAt:    0,
	}
}

func RequestAccess(clientID, clientSecret string) (*HWToken, error) {
	nowSeconds := time.Now().Unix()
	if tokenInstance.ExpireAt > nowSeconds && tokenInstance.AccessToken != "" {
		return tokenInstance, nil
	}
	form := url.Values{}
	form.Add("client_id", clientID)
	form.Add("client_secret", clientSecret)
	form.Add("grant_type", "client_credentials")
	bytes, err := doPost(context.Background(), accessTokenAPI, form)
	if err != nil {
		return nil, err
	}
	var newToken HWToken
	err = json.Unmarshal(bytes, &newToken)
	if err != nil {
		return nil, err
	}
	newToken.ExpireAt = nowSeconds + newToken.ExpireIn
	tokenInstance = &newToken
	// invalid the token
	time.AfterFunc(time.Second*time.Duration(tokenInstance.ExpireIn), func() {
		tokenInstance.AccessToken = ""
	})
	fmt.Println("huawei new token", tokenInstance)
	return tokenInstance, nil
}
