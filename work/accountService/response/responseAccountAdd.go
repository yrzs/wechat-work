package response

import (
	"github.com/yrzs/wechat-work/kernel/response"
)

type ResponseAccountAdd struct {
	*response.ResponseWork

	OpenKFID string `json:"open_kfid"`
}
