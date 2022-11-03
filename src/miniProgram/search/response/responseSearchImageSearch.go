package response

import (
	"github.com/yrzs/wechat-work/kernel/power"
	"github.com/yrzs/wechat-work/kernel/response"
)

type ResponseSearchImageSearch struct {
	*response.ResponseMiniProgram

	Items []*power.HashMap `json:"items"`
}
