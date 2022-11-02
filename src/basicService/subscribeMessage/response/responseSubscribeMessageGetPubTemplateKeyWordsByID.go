package response

import (
	"github.com/yrzs/wechat-work/kernel/power"
	"github.com/yrzs/wechat-work/kernel/response"
)

type ResponseSubscribeMessageGetPubTemplateKeyWordsByID struct {
	*response.ResponseMiniProgram
	Data []*power.HashMap `json:"data"`
}
