package response

import (
	"github.com/yrzs/wechat-work/kernel/response"
)

type ResponseSecurityMediaCheckASync struct {
	*response.ResponseMiniProgram
	TraceID string `json:"trace_id"`
}
