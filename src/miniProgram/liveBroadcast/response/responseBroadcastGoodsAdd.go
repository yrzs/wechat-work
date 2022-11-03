package response

import "github.com/yrzs/wechat-work/kernel/response"

type ResponseBroadcastGoodsAdd struct {
	*response.ResponseMiniProgram

	GoodsID string `json:"goodsId"`
	AuditID int64  `json:"auditId"`
}
