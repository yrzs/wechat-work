package response

import (
	"github.com/yrzs/wechat-work/kernel/response"
	"github.com/yrzs/wechat-work/work/menu/request"
)

type ResponseMenuGet struct {
	*response.ResponseWork

	// 企业微信官方：返回结果与请参考菜单创建接口
	// https://developer.work.weixin.qq.com/document/path/90232
	request.RequestMenuSet
}
