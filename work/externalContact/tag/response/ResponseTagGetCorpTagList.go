package response

import (
	"github.com/yrzs/wechat-work/kernel/response"
)

type ResponseTagGetCorpTagList struct {
	response.ResponseWork

	TagGroups []*CorpTagGroup `json:"tag_group"`
}
