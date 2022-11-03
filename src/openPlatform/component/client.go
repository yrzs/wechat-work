package component

import (
	"github.com/ArtisanCloud/PowerLibs/v2/object"
	"github.com/yrzs/wechat-work/kernel"
	response2 "github.com/yrzs/wechat-work/kernel/response"
	"github.com/yrzs/wechat-work/src/openPlatform/component/request"
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

// 快速注册企业小程序
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/Register_Mini_Programs/Fast_Registration_Interface_document.html#接口详情
func (comp *Client) RegisterMiniProgram(params *request.RequestRegisterMiniProgram) (*response2.ResponseOpenPlatform, error) {

	result := &response2.ResponseOpenPlatform{}

	_, err := comp.HttpPostJson("cgi-bin/component/fastregisterweapp", params, nil, nil, result)

	return result, err

}

// 查询创建任务状态
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/Register_Mini_Programs/Fast_Registration_Interface_document.html#接口详情
func (comp *Client) GetRegistrationStatus(companyName string, legalPersonaWechat string, legalPersonaName string) (*response2.ResponseOpenPlatform, error) {

	result := &response2.ResponseOpenPlatform{}

	params := &object.HashMap{
		"name":                 companyName,
		"legal_persona_wechat": legalPersonaWechat,
		"legal_persona_name":   legalPersonaName,
	}

	_, err := comp.HttpPostJson("cgi-bin/component/fastregisterweapp", params, nil, nil, result)

	return result, err

}
