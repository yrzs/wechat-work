package batchJobs

import (
	"github.com/ArtisanCloud/PowerLibs/v2/object"
	"github.com/yrzs/wechat-work/kernel"
	"github.com/yrzs/wechat-work/kernel/power"
	"github.com/yrzs/wechat-work/work/user/response"
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

// 增量更新成员
// https://developer.work.weixin.qq.com/document/path/90980
func (comp *Client) SyncUser(mediaID string, toInvite bool, callback *power.StringMap) (*response.ResponseUserBatchJobs, error) {

	result := &response.ResponseUserBatchJobs{}

	options := &object.HashMap{
		"media_id":  mediaID,
		"to_invite": toInvite,
		"callbacks": callback,
	}
	_, err := comp.HttpPostJson("cgi-bin/batch/syncuser", options, nil, nil, result)

	return result, err
}

// 全量覆盖成员
// https://developer.work.weixin.qq.com/document/path/90981
func (comp *Client) ReplaceUser(mediaID string, toInvite bool, callback *power.StringMap) (*response.ResponseUserBatchJobs, error) {

	result := &response.ResponseUserBatchJobs{}

	options := &object.HashMap{
		"media_id":  mediaID,
		"to_invite": toInvite,
		"callbacks": callback,
	}
	_, err := comp.HttpPostJson("cgi-bin/batch/replaceuser", options, nil, nil, result)

	return result, err
}

// 全量覆盖部门
// https://developer.work.weixin.qq.com/document/path/90982
func (comp *Client) ReplaceParty(mediaID string, toInvite bool, callback *power.StringMap) (*response.ResponseUserBatchJobs, error) {

	result := &response.ResponseUserBatchJobs{}

	options := &object.HashMap{
		"media_id":  mediaID,
		"to_invite": toInvite,
		"callbacks": callback,
	}
	_, err := comp.HttpPostJson("cgi-bin/batch/replaceparty", options, nil, nil, result)

	return result, err
}

// 获取异步任务结果
// https://developer.work.weixin.qq.com/document/path/90983
func (comp *Client) GetBatchResult(jobID string) (*response.ResponseUserBatchGetResult, error) {

	result := &response.ResponseUserBatchGetResult{}

	params := &object.StringMap{
		"jobid": jobID,
	}
	_, err := comp.HttpGet("cgi-bin/batch/getresult", params, nil, result)

	return result, err
}
