package response

import (
	"github.com/yrzs/wechat-work/kernel/power"
	"github.com/yrzs/wechat-work/kernel/response"
)

type ResponseLinkCorpGetUser struct {
	*response.ResponseWork

	UserInfo *power.HashMap `json:"user_info"`
}
