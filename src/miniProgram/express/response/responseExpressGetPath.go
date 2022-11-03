package response

import (
	"github.com/yrzs/wechat-work/kernel/power"
	"github.com/yrzs/wechat-work/kernel/response"
)

type ResponseExpressGetPath struct {
	*response.ResponseMiniProgram

	OpenID       string             `json:"openid"`        //   "openid",
	DeliveryID   string             `json:"delivery_id"`   //  "",
	WaybillID    string             `json:"waybill_id"`    //   "123456789",
	PathItemNum  int                `json:"path_item_num"` //  3,
	PathItemList []*power.StringMap `json:"path_item_list"`
}
