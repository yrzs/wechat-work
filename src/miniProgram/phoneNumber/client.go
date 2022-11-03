package phoneNumber

import (
	"github.com/ArtisanCloud/PowerLibs/v2/object"
	"github.com/yrzs/wechat-work/kernel"
	"github.com/yrzs/wechat-work/src/miniProgram/phoneNumber/response"
)

type Client struct {
	*kernel.BaseClient
}

// code换取用户手机号
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/phonenumber/phonenumber.getPhoneNumber.html
func (comp *Client) GetUserPhoneNumber(code string) (*response.ResponseGetUserPhoneNumber, error) {

	result := &response.ResponseGetUserPhoneNumber{}
	_, err := comp.HttpPostJson("wxa/business/getuserphonenumber", &object.HashMap{
		"code": code,
	}, nil, nil, result)

	return result, err
}
