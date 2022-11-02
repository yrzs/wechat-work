package response

import "github.com/yrzs/wechat-work/kernel/response"

type ResponseBind struct {
	response.ResponseOpenPlatform

	Userstr string `json:"userstr"`
}

type Member struct {
	UserStr string `json:"userstr"`
}

type ResponseList struct {
	response.ResponseOpenPlatform

	Members []*Member `json:"members"`
}
