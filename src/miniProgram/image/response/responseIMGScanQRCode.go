package response

import (
	"github.com/yrzs/wechat-work/kernel/power"
	"github.com/yrzs/wechat-work/kernel/response"
)

type ResponseIMGScanQRCode struct {
	*response.ResponseMiniProgram
	CodeResults []*power.HashMap `json:"code_results"`
	ImgSize     power.HashMap    `json:"img_size"`
}
