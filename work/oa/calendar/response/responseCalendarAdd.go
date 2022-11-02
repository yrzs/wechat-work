package response

import (
	"github.com/yrzs/wechat-work/kernel/response"
)

type ResponseCalendarAdd struct {
	*response.ResponseWork

	CalID string `json:"cal_id"`
}
