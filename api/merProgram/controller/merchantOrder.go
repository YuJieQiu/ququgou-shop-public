package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/ququgou-shop/api/merProgram/config"
	"github.com/ququgou-shop/service/orderService"
	"net/http"
)

func GetMerOrderList(c *gin.Context) {
	var (
		req orderService.GetOrderSmallInfoListModel
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

	err, data, count := orderService.GetOrderSmallInfoList(getConnDB(), &req, config.Config.ImgService.QiniuUrl)

	if err != nil {
		JSON(c, http.StatusBadRequest, "", err.Error())
		return
	}

	JSONPage(c, http.StatusOK, "", data, count)
	return
}

func GetOrderDetail(c *gin.Context) {
	var (
		req orderService.GetUserOrderDetailModel
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

	err, data := orderService.GetUserOrderDetail(getConnDB(), &req, u.UserInfo,config.Config.ImgService.QiniuUrl)

	if err != nil {
		JSON(c, http.StatusBadRequest, "", err.Error())
		return
	}

	JSON(c, http.StatusOK, "", data)
	return
}
//
//获取订单用户信息
func GetOrderUserInfo(c *gin.Context) {
	var (
		req  orderService.MerGetOrderUserInfoModel
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

	err, data := orderService.MerGetOrderUserInfo(getConnDB(), &req)

	if err != nil {
		JSON(c, http.StatusBadRequest, "", err.Error())
		return
	}

	JSON(c, http.StatusOK, "", data)
	return
}

//完成订单
func MerOrderSuccess(c *gin.Context) {
	var (
		req orderService.MerSuccessUserOrderModel
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

	req.MerId = u.MerInfo.MerId

	err = orderService.MerSuccessUserOrder(getConnDB(), &req, u.UserInfo)

	if err != nil {
		JSON(c, http.StatusBadRequest, "", err.Error())
		return
	}

	JSON(c, http.StatusOK, "", "")
	return

}
