package response

import (
	"github.com/yrzs/wechat-work/kernel/response"
)

type ResponseOperationGetPerformance struct {
	*response.ResponseMiniProgram
	DefaultTimeData string `json:"default_time_data"`
	CompareTimeData string `json:"compare_time_data"`
}
