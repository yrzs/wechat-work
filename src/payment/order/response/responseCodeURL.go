package response

import (
	"github.com/yrzs/wechat-work/kernel/response"
)

type ResponseCodeURL struct {
	response.ResponsePayment

	CodeURL string `json:"code_url"`
}
