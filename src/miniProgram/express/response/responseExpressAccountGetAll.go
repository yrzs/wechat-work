package response

import (
	"github.com/yrzs/wechat-work/kernel/power"
	response2 "github.com/yrzs/wechat-work/kernel/response"
)

type ResponseExpressAccountGetAll struct {
	*response2.ResponseMiniProgram

	Count int              `json:"count"`
	List  []*power.HashMap `json:"list"`
}
