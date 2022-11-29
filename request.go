package dps

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

func HTTPPostJsonWithClient(urlPath string, params interface{}, client *http.Client) (_ []byte, err error) {
	bytesParams, err := json.Marshal(params)
	if err != nil {
		return
	}
	reader := bytes.NewReader(bytesParams)

	req, err := http.NewRequest("POST", urlPath, reader)
	if err != nil {
		return
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	if resp == nil {
		return nil, errors.New("response is empty")
	}

	bytesBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	return bytesBody, nil
}
