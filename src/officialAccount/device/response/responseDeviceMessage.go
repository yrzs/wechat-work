package response

import "github.com/yrzs/wechat-work/kernel/response"

type ResponseDeviceMessage struct {
	*response.ResponseOfficialAccount

	Ret     int    `json:"ret"`
	RetInfo string `json:"ret_info"`
}
