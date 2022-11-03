package response

import (
	"github.com/yrzs/wechat-work/kernel/response"
)

type ResponseIMGSuperResolution struct {
	*response.ResponseMiniProgram
	MediaID string `json:"media_id"`
}
