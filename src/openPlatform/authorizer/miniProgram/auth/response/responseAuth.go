package response

import "github.com/yrzs/wechat-work/kernel/response"

type ResponseSession struct {
	response.ResponseOpenPlatform

	OpenID     string `json:"openid"`
	SessionKey string `json:"session_key"`
	UnionID    string `json:"unionid"`
}
