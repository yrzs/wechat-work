package ocr

import (
	"errors"
	"fmt"
	"github.com/ArtisanCloud/PowerLibs/v2/object"
	"github.com/yrzs/wechat-work/kernel"
	"github.com/yrzs/wechat-work/src/officialAccount/ocr/response"
)

type Client struct {
	*kernel.BaseClient

	AllowTypes []string
}

func NewClient(app kernel.ApplicationInterface) (*Client, error) {
	baseClient, err := kernel.NewBaseClient(&app, nil)
	if err != nil {
		return nil, err
	}
	return &Client{
		BaseClient: baseClient,
		AllowTypes: []string{"photo", "scan"},
	}, nil
}

// 身份证 OCR 识别接口
// https://developers.weixin.qq.com/doc/offiaccount/Intelligent_Interface/OCR.html
func (comp *Client) IDCard(path string, IDType string) (*response.ResponseOCRIDCard, error) {

	if !object.ContainsString(comp.AllowTypes, IDType) {
		return nil, errors.New(fmt.Sprintf("Unsupported media type: '%s'", IDType))
	}

	result := &response.ResponseOCRIDCard{}

	params := &object.HashMap{
		"type":    IDType,
		"img_url": path,
	}

	_, err := comp.HttpPost("cv/ocr/idcard", params, nil, result)

	return result, err
}

// 银行卡 OCR 识别接口
// https://developers.weixin.qq.com/doc/offiaccount/Intelligent_Interface/OCR.html
func (comp *Client) BankCard(path string) (*response.ResponseOCRBankCard, error) {

	result := &response.ResponseOCRBankCard{}

	params := &object.HashMap{
		"img_url": path,
	}

	_, err := comp.HttpPost("cv/ocr/bankcard", params, nil, result)

	return result, err
}

// 驾驶证 OCR 识别接口
// https://developers.weixin.qq.com/doc/offiaccount/Intelligent_Interface/OCR.html
func (comp *Client) VehicleLicense(path string) (*response.ResponseOCRVehicleLicense, error) {

	result := &response.ResponseOCRVehicleLicense{}

	params := &object.HashMap{
		"img_url": path,
	}

	_, err := comp.HttpPost("cv/ocr/drivinglicense", params, nil, result)

	return result, err
}

// 驾驶证 OCR 识别接口
// https://developers.weixin.qq.com/doc/offiaccount/Intelligent_Interface/OCR.html
func (comp *Client) Driving(path string) (*response.ResponseOCRDriving, error) {

	result := &response.ResponseOCRDriving{}

	params := &object.HashMap{
		"img_url": path,
	}

	_, err := comp.HttpPost("cv/ocr/driving", params, nil, result)

	return result, err
}

// 营业执照 OCR 识别接口
// https://developers.weixin.qq.com/doc/offiaccount/Intelligent_Interface/OCR.html
func (comp *Client) BizLicense(path string) (*response.ResponseOCRBizLicense, error) {

	result := &response.ResponseOCRBizLicense{}

	params := &object.HashMap{
		"img_url": path,
	}

	_, err := comp.HttpPost("cv/ocr/bizlicense", params, nil, result)

	return result, err
}

// 通用印刷体 OCR 识别接口
// https://developers.weixin.qq.com/doc/offiaccount/Intelligent_Interface/OCR.html
func (comp *Client) Common(path string) (*response.ResponseOCRCommon, error) {

	result := &response.ResponseOCRCommon{}

	params := &object.HashMap{
		"img_url": path,
	}

	_, err := comp.HttpPost("cv/ocr/comm", params, nil, result)

	return result, err
}

// 车牌识别接口
// https://developers.weixin.qq.com/doc/offiaccount/Intelligent_Interface/OCR.html
func (comp *Client) PlateNumber(path string) (*response.ResponseOCRPlateNumber, error) {

	result := &response.ResponseOCRPlateNumber{}

	params := &object.HashMap{
		"img_url": path,
	}

	_, err := comp.HttpPost("cv/ocr/platenum", params, nil, result)

	return result, err
}
