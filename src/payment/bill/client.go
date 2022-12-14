package bill

import (
	"github.com/ArtisanCloud/PowerLibs/v2/object"
	"github.com/yrzs/wechat-work/kernel/power"
	"github.com/yrzs/wechat-work/src/payment/bill/response"
	payment "github.com/yrzs/wechat-work/src/payment/kernel"
)

type Client struct {
	*payment.BaseClient
}

func NewClient(app *payment.ApplicationPaymentInterface) (*Client, error) {
	baseClient, err := payment.NewBaseClient(app)
	if err != nil {
		return nil, err
	}
	return &Client{
		baseClient,
	}, nil
}

// Get Trade Bill.
// https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter3_1_6.shtml
func (comp *Client) GetTradeBill(date string, billType string, tarType string) (*response.ResponseBillGet, error) {

	result := &response.ResponseBillGet{}

	params := &object.StringMap{
		"bill_date": date,
		"bill_type": billType,
		"tar_type":  tarType,
	}

	endpoint := comp.Wrap("/v3/bill/tradebill")
	_, err := comp.Request(endpoint, params, "GET", nil, false, nil, result)

	return result, err
}

// Get Flow Bill.
// https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter3_1_7.shtml
func (comp *Client) GetFlowBill(date string, accountType string, tarType string) (*response.ResponseBillGet, error) {

	result := &response.ResponseBillGet{}

	params := &object.StringMap{
		"bill_date":    date,
		"account_type": accountType,
		"tar_type":     tarType,
	}

	endpoint := comp.Wrap("/v3/bill/fundflowbill")
	_, err := comp.Request(endpoint, params, "GET", nil, false, nil, result)

	return result, err
}

// https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter3_1_8.shtml
func (comp *Client) DownloadBill(requestDownload *power.RequestDownload, filePath string) (int64, error) {
	return comp.StreamDownload(requestDownload, filePath)
}
