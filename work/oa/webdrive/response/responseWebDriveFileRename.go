package response

import (
	"github.com/yrzs/wechat-work/kernel/power"
	"github.com/yrzs/wechat-work/kernel/response"
)

type ResponseWebDriveFileRename struct {
	*response.ResponseWork

	File *power.HashMap `json:"file"`
}
