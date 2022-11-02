package user

import (
	"github.com/yrzs/wechat-work/kernel"
	"github.com/yrzs/wechat-work/work/user/batchJobs"
	"github.com/yrzs/wechat-work/work/user/exportJobs"
	"github.com/yrzs/wechat-work/work/user/linkedCorp"
	"github.com/yrzs/wechat-work/work/user/tag"
)

func RegisterProvider(app kernel.ApplicationInterface) (
	*Client,
	*batchJobs.Client,
	*exportJobs.Client,
	*linkedCorp.Client,
	*tag.Client,
	error,
) {
	//config := app.GetConfig()

	client, err := NewClient(app)
	if err != nil {
		return nil, nil, nil, nil, nil, err
	}
	UserBatchJobs, err := batchJobs.NewClient(app)
	if err != nil {
		return nil, nil, nil, nil, nil, err
	}
	UserExportJobs, err := exportJobs.NewClient(app)
	if err != nil {
		return nil, nil, nil, nil, nil, err
	}
	UserLinkedCorp, err := linkedCorp.NewClient(app)
	if err != nil {
		return nil, nil, nil, nil, nil, err
	}
	UserTag, err := tag.NewClient(app)
	if err != nil {
		return nil, nil, nil, nil, nil, err
	}
	return client,
		UserBatchJobs,
		UserExportJobs,
		UserLinkedCorp,
		UserTag,
		nil
}
