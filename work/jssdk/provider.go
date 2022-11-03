package jssdk

import (
	"github.com/yrzs/wechat-work/src/basicService/jssdk"
	"github.com/yrzs/wechat-work/kernel"
)

func RegisterProvider(app kernel.ApplicationInterface) (*Client, error) {

	jssdkClient, err := jssdk.NewClient(&app)
	if err != nil {
		return nil, err
	}
	return &Client{
		jssdkClient,
	}, nil

}
