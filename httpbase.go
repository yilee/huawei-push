package huaweipush

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func doPost(url string, form url.Values) ([]byte, error) {
	var result []byte
	var req *http.Request
	var resp *http.Response
	var err error
	req, err = http.NewRequest("POST", url, strings.NewReader(form.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	client := &http.Client{}
	tryTime := 0
tryAgain:
	resp, err = client.Do(req)
	if err != nil {
		fmt.Println("huawei push post err", err, tryTime)
		tryTime += 1
		if tryTime < 3 {
			goto tryAgain
		}
		return nil, err
	}
	defer func() {
		_ = resp.Body.Close()
	}()
	result, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	str := string(result)
	str, err = strconv.Unquote(str)
	if err != nil {
		str = string(result)
	}
	return []byte(str), nil
}
