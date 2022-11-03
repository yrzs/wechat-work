package uniformMessage

import (
	"github.com/yrzs/wechat-work/kernel"
	response2 "github.com/yrzs/wechat-work/kernel/response"
	"github.com/yrzs/wechat-work/src/miniProgram/uniformMessage/request"
)

type Client struct {
	*kernel.BaseClient
}

// 下发小程序和公众号统一的服务消息
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/uniform-message/uniformMessage.send.html
func (comp *Client) Send(options *request.RequestUniformMessageSend) (*response2.ResponseMiniProgram, error) {

	result := &response2.ResponseMiniProgram{}

	_, err := comp.HttpPostJson("cgi-bin/message/wxopen/template/uniform_send", options, nil, nil, result)

	return result, err
}
