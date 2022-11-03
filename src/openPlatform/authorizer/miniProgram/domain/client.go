package domain

import (
	"github.com/ArtisanCloud/PowerLibs/v2/object"
	"github.com/yrzs/wechat-work/kernel"
	response2 "github.com/yrzs/wechat-work/kernel/response"
	"github.com/yrzs/wechat-work/src/openPlatform/authorizer/miniProgram/domain/request"
	"github.com/yrzs/wechat-work/src/openPlatform/authorizer/miniProgram/domain/response"
)

type Client struct {
	*kernel.BaseClient
}

func NewClient(app *kernel.ApplicationInterface) (*Client, error) {
	baseClient, err := kernel.NewBaseClient(app, nil)
	if err != nil {
		return nil, err
	}
	return &Client{
		baseClient,
	}, nil
}

// 设置服务器域名
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/Mini_Program_Basic_Info/Server_Address_Configuration.html#请求地址
func (comp *Client) Modify(params *request.RequestModify) (*response.ResponseModify, error) {

	result := &response.ResponseModify{}

	_, err := comp.HttpPostJson("wxa/modify_domain", params, nil, nil, result)

	return result, err

}

// 设置业务域名
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/Mini_Program_Basic_Info/setwebviewdomain.html#请求地址
func (comp *Client) SetWebviewDomain(domains []string, action string) (*response2.ResponseOpenPlatform, error) {

	result := &response2.ResponseOpenPlatform{}

	params := &object.HashMap{
		"action":        action,
		"webviewdomain": domains,
	}
	_, err := comp.HttpPostJson("wxa/setwebviewdomain", params, nil, nil, result)

	return result, err

}
