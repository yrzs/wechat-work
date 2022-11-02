package response

import (
	"github.com/yrzs/wechat-work/kernel/response"
)

type ResponseUserExportJobs struct {
	*response.ResponseWork

	JobID string `json:"jobid"`
}
