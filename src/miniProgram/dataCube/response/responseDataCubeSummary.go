package response

import (
	"github.com/yrzs/wechat-work/kernel/power"
	"github.com/yrzs/wechat-work/kernel/response"
)

type ResponseDataCubeSummary struct {
	response.ResponseMiniProgram

	List []*power.HashMap `json:"list"`
}
