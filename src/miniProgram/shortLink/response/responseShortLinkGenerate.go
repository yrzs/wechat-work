package response

import (
	"github.com/yrzs/wechat-work/kernel/response"
)

type ResponseShortLinkGenerate struct {
	*response.ResponseMiniProgram

	Link string `json:"link"`
}
