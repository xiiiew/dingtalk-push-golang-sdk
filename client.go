package dps

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
)

type DpsClient struct {
	Endpoint    string
	Secret      string
	AccessToken string

	Client *http.Client
}

// create a DpsClient instance with default http client.
func NewDpsClient(endpoint, secret, accessToken string) (_ *DpsClient, err error) {
	return NewDpsClientWithHTTPClient(endpoint, secret, accessToken, &http.Client{})
}

// create a DpsClient instance with custom http client.
func NewDpsClientWithHTTPClient(endpoint, secret, accessToken string, client *http.Client) (_ *DpsClient, err error) {
	if len(secret) != 67 || !strings.HasPrefix(secret, "SEC") {
		err = errors.New("secret error")
		return
	}
	if len(accessToken) != 64 {
		err = errors.New("secret error")
		return
	}
	dpsClient := &DpsClient{
		Endpoint:    endpoint,
		Secret:      secret,
		AccessToken: accessToken,
		Client:      client,
	}
	return dpsClient, nil
}

// send message
func (self *DpsClient) Send(message IMessage) error {
	message.SetMessageType()
	dpRequest := DingtalkPushRequest{
		Secret:      self.Secret,
		AccessToken: self.AccessToken,
		Message:     message,
	}
	u := fmt.Sprintf("%s/dingtalk/send", self.Endpoint)
	result, err := HTTPPostJsonWithClient(u, dpRequest, self.Client)
	if err != nil {
		return err
	}
	dpResponse := DingtalkPushResponse{}
	err = json.Unmarshal(result, &dpResponse)
	if err != nil {
		return err
	}
	if dpResponse.Code != 200 {
		return errors.New(dpResponse.Result)
	}
	return nil
}
