package response

import (
	"github.com/yrzs/wechat-work/kernel/power"
	"github.com/yrzs/wechat-work/kernel/response"
)

type ResponseImmediateDeliveryBindAccount struct {
	*response.ResponseMiniProgram

	ShopList []*power.HashMap `json:"shop_list"`
}
