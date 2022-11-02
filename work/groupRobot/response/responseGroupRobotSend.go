package response

import (
	"github.com/yrzs/wechat-work/kernel/response"
)

type ResponseGroupRobotSend struct {
	*response.ResponseWork

	Type      string `json:"type"`
	MediaID   string `json:"media_id"`
	CreatedAt string `json:"created_at"`
}
