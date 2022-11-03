package response

import (
	"github.com/yrzs/wechat-work/kernel/response"
)

type ResponseAuthGetPaidUnionID struct {
	UnionID string `json:"unionid"`

	*response.ResponseMiniProgram
}
