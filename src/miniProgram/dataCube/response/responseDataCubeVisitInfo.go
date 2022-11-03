package response

import (
	"github.com/yrzs/wechat-work/kernel/power"
	"github.com/yrzs/wechat-work/kernel/response"
)

type ResponseDataCubeVisitInfo struct {
	*response.ResponseMiniProgram

	RefDate    string           `json:"ref_date"`
	VisitUVNew []*power.HashMap `json:"visit_uv_new"`
	VisitUV    []*power.HashMap `json:"visit_uv"`
}
