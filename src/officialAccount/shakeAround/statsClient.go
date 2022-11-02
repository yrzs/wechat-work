package shakeAround

import (
	"github.com/ArtisanCloud/PowerLibs/v2/object"
	"github.com/yrzs/wechat-work/kernel"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/officialAccount/shakeAround/request"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/officialAccount/shakeAround/response"
)

type StatsClient struct {
	*kernel.BaseClient
}

func NewStatsClient(app kernel.ApplicationInterface) (*StatsClient, error) {
	baseClient, err := kernel.NewBaseClient(&app, nil)
	if err != nil {
		return nil, err
	}
	return &StatsClient{
		baseClient,
	}, nil
}

// 查询单个设备进行摇周边操作的人数、次数，点击摇周边消息的人数、次数
// https://developers.weixin.qq.com/doc/offiaccount/Shake_Nearby/Analytics/Using_devices_as_a_dimension_for_the_data_statistics_interface.html
func (comp *StatsClient) DeviceSummary(data *request.RequestDeviceIdentifier, beginTime int, endTime int) (*response.ResponseStatsSummary, error) {

	result := &response.ResponseStatsSummary{}

	params := &object.HashMap{
		"device_identifier": data,
		"begin_date":        beginTime,
		"end_date":          endTime,
	}
	_, err := comp.HttpPostJson("shakearound/statistics/device", params, nil, nil, result)

	return result, err
}

// 查询指定时间商家帐号下的每个设备进行摇周边操作的人数、次数，点击摇周边消息的人数、次数
// https://developers.weixin.qq.com/doc/offiaccount/Shake_Nearby/Analytics/Batch_query_device_statistics_data_interface.html
func (comp *StatsClient) devicesSummary(timestamp int, pageIndex int) (*response.ResponseStatsDeviceList, error) {

	result := &response.ResponseStatsDeviceList{}

	params := &object.HashMap{
		"date":       timestamp,
		"page_index": pageIndex,
	}
	_, err := comp.HttpPostJson("shakearound/statistics/devicelist", params, nil, nil, result)

	return result, err
}

// 查询单个页面通过摇周边摇出来的人数、次数，点击摇周边页面的人数、次数
// https://developers.weixin.qq.com/doc/offiaccount/Shake_Nearby/Analytics/Using_pages_as_a_dimension_for_the_data_statistics_interface.html
func (comp *StatsClient) PageSummary(pageID int, beginTime int, endTime int) (*response.ResponseStatsPage, error) {

	result := &response.ResponseStatsPage{}

	params := &object.HashMap{
		"page_id":    pageID,
		"begin_date": beginTime,
		"end_date":   endTime,
	}
	_, err := comp.HttpPostJson("shakearound/statistics/page", params, nil, nil, result)

	return result, err
}

// 查询指定时间商家帐号下的每个设备进行摇周边操作的人数、次数，点击摇周边消息的人数、次数
// https://developers.weixin.qq.com/doc/offiaccount/Shake_Nearby/Analytics/Batch_Query_API_for_Page_Statistics.html
func (comp *StatsClient) pagesSummary(timestamp int, pageIndex int) (*response.ResponseStatsPageList, error) {

	result := &response.ResponseStatsPageList{}

	params := &object.HashMap{
		"date":       timestamp,
		"page_index": pageIndex,
	}
	_, err := comp.HttpPostJson("shakearound/statistics/pagelist", params, nil, nil, result)

	return result, err
}
