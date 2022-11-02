package response

import (
	"github.com/yrzs/wechat-work/kernel/power"
	"github.com/yrzs/wechat-work/kernel/response"
)

type ResponseMeetingRoomGetBookingInfo struct {
	*response.ResponseWork

	BookingList []*power.HashMap `json:"booking_list"`
}
