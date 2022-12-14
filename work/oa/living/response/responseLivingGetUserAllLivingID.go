package response

import (
	"github.com/yrzs/wechat-work/kernel/response"
)

type ResponseLivingGetUserAllLivingID struct {
	*response.ResponseWork

	NextCursor   string   `json:"next_cursor"`
	LivingIDList []string `json:"livingid_list"`
}
