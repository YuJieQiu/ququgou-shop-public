package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/ququgou-shop/api/userProgram/config"
	"github.com/ququgou-shop/api/userProgram/controller"
	"github.com/ququgou-shop/api/userProgram/middleware"
)

func API(g *gin.Engine) {

	apiG := g.Group(config.Config.API.RelativePath)

	//本地调用的一些 Api
	//apiLocal := g.Group(config.Config.API.RelativePath)

	v1 := apiG.Group("/" + config.Config.API.Version)
	{
		//v1.POST("/admin/AdminUserLogin",controller.AdminUserLogin)
	}

	//weChat
	v1.POST("/wechat/login", controller.WeChatLogin)

	v1.GET("/wechat/getHomeData", middleware.GetToken(), controller.GetHomeConfigData)

	//home
	v1.GET("/home/banner/getList", middleware.GetToken(), controller.GetBannerList)
	v1.GET("/home/config/list", controller.GetAppConfigList)                                            //获取首页配置列表
	v1.GET("/home/productConfig/getList", middleware.GetToken(), controller.GetHomeProductConfigList)   //获取首页配置产品列表
	v1.GET("/home/config/product/list", middleware.GetToken(), controller.GetHomeConfigProductInfoList) //获取首页配置产品信息列表

	v1.GET("/hot/search/get", middleware.GetToken(), controller.GetSearchHotList) //获取热门搜索

	v1.GET("/shop/product/info/getList", middleware.JWTAuth(), controller.GetProductInfoList)
	v1.GET("/shop/product/detail", middleware.GetToken(), controller.GetProductDetail)

	v1.GET("/shop/product/payment/type/list", middleware.JWTAuth(), controller.GetProductPaymentTypeList) //获取商品的支付方式
	v1.GET("/shop/payment/type/list", middleware.JWTAuth(), controller.GetPaymentTypeList)                //获取支付方式

	v1.POST("/shop/cart/add", middleware.JWTAuth(), controller.ShopCartAdd)               //购物车商品添加
	v1.POST("/shop/cart/remove", middleware.JWTAuth(), controller.ShopCartRemove)         //购物车商品删除
	v1.GET("/shop/cart/get", middleware.JWTAuth(), controller.GetShopCartProductInfoList) //购物车商品信息获取

	v1.POST("/shop/collection/add", middleware.JWTAuth(), controller.UserShopProductCollectionAdd)       //收藏商品添加
	v1.POST("/shop/collection/remove", middleware.JWTAuth(), controller.UserShopProductCollectionRemove) //收藏商品删除
	v1.GET("/shop/collection/get", middleware.JWTAuth(), controller.GetUserProductCollectionList)        //收藏商品获取

	//resource
	v1.POST("/resource/uploadFiles", middleware.JWTAuth(), controller.UploadFiles) //图片资源上传

	//order
	//g.POST("/order/befCreate", c.BefOrderCreate)//创建订单前 暂不使用
	v1.POST("/order/create", middleware.JWTAuth(), controller.UserOrderCreate)       //创建订单
	v1.POST("/order/cancel", middleware.JWTAuth(), controller.OrderCancel)           //订单取消
	v1.POST("/order/pay", middleware.JWTAuth(), controller.OrderPay)                 //订单支付
	v1.GET("/order/get/list", middleware.JWTAuth(), controller.GetUserOrderList)     //查询订单
	v1.GET("/order/get/detail", middleware.JWTAuth(), controller.GetUserOrderDetail) //查询订单详细信息
	//退款

	//address
	v1.POST("/address/create", middleware.JWTAuth(), controller.CreateAddress)
	v1.POST("/address/save", middleware.JWTAuth(), controller.SaveAddress)
	v1.POST("/address/user/delete", middleware.JWTAuth(), controller.DeleteAddress)
	v1.GET("/address/getList", middleware.JWTAuth(), controller.GetAddressList)
	v1.GET("/address/user/getList", middleware.JWTAuth(), controller.GetUserAddressList)
	v1.GET("/address/user/get", middleware.JWTAuth(), controller.GetAddress)
	v1.GET("/address/user/first", middleware.JWTAuth(), controller.GetUserAddressFirst)

	//merchant
	v1.GET("/merchant/get", middleware.GetToken(), controller.GetMerchantInfo)
	v1.GET("/mer/product/get/list", middleware.GetToken(), controller.GetMerProductInfoList)
	v1.POST("/mer/addresses", middleware.GetToken(), controller.GetMerAddressInfoList)
	v1.POST("/mer/list/search", middleware.GetToken(), controller.GetMerchantList)

	v1.GET("/category/get/list", middleware.GetToken(), controller.GetProductCategoryList)     //查询商品分类列表
	v1.GET("/category/product/list", middleware.GetToken(), controller.GetCategoryProductList) //查询分类商品列表

	//TODO:搜索
	v1.GET("/search/get", middleware.GetToken(), controller.TestProductGet)

	v1.GET("/search/query", middleware.GetToken(), controller.TestProductQuery)

	v1.GET("/search/sync", middleware.GetToken(), controller.TestProductSync)

	v1.POST("/product/search", middleware.GetToken(), controller.SearchProductList)
	v1.GET("/product/list/config", middleware.GetToken(), controller.GetProductSearchPageConfigData)

	//User 用户相关
	v1.GET("/user/center/get", middleware.JWTAuth(), controller.GetUserCenterInfo)

	//推荐
	v1.POST("/recommend/product/list", middleware.GetToken(), controller.RecommendProductList)
	//
	////通知
	//wechatNotify := controller.WechatNotifyController{}
	//v1.POST("/wechat/pay/notify", wechatNotify.PaySuccessNotify)
	////同步支付信息
	//v1.GET("/wechat/order/pay/sync", wechatNotify.PayResultSync)
	//
	//////////////本地调用的APi 接口
	//apiLocal.GET("/order/timeout/cancel", order.TimeoutNotPayOrderCancel) //取消超时待支付订单
}
