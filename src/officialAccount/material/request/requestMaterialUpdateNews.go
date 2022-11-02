package request

import "github.com/yrzs/wechat-work/kernel/power"

type RequestMaterialUpdateNews struct {
	MediaID  int64            `json:"media_id"`
	Index    int64            `json:"index"`
	Articles []*power.HashMap `json:"articles"`
}
