package response

import (
	"github.com/yrzs/wechat-work/kernel/response"
)

type ResponseGetAPIDomainIP struct {
	IPList []string `json:"ip_list"`
	*response.ResponseWork
}
