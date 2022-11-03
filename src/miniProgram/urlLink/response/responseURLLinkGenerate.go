package response

import (
	"github.com/yrzs/wechat-work/kernel/response"
)

type ResponseURLLinkGenerate struct {
	*response.ResponseMiniProgram
	URLLink string `json:"url_link"`
}
