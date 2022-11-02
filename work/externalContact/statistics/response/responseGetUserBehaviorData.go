package response

import (
	"github.com/yrzs/wechat-work/kernel/power"
	"github.com/yrzs/wechat-work/kernel/response"
)

type ResponseGetUserBehaviorData struct {
	*response.ResponseWork

	MomentList []*power.HashMap `json:"behavior_data"`
}
