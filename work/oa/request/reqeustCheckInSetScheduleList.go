package request

import "github.com/yrzs/wechat-work/kernel/power"

type RequestCheckInSetScheduleList struct {
	GroupID   int              `json:"groupid"`
	Items     []*power.HashMap `json:"items"`
	YearMonth int              `json:"yearmonth"`
}
