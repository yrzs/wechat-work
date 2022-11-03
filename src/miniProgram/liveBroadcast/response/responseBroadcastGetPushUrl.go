package response

import (
	"github.com/yrzs/wechat-work/kernel/response"
)

type ResponseBroadcastGetPushUrl struct {
	*response.ResponseMiniProgram

	PushAddr string `json:"pushAddr"`
}
