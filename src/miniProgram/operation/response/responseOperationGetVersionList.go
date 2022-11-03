package response

import (
	"github.com/yrzs/wechat-work/kernel/power"
	"github.com/yrzs/wechat-work/kernel/response"
)

type ResponseOperationGetVersionList struct {
	*response.ResponseMiniProgram
	CVList []*power.HashMap `json:"cvlist"`
}
