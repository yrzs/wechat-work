package response

import (
	"github.com/yrzs/wechat-work/kernel/power"
	"github.com/yrzs/wechat-work/kernel/response"
)

type ResponseLinkCorpGetDepartmentList struct {
	*response.ResponseWork

	department_list []*power.HashMap `json:"department_list"`
}
