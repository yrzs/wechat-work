package response

import (
	"github.com/yrzs/wechat-work/kernel/power"
	"github.com/yrzs/wechat-work/kernel/response"
)

type ResponseMomentStrategyGet struct {
	*response.ResponseWork

	Strategy *power.HashMap `json:"strategy"`
}
