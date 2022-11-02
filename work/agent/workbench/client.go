package workbench

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

// https://developer.work.weixin.qq.com/document/path/90205
func (comp *Client) SetWorkbenchTemplate(data *request.RequestSetWorkbenchTemplate) (*response.ResponseAgentSetWorkbenchTemplate, error) {

	result := &response.ResponseAgentSetWorkbenchTemplate{}

	_, err := comp.HttpPostJson("cgi-bin/agent/set_workbench_template", data, nil, nil, result)

	return result, err
}

// https://developer.work.weixin.qq.com/document/path/90206
func (comp Client) GetWorkbenchTemplate(agentID int) (*response.ResponseAgentGetWorkbenchTemplate, error) {

	result := &response.ResponseAgentGetWorkbenchTemplate{}

	_, err := comp.HttpGet("cgi-bin/agent/get_workbench_template", &object.StringMap{
		"agentid": fmt.Sprintf("%d", agentID),
	}, nil, result)

	return result, err
}

// https://developer.work.weixin.qq.com/document/path/92535
func (comp Client) SetWorkbenchData(data *request.RequestSetWorkBenchData) (*response.ResponseAgentSetWorkbenchData, error) {

	result := &response.ResponseAgentSetWorkbenchData{}

	params, err := object.StructToHashMap(data)
	if err != nil {
		return nil, err
	}
	_, err = comp.HttpPost("cgi-bin/agent/set_workbench_data", params, nil, result)

	return result, err
}
