package webdrive

import (
	"github.com/yrzs/wechat-work/kernel"
	response2 "github.com/yrzs/wechat-work/kernel/response"
	"github.com/yrzs/wechat-work/work/oa/webdrive/request"
	"github.com/yrzs/wechat-work/work/oa/webdrive/response"
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

// 新建空间
// https://developer.work.weixin.qq.com/document/path/93655#新建空间
func (comp *Client) SpaceCreate(options *request.RequestWebDriveSpaceCreate) (*response.ResponseWebDriveSpaceCreate, error) {

	result := &response.ResponseWebDriveSpaceCreate{}

	_, err := comp.HttpPostJson("cgi-bin/webdrive/space_create", options, nil, nil, result)

	return result, err
}

// 重命名空间
// https://developer.work.weixin.qq.com/document/path/93655#重命名空间
func (comp *Client) SpaceRename(options *request.RequestWebDriveSpaceRename) (*response2.ResponseWork, error) {

	result := &response2.ResponseWork{}

	_, err := comp.HttpPostJson("cgi-bin/webdrive/space_rename", options, nil, nil, result)

	return result, err
}

// 解散空间
// https://developer.work.weixin.qq.com/document/path/93655#解散空间
func (comp *Client) SpaceDismiss(options *request.RequestWebDriveSpaceDismiss) (*response2.ResponseWork, error) {

	result := &response2.ResponseWork{}

	_, err := comp.HttpPostJson("cgi-bin/webdrive/space_dismiss", options, nil, nil, result)

	return result, err
}

// 获取空间信息
// https://developer.work.weixin.qq.com/document/path/93655#获取空间相册信息
func (comp *Client) SpaceInfo(options *request.RequestWebDriveSpaceInfo) (*response.ResponseWebDriveSpaceInfo, error) {

	result := &response.ResponseWebDriveSpaceInfo{}

	_, err := comp.HttpPostJson("cgi-bin/webdrive/space_info", options, nil, nil, result)

	return result, err
}

// 添加成员/部门
// https://developer.work.weixin.qq.com/document/path/93656#添加成员部门
func (comp *Client) SpaceACLAdd(options *request.RequestWebDriveSpaceACLAdd) (*response2.ResponseWork, error) {

	result := &response2.ResponseWork{}

	_, err := comp.HttpPostJson("cgi-bin/webdrive/space_acl_add", options, nil, nil, result)

	return result, err
}

// 移除成员/部门
// https://developer.work.weixin.qq.com/document/path/93656#移除成员部门
func (comp *Client) SpaceACLDel(options *request.RequestWebDriveSpaceACLDel) (*response2.ResponseWork, error) {

	result := &response2.ResponseWork{}

	_, err := comp.HttpPostJson("cgi-bin/webdrive/space_acl_del", options, nil, nil, result)

	return result, err
}

// 安全设置
// https://developer.work.weixin.qq.com/document/path/93656#安全设置
func (comp *Client) SpaceSetting(options *request.RequestWebDriveSpaceSetting) (*response2.ResponseWork, error) {

	result := &response2.ResponseWork{}

	_, err := comp.HttpPostJson("cgi-bin/webdrive/space_setting", options, nil, nil, result)

	return result, err
}

// 获取邀请链接
// https://developer.work.weixin.qq.com/document/path/93656#获取邀请链接
func (comp *Client) SpaceShare(options *request.RequestWebDriveSpaceShare) (*response.ResponseWebDriveSpaceShare, error) {

	result := &response.ResponseWebDriveSpaceShare{}

	_, err := comp.HttpPostJson("cgi-bin/webdrive/space_share", options, nil, nil, result)

	return result, err
}

// 获取文件列表
// https://developer.work.weixin.qq.com/document/path/93657#获取文件列表
func (comp *Client) FileList(options *request.RequestWebDriveFileList) (*response.ResponseWebDriveFileList, error) {

	result := &response.ResponseWebDriveFileList{}

	_, err := comp.HttpPostJson("cgi-bin/webdrive/file_list", options, nil, nil, result)

	return result, err
}

// 上传文件
// https://developer.work.weixin.qq.com/document/path/93657#上传文件
func (comp *Client) FileUpload(options *request.RequestWebDriveFileUpload) (*response.ResponseWebDriveFileUpload, error) {

	result := &response.ResponseWebDriveFileUpload{}

	_, err := comp.HttpPostJson("cgi-bin/webdrive/file_upload", options, nil, nil, result)

	return result, err
}

// 下载文件
// https://developer.work.weixin.qq.com/document/path/93657#下载文件
func (comp *Client) FileDownload(options *request.RequestWebDriveFileDownload) (*response.ResponseWebDriveFileDownload, error) {

	result := &response.ResponseWebDriveFileDownload{}

	_, err := comp.HttpPostJson("cgi-bin/webdrive/file_download", options, nil, nil, result)

	return result, err
}

// 新建文件/微文档
// https://developer.work.weixin.qq.com/document/path/93657#新建文件文档
func (comp *Client) FileCreate(options *request.RequestWebDriveFileCreate) (*response.ResponseWebDriveFileCreate, error) {

	result := &response.ResponseWebDriveFileCreate{}

	_, err := comp.HttpPostJson("cgi-bin/webdrive/file_create", options, nil, nil, result)

	return result, err
}

// 重命名文件
// https://developer.work.weixin.qq.com/document/path/93657#重命名文件
func (comp *Client) FileRename(options *request.RequestWebDriveFileRename) (*response.ResponseWebDriveFileRename, error) {

	result := &response.ResponseWebDriveFileRename{}

	_, err := comp.HttpPostJson("cgi-bin/webdrive/file_rename", options, nil, nil, result)

	return result, err
}

// 移动文件
// https://developer.work.weixin.qq.com/document/path/93657#移动文件
func (comp *Client) FileMove(options *request.RequestWebDriveFileMove) (*response.ResponseWebDriveFileMove, error) {

	result := &response.ResponseWebDriveFileMove{}

	_, err := comp.HttpPostJson("cgi-bin/webdrive/file_move", options, nil, nil, result)

	return result, err
}

// 删除文件
// https://developer.work.weixin.qq.com/document/path/93657#删除文件
func (comp *Client) FileDelete(options *request.RequestWebDriveFileDelete) (*response2.ResponseWork, error) {

	result := &response2.ResponseWork{}

	_, err := comp.HttpPostJson("cgi-bin/webdrive/file_delete", options, nil, nil, result)

	return result, err
}

// 新增成员
// https://developer.work.weixin.qq.com/document/path/93658
func (comp *Client) FileACLAdd(options *request.RequestWebDriveFileACLAdd) (*response2.ResponseWork, error) {

	result := &response2.ResponseWork{}

	_, err := comp.HttpPostJson("cgi-bin/webdrive/file_acl_add", options, nil, nil, result)

	return result, err
}

// 删除成员
// https://developer.work.weixin.qq.com/document/path/93658#删除成员
func (comp *Client) FileACLDel(options *request.RequestWebDriveFileACLDel) (*response2.ResponseWork, error) {

	result := &response2.ResponseWork{}

	_, err := comp.HttpPostJson("cgi-bin/webdrive/file_acl_del", options, nil, nil, result)

	return result, err
}

// 分享设置
// https://developer.work.weixin.qq.com/document/path/93658#分享设置
func (comp *Client) FileSetting(options *request.RequestWebDriveFileSetting) (*response2.ResponseWork, error) {

	result := &response2.ResponseWork{}

	_, err := comp.HttpPostJson("cgi-bin/webdrive/file_setting", options, nil, nil, result)

	return result, err
}

// 获取分享链接
// https://developer.work.weixin.qq.com/document/path/95860#获取分享链接
func (comp *Client) FileShare(options *request.RequestWebDriveFileShare) (*response.ResponseWebDriveFileShare, error) {

	result := &response.ResponseWebDriveFileShare{}

	_, err := comp.HttpPostJson("cgi-bin/webdrive/file_share", options, nil, nil, result)

	return result, err
}
