package messages

import (
	"github.com/yrzs/wechat-work/kernel/power"
)

type Link struct {
	*Message
}

func NewLink(items *power.HashMap) *Link {
	m := &Link{
		NewMessage(items),
	}
	m.Type = "link"

	m.Properties = []string{
		"title",
		"description",
		"url",
	}

	return m
}
