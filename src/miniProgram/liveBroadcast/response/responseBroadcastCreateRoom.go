package response

import "github.com/yrzs/wechat-work/kernel/response"

type ResponseBroadcastCreateRoom struct {
	*response.ResponseMiniProgram

	RoomID    int    `json:"roomId"`
	QRCodeURL string `json:"qrcode_url"`
}
