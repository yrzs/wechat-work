package living

import (
	"github.com/ArtisanCloud/PowerLibs/v2/object"
	"github.com/yrzs/wechat-work/kernel"
	response2 "github.com/yrzs/wechat-work/kernel/response"
	"github.com/yrzs/wechat-work/work/oa/living/request"
	"github.com/yrzs/wechat-work/work/oa/living/response"
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

// 创建预约直播
// https://developer.work.weixin.qq.com/document/path/93637
func (comp *Client) Create(options *request.RequestLivingCreate) (*response.ResponseLivingCreate, error) {

	result := &response.ResponseLivingCreate{}

	_, err := comp.HttpPostJson("cgi-bin/living/create", options, nil, nil, result)

	return result, err
}

// 修改预约直播
// https://developer.work.weixin.qq.com/document/path/93640
func (comp *Client) Modify(options *request.RequestLivingModify) (*response2.ResponseWork, error) {

	result := &response2.ResponseWork{}

	_, err := comp.HttpPostJson("cgi-bin/living/modify", options, nil, nil, result)

	return result, err
}

// 取消预约直播
// https://developer.work.weixin.qq.com/document/path/93638
func (comp *Client) Cancel(livingID string) (*response2.ResponseWork, error) {

	result := &response2.ResponseWork{}

	options := &object.HashMap{
		"livingid": livingID,
	}

	_, err := comp.HttpPostJson("cgi-bin/living/cancel", options, nil, nil, result)

	return result, err
}

// 删除直播回放
// https://developer.work.weixin.qq.com/document/path/93874
func (comp *Client) DeleteReplayData(livingID string) (*response2.ResponseWork, error) {

	result := &response2.ResponseWork{}

	options := &object.HashMap{
		"livingid": livingID,
	}

	_, err := comp.HttpPostJson("cgi-bin/living/delete_replay_data", options, nil, nil, result)

	return result, err
}

// 在微信中观看直播或直播回放
// https://developer.work.weixin.qq.com/document/path/93641
func (comp *Client) GetLivingCode(livingID string, openID string) (*response.ResponseLivingGetLivingCode, error) {

	result := &response.ResponseLivingGetLivingCode{}

	options := &object.HashMap{
		"livingid": livingID,
		"openid":   openID,
	}

	_, err := comp.HttpPostJson("cgi-bin/living/get_living_code", options, nil, nil, result)

	return result, err
}

// 获取成员直播ID列表
// https://developer.work.weixin.qq.com/document/path/93634
func (comp *Client) GetUserAllLivingID(userID string, cursor string, limit int) (*response.ResponseLivingGetUserAllLivingID, error) {

	result := &response.ResponseLivingGetUserAllLivingID{}

	options := &object.HashMap{
		"userid": userID,
		"cursor": cursor,
		"limit":  limit,
	}

	_, err := comp.HttpPostJson("cgi-bin/living/get_user_all_livingid", options, nil, nil, result)

	return result, err
}

// 获取直播详情
// https://developer.work.weixin.qq.com/document/path/93635
func (comp *Client) GetLivingInfo(livingID string) (*response.ResponseLivingGetLivingInfo, error) {

	result := &response.ResponseLivingGetLivingInfo{}

	options := &object.HashMap{
		"livingid": livingID,
	}

	_, err := comp.HttpPostJson("cgi-bin/living/get_living_info", options, nil, nil, result)

	return result, err
}

// 获取直播观看明细
// https://developer.work.weixin.qq.com/document/path/93636
func (comp *Client) GetWatchStat(livingID string, nextKey string) (*response.ResponseLivingGetWatchStat, error) {

	result := &response.ResponseLivingGetWatchStat{}

	options := &object.HashMap{
		"livingid": livingID,
		"next_key": nextKey,
	}

	_, err := comp.HttpPostJson("cgi-bin/living/get_watch_stat", options, nil, nil, result)

	return result, err
}

// 获取跳转小程序商城的直播观众信息
// https://developer.work.weixin.qq.com/document/path/94442
func (comp *Client) GetLivingShareInfo(wwShareCode string) (*response.ResponseLivingGetLivingShareInfo, error) {

	result := &response.ResponseLivingGetLivingShareInfo{}

	options := &object.HashMap{
		"ww_share_code": wwShareCode,
	}

	_, err := comp.HttpPostJson("cgi-bin/living/get_living_share_info", options, nil, nil, result)

	return result, err
}
