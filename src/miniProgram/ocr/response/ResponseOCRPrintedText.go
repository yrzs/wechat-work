package response

import (
	"github.com/yrzs/wechat-work/kernel/power"
	"github.com/yrzs/wechat-work/kernel/response"
)

type ResponseOCRPrintedText struct {
	*response.ResponseMiniProgram
	Items   []*power.HashMap `json:"items"`
	ImgSize []*power.HashMap `json:"img_size"`
}
