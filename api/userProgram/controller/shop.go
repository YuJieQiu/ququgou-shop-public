package controller

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ququgou-shop/api/userProgram/cache"
	"github.com/ququgou-shop/api/userProgram/config"
	appConfigModel "github.com/ququgou-shop/modules/appConfig/model"
	paymentModel "github.com/ququgou-shop/modules/payment/model"
	"github.com/ququgou-shop/modules/product/model"
	"github.com/ququgou-shop/service/appConfigService"
	"github.com/ququgou-shop/service/paymentService"
	"github.com/ququgou-shop/service/productService"
	"github.com/ququgou-shop/service/shopCartService"
)

// GetProductInfoList 获取商品信息列表(用户商品列表展示使用)
// @Summary 获取商品信息列表
// @Description 用户商品列表展示使用
// @Accept json
// @Produce json
// @Param body body productService.GetProductSmallInfoListModel true "body参数"
// @Success 200 {object} Response{data=[]productService.ProductSmallInfoModel} "ok" "返回信息"
// @Failure 400 {object} Response "错误"
// @Failure 401 {object} Response "错误"
// @Failure 500 {object} Response "错误"
// @Router /shop/product/info/getList [get]
func GetProductInfoList(c *gin.Context) {
	var (
		req productService.GetProductSmallInfoListModel
	)
	if err := c.Bind(&req); err != nil {
		//解析失败
		JSON(c, http.StatusBadRequest, "参数解析失败", err)
		return
	}

	err, data, count := productService.GetProductSmallInfoList(getConnDB(), &req, true, "id", config.Config.ImgService.QiniuUrl)

	if err != nil {
		JSON(c, http.StatusBadRequest, "", err)
		return
	}

	JSONPage(c, http.StatusOK, "", data, count)
	return
}

// GetProductDetail 获取 单个商品详情 商品详细信息
// @Summary 获取 单个商品详情 商品详细信息
// @Description 获取 单个商品详情 商品详细信息
// @Accept json
// @Produce json
// @Param body body productService.GetProductDetailInfoSingleModel true "body参数"
// @Success 200 {object} Response{data=productService.ProductDetailInfoSingle} "ok" "返回信息"
// @Failure 400 {object} Response "错误"
// @Failure 401 {object} Response "错误"
// @Failure 500 {object} Response "错误"
// @Router /shop/product/detail [get]
func GetProductDetail(c *gin.Context) {
	var (
		req productService.GetProductDetailInfoSingleModel
	)
	if err := c.Bind(&req); err != nil {
		//解析失败
		JSON(c, http.StatusBadRequest, "参数解析失败", err.Error())
		return
	}

	//获取用户信息 后面会进行优化
	err, u := GetCustomUserInfo(c)
	if u != nil && u.ID > 0 {
		req.UserId = u.ID
	}

	err, data := productService.GetProductDetailInfoSingle(getConnDB(), &req, config.Config.ImgService.QiniuUrl)

	if err != nil {
		JSON(c, http.StatusInternalServerError, "", err.Error())
		return
	}

	JSON(c, http.StatusOK, "", data)
	return
}

//Product Info E

// ShopCartAdd 商品添加购物车
// @Summary 商品添加购物车
// @Description 商品添加购物车
//这里有措施防止同一个用户快速添加商品到购物车
// @Accept json
// @Produce json
// @Param body body shopCartService.AddShopCartItemModel true "body参数"
// @Success 200 {object} Response "ok" "返回信息"
// @Failure 400 {object} Response "错误"
// @Failure 401 {object} Response "错误"
// @Failure 500 {object} Response "错误"
// @Router /shop/cart/add [post]
func ShopCartAdd(c *gin.Context) {
	//TODO:购物车返回数据
	var (
		req shopCartService.AddShopCartItemModel
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

	timeoutDuration := 10 * time.Second
	//添加LOCK
	if b := cache.Lock(cache.CartAddLock+u.Guid, "1", timeoutDuration); !b {
		//防止快速 添加...
		JSON(c, http.StatusBadRequest, "", "")
		return
	}

	req.UserId = u.ID
	req.TotalPrice = float64(req.Number) * req.Price

	err, data := shopCartService.AddShopCartItem(getConnDB(), &req)

	if err != nil {
		cache.UnLock(cache.CartAddLock + u.Guid)
		JSON(c, http.StatusInternalServerError, "", err)
		return
	}

	cache.UnLock(cache.CartAddLock + u.Guid)
	JSON(c, http.StatusOK, "", data)
	return
}

// ShopCartRemove 购物车商品删除
// @Summary 购物车商品删除
// @Description 购物车商品删除
// @Accept json
// @Produce json
// @Param body body shopCartService.DeleteShopCartModel true "body参数"
// @Success 200 {object} Response "ok" "返回信息"
// @Failure 400 {object} Response "错误"
// @Failure 401 {object} Response "错误"
// @Failure 500 {object} Response "错误"
// @Router /shop/cart/remove [post]
func ShopCartRemove(c *gin.Context) {
	var (
		req shopCartService.DeleteShopCartModel
	)

	if err := c.BindJSON(&req); err != nil {
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

	req.UserId = u.ID

	err = shopCartService.DeleteShopCart(getConnDB(), &req)

	if err != nil {
		JSON(c, http.StatusInternalServerError, "", err)
		return
	}

	JSON(c, http.StatusOK, "", nil)
	return
}

// GetShopCartProductInfoList 获取购物车商品信息
// @Summary 获取购物车商品信息
// @Description 获取购物车商品信息
// @Accept json
// @Produce json
// @Param body body shopCartService.GetUserShopCartProductInfoModel true "body参数"
// @Success 200 {object} Response{data=[]shopCartService.UserShopCartMerModel} "ok" "返回信息"
// @Failure 400 {object} Response "错误"
// @Failure 401 {object} Response "错误"
// @Failure 500 {object} Response "错误"
// @Router /shop/cart/get [get]
func GetShopCartProductInfoList(c *gin.Context) {
	var (
		req shopCartService.GetUserShopCartProductInfoModel
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

	req.UserId = u.ID

	err, data := shopCartService.GetUserShopCartProductInfo(getConnDB(), &req, config.Config.ImgService.QiniuUrl)

	if err != nil {
		JSON(c, http.StatusBadRequest, "", err)
		return
	}

	JSON(c, http.StatusOK, "", data)
	return
}

// cart

//收藏商品 S

// UserShopProductCollectionAdd 收藏商品
// @Summary 收藏商品
// @Description 收藏商品
// @Accept json
// @Produce json
// @Param body body productService.UserProductCollectionAddModel true "body参数"
// @Success 200 {object} Response "ok" "返回信息"
// @Failure 400 {object} Response "错误"
// @Failure 401 {object} Response "错误"
// @Failure 500 {object} Response "错误"
// @Router /shop/collection/add [post]
func UserShopProductCollectionAdd(c *gin.Context) {
	var (
		req productService.UserProductCollectionAddModel
	)

	if err := c.BindJSON(&req); err != nil {
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

	req.UserId = u.ID

	err = productService.UserProductCollectionAdd(getConnDB(), &req)

	if err != nil {
		JSON(c, http.StatusInternalServerError, "", err)
		return
	}

	JSON(c, http.StatusOK, "", nil)
	return
}

// UserShopProductCollectionRemove 删除收藏商品
// @Summary 删除收藏商品
// @Description 删除收藏商品
// @Accept json
// @Produce json
// @Param body body productService.UserProductCollectionRemoveModel true "body参数"
// @Success 200 {object} Response "ok" "返回信息"
// @Failure 400 {object} Response "错误"
// @Failure 401 {object} Response "错误"
// @Failure 500 {object} Response "错误"
// @Router /shop/collection/remove [post]
func UserShopProductCollectionRemove(c *gin.Context) {
	var (
		req productService.UserProductCollectionRemoveModel
	)

	if err := c.BindJSON(&req); err != nil {
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

	req.UserId = u.ID

	err = productService.UserProductCollectionRemove(getConnDB(), &req)

	if err != nil {
		JSON(c, http.StatusInternalServerError, "", err)
		return
	}

	JSON(c, http.StatusOK, "", nil)
	return
}

// GetUserProductCollectionList 获取用户收藏商品列表
// @Summary 获取用户收藏商品列表
// @Description 获取用户收藏商品列表
// @Accept json
// @Produce json
// @Param body body productService.GetUserProductCollectionListModel true "body参数"
// @Success 200 {object} Response{data=[]productService.ProductSmallInfoModel} "ok" "返回信息"
// @Failure 400 {object} Response "错误"
// @Failure 401 {object} Response "错误"
// @Failure 500 {object} Response "错误"
// @Router /shop/collection/get [get]
func GetUserProductCollectionList(c *gin.Context) {
	var (
		req productService.GetUserProductCollectionListModel
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

	req.UserId = u.ID

	err, data, count := productService.GetUserProductCollectionList(getConnDB(), &req, config.Config.ImgService.QiniuUrl)

	if err != nil {
		JSON(c, http.StatusBadRequest, "", err)
		return
	}

	JSONPage(c, http.StatusOK, "", data, count)
	return
}

//收藏商品 E

// GetProductPaymentTypeList 获取商品支付方式
// @Summary 获取商品支付方式
// @Description 获取商品支付方式
// @Accept json
// @Produce json
// @Param body body productService.GetProductPaymentTypeListModel true "body参数"
// @Success 200 {object} Response "ok" "返回信息"
// @Failure 400 {object} Response "错误"
// @Failure 401 {object} Response "错误"
// @Failure 500 {object} Response "错误"
// @Router /shop/product/payment/type/list [get]
func GetProductPaymentTypeList(c *gin.Context) {
	var (
		req productService.GetProductPaymentTypeListModel
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

	req.UserId = u.ID

	err, data := productService.GetProductPaymentTypeList(getConnDB(), &req)

	if err != nil {
		JSON(c, http.StatusBadRequest, "", err)
		return
	}

	JSON(c, http.StatusOK, "", data)
	return
}

// GetPaymentTypeList 获取支付方式列表
// @Summary 简单描述
// @Description 描述
// @Accept json
// @Produce json
// @Param body body paymentService.GetPaymentTypeListModel true "body参数"
// @Success 200 {object} Response "ok" "返回信息"
// @Failure 400 {object} Response "错误"
// @Failure 401 {object} Response "错误"
// @Failure 500 {object} Response "错误"
// @Router /shop/payment/type/list [post]
func GetPaymentTypeList(c *gin.Context) {
	//TODO:返回参数 paymentModel.PaymentType
	var (
		req paymentService.GetPaymentTypeListModel
	)

	_ = paymentModel.PaymentType{}
	if err := c.Bind(&req); err != nil {
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
	//EEEEEEEE

	err, data := paymentService.GetPaymentTypeList(getConnDB(), &req)

	if err != nil {
		JSON(c, http.StatusBadRequest, "", err.Error())
		return
	}

	JSON(c, http.StatusOK, "", data)
	return
}

//Category

// CreateCategory 创建分类信息
// TODO: 暂无用接口 待处理
func CreateCategory(c *gin.Context) {
	var (
		req model.Category
	)

	if err := c.BindJSON(&req); err != nil {
		//解析失败
		JSON(c, http.StatusBadRequest, "参数解析失败", err)
		return
	}

	err, data := productService.CreateCategory(getConnDB(), &req)

	if err != nil {
		JSON(c, http.StatusBadRequest, "", err)
		return
	}

	JSON(c, http.StatusOK, "ok", data)
	return
}

// EditCategory 编辑分类信息
// TODO: 暂无用接口 待处理
func EditCategory(c *gin.Context) {
	var (
		req model.Category
	)

	if err := c.BindJSON(&req); err != nil {
		//解析失败
		JSON(c, http.StatusBadRequest, "参数解析失败", err)
		return
	}

	err, data := productService.EditCategory(getConnDB(), &req)

	if err != nil {
		JSON(c, http.StatusBadRequest, "", err)
		return
	}

	JSON(c, http.StatusOK, "ok", data)
	return
}

// GetCategoryList 获取分类列表
// TODO: 暂无用接口 待处理
func GetCategoryList(c *gin.Context) {
	var (
		req productService.GetCategoryListModel
	)

	if err := c.Bind(&req); err != nil {
		//解析失败
		JSON(c, http.StatusBadRequest, "参数解析失败", err)
		return
	}

	err, data, count := productService.GetCategoryList(getConnDB(), &req, true, config.Config.ImgService.QiniuUrl)

	if err != nil {
		JSON(c, http.StatusBadRequest, "", err)
		return
	}

	JSONPage(c, http.StatusOK, "", data, count)
	return

}

//Category E

//Attribute

// CreateAttribute 创建 Attribute
// TODO: 暂无用接口 待处理
func CreateAttribute(c *gin.Context) {
	var (
		req model.Attribute
	)

	if err := c.BindJSON(&req); err != nil {
		//解析失败
		JSON(c, http.StatusBadRequest, "参数解析失败", err)
		return
	}

	err, data := productService.CreateAttribute(getConnDB(), &req)

	if err != nil {
		JSON(c, http.StatusBadRequest, "", err)
		return
	}

	JSON(c, http.StatusOK, "ok", data)
	return
}

// EditAttribute 编辑 Attribute
// TODO: 暂无用接口 待处理
func EditAttribute(c *gin.Context) {
	var (
		req model.Attribute
	)

	if err := c.BindJSON(&req); err != nil {
		//解析失败
		JSON(c, http.StatusBadRequest, "参数解析失败", err)
		return
	}

	err, data := productService.EditAttribute(getConnDB(), &req)

	if err != nil {
		JSON(c, http.StatusBadRequest, "", err)
		return
	}

	JSON(c, http.StatusOK, "ok", data)
	return
}

// GetAttributeList 获取 Attribute 列表
// TODO: 暂无用接口 待处理
func GetAttributeList(c *gin.Context) {
	var (
		req productService.GetAttributeListModel
	)

	if err := c.Bind(&req); err != nil {
		//解析失败
		JSON(c, http.StatusBadRequest, "参数解析失败", err)
		return
	}

	err, data, count := productService.GetAttributeList(getConnDB(), &req, true, false)

	if err != nil {
		JSON(c, http.StatusBadRequest, "", err)
		return
	}

	JSONPage(c, http.StatusOK, "", data, count)
	return

}

//Attribute E

//Product Info

// CreateProductInfo 创建 产品信息
// TODO: 暂无用接口 待处理
func CreateProductInfo(c *gin.Context) {
	var (
		req productService.CreateProductInfoModel
	)

	if err := c.BindJSON(&req); err != nil {
		//解析失败
		JSON(c, http.StatusBadRequest, "参数解析失败", err)
		return
	}

	err, data := productService.CreateProductInfo(getConnDB(), &req)

	if err != nil {
		JSON(c, http.StatusBadRequest, "", err)
		return
	}

	//创建商品信息成功后，添加到首页展示中去 Test
	//测试使用
	//appConfig.GetAppConfigList(getConnDB(),&service.GetAppConfigListModel{},false)

	_, aList, _ := appConfigService.GetAppConfigList(getConnDB(), &appConfigService.GetAppConfigListModel{}, true, config.Config.ImgService.QiniuUrl)
	for _, v := range aList {
		appConfigService.CreateHomeProductConfig(getConnDB(), &appConfigModel.HomeProductConfig{
			AppConfigId: v.ID,
			ProductId:   data.ID,
		})
	}

	JSON(c, http.StatusOK, "ok", data)
	return
}

// GetProductList 获取产品列表
// TODO: 暂无用接口 待处理
func GetProductList(c *gin.Context) {
	var (
		req productService.GetProductListModel
	)
	if err := c.Bind(&req); err != nil {
		//解析失败
		JSON(c, http.StatusBadRequest, "参数解析失败", err)
		return
	}

	err, data, count := productService.GetProductList(getConnDB(), &req, true, config.Config.ImgService.QiniuUrl)

	if err != nil {
		JSON(c, http.StatusBadRequest, "", err)
		return
	}

	JSONPage(c, http.StatusOK, "", data, count)
	return

}
