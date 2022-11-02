package response

import (
	"github.com/yrzs/wechat-work/kernel/power"
	"github.com/yrzs/wechat-work/kernel/response"
)

type ResponseCustomerStrategyGetRange struct {
	*response.ResponseWork

	Range []*power.HashMap `json:"range"`
}
