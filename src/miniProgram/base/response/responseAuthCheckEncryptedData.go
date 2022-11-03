package response

import (
	"github.com/yrzs/wechat-work/kernel/response"
)

type ResponseAuthCheckEncryptedData struct {
	*response.ResponseMiniProgram

	Valid      bool    `json:"vaild"`
	CreateTime float64 `json:"create_time"`
}
