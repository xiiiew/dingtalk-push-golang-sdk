package dps

type DingtalkPushRequest struct {
	Secret      string      `json:"secret"`
	AccessToken string      `json:"access_token"`
	Message     interface{} `json:"message"`
}

type DingtalkPushResponse struct {
	Code   int    `json:"code"`
	Result string `json:"result"`
}
