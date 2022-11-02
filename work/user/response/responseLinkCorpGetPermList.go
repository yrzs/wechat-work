package response

import (
	"github.com/yrzs/wechat-work/kernel/response"
)

type ResponseLinkCorpGetPermList struct {
	*response.ResponseWork

	UserIDs       []string `json:"userids"`
	DepartmentIDs []string `json:"department_ids"`
}
