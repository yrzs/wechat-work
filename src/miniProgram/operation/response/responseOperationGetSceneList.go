package response

import (
	"github.com/yrzs/wechat-work/kernel/power"
	"github.com/yrzs/wechat-work/kernel/response"
)

type ResponseOperationGetSceneList struct {
	*response.ResponseMiniProgram
	Scene []*power.HashMap `json:"scene"`
}
