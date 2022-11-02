package models

import (
	"github.com/yrzs/wechat-work/kernel/contract"
	"github.com/yrzs/wechat-work/kernel/models"
)

const (
	CALLBACK_EVENT_CHANGE_TYPE_CREATE_PARTY = "create_party"
	CALLBACK_EVENT_CHANGE_TYPE_UPDATE_PARTY = "update_party"
	CALLBACK_EVENT_CHANGE_TYPE_DELETE_PARTY = "delete_party"
)

type EventPartyCreate struct {
	contract.EventInterface
	models.CallbackMessageHeader
	ID       string `xml:"Id"`
	Name     string `xml:"Name"`
	ParentID string `xml:"ParentId"`
	Order    string `xml:"Order"`
}

type EventPartyUpdate struct {
	contract.EventInterface
	models.CallbackMessageHeader
	ID       string `xml:"Id"`
	Name     string `xml:"Name"`
	ParentID string `xml:"ParentId"`
}

type EventPartyDelete struct {
	contract.EventInterface
	models.CallbackMessageHeader
	ID string `xml:"Id"`
}
