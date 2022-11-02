package response

import (
	"github.com/yrzs/wechat-work/kernel/power"
	"github.com/yrzs/wechat-work/kernel/response"
)

type ResponseCustomerBatchGet struct {
	*response.ResponseWork

	CustomerList          []*power.HashMap `json:"customer_list"`
	InvalidExternalUserID []*string        `json:"invalid_external_userid"`
}
