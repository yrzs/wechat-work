package response

import (
	"github.com/yrzs/wechat-work/kernel/response"
)

type ResponseBroadcastGetSubAnchor struct {
	*response.ResponseMiniProgram

	UserName string `json:"username"`
}
