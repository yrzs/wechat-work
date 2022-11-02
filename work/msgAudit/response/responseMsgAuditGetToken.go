package response

import (
	"github.com/yrzs/wechat-work/kernel/response"
)

type ResponseCorpGroupGetToken struct {
	*response.ResponseWork
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}
