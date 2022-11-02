package request

import "github.com/yrzs/wechat-work/kernel/power"

type RequestSubscribeMessageSend struct {
	ToUser           string         `json:"touser"`
	TemplateID       string         `json:"template_id"`
	Page             string         `json:"page,omitempty"`
	MiniProgramState string         `json:"miniprogram_state,omitempty"`
	Lang             string         `json:"lang,omitempty"`
	Data             *power.HashMap `json:"data"`
}
