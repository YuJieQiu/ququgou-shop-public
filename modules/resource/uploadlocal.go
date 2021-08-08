//上传文件到本地服务器下
package resource

import (
	"bytes"
	"github.com/jinzhu/gorm"
	"github.com/ququgou-shop/library/utils"
	"github.com/ququgou-shop/modules/resource/model"
	"github.com/ququgou-shop/modules/resource/resourceEnum"
	"io"
	"os"
)

type FileUploadStoreLocal struct {
	DB           *gorm.DB                  //资源配置信息
	Path         string                    //保存路径
	Ext          string                    //文件扩展名
	ResourceType resourceEnum.ResourceType //类型
	FileSize     int64
	Md5code      string
}

//file upload
func (u FileUploadStoreLocal) UploadFile(reader *bytes.Reader) (error, *model.Resource) {

	var (
		r   model.Resource
		err error
	)

	//check path exist
	_, err = os.Stat(u.Path)

	if err != nil {
		if !os.IsExist(err) {
			err = os.Mkdir(u.Path, os.ModePerm)
			if err != nil {
				return &UploadError{Err: err}, nil
			}
		}
	}

	//get file guid
	guid := utils.CreateUUID()

	fileName := guid + u.Ext

	saveFilePath := u.Path + fileName

	out, err := os.Create(saveFilePath)
	defer out.Close()

	if err != nil {
		return &UploadError{Err: err}, nil
	}

	_, err = io.Copy(out, reader)
	if err != nil {
		return &UploadError{Err: err}, nil
	}

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
