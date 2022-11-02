package message

import (
	"github.com/ArtisanCloud/PowerLibs/v2/object"
	"github.com/yrzs/wechat-work/kernel"
	"github.com/yrzs/wechat-work/work/accountService/message/request"
	"github.com/yrzs/wechat-work/work/accountService/message/response"
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

// 读取消息
// https://developer.work.weixin.qq.com/document/path/94670
func (comp *Client) SyncMsg(cursor string, token string, limit int) (*response.ResponseMessageSyncMsg, error) {

	result := &response.ResponseMessageSyncMsg{}

	options := &object.HashMap{
		"cursor": cursor,
		"token":  token,
		"limit":  limit,
	}

	_, err := comp.HttpPostJson("cgi-bin/kf/sync_msg", options, nil, nil, result)

	return result, err
}

// 发送消息
// https://developer.work.weixin.qq.com/document/path/90236
func (comp *Client) SendMsg(messages *request.RequestAccountServiceSendMsg) (*response.ResponseAccountServiceSendMsg, error) {

	result := &response.ResponseAccountServiceSendMsg{}

	_, err := comp.HttpPostJson("cgi-bin/kf/send_msg", messages, nil, nil, result)

	return result, err
}

// 发送消息
// https://developer.work.weixin.qq.com/document/path/95122
func (comp *Client) SendMsgOnEvent(messages *request.RequestAccountServiceSendMsgOnEvent) (*response.ResponseAccountServiceSendMsgOnEvent, error) {

	result := &response.ResponseAccountServiceSendMsgOnEvent{}

	_, err := comp.HttpPostJson("cgi-bin/kf/send_msg_on_event", messages, nil, nil, result)

	return result, err
}
