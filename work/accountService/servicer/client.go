package servicer

import (
	"github.com/ArtisanCloud/PowerLibs/v2/object"
	"github.com/yrzs/wechat-work/kernel"
	"github.com/yrzs/wechat-work/work/accountService/servicer/response"
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

// 添加接待人员
// https://developer.work.weixin.qq.com/document/path/94646
func (comp *Client) Add(openKFID string, userIDList []string) (*response.ResponseServicerAdd, error) {

	result := &response.ResponseServicerAdd{}

	options := &object.HashMap{
		"open_kfid":   openKFID,
		"userid_list": userIDList,
	}

	_, err := comp.HttpPostJson("cgi-bin/kf/servicer/add", options, nil, nil, result)

	return result, err
}

// 删除接待人员
// https://developer.work.weixin.qq.com/document/path/94647
func (comp *Client) Del(openKFID string, userIDList []string) (*response.ResponseServicerDel, error) {

	result := &response.ResponseServicerDel{}

	options := &object.HashMap{
		"open_kfid":   openKFID,
		"userid_list": userIDList,
	}

	_, err := comp.HttpPostJson("cgi-bin/kf/servicer/del", options, nil, nil, result)

	return result, err
}

// 获取接待人员列表
// https://developer.work.weixin.qq.com/document/path/94645
func (comp *Client) List(openKFID string) (*response.ResponseServicerList, error) {

	result := &response.ResponseServicerList{}

	params := &object.StringMap{
		"open_kfid": openKFID,
	}

	_, err := comp.HttpPostJson("cgi-bin/kf/servicer/list", nil, params, nil, result)

	return result, err
}
