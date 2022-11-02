package customerStrategy

import (
	"github.com/ArtisanCloud/PowerLibs/v2/object"
	"github.com/yrzs/wechat-work/kernel"
	response2 "github.com/yrzs/wechat-work/kernel/response"
	"github.com/yrzs/wechat-work/work/externalContact/customerStrategy/request"
	"github.com/yrzs/wechat-work/work/externalContact/customerStrategy/response"
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

// 获取规则组列表
// https://developer.work.weixin.qq.com/document/path/94883
func (comp *Client) List(cursor string, limit int64) (*response.ResponseCustomerStrategyList, error) {

	result := &response.ResponseCustomerStrategyList{}

	options := &object.HashMap{
		"cursor": cursor,
		"limit":  limit,
	}

	_, err := comp.HttpPostJson("cgi-bin/externalcontact/customer_strategy/list", options, nil, nil, result)

	return result, err
}

// 获取规则组详情
// https://developer.work.weixin.qq.com/document/path/94883
func (comp *Client) Get(strategyID int64) (*response.ResponseCustomerStrategyGet, error) {

	result := &response.ResponseCustomerStrategyGet{}

	options := &object.HashMap{
		"strategy_id": strategyID,
	}

	_, err := comp.HttpPostJson("cgi-bin/externalcontact/customer_strategy/get", options, nil, nil, result)

	return result, err
}

// 获取规则组管理范围
// https://developer.work.weixin.qq.com/document/path/94883
func (comp *Client) GetRange(strategyID int64, cursor string, limit int64) (*response.ResponseCustomerStrategyGetRange, error) {

	result := &response.ResponseCustomerStrategyGetRange{}

	options := &object.HashMap{
		"strategy_id": strategyID,
		"cursor":      cursor,
		"limit":       limit,
	}

	_, err := comp.HttpPostJson("cgi-bin/externalcontact/customer_strategy/get_range", options, nil, nil, result)

	return result, err
}

// 创建新的规则组
// https://developer.work.weixin.qq.com/document/path/94883
func (comp *Client) Create(options *request.RequestCustomerStrategyCreate) (*response.ResponseCustomerStrategyCreate, error) {

	result := &response.ResponseCustomerStrategyCreate{}

	_, err := comp.HttpPostJson("cgi-bin/externalcontact/customer_strategy/create", options, nil, nil, result)

	return result, err
}

// 编辑规则组及其管理范围
// https://developer.work.weixin.qq.com/document/path/94883
func (comp *Client) Edit(options *request.RequestCustomerStrategyEdit) (*response2.ResponseWork, error) {

	result := &response2.ResponseWork{}

	_, err := comp.HttpPostJson("cgi-bin/externalcontact/customer_strategy/edit", options, nil, nil, result)

	return result, err
}

// 删除规则组
// https://developer.work.weixin.qq.com/document/path/94883
func (comp *Client) Del(strategyID int64) (*response2.ResponseWork, error) {

	result := &response2.ResponseWork{}

	options := &object.HashMap{
		"strategy_id": strategyID,
	}

	_, err := comp.HttpPostJson("cgi-bin/externalcontact/customer_strategy/del", options, nil, nil, result)

	return result, err
}
