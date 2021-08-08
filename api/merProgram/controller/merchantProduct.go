package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/ququgou-shop/api/merProgram/config"
	"github.com/ququgou-shop/service/paymentService"
	"github.com/ququgou-shop/service/productService"
	"net/http"
)

//获取支付方式列表
func GetPaymentTypeList(c *gin.Context) {
	var (
		req paymentService.GetPaymentTypeListModel
	)

	if err := c.Bind(&req); err != nil {
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

	err, data := paymentService.GetPaymentTypeList(getConnDB(), &req)

	if err != nil {
		JSON(c, http.StatusBadRequest, "", err.Error())
		return
	}

	JSON(c, http.StatusOK, "", data)
	return
}

//创建
func CreateProductMerchant(c *gin.Context) {
	var (
		req productService.CreateProductInfoModel
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

	err, data := productService.CreateProductInfo(getConnDB(), &req)
	if err != nil {
		JSON(c, http.StatusInternalServerError, "", err.Error())
		return
	}

	JSON(c, http.StatusOK, "ok", data)
	return
}

//更新
func UpdateProductMerchant(c *gin.Context) {
	var (
		req productService.UpdateProductInfoModel
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

	err, data := productService.UpdateProductInfo(getConnDB(), &req)
	if err != nil {
		JSON(c, http.StatusInternalServerError, "", err.Error())
		return
	}

	JSON(c, http.StatusOK, "ok", data)
	return
}

//获取商户商品List
func GetMerProductList(c *gin.Context) {
	var (
		req productService.GetProductListModel
	)

	if err := c.Bind(&req); err != nil {
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

	err, data, _ := productService.GetProductList(getConnDB(), &req, true, config.Config.ImgService.QiniuUrl)

	if err != nil {
		JSON(c, http.StatusBadRequest, "", err.Error())
		return
	}

	JSON(c, http.StatusOK, "", data)
	return
}

//获取商户商品信息
func GetMerProductInfo(c *gin.Context) {
	var (
		req productService.GetProductInfoModel
	)

	if err := c.Bind(&req); err != nil {
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

	err, data := productService.GetProductInfo(getConnDB(), &req, config.Config.ImgService.QiniuUrl)

	if err != nil {
		JSON(c, http.StatusBadRequest, "", err.Error())
		return
	}

	JSON(c, http.StatusOK, "", data)
	return
}

//更新商品状态
func UpdateProductStatusMerchant(c *gin.Context) {
	var (
		req productService.UpdateProductStatusModel
	)

	if err := c.Bind(&req); err != nil {
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
	req.UserId = u.UserInfo.ID

	err = productService.UpdateProductStatus(getConnDB(), &req)

	if err != nil {
		JSON(c, http.StatusBadRequest, "", err.Error())
		return
	}

	JSON(c, http.StatusOK, "", "")
	return

}
