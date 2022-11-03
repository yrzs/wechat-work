package response

import (
	"github.com/yrzs/wechat-work/kernel/power"
	"github.com/yrzs/wechat-work/kernel/response"
)

type ResponseOperationGetGrayReleasePlan struct {
	*response.ResponseMiniProgram
	GrayReleasePlan *power.HashMap `json:"gray_release_plan"`
}
