package card

import "github.com/yrzs/wechat-work/kernel"

func RegisterProvider(app kernel.ApplicationInterface) (*Client, error) {

	baseClient, err := kernel.NewBaseClient(&app, nil)
	if err != nil {
		return nil, err
	}
	client := &Client{
		BaseClient: baseClient,
	}

	return client, nil
}
