package response

import (
	"github.com/yrzs/wechat-work/kernel/power"
	"github.com/yrzs/wechat-work/kernel/response"
)

type ResponseMeetingRoomList struct {
	*response.ResponseWork

	MeetingRoomList []*power.HashMap `json:"meetingroom_list"`
}
