package response

import (
	"github.com/yrzs/wechat-work/kernel/response"
	"github.com/yrzs/wechat-work/work/user/request"
)

type ResponseTagList struct {
	*response.ResponseWork

	TagName string                `json:"tagname"`
	TagList []*request.RequestTag `json:"taglist"`
}
