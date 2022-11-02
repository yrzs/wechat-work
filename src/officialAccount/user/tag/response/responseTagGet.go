package response

import "github.com/yrzs/wechat-work/kernel/response"

type Tag struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Count int    `json:"count"`
}

type ResponseTagGet struct {
	*response.ResponseOfficialAccount

	Tag *Tag `json:"tag"`
}

//---------------------------------------------

type ResponseTagGetList struct {
	*response.ResponseOfficialAccount
	Tags []*Tag `json:"tags"`
}
