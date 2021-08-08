package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ququgou-shop/modules/order/orderEnum"
	"github.com/ququgou-shop/service/orderService"
	"github.com/ququgou-shop/service/productService"
)

type UserCenterInfoModel struct {
	OrderWaitCount         int `json:"orderWaitCount"`         //待完成订单
	OrderWaitPayCount      int `json:"orderWaitPayCount"`      //待支付订单
	ProductCollectionCount int `json:"productCollectionCount"` //收藏商品
}

// GetUserCenterInfo 获取用户个人中心信息
// @Summary 获取用户个人中心信息
// @Description 获取用户个人中心信息
// @Accept json
// @Produce json
// @Success 200 {object} Response{data=UserCenterInfoModel} "ok" "返回信息"
// @Failure 400 {object} Response "错误"
// @Failure 401 {object} Response "错误"
// @Failure 500 {object} Response "错误"
// @Router /user/center/get [get]
func GetUserCenterInfo(c *gin.Context) {

	//获取用户信息 后面会进行优化
	err, u := GetCustomUserInfo(c)
	if err != nil {
		JSON(c, http.StatusInternalServerError, "", err.Error())
		return
	}

	data := UserCenterInfoModel{}
	err, waitCount := orderService.GetOrderCountByStatus(getConnDB(), orderEnum.OrderBusinessStatusWaitProcess, 0, u.ID)

	if err != nil {
		if err != nil {
			JSON(c, http.StatusInternalServerError, "", err.Error())
			return
		}
	}

	err, waitPayCount := orderService.GetOrderCountByStatus(getConnDB(), orderEnum.OrderBusinessStatusWaitPay, 0, u.ID)
	if err != nil {
		if err != nil {
			JSON(c, http.StatusInternalServerError, "", err.Error())
			return
		}
	}

	data.OrderWaitCount = waitCount
	data.OrderWaitPayCount = waitPayCount

	err, productCount := productService.GetUserProductCollectionCount(getConnDB(), u.ID)

	data.ProductCollectionCount = productCount

	JSON(c, http.StatusOK, "", data)
	return
}
