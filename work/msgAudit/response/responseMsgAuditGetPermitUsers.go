package response

import (
	"github.com/yrzs/wechat-work/kernel/response"
)

type ResponseMsgAuditGetPermitUsers struct {
	*response.ResponseWork
	IDs []string `json:"ids"`
}
