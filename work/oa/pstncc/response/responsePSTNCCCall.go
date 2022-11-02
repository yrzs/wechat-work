package response

import (
	"github.com/yrzs/wechat-work/kernel/power"
	"github.com/yrzs/wechat-work/kernel/response"
)

type ResponsePSTNCCCall struct {
	*response.ResponseWork

	States []*power.HashMap `json:"states"`
}
