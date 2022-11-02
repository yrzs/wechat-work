package response

import (
	"github.com/yrzs/wechat-work/kernel/power"
	"github.com/yrzs/wechat-work/kernel/response"
)

type ResponseGetAPIDomainIP struct {
	*response.ResponseOfficialAccount

	DNS  []power.StringMap `json:"dns"`
	Ping []power.StringMap `json:"ping"`
}
