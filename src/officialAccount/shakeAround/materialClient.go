package shakeAround

import (
	"errors"
	"fmt"
	"github.com/ArtisanCloud/PowerLibs/v2/object"
	"github.com/yrzs/wechat-work/kernel"
	"github.com/yrzs/wechat-work/src/officialAccount/shakeAround/response"
	"os"
	"strings"
)

type MaterialClient struct {
	*kernel.BaseClient
}

func NewMaterialClient(app kernel.ApplicationInterface) (*MaterialClient, error) {
	baseClient, err := kernel.NewBaseClient(&app, nil)
	if err != nil {
		return nil, err
	}
	return &MaterialClient{
		baseClient,
	}, nil
}

// 上传在摇一摇功能中需使用到的图片素材
// https://developers.weixin.qq.com/doc/offiaccount/Shake_Nearby/Uploading_Image_Assets.html
func (comp *MaterialClient) UploadImage(path string, Type string) (*response.ResponseMaterialUpload, error) {

	result := &response.ResponseMaterialUpload{}

	_, err := os.Stat(path)
	if (err != nil && os.IsExist(err)) && (err != nil && os.IsPermission(err)) {
		return nil, errors.New(fmt.Sprintf("File does not exist, or the file is unreadable: \"%s\"", path))
	}

	var files *object.HashMap
	if path != "" {
		files = &object.HashMap{
			"media": path,
		}
	} else {
		return nil, errors.New("path is empty")
		
	}

	_, err = comp.HttpUpload("shakearound/material/add", files, nil, &object.StringMap{"type": strings.ToLower(Type)}, nil, result)
	return result, err
}
