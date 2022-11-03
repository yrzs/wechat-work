package response

import (
	"github.com/yrzs/wechat-work/kernel/power"
	"github.com/yrzs/wechat-work/kernel/response"
)

type ResponseOperationGetJsErrList struct {
	*response.ResponseMiniProgram
	Success    bool             `json:"success"`
	OpenID     string           `json:"openid"`
	Data       []*power.HashMap `json:"data"`
	TotalCount int64            `json:"totalCount"`
}
