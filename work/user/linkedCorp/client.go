package linkedCorp

import (
	"github.com/ArtisanCloud/PowerLibs/v2/object"
	"github.com/yrzs/wechat-work/kernel"
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

// 获取应用的可见范围
// https://developer.work.weixin.qq.com/document/path/93172
func (comp *Client) GetPermList() (*response.ResponseLinkCorpGetPermList, error) {

	result := &response.ResponseLinkCorpGetPermList{}

	_, err := comp.HttpPostJson("cgi-bin/linkedcorp/agent/get_perm_list", nil, nil, nil, result)

	return result, err
}

// 获取互联企业成员详细信息
// https://developer.work.weixin.qq.com/document/path/93171
func (comp *Client) GetUser(userID string) (*response.ResponseLinkCorpGetUser, error) {

	result := &response.ResponseLinkCorpGetUser{}

	options := &object.HashMap{
		"userid": userID,
	}

	_, err := comp.HttpPostJson("cgi-bin/linkedcorp/user/get", options, nil, nil, result)

	return result, err
}

// 获取互联企业部门成员
// https://developer.work.weixin.qq.com/document/path/93168
func (comp *Client) GetUserSimpleList(departmentID string, fetchChild bool) (*response.ResponseLinkCorpGetUserSimpleList, error) {

	result := &response.ResponseLinkCorpGetUserSimpleList{}

	options := &object.HashMap{
		"department_id": departmentID,
		"fetch_child":   fetchChild,
	}

	_, err := comp.HttpPostJson("cgi-bin/linkedcorp/user/simplelist", options, nil, nil, result)

	return result, err
}

// 获取互联企业部门成员详情
// https://developer.work.weixin.qq.com/document/path/93169
func (comp *Client) GetUserList(departmentID string, fetchChild bool) (*response.ResponseLinkCorpGetUserList, error) {

	result := &response.ResponseLinkCorpGetUserList{}

	options := &object.HashMap{
		"department_id": departmentID,
		"fetch_child":   fetchChild,
	}

	_, err := comp.HttpPostJson("cgi-bin/linkedcorp/user/list?", options, nil, nil, result)

	return result, err
}

// 获取互联企业部门列表
// https://developer.work.weixin.qq.com/document/path/93170
func (comp *Client) GetDepartmentList(departmentID string) (*response.ResponseLinkCorpGetDepartmentList, error) {

	result := &response.ResponseLinkCorpGetDepartmentList{}

	options := &object.HashMap{
		"department_id": departmentID,
	}

	_, err := comp.HttpPostJson("cgi-bin/linkedcorp/department/list?", options, nil, nil, result)

	return result, err
}
