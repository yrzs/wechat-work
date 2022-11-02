package customerService

import (
	"github.com/yrzs/wechat-work/kernel"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/officialAccount/customerService/session"
)

func RegisterProvider(app kernel.ApplicationInterface) (*Client, *session.Client, error) {

	client, err := NewClient(app)
	if err != nil {
		return nil, nil, err
	}

	sessionClient, err := session.NewClient(app)
	if err != nil {
		return nil, nil, err
	}

	return client, sessionClient, nil
}
