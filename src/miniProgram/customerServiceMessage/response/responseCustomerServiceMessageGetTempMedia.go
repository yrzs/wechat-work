package response

import (
	"github.com/yrzs/wechat-work/kernel/response"
)

type ResponseCustomerServiceMessageGetTempMedia struct {
	*response.ResponseMiniProgram
	ContentType string `json:"contentType"`
	Buffer      []byte `json:"buffer"`
}
