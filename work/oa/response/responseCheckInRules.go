package response

import (
	"github.com/yrzs/wechat-work/kernel/power"
	"github.com/yrzs/wechat-work/kernel/response"
)

type ResponseCheckInRules struct {
	*response.ResponseWork
	Info []*power.HashMap `json:"info"`
}
