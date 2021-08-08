package controller

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ququgou-shop/api/userProgram/cache"
	"github.com/ququgou-shop/api/userProgram/config"
	"github.com/ququgou-shop/api/userProgram/enum"
	"github.com/ququgou-shop/modules/order/orderEnum"
	"github.com/ququgou-shop/modules/payment/paymentEnum"
	"github.com/ququgou-shop/service/orderService"
)

// UserOrderCreate 创建用户订单
// @Summary 创建用户订单
// @Description 创建用户订单
// @Accept json
// @Produce json
// @Param body body orderService.CreateOrderInfoModel true "body参数"
// @Success 200 {object} Response "ok" "返回信息"
// @Failure 400 {object} Response "错误"
// @Failure 401 {object} Response "错误"
// @Failure 500 {object} Response "错误"
// @Router /order/create [post]
func UserOrderCreate(c *gin.Context) {
	var (
		req orderService.CreateOrderInfoModel
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

	timeoutDuration := 60 * time.Second

	//添加LOCK
	if b := cache.Lock(cache.OrderCreateLock+u.Guid, "1", timeoutDuration); !b {
		//该用户有订单正在创建中...
		JSON(c, http.StatusOK, "订单创建中，请稍后", "")
		return
	}

	err, data := orderService.CreateOrderInfo(getConnDB(), &req, u)

	if err != nil {
		JSON(c, http.StatusBadRequest, "", err.Error())
		cache.UnLock(cache.OrderCreateLock + u.Guid)
		return
	}

	cache.UnLock(cache.OrderCreateLock + u.Guid)

	//订单创建成功，调用通知接口
	//go messageNotifi.OrderCreateSuccessNotifi(
	//	&messageNotifi.OrderCreateSuccessNotifiModel{
	//		OrderNo: data.OrderNo,
	//		MerId:   data.UserId,
	//	})

	//订单创建成功 支付方式 为在线支付的，调用在线支付接口
	if data.PaymentTypeCode != paymentEnum.PaymentTypeOfflinePay {
		err, payResult := orderService.OrderMasterPaymentCreate(getConnDB(), &orderService.UserOrderPayModel{
			UserId:  u.ID,
			OrderId: data.OrderId,
		})
		//err, payResult := payment.CreatePayment(db.MysqlConn(), &payment.CreatePaymentModel{
		//	UserId:        u.ID,
		//	PaymentTypeId: data.PaymentTypeId,
		//	OrderId:       data.OrderId,
		//	Source:        0,
		//	Note:          "",
		//	ClientInfo: payment.ClientInfoModel{
		//		Ip:         "123.12.12.123",
		//		DeviceInfo: "",
		//	},
		//})

		if err != nil {
			JSON(c, http.StatusInternalServerError, "", err.Error())
			return
		}

		JSON(c, http.StatusOK, "", payResult)
		return
	}

	JSON(c, http.StatusOK, "", data)
	return

}

// BefOrderCreate 创建订单前(提交订单信息); 用来做校验  (失效接口)
// @Summary    (失效接口)
// @Description   (失效接口)
// @Accept json
// @Produce json
// @Param body body orderService.GetBeforeOrderCreateInfoModel true "body参数"
// @Success 200 {object} Response{data=orderService.BefOrderInfo} "ok" "返回信息"
// @Failure 400 {object} Response "错误"
// @Failure 401 {object} Response "错误"
// @Failure 500 {object} Response "错误"
// @Router /order/befCreate [post]
func BefOrderCreate(c *gin.Context) {
	var (
		req orderService.GetBeforeOrderCreateInfoModel
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

	err, data := orderService.GetBeforeOrderCreateInfo(getConnDB(), &req, u)

	if err != nil {
		JSON(c, http.StatusInternalServerError, "", err.Error())
		return
	}

	JSON(c, http.StatusOK, "ok", data)
	return

}

// OrderPay 订单支付
// @Summary 订单支付
// @Description 这个步骤在创建订单以后，订单支付 成功后会返回数据到微信小程序，
//微信收到后会调用起支付密码输入框，成功后微信支付回调
// @Accept json
// @Produce json
// @Param body body orderService.UserOrderPayModel true "body参数"
// @Success 200 {object} Response "ok" "返回信息"
// @Failure 400 {object} Response "错误"
// @Failure 401 {object} Response "错误"
// @Failure 500 {object} Response "错误"
// @Router /order/pay [post]
func OrderPay(c *gin.Context) {
	var (
		req orderService.UserOrderPayModel
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

	err, payResult := orderService.OrderPaymentCreate(getConnDB(), &orderService.UserOrderPayModel{
		UserId:  u.ID,
		OrderNo: req.OrderNo,
	})

	if err != nil {
		JSON(c, http.StatusInternalServerError, "", err.Error())
		return
	}

	JSON(c, http.StatusOK, "", payResult)
	return

}

// OrderCancel 取消订单
// @Summary 取消订单
// @Description 根据订单号取消订单 TODO: 逻辑待补充
// @Accept json
// @Produce json
// @Param body body orderService.CancelUserOrderModel true "body参数"
// @Success 200 {object} Response "ok" "返回信息"
// @Failure 400 {object} Response "错误"
// @Failure 401 {object} Response "错误"
// @Failure 500 {object} Response "错误"
// @Router /order/cancel [post]
func OrderCancel(c *gin.Context) {
	var (
		req orderService.CancelUserOrderModel
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

	err = orderService.OrderCancel(getConnDB(), &req, u)

	if err != nil {
		JSON(c, http.StatusBadRequest, "", err.Error())
		return
	}

	JSON(c, http.StatusOK, "ok", "")
	return

}

// GetUserOrderList 获取用户订单列表
// @Summary 获取用户订单列表
// @Description 获取用户订单列表
// @Accept json
// @Produce json
// @Param body body orderService.GetOrderSmallInfoListModel true "body参数"
// @Success 200 {object} Response{data=[]orderService.SmallOrderInfModel} "ok" "返回信息"
// @Failure 400 {object} Response "错误"
// @Failure 401 {object} Response "错误"
// @Failure 500 {object} Response "错误"
// @Router /order/get/list [get]
func GetUserOrderList(c *gin.Context) {
	var (
		req orderService.GetOrderSmallInfoListModel
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

	//前台显示订单状态 转化为 业务订单状态
	status := enum.OrderDisplayStatus(req.DisplayStatus)

	req.All = false

	switch status {
	case enum.OrderDisplayStatusAll:
		req.All = true
		break
	case enum.OrderDisplayStatusWaitPay:
		req.BusinessStatus = append(req.BusinessStatus, string(orderEnum.OrderBusinessStatusWaitPay))
		break
	case enum.OrderDisplayStatusWaitProcess:
		req.BusinessStatus = append(req.BusinessStatus, string(orderEnum.OrderBusinessStatusWaitProcess))
		req.BusinessStatus = append(req.BusinessStatus, string(orderEnum.OrderBusinessStatusPaySuccess))
		req.BusinessStatus = append(req.BusinessStatus, string(orderEnum.OrderBusinessStatusShip))
		req.BusinessStatus = append(req.BusinessStatus, string(orderEnum.OrderBusinessStatusDelivered))
		break
	case enum.OrderDisplayStatusFinish:
		req.BusinessStatus = append(req.BusinessStatus, string(orderEnum.OrderBusinessStatusFinish))
		break
	case enum.OrderDisplayStatusCancel:
		req.BusinessStatus = append(req.BusinessStatus, string(orderEnum.OrderBusinessStatusOrderCancel))
		break
	}

	err, data, count := orderService.GetOrderSmallInfoList(getConnDB(), &req, config.Config.ImgService.QiniuUrl)

	if err != nil {
		JSON(c, http.StatusBadRequest, "", err.Error())
		return
	}

	JSONPage(c, http.StatusOK, "ok", data, count)
	return
}

// GetUserOrderDetail 获取用户订单详情
// @Summary 获取用户订单详情
// @Description 获取用户订单详情
// @Accept json
// @Produce json
// @Param body body orderService.GetUserOrderDetailModel true "body参数"
// @Success 200 {object} Response{data=orderService.UserOrderDetailModel} "ok" "返回信息"
// @Failure 400 {object} Response "错误"
// @Failure 401 {object} Response "错误"
// @Failure 500 {object} Response "错误"
// @Router /order/get/detail [get]
func GetUserOrderDetail(c *gin.Context) {
	var (
		req orderService.GetUserOrderDetailModel
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

	err, data := orderService.GetUserOrderDetail(getConnDB(), &req, u, config.Config.ImgService.QiniuUrl)

	if err != nil {
		JSON(c, http.StatusBadRequest, "", err.Error())
		return
	}

	JSON(c, http.StatusOK, "ok", data)
	return
}

// TimeoutNotPayOrderCancel 超时未支付订单取消 (接口无效)
// @Summary 超时未支付订单取消
// @Description 脚本自动执行
// @Accept json
// @Produce json
// @Success 200 {object} Response "ok" "返回信息"
// @Failure 400 {object} Response "错误"
// @Failure 401 {object} Response "错误"
// @Failure 500 {object} Response "错误"
// @Router /order/get/detail [get]
func TimeoutNotPayOrderCancel(c *gin.Context) {

	timeoutDuration := 120 * time.Second
	//添加LOCK
	if b := cache.Lock(cache.TimeoutNotPayOrderCancelLock, "1", timeoutDuration); !b {

		JSON(c, http.StatusOK, "处理中", "")
		return
	}

	err := orderService.TimeoutNotPayOrderCancel(getConnDB(), -20)

	cache.UnLock(cache.TimeoutNotPayOrderCancelLock)

	if err != nil {
		JSON(c, http.StatusBadRequest, "", err.Error())
		return
	}

	JSON(c, http.StatusOK, "ok", "")
	return
}

//UserOrderRefund 支付订单取消并退款 申请
//TODO:支付订单取消并退款 申请
func UserOrderRefund(c *gin.Context) {
}
