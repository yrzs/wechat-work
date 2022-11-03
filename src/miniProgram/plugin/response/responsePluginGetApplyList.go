package response

import (
	"github.com/yrzs/wechat-work/kernel/power"
	"github.com/yrzs/wechat-work/kernel/response"
)

type ResponsePluginGetPluginList struct {
	*response.ResponseMiniProgram
	PluginList []*power.HashMap `json:"plugin_list"`
}
