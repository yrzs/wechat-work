package response

import (
	"github.com/yrzs/wechat-work/kernel/response"
)

type ResponseBroadcastGoodsAudit struct {
	*response.ResponseMiniProgram

	AuditID int `json:"auditId"`
}
