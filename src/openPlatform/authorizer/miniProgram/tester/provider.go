package tester

import "github.com/yrzs/wechat-work/kernel"

func RegisterProvider(app kernel.ApplicationInterface) (*Client, error) {

	code, err := NewClient(&app)
	if err != nil {
		return nil, err
	}

	return code, nil
}
