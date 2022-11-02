package meeting

import (
	"github.com/ArtisanCloud/PowerLibs/v2/object"
	"github.com/yrzs/wechat-work/kernel"
	response2 "github.com/yrzs/wechat-work/kernel/response"
	"github.com/yrzs/wechat-work/work/oa/meeting/request"
	"github.com/yrzs/wechat-work/work/oa/meeting/response"
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

// 创建预约会议
// https://developer.work.weixin.qq.com/document/path/93627
func (comp *Client) Create(options *request.RequestMeetingCreate) (*response.ResponseMeetingCreate, error) {

	result := &response.ResponseMeetingCreate{}

	_, err := comp.HttpPostJson("cgi-bin/meeting/create", options, nil, nil, result)

	return result, err
}

// 修改预约会议
// https://developer.work.weixin.qq.com/document/path/93631
func (comp *Client) Update(options *request.RequestMeetingUpdate) (*response2.ResponseWork, error) {

	result := &response2.ResponseWork{}

	_, err := comp.HttpPostJson("cgi-bin/meeting/update", options, nil, nil, result)

	return result, err
}

// 取消预约会议
// https://developer.work.weixin.qq.com/document/path/93630
func (comp *Client) Cancel(meetingID string) (*response2.ResponseWork, error) {

	result := &response2.ResponseWork{}

	options := &object.HashMap{
		"meetingid": meetingID,
	}

	_, err := comp.HttpPostJson("cgi-bin/meeting/cancel", options, nil, nil, result)

	return result, err
}

// 获取成员会议ID列表
// https://developer.work.weixin.qq.com/document/path/93628
func (comp *Client) GetUserMeetingID(userID string, cursor string, beginTime int64, endTime int64, limit int) (*response.ResponseMeetingGetUserMeetingID, error) {

	result := &response.ResponseMeetingGetUserMeetingID{}

	options := &object.HashMap{
		"userid":     userID,
		"cursor":     cursor,
		"begin_time": beginTime,
		"end_time":   endTime,
		"limit":      limit,
	}

	_, err := comp.HttpPostJson("cgi-bin/meeting/get_user_meetingid", options, nil, nil, result)

	return result, err
}

// 获取会议详情
// https://developer.work.weixin.qq.com/document/path/93629
func (comp *Client) GetBookingInfo(meetingID string) (*response.ResponseMeetingGetBookingInfo, error) {

	result := &response.ResponseMeetingGetBookingInfo{}

	options := &object.HashMap{
		"meetingid": meetingID,
	}

	_, err := comp.HttpPostJson("cgi-bin/meeting/get_info", options, nil, nil, result)

	return result, err
}
