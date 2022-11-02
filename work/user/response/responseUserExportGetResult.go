package response

import (
	"github.com/ArtisanCloud/PowerLibs/v2/object"
	"github.com/yrzs/wechat-work/kernel/response"
)

type ResponseUserExportGetResult struct {
	*response.ResponseWork

	Status   int               `json:"status"`
	DataList []*object.HashMap `json:"data_list"`
}
