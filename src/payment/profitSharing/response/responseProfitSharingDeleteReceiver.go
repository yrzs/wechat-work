package response

import "github.com/yrzs/wechat-work/kernel/response"

// https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter8_1_9.shtml

type ResponseProfitSharingDeleteReceiver struct {
	response.ResponsePayment

	Type    string `json:"type"`
	Account string `json:"accountService"`
}
