package response

import (
	"github.com/yrzs/wechat-work/kernel/response"
)

type ResponseUnitfy struct {
	response.ResponsePayment

	PrepayID string `json:"prepay_id"`
}
