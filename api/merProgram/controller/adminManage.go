package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/ququgou-shop/api/merProgram/cache"
	"github.com/ququgou-shop/api/merProgram/config"
	"github.com/ququgou-shop/service/appConfigService"
	"github.com/ququgou-shop/service/productService"
	"net/http"
)

func BannerSave(c *gin.Context) {
	var (
		req appConfigService.BannerSaveModel
	)

	if err := c.BindJSON(&req); err != nil {
		//解析失败
		JSON(c, http.StatusBadRequest, "参数解析失败", err.Error())
		return
	}

	err := appConfigService.SaveBanner(getConnDB(),&req)

	if err != nil {
		JSON(c, http.StatusBadRequest, "", err.Error())
		return
	}

	JSON(c, http.StatusOK, "", "")
	return
}

//获取Banner列表
func GetBannerList(c *gin.Context) {
	var (
		req appConfigService.GetBannerListModel
	)

	//if err := c.BindJSON(&req); err != nil {
	//	//解析失败
	//	JSON(c, http.StatusBadRequest, "参数解析失败", err.Error())
	//	return
	//}
	req = appConfigService.GetBannerListModel{}
	err, data, _ := appConfigService.GetBannerList(getConnDB(),&req, false, config.Config.ImgService.QiniuUrl)

	if err != nil {
		JSON(c, http.StatusBadRequest, "", err.Error())
		return
	}

	JSON(c, http.StatusOK, "ok", data)
	return
}

func GetAppConfigList(c *gin.Context) {
	var (
		req appConfigService.GetAppConfigListModel
	)

	if err := c.Bind(&req); err != nil {
		//解析失败
		JSON(c, http.StatusBadRequest, "参数解析失败", err.Error())
		return
	}
	//config.Config.ImgService.QiniuUrl
	err, data, _ := appConfigService.GetAppConfigList(getConnDB(),&req, false, config.Config.ImgService.QiniuUrl)

	if err != nil {
		JSON(c, http.StatusBadRequest, "", err.Error())
		return
	}

	JSON(c, http.StatusOK, "", data)
	return
}

func AppConfigListSave(c *gin.Context) {
	var (
		req appConfigService.AppConfigSaveModel
	)

	if err := c.BindJSON(&req); err != nil {
		//解析失败
		JSON(c, http.StatusBadRequest, "参数解析失败", err.Error())
		return
	}

	err := appConfigService.SaveAppConfig(getConnDB(),&req)

	if err != nil {
		JSON(c, http.StatusBadRequest, "", err.Error())
		return
	}

	JSON(c, http.StatusOK, "", "")
	return
}

//获取系统分类列表
func GetMerchantCategoryList(c *gin.Context) {
	var (
		req productService.GetMerchantCategoryListModel
	)

	if err := c.Bind(&req); err != nil {
		//解析失败
		JSON(c, http.StatusBadRequest, "参数解析失败", err)
		return
	}

	//获取用户信息 后面会进行优化
	err, u := GetCustomUserInfo(c)
	if err != nil {
		JSON(c, http.StatusInternalServerError, "", err.Error())
		return
	}
	//EEEEEEEE

	req.MerId = u.MerInfo.MerId

	err, data, count := productService.GetMerchantCategoryList(getConnDB(),&req, false, config.Config.ImgService.QiniuUrl)

	if err != nil {
		JSON(c, http.StatusBadRequest, "", err)
		return
	}

	res := productService.CategoryChildLoad(getConnDB(),data)

	JSONPage(c, http.StatusOK, "", res, count)
	return
}

//商家商品分类修改保存
func MerchantProductCategoryListSave(c *gin.Context) {
	var (
		req productService.MerchantCategoryListSaveModel
	)

	if err := c.BindJSON(&req); err != nil {
		//解析失败
		JSON(c, http.StatusBadRequest, "参数解析失败", err.Error())
		return
	}

	//获取用户信息 后面会进行优化
	err, u := GetCustomUserInfo(c)
	if err != nil {
		JSON(c, http.StatusInternalServerError, "", err.Error())
		return
	}
	//EEEEEEEE

	req.MerId = u.MerInfo.MerId

	err = productService.MerchantCategoryListSave(getConnDB(),&req)

	if err != nil {
		JSON(c, http.StatusBadRequest, "", err.Error())
		return
	}

	JSON(c, http.StatusOK, "", "")
	return
}

//获取系统分类列表
func GetSystemProductCategoryList(c *gin.Context) {
	var (
		req productService.GetCategoryListModel
	)
	//
	//if err := c.Bind(&req); err != nil {
	//	//解析失败
	//	JSON(c, http.StatusBadRequest, "参数解析失败", err)
	//	return
	//}
	req = productService.GetCategoryListModel{}

	//获取用户信息 后面会进行优化

	//EEEEEEEE

	req.IsSystem = true

	err, data, count := productService.GetCategoryList(getConnDB(),&req, false, config.Config.ImgService.QiniuUrl)

	if err != nil {
		JSON(c, http.StatusBadRequest, "", err)
		return
	}

	res := productService.CategoryChildLoad(getConnDB(),data)

	JSONPage(c, http.StatusOK, "", res, count)
	return
}

//系统分类修改保存

func SystemProductCategoryListSave(c *gin.Context) {
	var (
		req productService.SystemCategoryListSaveModel
	)

	if err := c.BindJSON(&req); err != nil {
		//解析失败
		JSON(c, http.StatusBadRequest, "参数解析失败", err.Error())
		return
	}

	err := productService.SystemCategoryListSave(getConnDB(),&req)

	if err != nil {
		JSON(c, http.StatusBadRequest, "", err.Error())
		return
	}

	JSON(c, http.StatusOK, "", "")
	return
}

//获取热搜配置列表
func GetHotSearchConfigList(c *gin.Context) {

	var (
		req appConfigService.GetHotSearchConfigListModel
	)

	err, data, _ := appConfigService.GetHotSearchConfigList(getConnDB(),&req, false)

	if err != nil {
		JSON(c, http.StatusBadRequest, "", err.Error())
		return
	}

	JSON(c, http.StatusOK, "", data)
	return
}

//保存热搜配置列表
func HotSearchConfigListSave(c *gin.Context) {
	var (
		req appConfigService.HotSearchConfigListSaveModel
	)

	if err := c.BindJSON(&req); err != nil {
		//解析失败
		JSON(c, http.StatusBadRequest, "参数解析失败", err.Error())
		return
	}

	err := appConfigService.HotSearchConfigListSave(getConnDB(),&req)

	if err != nil {
		JSON(c, http.StatusBadRequest, "", err.Error())
		return
	}

	//删除缓存
	cache.WebCache.Del(cache.HomeHotSearch)

	JSON(c, http.StatusOK, "", nil)
	return
}
