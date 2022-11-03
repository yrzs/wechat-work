package response

import "github.com/yrzs/wechat-work/kernel/response"

type ResponseImmediateDeliveryCancelOrder struct {
	*response.ResponseMiniProgram

	DeductFee int64  `json:"deduct_fee"`
	Desc      string `json:"desc"`
}
