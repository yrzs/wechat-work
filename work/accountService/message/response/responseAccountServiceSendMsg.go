package response

import (
	"github.com/yrzs/wechat-work/kernel/response"
)

type ResponseAccountServiceSendMsg struct {
	*response.ResponseWork

	MsgID string `json:"msgid"`
}
