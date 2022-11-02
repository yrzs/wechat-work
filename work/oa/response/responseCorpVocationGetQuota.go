package response

import (
	"github.com/yrzs/wechat-work/kernel/power"
	"github.com/yrzs/wechat-work/kernel/response"
)

type ResponseCorpVacationGetQuota struct {
	*response.ResponseWork
	Lists []*power.HashMap `json:"lists"`
}
