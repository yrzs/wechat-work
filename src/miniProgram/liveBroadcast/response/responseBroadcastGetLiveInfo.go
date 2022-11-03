package response

import (
	"github.com/yrzs/wechat-work/kernel/power"
	"github.com/yrzs/wechat-work/kernel/response"
)

type ResponseBroadcastGetLiveInfo struct {
	*response.ResponseMiniProgram

	Total    int              `json:"total"`
	RoomInfo []*power.HashMap `json:"room_info"`
}
