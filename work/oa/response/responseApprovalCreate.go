package response

import (
	"github.com/yrzs/wechat-work/kernel/response"
)

type ResponseApprovalCreate struct {
	*response.ResponseWork
	SpNo string `json:"sp_no"`
}
