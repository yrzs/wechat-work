package response

import (
	"github.com/yrzs/wechat-work/kernel/response"
)

type ResponseUserActiveCount struct {
	*response.ResponseWork

	ActiveCount string `json:"active_cnt"`
}
