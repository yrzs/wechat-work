package response

import (
	"github.com/yrzs/wechat-work/kernel/power"
	"github.com/yrzs/wechat-work/kernel/response"
)

type ResponseMomentStrategyGetRange struct {
	*response.ResponseWork

	Range      []*power.HashMap `json:"range"`
	NextCursor string           `json:"next_cursor"`
}
