package response

import (
	"github.com/yrzs/wechat-work/kernel/response"
)

type ResponseDepartmentCreate struct {
	*response.ResponseWork
	ID int `json:"id"`
}
