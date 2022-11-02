package response

import (
	"github.com/yrzs/wechat-work/kernel/power"
	"github.com/yrzs/wechat-work/kernel/response"
)

type ResponseCheckInRecords struct {
	*response.ResponseWork

	CheckInData []*power.HashMap `json:"checkindata"`
}
