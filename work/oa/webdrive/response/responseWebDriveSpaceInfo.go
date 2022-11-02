package response

import (
	"github.com/yrzs/wechat-work/kernel/power"
	"github.com/yrzs/wechat-work/kernel/response"
)

type ResponseWebDriveSpaceInfo struct {
	*response.ResponseWork

	SpaceInfo *power.HashMap `json:"space_info"`
}
