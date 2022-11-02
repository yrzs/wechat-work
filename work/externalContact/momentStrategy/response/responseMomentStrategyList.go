package response

import (
	"github.com/yrzs/wechat-work/kernel/response"
)

type ResponseMomentStrategyList struct {
	*response.ResponseWork

	StrategyID int64 `json:"strategy_id"`
}
