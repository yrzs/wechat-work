package response

import (
	"github.com/yrzs/wechat-work/kernel/power"
	"github.com/yrzs/wechat-work/kernel/response"
)

type ResponseMomentGetMomentTask struct {
	*response.ResponseWork

	TaskList   []*power.HashMap `json:"task_list"`
	NextCursor string           `json:"next_cursor"`
}
