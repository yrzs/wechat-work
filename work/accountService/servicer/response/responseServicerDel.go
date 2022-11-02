package response

import (
	"github.com/yrzs/wechat-work/kernel/power"
	"github.com/yrzs/wechat-work/kernel/response"
)

type ResponseServicerDel struct {
	*response.ResponseWork

	ResultList []*power.HashMap `json:"result_list"`
}
