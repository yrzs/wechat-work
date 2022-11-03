package response

import (
	"github.com/yrzs/wechat-work/kernel/response"
)

type ResponseSoterVerifySignature struct {
	*response.ResponseMiniProgram

	IsOK bool `json:"is_ok"`
}
