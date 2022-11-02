package response

import (
	"github.com/yrzs/wechat-work/kernel/response"
)

type ResponseApprovalNoList struct {
	*response.ResponseWork
	SpNoList []string `json:"sp_no_list"`
}
