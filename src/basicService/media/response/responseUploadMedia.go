package response

import (
	"github.com/yrzs/wechat-work/kernel/power"
	"github.com/yrzs/wechat-work/kernel/response"
)

type ResponseUploadMedia struct {
	*response.ResponseOfficialAccount

	Item         []*power.HashMap `json:"item"`
	Type         string           `json:"type"`
	MediaID      string           `json:"media_id"`
	ThumbMediaID string           `json:"thumb_media_id,omitempty"`
	CreatedAt    int              `json:"created_at"`
}
