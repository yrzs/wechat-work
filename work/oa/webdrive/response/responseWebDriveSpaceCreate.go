package response

import (
	"github.com/yrzs/wechat-work/kernel/response"
)

type ResponseWebDriveSpaceCreate struct {
	*response.ResponseWork

	SpaceID string `json:"spaceid"`
}
