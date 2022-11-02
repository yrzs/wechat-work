package linkedCorp

import (
	"github.com/yrzs/wechat-work/kernel"
	"github.com/yrzs/wechat-work/kernel/power"
	"github.com/yrzs/wechat-work/work/message/linkedCorp/response"
)

type Client struct {
	*kernel.BaseClient
}

func NewClient(app kernel.ApplicationInterface) (*Client, error) {
	baseClient, err := kernel.NewBaseClient(&app, nil)
	if err != nil {
		return nil, err
	}
	return &Client{
		baseClient,
	}, nil
}

// 应用推送消息
// https://developer.work.weixin.qq.com/document/path/90250
func (comp *Client) Send(messages *power.HashMap) (*response.ResponseLinkCorpMessageSend, error) {

	result := &response.ResponseLinkCorpMessageSend{}

	_, err := comp.HttpPostJson("cgi-bin/linkedcorp/message/send", messages, nil, nil, result)

	return result, err
}
