package response

import (
	"github.com/yrzs/wechat-work/kernel/models"
	"github.com/yrzs/wechat-work/kernel/response"
)

type ResponseDepartmentGet struct {
	*response.ResponseWork
	Department *models.Department `json:"department"`
}
