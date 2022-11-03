package reverse

import (
	"github.com/yrzs/wechat-work/src/payment/kernel"
)

func RegisterProvider(app kernel.ApplicationPaymentInterface) (*Client, error) {

	return NewClient(&app)

}
