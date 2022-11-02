package response

import (
	"github.com/yrzs/wechat-work/kernel/power"
	"github.com/yrzs/wechat-work/kernel/response"
)

type ResponseWebDriveFileMove struct {
	*response.ResponseWork

	FileList *power.HashMap `json:"file_list"`
}
