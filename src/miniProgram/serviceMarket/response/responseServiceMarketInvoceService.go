package response

import (
	"github.com/yrzs/wechat-work/kernel/response"
)

type ResponseServiceMarketInvoceService struct {
	*response.ResponseMiniProgram

	Data string `json:"data"`
}
