package response

import (
	"github.com/yrzs/wechat-work/kernel/response"
)

type ResponseWebDriveFileUpload struct {
	*response.ResponseWork

	FileID string `json:"fileid"`
}
