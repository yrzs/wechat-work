package response

import (
	"github.com/yrzs/wechat-work/kernel/response"
)

type ResponseMaterialAddMaterial struct {
	*response.ResponseOfficialAccount

	MediaID string `json:"media_id"`
	URL     string `json:"url"`
}
