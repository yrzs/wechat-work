package response

import (
	"github.com/yrzs/wechat-work/kernel/power"
	"github.com/yrzs/wechat-work/kernel/response"
)

type ResponseResignedTransferResult struct {
	*response.ResponseWork

	Customer []*power.HashMap `json:"customer"`
}
