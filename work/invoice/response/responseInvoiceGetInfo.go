package response

import (
	"github.com/yrzs/wechat-work/kernel/power"
	"github.com/yrzs/wechat-work/kernel/response"
)

type ResponseInvoiceGetInfo struct {
	*response.ResponseWork

	CardID    string         `json:"card_id"`
	BeginTime string         `json:"begin_time"`
	EndTime   string         `json:"end_time"`
	OpenID    string         `json:"openid"`
	Type      string         `json:"type"`
	Payee     string         `json:"payee"`
	Detail    string         `json:"detail"`
	UserInfo  *power.HashMap `json:"user_info"`
}
