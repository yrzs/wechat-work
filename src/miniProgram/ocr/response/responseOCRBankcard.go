package response

import (
	"github.com/yrzs/wechat-work/kernel/response"
)

type ResponseOCRBankcard struct {
	*response.ResponseMiniProgram
	Number string `json:"number"`
}
