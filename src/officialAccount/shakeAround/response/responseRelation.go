package response

import "github.com/yrzs/wechat-work/kernel/response"

type Relation struct {
	DeviceID int    `json:"device_id"`
	Major    int    `json:"major"`
	Minor    int    `json:"minor"`
	PageID   int    `json:"page_id"`
	UUID     string `json:"uuid"`
}

type DataRelationSearch struct {
	Relations  []*Relation `json:"relations"`
	TotalCount int         `json:"total_count"`
}

type ResponseRelationSearch struct {
	*response.ResponseOfficialAccount

	Data *DataRelationSearch `json:"data"`
}
