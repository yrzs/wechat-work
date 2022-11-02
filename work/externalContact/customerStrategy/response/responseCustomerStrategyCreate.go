package response

import (
	"github.com/yrzs/wechat-work/kernel/power"
	"github.com/yrzs/wechat-work/kernel/response"
)

type ResponseCustomerStrategyCreate struct {
	*response.ResponseWork

	StrategyID *power.HashMap `json:"strategy_id"`
}
