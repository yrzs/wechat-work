package response

import (
	"github.com/yrzs/wechat-work/kernel/response"
)

type ResponseAccountServiceSendMsgOnEvent struct {
	*response.ResponseWork

	MsgID string `json:"msgid"`
}
