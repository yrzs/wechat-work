package response

import (
	"github.com/yrzs/wechat-work/kernel/power"
	"github.com/yrzs/wechat-work/kernel/response"
)

type ResponseCustomerStrategyGet struct {
	*response.ResponseWork

	Strategy *power.HashMap `json:"momentStrategy"`
}
