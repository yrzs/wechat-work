package response

import (
	"github.com/yrzs/wechat-work/kernel/response"
)

type ResponseOpenIDToUserID struct {
	*response.ResponseWork

	UserID string `json:"userid"`
}
