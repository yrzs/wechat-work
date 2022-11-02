package response

import (
	"github.com/yrzs/wechat-work/kernel/power"
	"github.com/yrzs/wechat-work/kernel/response"
)

type ResponseStatistic struct {
	*response.ResponseWork

	Total      int `json:"total"`
	NextOffset int `json:"next_offset"`

	Items []*power.HashMap `json:"items"`
}
