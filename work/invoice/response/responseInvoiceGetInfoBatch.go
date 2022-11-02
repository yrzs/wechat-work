package response

import (
	"github.com/yrzs/wechat-work/kernel/power"
	"github.com/yrzs/wechat-work/kernel/response"
)

type ResponseInvoiceGetInfoBatch struct {
	*response.ResponseWork

	ItemList []*power.HashMap `json:"item_list"`
}
