package response

import (
	"github.com/yrzs/wechat-work/kernel/power"
	"github.com/yrzs/wechat-work/kernel/response"
)

type ResponseCustomerGetUpgradeServiceConfig struct {
	*response.ResponseWork

	MemberRange    *power.HashMap `json:"member_range"`
	GroupChatRange *power.HashMap `json:"groupchat_range"`
}
