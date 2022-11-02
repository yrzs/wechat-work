package response

import (
	"github.com/yrzs/wechat-work/kernel/power"
	"github.com/yrzs/wechat-work/kernel/response"
)

type ResponseLivingGetLivingInfo struct {
	*response.ResponseWork

	LivingInfo *power.HashMap `json:"living_info"`
}
