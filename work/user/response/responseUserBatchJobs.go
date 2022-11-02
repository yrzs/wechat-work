package response

import (
	"github.com/yrzs/wechat-work/kernel/response"
)

type ResponseUserBatchJobs struct {
	*response.ResponseWork

	JobID string `json:"jobid"`
}
