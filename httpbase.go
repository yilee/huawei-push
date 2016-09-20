package huaweipush

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func doPost(url string, params map[string]interface{}) ([]byte, error) {
	var result []byte
	var req *http.Request
	var resp *http.Response
	var err error
	data, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}
	req, err = http.NewRequest("POST", url, bytes.NewBuffer(data))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	client := &http.Client{}
	resp, err = client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = resp.Body.Close()
	}()
	result, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return result, nil
}
