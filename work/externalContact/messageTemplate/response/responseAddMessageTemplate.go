package response

import "github.com/yrzs/wechat-work/kernel/response"

type ResponseAddMessageTemplate struct {
	*response.ResponseWork
	FailList []string `json:"fail_list"`
	MsgID    string   `json:"msgid"`
}
