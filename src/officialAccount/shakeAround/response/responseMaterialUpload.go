package response

import "github.com/yrzs/wechat-work/kernel/response"

type DataMaterialUpload struct {
	PicURL string `json:"pic_url"`
}

type ResponseMaterialUpload struct {
	*response.ResponseOfficialAccount

	Data DataMaterialUpload `json:"data"`
}
