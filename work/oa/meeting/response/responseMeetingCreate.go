package response

import (
	"github.com/yrzs/wechat-work/kernel/response"
)

type ResponseMeetingCreate struct {
	*response.ResponseWork

	MeetingID int `json:"meetingid"`
}
