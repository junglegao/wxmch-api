package wxmch_api

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"image"
	"io"
	"io/ioutil"

	_ "image/jpeg"
	_ "image/png"
)

type MediaUploadRequest struct {
	file io.Reader
}

type MediaUploadResponse struct {
	// 媒体文件标识 Id
	MediaID string `json:"media_id"`
}

type ImageFileSuffix string

const JPGSuffix ImageFileSuffix = "jpg"
const PNGSuffix ImageFileSuffix = "png"
const BMPSuffix ImageFileSuffix = "bmp"

// 图片上传API
func (c MerchantApiClient) MediaUpload (req MediaUploadRequest) (resp MediaUploadResponse, err error) {
	// 图片大小不能超过2M，只支持JPG、BMP、PNG格式,
	fBytes, err := ioutil.ReadAll(req.file)
	if err != nil {
		return
	}
	if len(fBytes)> 2*1024*1024 {
		err = errors.New("图片大小不能超过2M")
	}
	// io reader cannot read multiple times
	_, format, err := image.DecodeConfig(bytes.NewReader(fBytes))
	if err != nil {
		return
	}
	var ctype ContentType
	var suffix ImageFileSuffix
	switch format {
	case "jpg":
		ctype = ContentTypeJPG
		suffix = JPGSuffix
		break
	case "bmp":
		ctype = ContentTypeBMP
		suffix = BMPSuffix
		break
	case "png":
		ctype = ContentTypePNG
		suffix = PNGSuffix
		break
	default:
		err = errors.New(fmt.Sprintf("不支持的图片格式:%s", format))
		return
	}
	res, err := c.doFormUpload(context.Background(),  "/v3/merchant/media/upload", fBytes, "image."+string(suffix), ctype)
	if err != nil {
		return
	}
	err = json.Unmarshal(res, &resp)
	return
}