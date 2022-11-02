package contentSecurity

import (
	"github.com/yrzs/wechat-work/kernel"
)

func RegisterProvider(app kernel.ApplicationInterface) (*Client, error) {

	return NewClient(&app)

}
