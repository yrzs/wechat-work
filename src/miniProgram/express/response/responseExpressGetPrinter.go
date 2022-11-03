package response

import "github.com/yrzs/wechat-work/kernel/response"

type ResponseExpressGetPrinter struct {
	*response.ResponseMiniProgram

	Count     string   `json:"count"`
	OpenID    []string `json:"openid"`
	TagIDList []string `json:"tagid_list"`
}
