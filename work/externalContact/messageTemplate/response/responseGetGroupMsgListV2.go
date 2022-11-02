package response

import (
	"github.com/yrzs/wechat-work/kernel/response"
	"github.com/yrzs/wechat-work/work/externalContact/messageTemplate/request"
)

type GroupMsg struct {
	MsgID       string                `json:"msgid"`
	Creator     string                `json:"creator"`
	CreateTime  int                   `json:"create_time"`
	CreateType  int                   `json:"create_type"`
	Text        request.TextOfMessage `json:"text"`
	Attachments []*request.Attachment `json:"attachments"`
}

type ResponseGetGroupMsgListV2 struct {
	*response.ResponseWork

	NextCursor   string      `json:"next_cursor"`
	GroupMsgList []*GroupMsg `json:"group_msg_list"`
}
