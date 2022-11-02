package response

import (
	"github.com/yrzs/wechat-work/kernel/power"
	"github.com/yrzs/wechat-work/kernel/response"
)

type ResponseExternalPayGetBillList struct {
	*response.ResponseWork

	NextCursor string           `json:"next_cursor"`
	BillList   []*power.HashMap `json:"bill_list"`
}
