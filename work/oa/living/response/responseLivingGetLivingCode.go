package response

import (
	"github.com/yrzs/wechat-work/kernel/response"
)

type ResponseLivingGetLivingCode struct {
	*response.ResponseWork

	LivingCode int `json:"living_code"`
}
