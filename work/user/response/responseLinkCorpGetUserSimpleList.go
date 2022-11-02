package response

import (
	"github.com/yrzs/wechat-work/kernel/power"
	"github.com/yrzs/wechat-work/kernel/response"
)

type ResponseLinkCorpGetUserSimpleList struct {
	*response.ResponseWork

	UserList []*power.HashMap `json:"userlist"`
}
