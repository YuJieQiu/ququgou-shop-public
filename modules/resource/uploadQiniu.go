package resource

import (
	"bytes"
	"context"
	"github.com/jinzhu/gorm"
	"github.com/qiniu/api.v7/auth/qbox"
	"github.com/qiniu/api.v7/storage"
	"github.com/ququgou-shop/library/utils"
	"github.com/ququgou-shop/modules/resource/model"
	"github.com/ququgou-shop/modules/resource/resourceEnum"
)

type FileUploadStoreQiniu struct {
	DB           *gorm.DB                  //资源配置信息
	Path         string                    //保存路径
	Ext          string                    //文件扩展名
	ResourceType resourceEnum.ResourceType //类型
	FileSize     int64
	Md5code      string
	AccessKey    string `json:"accessKey"`
	SecretKey    string `json:"secretKey"`
	Bucket       string `json:"bucket"`
}

func (u FileUploadStoreQiniu) UploadFile(reader *bytes.Reader) (error, *model.Resource) {

	guid := utils.CreateUUID()
	fileName := guid + u.Ext
	saveFilePath := u.Path + fileName

	putPolicy := storage.PutPolicy{
		Scope: u.Bucket,
	}
	mac := qbox.NewMac(u.AccessKey, u.SecretKey)
	upToken := putPolicy.UploadToken(mac)
	cfg := storage.Config{}

	// 空间对应的机房
	cfg.Zone = &storage.ZoneHuadong
	// 是否使用https域名
	cfg.UseHTTPS = false
	// 上传是否使用CDN上传加速
	cfg.UseCdnDomains = false

	// 构建表单上传的对象
	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}

	// 可选配置
	putExtra := storage.PutExtra{
		Params: map[string]string{
			"x:name": fileName,
		},
	}

	err := formUploader.Put(context.Background(), &ret, upToken, saveFilePath, reader, u.FileSize, &putExtra)
	if err != nil {
		return &UploadError{Err: err}, nil
	}

	var r model.Resource

	r.Guid = guid
	r.Type = int(u.ResourceType)
	r.Path = u.Path
	r.Size = u.FileSize
	r.FileName = fileName
	r.Ext = u.Ext
	r.Url = saveFilePath
	r.ContentType = ""
	r.Hosted = ""
	r.HashCode = u.Md5code

	return nil, &r
}
