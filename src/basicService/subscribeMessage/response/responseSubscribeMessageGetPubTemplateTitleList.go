package response

import (
	"github.com/yrzs/wechat-work/kernel/power"
	"github.com/yrzs/wechat-work/kernel/response"
)

type ResponseSubscribeMessageGetPubTemplateTitleList struct {
	*response.ResponseMiniProgram
	Count int              `json:"count"`
	Data  []*power.HashMap `json:"data"`
}
