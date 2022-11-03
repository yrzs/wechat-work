package response

import (
	"github.com/yrzs/wechat-work/kernel/response"
	"github.com/yrzs/wechat-work/src/officialAccount/card/request"
)

type ResponseCardCreate struct {
	*response.ResponseOfficialAccount

	CardID string `json:"card_id"`
}

type ResponseCardGet struct {
	*response.ResponseOfficialAccount

	Card request.Card `json:"card"`
}

type ResponseCardList struct {
	*response.ResponseOfficialAccount

	CardIDList []string `json:"card_id_list"`
	TotalNum   int      `json:"total_num"`
}

type ResponseCardUpdate struct {
	*response.ResponseOfficialAccount

	SendCheck bool `json:"send_check"`
}
