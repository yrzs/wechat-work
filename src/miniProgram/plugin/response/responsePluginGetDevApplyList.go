package response

import (
	"github.com/yrzs/wechat-work/kernel/power"
	"github.com/yrzs/wechat-work/kernel/response"
)

type ResponsePluginGetDevApplyList struct {
	*response.ResponseMiniProgram
	ApplyList []*power.HashMap `json:"apply_list"`
}
