package response

import (
	"github.com/yrzs/wechat-work/kernel/power"
	"github.com/yrzs/wechat-work/kernel/response"
)

type ResponseScheduleGet struct {
	*response.ResponseWork

	ScheduleList []*power.HashMap `json:"schedule_list"`
}
