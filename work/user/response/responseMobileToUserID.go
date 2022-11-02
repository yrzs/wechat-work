package response

import (
	"github.com/yrzs/wechat-work/kernel/response"
)

type ResponseMobileToUserID struct {
	*response.ResponseWork

	UserID string `json:"userid"`
}

type ResponseConvertToUserID struct {
	*response.ResponseWork

	UserID string `json:"userid"`
}
