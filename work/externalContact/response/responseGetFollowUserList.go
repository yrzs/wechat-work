package response

import "github.com/yrzs/wechat-work/kernel/response"

type ResponseGetFollowUserList struct {
	*response.ResponseWork

	FollowUser []string `json:"follow_user"` // ["zhangsan","tagid2"]
}
