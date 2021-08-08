package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ququgou-shop/api/userProgram/config"
	"github.com/ququgou-shop/modules/user/model"
	"github.com/ququgou-shop/service/productService"
	"github.com/ququgou-shop/service/searchService"
)

//产品推荐 查询

// RecommendProductList 推荐商品列表
// @Summary 推荐商品列表
// @Description 推荐基本逻辑
//1、根据用户的经纬度
//2、根据用户平时浏览的分类类型
//3、根据商品的推荐权重
//4、根据价格、销量
// @Accept json
// @Produce json
// @Param body body searchService.QueryRecommendProductListModel true "body参数"
// @Success 200 {object} Response{data=[]productService.ProductSmallInfoModel} "ok" "返回信息"
// @Failure 400 {object} Response "错误"
// @Failure 401 {object} Response "错误"
// @Failure 500 {object} Response "错误"
// @Router /recommend/product/list [post]
func RecommendProductList(c *gin.Context) {
	//推荐基本逻辑
	//1、根据用户的经纬度
	//2、根据用户平时浏览的分类类型
	//3、根据商品的推荐权重
	//4、根据价格、销量
	var (
		req searchService.QueryRecommendProductListModel
		err error
		u   *model.User
	)
	_ = productService.ProductSmallInfoModel{}

	if err := c.BindJSON(&req); err != nil {
		//解析失败
		JSON(c, http.StatusBadRequest, "参数解析失败", err.Error())
		return
	}

	if req.Source == 2 || req.Source == 3 {
		//获取用户信息 后面会进行优化
		_, u = GetCustomUserInfo(c)
	}

	err, data, count := searchService.GetRecommendProductListByDB(getConnDB(), &req, u, config.Config.ImgService.QiniuUrl)

	if err != nil {
		JSON(c, http.StatusInternalServerError, "", err)
		return
	}

	JSONPage(c, http.StatusOK, "", data, count)
	return
}
