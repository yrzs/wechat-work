package aggregate

import (
	"github.com/yrzs/wechat-work/kernel"
	"github.com/yrzs/wechat-work/src/openPlatform/authorizer/aggregate/account"
)

func RegisterProvider(app kernel.ApplicationInterface) (*account.Client, error) {

	//account, err := account.NewClient(app)
	//if err != nil {
	//	return nil, err
	//}
	//
	//return account, nil

	return nil, nil
}
