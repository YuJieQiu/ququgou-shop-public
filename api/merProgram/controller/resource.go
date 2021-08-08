package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/ququgou-shop/api/merProgram/config"
	"github.com/ququgou-shop/modules/resource"
	"github.com/ququgou-shop/modules/resource/model"
	"net/http"
)

var uploadResource=resource.UploadResource{
	DB:getConnDB(),
	Config:&config.Config.UploadConfigForImage,
}

var moduleResource=resource.ModuleResource{
	DB:getConnDB(),
}

//上传文件资源 支持 图片和视频 资源
func UploadFiles(c *gin.Context) {

	var (
		res []model.Resource
	)

	form, _ := c.MultipartForm()


	if form==nil {
		JSON(c, http.StatusBadRequest, "参数解析失败", nil)
		return
	}

	files := form.File["file"]

	if len(files)<=0 {
		JSON(c, http.StatusBadRequest, "参数解析失败", nil)
		return
	}

	for _, file := range files {

		err,data,bl:=uploadResource.UploadFile(file)

		if err!=nil {
			JSON(c, http.StatusBadRequest, "", err.Error())
			return
		}

		//if bl == true ,resource exist
		if !bl {
			if err:=moduleResource.Create(data);err!=nil{
				JSON(c, http.StatusBadRequest, "", err.Error())
				return
			}
		}

		if resource.UploadStoreType(config.Config.UploadConfigForImage.Type)==resource.QiNiuYunStorage {//qiniuyun
			data.Url=config.Config.ImgService.QiniuUrl+data.Url
		}else {
			data.Url=config.Config.ImgService.Url+data.Url
		}

		data.Path=data.Path+data.FileName
		res= append(res, *data)
	}
	JSON(c, http.StatusOK, "", res)
	return
}