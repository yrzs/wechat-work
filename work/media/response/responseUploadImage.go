package response

import (
	"github.com/yrzs/wechat-work/kernel/response"
)

type ResponseUploadImage struct {
	*response.ResponseWork

	URL string `json:"url"`
}
