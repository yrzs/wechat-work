package response

import (
	"github.com/yrzs/wechat-work/kernel/response"
)

type ResponseSubscribeMessageAddTemplate struct {
	*response.ResponseMiniProgram
	PriTmplID string `json:"priTmplId"`
}
