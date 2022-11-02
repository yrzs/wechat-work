package response

import (
	"github.com/yrzs/wechat-work/kernel/response"
)

type ResponseUserIDToOpenID struct {
	*response.ResponseWork

	OpenID string `json:"openid"`
}
