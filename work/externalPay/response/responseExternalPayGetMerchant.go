package response

import (
	"github.com/yrzs/wechat-work/kernel/power"
	"github.com/yrzs/wechat-work/kernel/response"
)

type ResponseExternalPayGetMerchant struct {
	*response.ResponseWork

	BindStatus    string           `json:"bind_status"`
	MchID         string           `json:"mch_id"`
	MerchantName  string           `json:"merchant_name"`
	AllowUseScope []*power.HashMap `json:"allow_use_scope"`
}
