package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ququgou-shop/api/userProgram/config"
	"github.com/ququgou-shop/modules/merchant/model"
	"github.com/ququgou-shop/service/merchantService"
	"github.com/ququgou-shop/service/productService"
)

// GetMerchantInfo 获取商户信息
// @Summary 获取商户信息
// @Description 获取商户信息
// @Accept json
// @Produce json
// @Param body body merchantService.GetMerchantInfoModel true "body参数"
// @Success 200 {object} Response{data=merchantService.MerInfoModel} "ok" "返回信息"
// @Failure 400 {object} Response "错误"
// @Failure 401 {object} Response "错误"
// @Failure 500 {object} Response "错误"
// @Router /merchant/get [get]
func GetMerchantInfo(c *gin.Context) {
	var (
		req merchantService.GetMerchantInfoModel
	)

	if err := c.Bind(&req); err != nil {
		//解析失败
		JSON(c, http.StatusBadRequest, "参数解析失败", err.Error())
		return
	}

	err, data := merchantService.GetMerchantInfo(getConnDB(), &req, config.Config.ImgService.QiniuUrl)

	if err != nil {
		JSON(c, http.StatusBadRequest, "", err.Error())
		return
	}

	JSON(c, http.StatusOK, "", data)
	return
}

// GetMerProductInfoList 获取商品列表
// @Summary 获取商品列表
// @Description 获取商品列表
// @Accept json
// @Produce json
// @Param body body GetMerProductInfoListModel true "body参数"
// @Success 200 {object} Response{data=[]productService.ProductSmallInfoModel} "ok" "返回信息"
// @Failure 400 {object} Response "错误"
// @Failure 401 {object} Response "错误"
// @Failure 500 {object} Response "错误"
// @Router /mer/product/get/list [get]
func GetMerProductInfoList(c *gin.Context) {
	var (
		req     GetMerProductInfoListModel
		shopReq productService.GetProductSmallInfoListModel
		mer     model.Merchant
		err     error
	)

	if err = c.Bind(&req); err != nil {
		//解析失败
		JSON(c, http.StatusBadRequest, "参数解析失败", err)
		return
	}

	if len(req.MerCode) > 0 {
		err = getConnDB().Where("guid = ?", req.MerCode).First(&mer).Error
	} else if req.MerId > 0 {
		err = getConnDB().Where("id = ?", req.MerId).First(&mer).Error
	} else {
		//返回
		JSONPage(c, http.StatusOK, "", nil, 0)
		return
	}

	if err != nil {
		JSON(c, http.StatusInternalServerError, "", err)
		return
	}

	shopReq = productService.GetProductSmallInfoListModel{
		MerId: mer.ID,
	}
	shopReq.Page = req.Page
	shopReq.Limit = req.Limit
	shopReq.Offset = req.Offset

	err, data, count := productService.GetProductSmallInfoList(getConnDB(), &shopReq, true, "id", config.Config.ImgService.QiniuUrl)

	if err != nil {
		JSON(c, http.StatusInternalServerError, "", err)
		return
	}

	JSONPage(c, http.StatusOK, "", data, count)
	return
}

// GetMerAddressInfoList 获取商户地址信息
// @Summary 获取商户地址信息
// @Description 获取商户地址信息
// @Accept json
// @Produce json
// @Param body body merchantService.QueryMerchantAddressInfoListModel true "body参数"
// @Success 200 {object} Response{data=[]merchantService.MerchantAddressInfoModel} "ok" "返回信息"
// @Failure 400 {object} Response "错误"
// @Failure 401 {object} Response "错误"
// @Failure 500 {object} Response "错误"
// @Router /mer/addresses [post]
func GetMerAddressInfoList(c *gin.Context) {
	var (
		req merchantService.QueryMerchantAddressInfoListModel
	)

	if err := c.BindJSON(&req); err != nil {
		//解析失败
		JSON(c, http.StatusBadRequest, "参数解析失败", err.Error())
		return
	}

	//获取用户信息 后面会进行优化
	//err, u := GetCustomUserInfo(c)
	//if err != nil {
	//	JSON(c, http.StatusInternalServerError, "", err.Error())
	//	return
	//}

	err, data := merchantService.QueryMerchantAddressInfoList(getConnDB(), &req)

	if err != nil {
		JSON(c, http.StatusInternalServerError, "", err.Error())
		return
	}

	JSON(c, http.StatusOK, "", data)
	return
}

// GetMerchantList 获取商户列表
// @Summary 获取商户列表
// @Description 获取商户列表
// @Accept json
// @Produce json
// @Param body body merchantService.GetMerchantListModel true "body参数"
// @Success 200 {object} Response{data=merchantService.MerchantListModel} "ok" "返回信息"
// @Failure 400 {object} Response "错误"
// @Failure 401 {object} Response "错误"
// @Failure 500 {object} Response "错误"
// @Router /mer/list/search [post]
func GetMerchantList(c *gin.Context) {
	var (
		req merchantService.GetMerchantListModel
	)

	if err := c.BindJSON(&req); err != nil {
		//解析失败
		JSON(c, http.StatusBadRequest, "参数解析失败", err.Error())
		return
	}

	//获取用户信息 后面会进行优化
	//err, u := GetCustomUserInfo(c)
	//	//if err != nil {
	//	//	JSON(c, http.StatusInternalServerError, "", err.Error())
	//	//	return
	//	//}
	err, data, count := merchantService.GetMerchantList(getConnDB(), &req, config.Config.ImgService.QiniuUrl)

	if err != nil {
		JSON(c, http.StatusInternalServerError, "", err.Error())
		return
	}

	JSONPage(c, http.StatusOK, "", data, count)
	return
}

type GetMerchantShopQRCodeModel struct {
	MerId   uint64 `json:"merId" form:"merId"`
	MerCode string `json:"merCode" form:"merCode"`
}

// GetMerchantShopQRCode 获取商户店铺二维码
// @Summary 获取商户店铺二维码
// @Description 描述
// @Accept json
// @Produce json
// @Param body body GetMerchantShopQRCodeModel true "body参数"
// @Success 200 {object} Response{data=GetMerchantShopQRCodeModel} "ok" "返回信息"
// @Failure 400 {object} Response "错误"
// @Failure 401 {object} Response "错误"
// @Failure 500 {object} Response "错误"
// @Router /mer/shopQRCode/get [post]
func GetMerchantShopQRCode(c *gin.Context) {
	//TODO:获取商户店铺二维码
	var (
		req GetMerchantShopQRCodeModel
	)

	if err := c.BindJSON(&req); err != nil {
		//解析失败
		JSON(c, http.StatusBadRequest, "参数解析失败", err.Error())
		return
	}

}
