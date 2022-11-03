package response

import (
	"github.com/yrzs/wechat-work/kernel/power"
	response2 "github.com/yrzs/wechat-work/kernel/response"
)

type ResponseExpressGetContact struct {
	*response2.ResponseMiniProgram

	WayBillID int              `json:"waybill_id"`
	Sender    []*power.HashMap `json:"sender"`
	Receiver  []*power.HashMap `json:"receiver"`
}
