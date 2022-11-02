package response

import (
	"github.com/yrzs/wechat-work/kernel/response"
)

type ResponseTagCreate struct {
	*response.ResponseWork

	TagID string `json:"tagid"`
}
