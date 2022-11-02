package response

import (
	"github.com/yrzs/wechat-work/kernel/power"
	"github.com/yrzs/wechat-work/kernel/response"
)

type ResponseCalendarGet struct {
	*response.ResponseWork

	CalendarList []*power.HashMap `json:"calendar_list"`
}
