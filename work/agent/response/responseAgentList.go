package response

import (
	"github.com/yrzs/wechat-work/kernel/response"
)

type ResponseAgentList struct {
	*response.ResponseWork
	AgentList []ResponseAgentGet `json:"agentlist"`
}
