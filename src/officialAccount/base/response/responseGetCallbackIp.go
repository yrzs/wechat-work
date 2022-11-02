package response

import (
	"github.com/yrzs/wechat-work/kernel/response"
)

type ResponseGetCallBackIP struct {
	*response.ResponseOfficialAccount

	IPList []string `json:"ip_list"`
}
