package response

import (
	"github.com/yrzs/wechat-work/kernel/response"
)

type ResponseAccountServiceAddContactWay struct {
	*response.ResponseWork

	URL string `json:"url"`
}
