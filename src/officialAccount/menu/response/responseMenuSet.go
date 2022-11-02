package response

import (
	"github.com/yrzs/wechat-work/kernel/response"
)

type ResponseMenuCreate struct {
	*response.ResponseOfficialAccount
}

type ResponseMenuCreateConditional struct {
	*response.ResponseOfficialAccount

	MenuID string `json:"menuid"`
}
