package response

import (
	"github.com/yrzs/wechat-work/kernel/power"
	"github.com/yrzs/wechat-work/kernel/response"
)

type ResponseBroadcastGetRoleList struct {
	*response.ResponseMiniProgram

	Total int              `json:"total,omitempty"`
	List  []*power.HashMap `json:"list,omitempty"`
}
