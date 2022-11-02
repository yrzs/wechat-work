package response

import "github.com/yrzs/wechat-work/kernel/response"

type ResponseAddContactWay struct {
	response.ResponseWork

	ConfigID string `json:"config_id"`
	QRCode   string `json:"qr_code"`
}
