package response

import (
	"github.com/yrzs/wechat-work/kernel/response"
)

type ResponseOCRIDCard struct {
	*response.ResponseMiniProgram
	Type        string `json:"type"`
	Name        string `json:"name"`
	ID          string `json:"id"`
	Addr        string `json:"addr"`
	Gender      string `json:"gender"`
	Nationality string `json:"nationality"`
	ValidDate   string `json:"valid_date"`
}
