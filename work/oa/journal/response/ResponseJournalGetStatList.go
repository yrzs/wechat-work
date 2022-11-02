package response

import (
	"github.com/yrzs/wechat-work/kernel/power"
	"github.com/yrzs/wechat-work/kernel/response"
)

type ResponseJournalGetStatList struct {
	*response.ResponseWork

	StatList *power.HashMap `json:"stat_list"`
}
