package models

import (
	"github.com/yrzs/wechat-work/kernel/contract"
	"github.com/yrzs/wechat-work/kernel/models"
)

const CALLBACK_EVENT_MSGAUDIT_NOTIFY = "msgaudit_notify"

type EventMsgAuditNotify struct {
	contract.EventInterface
	models.CallbackMessageHeader
	AgentID string `xml:"AgentID"`
}
