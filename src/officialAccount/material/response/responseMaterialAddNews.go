package response

import (
	"github.com/yrzs/wechat-work/kernel/response"
)

type ResponseMaterialAddNews struct {
	*response.ResponseOfficialAccount

	MediaID string `json:"media_id"`
}
