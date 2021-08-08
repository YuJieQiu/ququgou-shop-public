package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ququgou-shop/api/userProgram/config"
	"github.com/ququgou-shop/modules/appConfig/model"
	bannerModule "github.com/ququgou-shop/modules/banner/model"
	"github.com/ququgou-shop/service/appConfigService"
	"github.com/ququgou-shop/service/productService"
)

// GetBannerList 获取首页Banner列表
// @Summary 获取首页Banner列表
// @Description 获取首页Banner列表 (TODO：返回参数)
// @Accept json
// @Produce json
// @Param body body appConfigService.GetBannerListModel true "body参数"
// @Success 200 {object} Response "ok" "返回信息"
// @Failure 400 {object} Response "错误"
// @Failure 401 {object} Response "错误"
// @Failure 500 {object} Response "错误"
// @Router /home/banner/getList [get]
func GetBannerList(c *gin.Context) {
	var (
		req appConfigService.GetBannerListModel
	)
	_ = bannerModule.Banner{}

	if err := c.Bind(&req); err != nil {
		//解析失败
		JSON(c, http.StatusBadRequest, "参数解析失败", err.Error())
		return
	}

	err, data, _ := appConfigService.GetBannerList(getConnDB(), &req, false, config.Config.ImgService.QiniuUrl)

	if err != nil {
		JSON(c, http.StatusInternalServerError, "", err.Error())
		return
	}

	JSON(c, http.StatusOK, "", data)
	return
}

type GetAppConfigListResponseModel struct {
	List     map[string][]interface{} `json:"list"`
	TextList map[string]string        `json:"textList"`
}

// GetAppConfigList 获取app配置列表
// @Summary 获取app配置列表
// @Description app 配置比如首页分类相关等配置
// @Accept json
// @Produce json
// @Param body body appConfigService.GetAppConfigListModel true "body参数"
// @Success 200 {object} Response{data=GetAppConfigListResponseModel} "ok" "返回信息"
// @Failure 400 {object} Response "错误"
// @Failure 401 {object} Response "错误"
// @Failure 500 {object} Response "错误"
// @Router /home/config/list [get]
func GetAppConfigList(c *gin.Context) {

	var (
		req               appConfigService.GetAppConfigListModel
		appConfigTypeList *[]model.AppConfigType
	)

	res := GetAppConfigListResponseModel{}
	res.List = make(map[string][]interface{})
	res.TextList = make(map[string]string)

	if err := c.Bind(&req); err != nil {
		//解析失败
		JSON(c, http.StatusBadRequest, "参数解析失败", err.Error())
		return
	}

	err, data, count := appConfigService.GetAppConfigMapList(getConnDB(), &req, true, config.Config.ImgService.QiniuUrl)

	if err != nil {
		JSON(c, http.StatusInternalServerError, "", err.Error())
		return
	}

	err, appConfigTypeList = appConfigService.GetAppConfigTypeList(getConnDB())
	if err == nil && appConfigTypeList != nil && len(*appConfigTypeList) > 0 {
		for _, i := range *appConfigTypeList {
			res.TextList[i.Name] = i.Text
		}
	}

	//设置第一个hometba的商品数据
	if _, ok := data["hometba"]; ok && len(data["hometba"]) > 0 {

		c := make(chan *[]productService.ProductSmallInfoModel)
		cErr := make(chan error)
		defer close(c)
		defer close(cErr)

		go appConfigService.GetProductInfoListForConfigId(getConnDB(), data["hometba"][0].ID, c, cErr, config.Config.ImgService.QiniuUrl)

		data["hometba"][0].ProductSmallInfos = <-c
		err = <-cErr
	}

	for i, v := range data {
		ee := make([]interface{}, len(v))
		for j, k := range v {
			ee[j] = k
		}
		res.List[i] = ee
	}

	//获取首页热搜配置

	err, hotSearchList := appConfigService.GetHomeHotSearchList(getConnDB())

	if hotSearchList != nil && len(hotSearchList) > 0 {
		hotSearchText := make([]interface{}, len(hotSearchList))
		for i, v := range hotSearchList {
			hotSearchText[i] = v.Text
		}
		res.List["hot_search_text"] = hotSearchText
	}

	JSONPage(c, http.StatusOK, "", res, count)
	return
}

// GetHomeProductConfigList 获取首页产品配置列表
// @Summary 获取首页产品配置列表
// @Description 获取首页产品配置列表
// @Accept json
// @Produce json
// @Param body body appConfigService.GetHomeProductConfigListModel true "body参数"
// @Success 200 {object} Response{data=[]model.HomeProductConfig} "ok" "返回信息"
// @Failure 400 {object} Response "错误"
// @Failure 401 {object} Response "错误"
// @Failure 500 {object} Response "错误"
// @Router /home/productConfig/getList [get]
func GetHomeProductConfigList(c *gin.Context) {

	var (
		req appConfigService.GetHomeProductConfigListModel
	)

	if err := c.Bind(&req); err != nil {
		//解析失败
		JSON(c, http.StatusBadRequest, "参数解析失败", err.Error())
		return
	}

	err, data, count := appConfigService.GetHomeProductConfigList(getConnDB(), &req, true)

	if err != nil {
		JSON(c, http.StatusInternalServerError, "", err.Error())
		return
	}

	JSONPage(c, http.StatusOK, "", data, count)
	return
}

// GetHomeConfigProductInfoList 获取首页配置产品信息列表
// @Summary 获取首页配置产品信息列表
// @Description 获取首页配置产品信息列表
// @Accept json
// @Produce json
// @Param body body appConfigService.GetHomeConfigProductInfoModel true "body参数"
// @Success 200 {object} Response{data=productService.ProductSmallInfoModel} "ok" "返回信息"
// @Failure 400 {object} Response "错误"
// @Failure 401 {object} Response "错误"
// @Failure 500 {object} Response "错误"
// @Router /home/config/product/list [get]
func GetHomeConfigProductInfoList(c *gin.Context) {

	var (
		req appConfigService.GetHomeConfigProductInfoModel
	)

	if err := c.Bind(&req); err != nil {
		//解析失败
		JSON(c, http.StatusBadRequest, "参数解析失败", err.Error())
		return
	}

	err, data, count := appConfigService.GetHomeConfigProductInfoList(getConnDB(), &req, true, config.Config.ImgService.QiniuUrl)

	if err != nil {
		JSON(c, http.StatusInternalServerError, "", err.Error())
		return
	}

	JSONPage(c, http.StatusOK, "", data, count)
	return
}

type GetSearchHotListModel struct {
	Text string `json:"text"`
	Id   uint64 `json:"id"`
}

// GetSearchHotList 获取热门搜索列表
// @Summary 简单描述
// @Description 描述
// @Accept json
// @Produce json
// @Success 200 {object} Response{data=[]GetSearchHotListModel} "ok" "返回信息"
// @Failure 400 {object} Response "错误"
// @Failure 401 {object} Response "错误"
// @Failure 500 {object} Response "错误"
// @Router /hot/search/get [get]
func GetSearchHotList(c *gin.Context) {

	var data []GetSearchHotListModel

	//获取热搜词
	_, hotSearchList := appConfigService.GetHotSearchList(getConnDB())

	if hotSearchList != nil && len(hotSearchList) > 0 {
		for _, v := range hotSearchList {
			data = append(data, GetSearchHotListModel{Text: v.Text, Id: v.ID})
		}
	}

	JSON(c, http.StatusOK, "", data)
	return
}

// GetHomeConfigData 获取首页 hometba 配置 ?
// @Summary 获取首页 hometba 配置
// @Description 待确定
// @Accept json
// @Produce json
// @Success 200 {object} Response{data=[]model.AppConfig} "ok" "返回信息"
// @Failure 400 {object} Response "错误"
// @Failure 401 {object} Response "错误"
// @Failure 500 {object} Response "错误"
// @Router /wechat/getHomeData [get]
func GetHomeConfigData(c *gin.Context) {

	var (
		parm appConfigService.GetAppConfigListModel
	)
	//get home-tab config
	parm.Type = "hometba"
	err, list, _ := appConfigService.GetAppConfigList(getConnDB(), &parm, false, config.Config.ImgService.QiniuUrl)
	if err != nil {
		JSON(c, http.StatusBadRequest, "", err)
		return
	}

	JSON(c, http.StatusOK, "", list)
	return
}
