package response

import (
	"github.com/yrzs/wechat-work/kernel/response"
)

type ResponseWebDriveSpaceShare struct {
	*response.ResponseWork

	SpaceShareURL string `json:"space_share_url"`
}
