package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ququgou-shop/api/userProgram/config"
	"github.com/ququgou-shop/modules/product/model"
	"github.com/ququgou-shop/service/productService"
)

// GetProductCategoryList 获取商品分类列表
// @Summary 获取商品分类列表
// @Description 获取商品分类列表
// @Accept json
// @Produce json
// @Param body body productService.GetCategoryListModel true "body参数"
// @Success 200 {object} Response{data=[]model.Category} "ok" "返回信息"
// @Failure 400 {object} Response "错误"
// @Failure 401 {object} Response "错误"
// @Failure 500 {object} Response "错误"
// @Router /category/get/list [get]
func GetProductCategoryList(c *gin.Context) {
	var (
		req = productService.GetCategoryListModel{}
	)

	req.IsSystem = true

	_ = model.Category{}

	err, data, count := productService.GetCategoryList(getConnDB(), &req, false, config.Config.ImgService.QiniuUrl)

	if err != nil {
		JSON(c, http.StatusBadRequest, "", err)
		return
	}

	res := productService.CategoryChildLoad(getConnDB(), data)

	JSONPage(c, http.StatusOK, "", res, count)
	return
}

// GetCategoryProductList 获取分类商品列表
// @Summary 简单描述
// @Description 描述
// @Accept json
// @Produce json
// @Param body body productService.GetCategoryProductListModel true "body参数"
// @Success 200 {object} Response{data=[]productService.ProductSmallInfoModel} "ok" "返回信息"
// @Failure 400 {object} Response "错误"
// @Failure 401 {object} Response "错误"
// @Failure 500 {object} Response "错误"
// @Router /category/product/list [get]
func GetCategoryProductList(c *gin.Context) {
	var (
		req productService.GetCategoryProductListModel
	)

	if err := c.Bind(&req); err != nil {
		//解析失败
		JSON(c, http.StatusBadRequest, "参数解析失败", err)
		return
	}

	err, data, count := productService.GetCategoryProductList(getConnDB(), &req, config.Config.ImgService.QiniuUrl)

	if err != nil {
		JSON(c, http.StatusInternalServerError, "", err)
		return
	}

	JSONPage(c, http.StatusOK, "", data, count)
	return
}
