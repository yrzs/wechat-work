package response

import (
	"github.com/yrzs/wechat-work/kernel/power"
	"github.com/yrzs/wechat-work/kernel/response"
)

type ResponseServicerList struct {
	*response.ResponseWork

	ServicerList []*power.HashMap `json:"servicer_list"`
}
