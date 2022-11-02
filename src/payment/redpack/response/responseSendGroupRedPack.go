package response

import (
	"encoding/xml"
	"github.com/yrzs/wechat-work/kernel/response"
)

type ResponseSendGroupRedPack struct {
	response.ResponsePayment

	XMLName     xml.Name `xml:"xml"`
	Text        string   `xml:",chardata"`
	ReturnCode  string   `xml:"return_code"`
	ReturnMsg   string   `xml:"return_msg"`
	ResultCode  string   `xml:"result_code"`
	ErrCode     string   `xml:"err_code"`
	ErrCodeDes  string   `xml:"err_code_des"`
	MchBillNO   string   `xml:"mch_billno"`
	MchID       string   `xml:"mch_id"`
	WXAppID     string   `xml:"wxappid"`
	ReOpenID    string   `xml:"re_openid"`
	TotalAmount string   `xml:"total_amount"`
	SendTime    string   `xml:"send_time"`
	SendListID  string   `xml:"send_listid"`
}
