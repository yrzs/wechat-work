package response

import (
	"github.com/yrzs/wechat-work/kernel/response"
)

type ResponseAppChatCreate struct {
	*response.ResponseWork

	ChatID string `json:"chatid"`
}
