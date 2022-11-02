package agent

import (
	"fmt"
	"github.com/ArtisanCloud/PowerLibs/v2/object"
	"github.com/yrzs/wechat-work/kernel"
	"github.com/yrzs/wechat-work/work/agent/request"
	"github.com/yrzs/wechat-work/work/agent/response"
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

// https://developer.work.weixin.qq.com/document/path/90227
func (comp *Client) Get(agentID int) (*response.ResponseAgentGet, error) {

	result := &response.ResponseAgentGet{}

	_, err := comp.HttpPostJson("cgi-bin/agent/get", nil, &object.StringMap{
		"agentid": fmt.Sprintf("%d", agentID),
	}, nil, result)

	return result, err
}

// https://developer.work.weixin.qq.com/document/path/90228
func (comp *Client) Set(data *request.RequestAgentSet) (*response.ResponseAgentSet, error) {

	result := &response.ResponseAgentSet{}

	_, err := comp.HttpPostJson("cgi-bin/agent/set", data, nil, nil, result)

	return result, err
}

// https://developer.work.weixin.qq.com/document/path/90227
func (comp *Client) List() (*response.ResponseAgentList, error) {

	result := &response.ResponseAgentList{}

	_, err := comp.HttpGet("cgi-bin/agent/list", nil, nil, result)

	return result, err
}
