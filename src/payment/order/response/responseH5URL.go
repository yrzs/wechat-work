package response

import (
	"github.com/yrzs/wechat-work/kernel/response"
)

type ResponseH5URL struct {
	response.ResponsePayment

	H5URL string `json:"h5_url"`
}
