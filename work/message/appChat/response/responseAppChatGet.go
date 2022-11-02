package response

import (
	"github.com/yrzs/wechat-work/kernel/power"
	"github.com/yrzs/wechat-work/kernel/response"
)

type ResponseAppChatGet struct {
	*response.ResponseWork

	ChatInfo *power.HashMap `json:"chat_info"`
}
