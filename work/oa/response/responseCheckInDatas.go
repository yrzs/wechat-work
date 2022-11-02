package response

import (
	"github.com/yrzs/wechat-work/kernel/power"
	"github.com/yrzs/wechat-work/kernel/response"
)

type ResponseCheckInDatas struct {
	*response.ResponseWork

	Datas []*power.HashMap `json:"datas"`
}
