package response

import (
	"github.com/yrzs/wechat-work/kernel/power"
	"github.com/yrzs/wechat-work/kernel/response"
)

type ResponseMenuCreate struct {
	*response.ResponseWork
	Button []*power.HashMap `json:"button"`
}
