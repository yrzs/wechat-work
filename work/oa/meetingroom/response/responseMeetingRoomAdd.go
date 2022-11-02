package response

import (
	"github.com/yrzs/wechat-work/kernel/response"
)

type ResponseMeetingRoomAdd struct {
	*response.ResponseWork

	MeetingRoomID int `json:"meetingroom_id"`
}
