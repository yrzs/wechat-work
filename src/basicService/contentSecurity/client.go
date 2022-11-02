package contentSecurity

import "github.com/yrzs/wechat-work/kernel"

type Client struct {
	*kernel.BaseClient

	BaseUri string
}

func NewClient(app *kernel.ApplicationInterface) (*Client, error) {

	baseClient, err := kernel.NewBaseClient(app, nil)
	if err != nil {
		return nil, err
	}
	client := &Client{
		BaseClient: baseClient,
	}

	client.BaseClient.HttpRequest.BaseURI = "https://api.weixin.qq.com/wxa/"

	return client, nil
}
