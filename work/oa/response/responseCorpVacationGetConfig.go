package response

import (
	"github.com/yrzs/wechat-work/kernel/power"
	"github.com/yrzs/wechat-work/kernel/response"
)

type ResponseCorpVacationGetConfig struct {
	*response.ResponseWork
	Lists []*power.HashMap `json:"lists"`
}
