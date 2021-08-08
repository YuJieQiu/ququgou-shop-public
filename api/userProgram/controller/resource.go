package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ququgou-shop/api/merProgram/config"
	"github.com/ququgou-shop/modules/resource"
	"github.com/ququgou-shop/modules/resource/model"
)

var uploadResource = resource.UploadResource{
	DB:     getConnDB(),
	Config: &config.Config.UploadConfigForImage,
}

var moduleResource = resource.ModuleResource{
	DB: getConnDB(),
}

// UploadFiles 上传文件资源 支持 图片和视频 资源(暂不使用视频)
// @Summary 上传文件资源 支持 图片和视频 资源
// @Description 上传文件资源 支持 图片和视频 资源
// @Accept  multipart/form-data
// @Produce json
// @Param file formData file true "array files"
// @Success 200 {object} Response{data=[]model.Resource} "ok" "返回信息"
// @Failure 400 {object} Response "错误"
// @Failure 401 {object} Response "错误"
// @Failure 500 {object} Response "错误"
// @Router /resource/uploadFiles [post]
func UploadFiles(c *gin.Context) {
	//TODO:单独拆分成service
	var (
		res []model.Resource
	)

	var uploadResource = resource.UploadResource{}

	//获取用户信息 后面会进行优化
	err, u := GetCustomUserInfo(c)
	if err != nil {
		JSON(c, http.StatusInternalServerError, "", err.Error())
		return
	}

	form, _ := c.MultipartForm()

	if form == nil {
		JSON(c, http.StatusBadRequest, "参数解析失败", nil)
		return
	}

	files := form.File["file"]

	if len(files) <= 0 {
		JSON(c, http.StatusBadRequest, "参数解析失败", nil)
		return
	}

	//conf := config.Config.UploadConfigForImage

	for _, file := range files {

		err, data, bl := uploadResource.UploadFile(file)

		if err != nil {
			JSON(c, http.StatusBadRequest, "", err.Error())
			return
		}

		//if bl == true ,resource exist
		if !bl {
			data.UserId = u.ID
			if err := moduleResource.Create(data); err != nil {
				JSON(c, http.StatusBadRequest, "", err.Error())
				return
			}
		}

		if data.Type == 2 { //qiniuyun  1、locality service (本地服务器) 2、七牛云存储
			data.Url = config.Config.ImgService.QiniuUrl + data.Url
		} else {
			data.Url = config.Config.ImgService.Url + data.Url
		}

		data.Path = data.Path + data.FileName
		res = append(res, *data)
	}

	JSON(c, http.StatusOK, "", res)
	return
}
