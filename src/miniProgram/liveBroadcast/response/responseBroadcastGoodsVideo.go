package response

import (
	"github.com/yrzs/wechat-work/kernel/response"
)

type ResponseBroadcastGoodsVideo struct {
	*response.ResponseMiniProgram

	URL int `json:"url"`
}
