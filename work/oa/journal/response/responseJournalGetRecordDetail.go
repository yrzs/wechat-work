package response

import (
	"github.com/yrzs/wechat-work/kernel/power"
	"github.com/yrzs/wechat-work/kernel/response"
)

type ResponseJournalGetRecordDetail struct {
	*response.ResponseWork

	Info *power.HashMap `json:"info"`
}
