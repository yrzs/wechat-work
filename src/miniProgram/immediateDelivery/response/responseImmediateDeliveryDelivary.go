package response

import (
	"github.com/ArtisanCloud/PowerLibs/v2/object"
	"github.com/yrzs/wechat-work/kernel/response"
)

type ResponseImmediateDeliveryDelivery struct {
	*response.ResponseMiniProgram

	List []*object.HashMap `json:"list"`
}
