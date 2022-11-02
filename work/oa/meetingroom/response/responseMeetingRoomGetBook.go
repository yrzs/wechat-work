package response

import (
	"github.com/yrzs/wechat-work/kernel/response"
)

type ResponseMeetingRoomGetBook struct {
	*response.ResponseWork

	MeetingID  int `json:"meeting_id"`
	ScheduleID int `json:"schedule_id"`
}
