package response

import (
	"github.com/yrzs/wechat-work/kernel/power"
	"github.com/yrzs/wechat-work/kernel/response"
)

type ResponseCorpCheckInRules struct {
	*response.ResponseWork
	Group []*power.HashMap `json:"group"`
}
