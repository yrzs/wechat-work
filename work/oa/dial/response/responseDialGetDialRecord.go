package response

import (
	"github.com/yrzs/wechat-work/kernel/power"
	"github.com/yrzs/wechat-work/kernel/response"
)

type ResponseDialGetDialRecord struct {
	*response.ResponseWork

	MeetingidList []*power.HashMap `json:"record"`
}
