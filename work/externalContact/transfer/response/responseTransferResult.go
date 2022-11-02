package response

import (
	"github.com/yrzs/wechat-work/kernel/power"
	"github.com/yrzs/wechat-work/kernel/response"
)

type ResponseTransferResult struct {
	*response.ResponseWork

	Customer   []*power.HashMap `json:"customer"`
	NextCursor string           `json:"next_cursor"`
}
