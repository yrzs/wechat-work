package response

import (
	"github.com/yrzs/wechat-work/kernel/response"
)

type ResponseRiskControlGetUserRiskRank struct {
	*response.ResponseMiniProgram
	RiskRank int64 `json:"risk_rank"`
	UnoinID  int64 `json:"unoin_id"`
}
