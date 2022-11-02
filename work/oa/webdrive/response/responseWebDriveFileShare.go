package response

import (
	"github.com/yrzs/wechat-work/kernel/response"
)

type ResponseWebDriveFileShare struct {
	*response.ResponseWork

	ShareURL string `json:"share_url"`
}
