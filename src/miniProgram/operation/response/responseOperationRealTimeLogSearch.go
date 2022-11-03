package response

import (
	"github.com/yrzs/wechat-work/kernel/power"
	"github.com/yrzs/wechat-work/kernel/response"
)

type ResponseOperationRealTimeLogSearch struct {
	*response.ResponseMiniProgram
	Data *power.HashMap   `json:"data"`
	List []*power.HashMap `json:"list"`
}
