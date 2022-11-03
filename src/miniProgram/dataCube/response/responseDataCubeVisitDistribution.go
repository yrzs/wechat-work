package response

import (
	"github.com/yrzs/wechat-work/kernel/power"
	"github.com/yrzs/wechat-work/kernel/response"
)

type ResponseDataCubeVisit struct {
	*response.ResponseMiniProgram

	RefDate string           `json:"ref_date"`
	List    []*power.HashMap `json:"list"`
}
