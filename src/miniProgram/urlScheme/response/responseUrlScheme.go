package response

import "github.com/yrzs/wechat-work/kernel/response"

type ResponseUrlScheme struct {
	*response.ResponseMiniProgram

	OpenLink string `json:"openlink"`
}
