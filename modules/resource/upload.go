//资源上传
package resource

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"path/filepath"
	"strings"

	"github.com/jinzhu/gorm"
	"github.com/ququgou-shop/modules/resource/model"
	"github.com/ququgou-shop/modules/resource/resourceEnum"
)

type FileUpload interface {
	//文件上传
	UploadFile(reader *bytes.Reader) (error, *model.Resource)
}

//上传存储空间类型
type UploadStoreType int

const (
	LocalityService UploadStoreType = iota + 1 //从1 开始 增长  本地服务器

	QiNiuYunStorage //2			 七牛云存储
)

//上传配置
type UploadConfigModel struct {
	Path             string          `json:"path"`         //保存路径
	Type             UploadStoreType `json:"type"`         // 1、locality service (本地服务器) 2、七牛云存储
	HostAddress      string          `json:"host_address"` //远程主机地址
	LimitSize        int64           `json:"limit_size"`   //限制大小
	LimitTypes       []string        `json:"limit_types"`  //限制类型
	QiNiuConfigModel UploadConfigQiNiuModel
}

//七牛云上传文件配置
type UploadConfigQiNiuModel struct {
	AccessKey string `json:"accessKey"`
	SecretKey string `json:"secretKey"`
	Bucket    string `json:"bucket"`
}

type UploadResource struct {
	Config *UploadConfigModel
	DB     *gorm.DB
}

func NewUploadResource(conf *UploadConfigModel) *UploadResource {
	return &UploadResource{Config: conf}
}

//检查是否实现接口
var _ FileUpload = (*FileUploadStoreLocal)(nil)

//TODO:file compression
func (u *UploadResource) UploadFile(file *multipart.FileHeader) (error, *model.Resource, bool) {
	var up FileUpload

	//file check
	if file == nil {
		return ErrFileEmpty, nil, false
	}

	if u.Config.LimitSize > 0 && file.Size > u.Config.LimitSize {
		return ErrFileExceedSize, nil, false
	}

	ext := filepath.Ext(file.Filename)

	ext = strings.Replace(ext, " ", "", -1)

	var ex bool = false

	if len(u.Config.LimitTypes) > 0 {
		for _, v := range u.Config.LimitTypes {
			if v == ext {
				ex = true
				break
			}
		}
		if !ex {
			return ErrFileTypeError, nil, false
		}
	}

	//file check exist
	fo, err := file.Open()
	defer fo.Close()

	if err != nil {
		return &UploadError{Err: err}, nil, false
	}

	newBytes, _ := ioutil.ReadAll(fo)
	mReader := bytes.NewReader(newBytes)
	cReader := bytes.NewReader(newBytes)

	var md5code string
	md5h := md5.New()
	_, err = io.Copy(md5h, mReader)
	if err != nil {
		return &UploadError{Err: err}, nil, false
	}
	md5code = fmt.Sprintf("%x", md5h.Sum([]byte("")))
	/*hasher:= sha256.New()
	io.Copy(hasher, b)
	fmt.Printf("%x", hasher.Sum([]byte(""))) //hasher*/

	//md5 [查询数据库]是否存在
	bl, err, rr := exist(u.DB, md5code)
	if err != nil {
		return &UploadError{Err: err}, nil, false
	} else if bl {
		return nil, rr, true
	}

	switch UploadStoreType(u.Config.Type) {
	case LocalityService:
		up = FileUploadStoreLocal{
			DB:           u.DB,
			Path:         u.Config.Path,
			Ext:          ext,
			ResourceType: resourceEnum.ResourceTypeImage,
			Md5code:      md5code,
			FileSize:     file.Size,
		}
		break
	case QiNiuYunStorage:
		up = FileUploadStoreQiniu{
			DB:           u.DB,
			Path:         u.Config.Path,
			Ext:          ext,
			ResourceType: resourceEnum.ResourceTypeImage,
			Md5code:      md5code,
			FileSize:     file.Size,
			AccessKey:    u.Config.QiNiuConfigModel.AccessKey,
			SecretKey:    u.Config.QiNiuConfigModel.SecretKey,
			Bucket:       u.Config.QiNiuConfigModel.Bucket,
		}
		break
	default:
		return ErrInvalidStorageType, nil, false
	}

	err, res := up.UploadFile(cReader)
	return err, res, false

}

func (u *UploadResource) UploadFileForBytes(b []byte, ext string) (error, *model.Resource) {
	var up FileUpload
	if len(b) <= 0 {
		return ErrByteNullError, nil
	}
	size := int64(len(b))
	var ex bool = false
	if len(u.Config.LimitTypes) > 0 {
		for _, v := range u.Config.LimitTypes {
			if v == ext {
				ex = true
				break
			}
		}
		if !ex {
			return ErrFileTypeError, nil
		}
	}

	mReader := bytes.NewReader(b)
	cReader := bytes.NewReader(b)

	var md5code string
	md5h := md5.New()
	_, err := io.Copy(md5h, mReader)
	if err != nil {
		return &UploadError{Err: err}, nil
	}

	md5code = fmt.Sprintf("%x", md5h.Sum([]byte("")))

	//md5 [查询数据库]是否存在
	bl, err, rr := exist(u.DB, md5code)
	if err != nil {
		return &UploadError{Err: err}, nil
	} else if bl {
		return nil, rr
	}

	switch UploadStoreType(u.Config.Type) {
	case LocalityService:
		up = FileUploadStoreLocal{
			DB:           u.DB,
			Path:         u.Config.Path,
			Ext:          ext,
			ResourceType: resourceEnum.ResourceTypeImage,
			Md5code:      md5code,
			FileSize:     size,
		}
		break
	case QiNiuYunStorage:
		up = FileUploadStoreQiniu{
			DB:           u.DB,
			Path:         u.Config.Path,
			Ext:          ext,
			ResourceType: resourceEnum.ResourceTypeImage,
			Md5code:      md5code,
			FileSize:     size,
			AccessKey:    u.Config.QiNiuConfigModel.AccessKey,
			SecretKey:    u.Config.QiNiuConfigModel.SecretKey,
			Bucket:       u.Config.QiNiuConfigModel.Bucket,
		}
		break
	default:
		return ErrInvalidStorageType, nil
	}

	err, res := up.UploadFile(cReader)
	return err, res
}
