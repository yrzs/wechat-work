package response

import (
	"github.com/yrzs/wechat-work/kernel/response"
)

type ResponseHeaderCloseOrdr struct {
	response.ResponsePayment

	Status string `header:"status"`
}
