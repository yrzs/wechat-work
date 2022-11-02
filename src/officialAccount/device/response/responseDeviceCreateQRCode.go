package response

import "github.com/yrzs/wechat-work/kernel/response"

type Code struct {
	DeviceID string `json:"device_id"`
	Ticket   string `json:"ticket"`
}

type ResponseDeviceCreateQRCode struct {
	*response.ResponseOfficialAccount

	DeviceNum int     `json:"device_num"`
	CodeList  []*Code `json:"code_list"`
}
