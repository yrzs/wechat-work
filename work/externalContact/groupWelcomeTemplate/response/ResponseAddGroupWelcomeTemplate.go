package response

import "github.com/yrzs/wechat-work/kernel/response"

type ResponseAddGroupWelcomeTemplate struct {
	*response.ResponseWork

	TemplateID string `json:"template_id"`
}
