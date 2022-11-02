package meetingroom

import (
	"github.com/yrzs/wechat-work/kernel"
	response2 "github.com/yrzs/wechat-work/kernel/response"
	"github.com/yrzs/wechat-work/work/oa/meetingroom/request"
	"github.com/yrzs/wechat-work/work/oa/meetingroom/response"
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

// 添加会议室
// https://developer.work.weixin.qq.com/document/path/93619
func (comp *Client) Add(options *request.RequestMeetingRoomAdd) (*response.ResponseMeetingRoomAdd, error) {

	result := &response.ResponseMeetingRoomAdd{}

	_, err := comp.HttpPostJson("cgi-bin/oa/meetingroom/add", options, nil, nil, result)

	return result, err
}

// 查询会议室
// https://developer.work.weixin.qq.com/document/path/93619
func (comp *Client) List(options *request.RequestMeetingRoomList) (*response.ResponseMeetingRoomList, error) {

	result := &response.ResponseMeetingRoomList{}

	_, err := comp.HttpPostJson("cgi-bin/oa/meetingroom/list", options, nil, nil, result)

	return result, err
}

// 编辑会议室
// https://developer.work.weixin.qq.com/document/path/93619
func (comp *Client) Edit(options *request.RequestMeetingRoomEdit) (*response2.ResponseWork, error) {

	result := &response2.ResponseWork{}

	_, err := comp.HttpPostJson("cgi-bin/oa/meetingroom/edit", options, nil, nil, result)

	return result, err
}

// 删除会议室
// https://developer.work.weixin.qq.com/document/path/93619
func (comp *Client) Del(options *request.RequestMeetingRoomDel) (*response2.ResponseWork, error) {

	result := &response2.ResponseWork{}

	_, err := comp.HttpPostJson("cgi-bin/oa/meetingroom/del", options, nil, nil, result)

	return result, err
}

// 查询会议室的预定信息
// https://developer.work.weixin.qq.com/document/path/93620#%E6%9F%A5%E8%AF%A2%E4%BC%9A%E8%AE%AE%E5%AE%A4%E7%9A%84%E9%A2%84%E5%AE%9A%E4%BF%A1%E6%81%AF
func (comp *Client) GetBookingInfo(options *request.RequestMeetingRoomGetBookingInfo) (*response.ResponseMeetingRoomGetBookingInfo, error) {

	result := &response.ResponseMeetingRoomGetBookingInfo{}

	_, err := comp.HttpPostJson("cgi-bin/oa/meetingroom/get_booking_info", options, nil, nil, result)

	return result, err
}

// 预定会议室
// https://developer.work.weixin.qq.com/document/path/93620
func (comp *Client) Book(options *request.RequestMeetingRoomBook) (*response.ResponseMeetingRoomGetBook, error) {

	result := &response.ResponseMeetingRoomGetBook{}

	_, err := comp.HttpPostJson("cgi-bin/oa/meetingroom/book", options, nil, nil, result)

	return result, err
}

// 取消预定会议室
// https://developer.work.weixin.qq.com/document/path/93620
func (comp *Client) CancelBook(options *request.RequestMeetingRoomCancelBook) (*response2.ResponseWork, error) {

	result := &response2.ResponseWork{}

	_, err := comp.HttpPostJson("cgi-bin/oa/meetingroom/cancel_book", options, nil, nil, result)

	return result, err
}
