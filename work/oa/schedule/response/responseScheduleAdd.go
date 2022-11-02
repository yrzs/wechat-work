package response

import (
	"github.com/yrzs/wechat-work/kernel/response"
)

type ResponseScheduleAdd struct {
	*response.ResponseWork

	ScheduleID string `json:"schedule_id"`
}
