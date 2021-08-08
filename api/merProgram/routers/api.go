package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/ququgou-shop/api/merProgram/config"
	"github.com/ququgou-shop/api/merProgram/controller"
	"github.com/ququgou-shop/api/merProgram/middleware"
)

func API(g *gin.Engine) {
	apiG := g.Group(config.Config.API.RelativePath)

	v1 := apiG.Group("/" + config.Config.API.Version)
	{
		//v1.POST("/admin/AdminUserLogin",controller.AdminUserLogin)
	}

	v1.GET("/get/home/info", middleware.JWTAuth(), middleware.UserMerAuth(), controller.GetMerHomeInfo)

	v1.GET("/get/mer/info", middleware.JWTAuth(), middleware.UserMerAuth(), controller.GetMerchantInfo)

	v1.POST("/mer/info/update", middleware.JWTAuth(), middleware.UserMerAuth(), controller.UpdateMerchantInfo)
	v1.POST("/mer/apply/create", middleware.JWTAuth(), controller.CreateMerApplyInfo)
	v1.GET("/mer/apply", middleware.JWTAuth(), controller.GetMerApplyInfo)
	v1.POST("/mer/apply/auto/verified", middleware.JWTAuth(), controller.AutoMerApplyVerified)

	v1.POST("/file/uploadFiles", middleware.JWTAuth(), controller.UploadFiles)

	v1.POST("/wechat/login", controller.WeChatLogin)

	////label
	v1.GET("/label/get/list", middleware.JWTAuth(), middleware.UserMerAuth(), controller.GetLabelList)
	v1.POST("/label/create", middleware.JWTAuth(), middleware.UserMerAuth(), controller.CreateLabel)

	//product
	v1.POST("/product/create", middleware.JWTAuth(), middleware.UserMerAuth(), controller.CreateProductMerchant)
	v1.POST("/product/update", middleware.JWTAuth(), middleware.UserMerAuth(), controller.UpdateProductMerchant)
	v1.GET("/product/get/list", middleware.JWTAuth(), middleware.UserMerAuth(), controller.GetMerProductList)
	v1.GET("/product/get", middleware.JWTAuth(), middleware.UserMerAuth(), controller.GetMerProductInfo)
	v1.POST("/product/update/status", middleware.JWTAuth(), middleware.UserMerAuth(), controller.UpdateProductStatusMerchant)

	//order
	v1.GET("/order/get/list", middleware.JWTAuth(), middleware.UserMerAuth(), controller.GetMerOrderList)
	v1.GET("/order/get/detail", middleware.JWTAuth(), middleware.UserMerAuth(), controller.GetOrderDetail)
	v1.GET("/order/get/user/info", middleware.JWTAuth(), middleware.UserMerAuth(), controller.GetOrderUserInfo)
	v1.POST("/order/user/success", middleware.JWTAuth(), middleware.UserMerAuth(), controller.MerOrderSuccess)

	//admin Manage
	v1.GET("/banner/get", middleware.JWTAuth(), middleware.UserMerAuth(), controller.GetBannerList)
	v1.POST("/banner/save", middleware.JWTAuth(), middleware.UserMerAuth(), controller.BannerSave)
	v1.GET("/app/module/get", middleware.JWTAuth(), middleware.UserMerAuth(), controller.GetAppConfigList)
	v1.POST("/app/module/save", middleware.JWTAuth(), middleware.UserMerAuth(), controller.AppConfigListSave)
	v1.GET("/product/category/get", middleware.JWTAuth(), middleware.UserMerAuth(), controller.GetSystemProductCategoryList)

	v1.GET("/mer/product/category/get", middleware.JWTAuth(), middleware.UserMerAuth(), controller.GetMerchantCategoryList)
	v1.POST("/mer/product/category/save", middleware.JWTAuth(), middleware.UserMerAuth(), controller.MerchantProductCategoryListSave)

	v1.POST("/product/category/save", middleware.JWTAuth(), middleware.UserMerAuth(), controller.SystemProductCategoryListSave)
	v1.GET("/get/product/pay/type/list", middleware.JWTAuth(), middleware.UserMerAuth(), controller.GetPaymentTypeList)

	v1.GET("/hot/search/get", middleware.JWTAuth(), middleware.UserMerAuth(), controller.GetHotSearchConfigList)
	v1.POST("/hot/search/save", middleware.JWTAuth(), middleware.UserMerAuth(), controller.HotSearchConfigListSave)

	//material 产品素材获取
	//v1.GET("/product/material/search", middleware.JWTAuth(), middleware.UserMerAuth(), controller.SearchProductMaterial)

}
