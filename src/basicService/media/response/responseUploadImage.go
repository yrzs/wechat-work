package response

import (
	"github.com/yrzs/wechat-work/kernel/response"
)

type ResponseUploadImage struct {
	*response.ResponseOfficialAccount

	URL string `json:"url"`
}
