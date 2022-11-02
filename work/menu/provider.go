package menu

import (
	"github.com/yrzs/wechat-work/kernel"
)

func RegisterProvider(app kernel.ApplicationInterface) (*Client, error) {

	client, err := NewClient(app)

	return client, err

}
