package request

import "github.com/yrzs/wechat-work/kernel/power"

type RequestGroupChatList struct {
	StatusFilter int            `json:"status_filter"`
	OwnerFilter  *power.HashMap `json:"owner_filter"`
	Cursor       string         `json:"cursor"`
	Limit        int            `json:"limit"`
}
