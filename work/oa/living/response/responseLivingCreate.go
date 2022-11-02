package response

import (
	"github.com/yrzs/wechat-work/kernel/response"
)

type ResponseLivingCreate struct {
	*response.ResponseWork

	LivingID int `json:"livingid"`
}
