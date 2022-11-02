package response

import (
	"github.com/yrzs/wechat-work/kernel/response"
)

type ResponseWebDriveFileCreate struct {
	*response.ResponseWork

	FileID string `json:"fileid"`
	Url    string `json:"url"`
}
