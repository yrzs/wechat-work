package response

import (
	"github.com/yrzs/wechat-work/kernel/response"
)

type ResponseJoinCode struct {
	*response.ResponseWork

	JoinCode string `json:"join_qrcode"`
}
