package messages

import (
	"github.com/yrzs/wechat-work/kernel/power"
)

type Image struct {
	*Media
}

func NewImage(mediaID string, items *power.HashMap) *Image {
	m := &Image{
		NewMedia(mediaID, "image", items),
	}

	return m
}
