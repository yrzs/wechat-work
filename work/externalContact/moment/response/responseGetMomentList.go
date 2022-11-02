package response

import (
	"github.com/yrzs/wechat-work/kernel/power"
	"github.com/yrzs/wechat-work/kernel/response"
)

type ResponseGetMomentList struct {
	*response.ResponseWork
	NextCursor string         `json:"next_cursor"`
	MomentList *power.HashMap `json:"moment_list"`
}
