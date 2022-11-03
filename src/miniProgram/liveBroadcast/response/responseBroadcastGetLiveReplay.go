package response

import (
	"github.com/yrzs/wechat-work/kernel/power"
	"github.com/yrzs/wechat-work/kernel/response"
)

type ResponseBroadcastGetLiveReplay struct {
	*response.ResponseMiniProgram

	Total      int              `json:"total"`
	LiveReplay []*power.HashMap `json:"live_replay"`
}
