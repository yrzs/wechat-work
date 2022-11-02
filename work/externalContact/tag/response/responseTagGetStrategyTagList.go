package response

import (
	"github.com/yrzs/wechat-work/kernel/response"
)

type ResponseTagGetStrategyTagList struct {
	response.ResponseWork

	TagGroups []*StrategyTagGroup `json:"tag_group"`
}
