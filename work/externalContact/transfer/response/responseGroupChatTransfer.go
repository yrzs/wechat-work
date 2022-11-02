package response

import (
	"github.com/yrzs/wechat-work/kernel/power"
	"github.com/yrzs/wechat-work/kernel/response"
)

type ResponseGroupChatTransfer struct {
	*response.ResponseWork

	Customer []*power.HashMap `json:"failed_chat_list"`
}
