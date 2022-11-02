package response

import "github.com/yrzs/wechat-work/kernel/response"

type ResponseGetList struct {
	*response.ResponseWork

	ExternalUserID []string `json:"external_userid"` // ["woAJ2GCAAAXtWyujaWJHDDGi0mACAAA","wmqfasd1e1927831291723123109rAAA"]
}
