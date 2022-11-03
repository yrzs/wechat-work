package response

import (
	"github.com/yrzs/wechat-work/kernel/power"
	"github.com/yrzs/wechat-work/kernel/response"
)

type ResponseBroadcastGetFollowers struct {
	*response.ResponseMiniProgram

	Followers []*power.HashMap `json:"followers"`
	PageBreak string           `json:"page_break,omitempty"`
}
