package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/ququgou-shop/modules/order/orderEnum"
	"github.com/ququgou-shop/service/orderService"
	"net/http"
)

type MerHomeBasicInfoModel struct {
	OrderWaitCount int `json:"orderWaitCount"`
}

//get mer home info
func GetMerHomeInfo(c *gin.Context) {

	//获取用户信息 后面会进行优化
	err, u := GetCustomUserInfo(c)
	if err != nil {
		JSON(c, http.StatusInternalServerError, "", err.Error())
		return
	}

	data := MerHomeBasicInfoModel{}
	err, waitCount := orderService.GetOrderCountByStatus(getConnDB(), orderEnum.OrderBusinessStatusWaitProcess, u.MerInfo.MerId, u.UserInfo.ID)
	err, waitShipCount := orderService.GetOrderCountByStatus(getConnDB(), orderEnum.OrderBusinessStatusPaySuccess, u.MerInfo.MerId, u.UserInfo.ID)

	if err != nil {
		if err != nil {
			JSON(c, http.StatusBadRequest, "", err.Error())
			return
		}

	}
	data.OrderWaitCount = waitCount + waitShipCount

	JSON(c, http.StatusOK, "", nil)
	return
}

////TODO:获取商家店铺二维码
//func GetMerShopQRCode(c *gin.Context) {
//
//	//获取用户信息
//	err, _ := GetCustomUserInfo(c)
//	if err != nil {
//		JSON(c, http.StatusInternalServerError, "", err.Error())
//		return
//	}
//
//}
