package response

import "github.com/yrzs/wechat-work/kernel/response"

type ResponseExpressGetQuota struct {
	*response.ResponseMiniProgram

	QuotaNum string `json:"quota_num"`
}
