package response

import (
	"github.com/yrzs/wechat-work/kernel/power"
	"github.com/yrzs/wechat-work/kernel/response"
)

type ResponseGroupChatOpenGIDToChatID struct {
	*response.ResponseWork
	ChatID *power.HashMap `json:"chat_id"`
}
