package response

import (
	"github.com/yrzs/wechat-work/kernel/power"
	"github.com/yrzs/wechat-work/kernel/response"
)

type ResponseAccountList struct {
	*response.ResponseWork

	AccountList []*power.HashMap `json:"account_list"`
}
