package response

import (
	"github.com/yrzs/wechat-work/kernel/response"
)

type ResponseCorpGroupTransferSession struct {
	*response.ResponseWork
	Userid     string `json:"userid"`
	SessionKey string `json:"session_key"`
}
