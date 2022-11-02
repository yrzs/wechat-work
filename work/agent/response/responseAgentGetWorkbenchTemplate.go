package response

import (
	"github.com/yrzs/wechat-work/kernel/response"
	"github.com/yrzs/wechat-work/work/agent/request"
)

type ResponseAgentGetWorkbenchTemplate struct {
	*response.ResponseWork

	TemplateType    string                 `json:"type"`
	Image           request.WorkBenchImage `json:"image"`
	ReplaceUserData bool                   `json:"replace_user_data"`
}
