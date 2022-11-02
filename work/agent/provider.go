package agent

import (
	"github.com/yrzs/wechat-work/kernel"
	"github.com/yrzs/wechat-work/work/agent/workbench"
)

func RegisterProvider(app kernel.ApplicationInterface) (*Client, *workbench.Client, error) {

	client, err := NewClient(app)
	if err != nil {
		return nil, nil, err
	}

	Workbench, err := workbench.NewClient(app)
	if err != nil {
		return nil, nil, err
	}

	return client, Workbench, nil

}
