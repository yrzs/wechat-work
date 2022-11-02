package response

import (
	"github.com/yrzs/wechat-work/kernel/power"
	"github.com/yrzs/wechat-work/kernel/response"
)

type ResponseLivingGetWatchStat struct {
	*response.ResponseWork

	NextKey  string         `json:"next_key"`
	StatInfo *power.HashMap `json:"stat_info"`
}
