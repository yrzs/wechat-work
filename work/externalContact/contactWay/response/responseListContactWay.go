package response

import (
	"github.com/yrzs/wechat-work/kernel/response"
)

// ------------------------------------------------------------------------------------
type ContactWayID struct {
	ConfigID string `json:"config_id"`
}

type ResponseListContactWay struct {
	response.ResponseWork

	ContactWayIDs []*ContactWayID `json:"contact_way"`
	NextCursor    string          `json:"next_cursor"`
}
