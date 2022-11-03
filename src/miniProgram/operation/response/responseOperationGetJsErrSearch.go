package response

import (
	"github.com/yrzs/wechat-work/kernel/power"
	"github.com/yrzs/wechat-work/kernel/response"
)

type ResponseOperationGetJsErrSearch struct {
	*response.ResponseMiniProgram
	results *power.HashMap `json:"results"`
	Total   int64          `json:"total"`
}
