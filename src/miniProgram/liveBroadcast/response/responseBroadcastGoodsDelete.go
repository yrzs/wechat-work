package response

import (
	"github.com/yrzs/wechat-work/kernel/power"
	"github.com/yrzs/wechat-work/kernel/response"
)

type ResponseBroadcastGoodsDelete struct {
	*response.ResponseMiniProgram

	Total int              `json:"total"`
	Goods []*power.HashMap `json:"goods"`
}