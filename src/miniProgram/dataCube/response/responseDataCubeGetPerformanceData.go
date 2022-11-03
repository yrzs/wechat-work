package response

import (
	"github.com/yrzs/wechat-work/kernel/power"
	"github.com/yrzs/wechat-work/kernel/response"
)

type ResponseDataCubeGetPerformanceData struct {
	Data *power.HashMap `json:"data"`

	*response.ResponseMiniProgram
}
