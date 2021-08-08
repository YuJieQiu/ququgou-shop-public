package orderService

import (
	"errors"
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/ququgou-shop/service/paymentService"

	"github.com/ququgou-shop/modules/order/model"
	"github.com/ququgou-shop/modules/order/orderEnum"
	"github.com/ququgou-shop/modules/payment/paymentEnum"
	userModel "github.com/ququgou-shop/modules/user/model"
)

//原始订单支付创建
//这里的订单orderNo 为订单订单Id
//除了在创建订单的时候使用的是 主订单 orderNo ，其它时候使用的都是 subOrderNo
//这里自需要调用交易的方法，根据传入的订单编号更新
//每次交易都会产生不同的交易流水单号
//问题是 如果 根据交易流水 查询订单
//因为订单可能是主订单编号  也可能是子订单编号，这里需要加个标记
//然后查询主订单是否在交易表有记录
//然后主订单的支付和子订单的支付记录 区分问题，
//因为涉及到后面的分润等问题
//查询子订单是否在交易表有记录
//如果有记录的话查询 支付记录表中是否有成功的记录
//如果有成功的记录 直接返回前面的支付信息 进行支付，
//这里需要注意的是，如果 时间是否过期 ，一般是10--15分钟，具体的话需要看当前支付类型配置的过期时间
func OrderMasterPaymentCreate(db *gorm.DB, q *UserOrderPayModel) (error, interface{}) {
	var (
		err         error
		orderMaster model.OrderMaster
		orders      []model.Order
		user        userModel.User
		orderTrade  model.OrderTrade
	)

	if q.OrderMasterId > 0 {
		err = db.Where("id =?", q.OrderMasterId).First(&orderMaster).Error
	} else {
		err = db.Where("order_master_no =?", q.OrderMasterNo).First(&orderMaster).Error
	}

	if err != nil {
		return err, nil
	}

	err = db.Where("id =?", q.UserId).First(&user).Error
	if err != nil {
		return err, nil
	}

	req := paymentService.CreatePaymentModel{
		UserId:        user.ID,
		PaymentTypeId: orderMaster.PaymentTypeId,
		BusinessNo:    orderMaster.OrderMasterNo,
		BusinessType:  int(paymentEnum.BusinessTypeOrderMaster),
		Source:        0,
		Note:          "",
		ClientInfo: paymentService.ClientInfoModel{
			Ip:         "123.12.12.123", //TODO:客户端IP地址
			DeviceInfo: "",
		},
		Amount: orderMaster.OrderAmountTotal,
	}

	err, res := paymentService.CreatePayment(db, &req)
	if err != nil {
		return err, nil
	}

	if !res.Success {
		return errors.New("pay error " + res.ErrCodeDes), nil
	}

	//获取主订单下所以订单
	db.Where("order_master_no=?", orderMaster.OrderMasterNo).Find(&orders)

	for _, v := range orders {
		err = db.Where("order_master_no=? and order_no=? and trade_no=?",
			orderMaster.OrderMasterNo,
			v.OrderNo,
			res.TradeNo).First(&orderTrade).Error

		if err != nil {
			if err == gorm.ErrRecordNotFound {
				//不存在新增
				orderTrade.TradeNo = res.TradeNo
				orderTrade.OrderMasterNo = orderMaster.OrderMasterNo
				orderTrade.OrderNo = v.OrderNo
				orderTrade.PaymentTypeId = req.PaymentTypeId
				orderTrade.Amount = req.Amount
				orderTrade.Status = 0
				orderTrade.Remark = req.Note

				err = db.Create(&orderTrade).Error
				if err != nil {
					fmt.Println(err.Error())
				}

			}
		}
		//存在跳过
	}

	return nil, res.Data
}

//订单支付创建
func OrderPaymentCreate(db *gorm.DB, q *UserOrderPayModel) (error, interface{}) {
	var (
		err         error
		order       model.Order
		orderMaster model.OrderMaster
		orderTrade  model.OrderTrade
		user        userModel.User
	)

	if q.OrderId > 0 {
		err = db.Where("id =?", q.OrderId).First(&order).Error
	} else {
		err = db.Where("order_no =?", q.OrderNo).First(&order).Error
	}

	if err != nil {
		return err, nil
	}

	err = db.Where("id =?", order.OrderMasterId).First(&orderMaster).Error
	if err != nil {
		return err, nil
	}

	err = db.Where("id =?", q.UserId).First(&user).Error
	if err != nil {
		return err, nil
	}

	req := paymentService.CreatePaymentModel{
		UserId:        user.ID,
		PaymentTypeId: orderMaster.PaymentTypeId,
		BusinessNo:    order.OrderNo,
		BusinessType:  int(paymentEnum.BusinessTypeOrder),
		Source:        0,
		Note:          "",
		ClientInfo: paymentService.ClientInfoModel{
			Ip:         "123.12.12.123",
			DeviceInfo: "",
		},
		Amount: order.OrderAmountTotal,
	}

	err, res := paymentService.CreatePayment(db, &req)
	if err != nil {
		return err, nil
	}

	if !res.Success {
		return errors.New("pay error " + res.ErrCodeDes), nil
	}

	//订单交易创建、获取下面所以子订单
	err = db.Where("order_no=? and order_master_no=? and trade_no=?",
		orderMaster.OrderMasterNo,
		order.OrderNo,
		res.TradeNo).First(&orderTrade).Error

	//存在跳过 //不存在新增
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			//不存在新增
			orderTrade.TradeNo = res.TradeNo
			orderTrade.OrderMasterNo = orderMaster.OrderMasterNo
			orderTrade.OrderNo = order.OrderNo
			orderTrade.PaymentTypeId = req.PaymentTypeId
			orderTrade.Amount = req.Amount
			orderTrade.Status = 0
			orderTrade.Remark = req.Note
			db.Create(&orderTrade)
		}
	}

	return nil, res.Data
}

//订单支付成功状态更新
func OrderPaymentSuccessUpdate(db *gorm.DB, q *OrderPaymentSuccessUpdateModel) error {

	if q.BusinessType == int(paymentEnum.BusinessTypeOrderMaster) {
		return orderMasterPaymentSuccessUpdate(db, q)
	} else if q.BusinessType == int(paymentEnum.BusinessTypeOrder) {
		return orderPaymentSuccessUpdate(db, q)
	} else {
		return errors.New("BusinessType Error")
	}

}

//原始订单支付成功状态更新
func orderMasterPaymentSuccessUpdate(db *gorm.DB, q *OrderPaymentSuccessUpdateModel) error {
	var (
		err         error
		orderMaster model.OrderMaster
		orders      []model.Order
	)

	err = db.Where("order_master_no=?", q.BusinessNo).First(&orderMaster).Error
	if err != nil {
		return err
	}

	err = db.Where("order_master_id=?", orderMaster.ID).Find(&orders).Error
	if err != nil {
		return err
	}

	tx := db.Begin()

	//订单支付状态更新
	//order.PayStatus = int(orderEnum.PaySuccess)
	//err = tx.Save(order).Error
	err = db.Model(&orderMaster).Update(model.Order{PayStatus: int(orderEnum.PaySuccess)}).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	//所有 子订单状态更新
	for i := 0; i < len(orders); i++ {
		err = db.Model(&orders[i]).Update(model.Order{PayStatus: int(orderEnum.PaySuccess), DeliveryStatus: int(orderEnum.DeliveryWait)}).Error

		//orderSubs[i].PayStatus = int(orderEnum.PaySuccess)
		//tx.Save(orderSubs[i])
		if err != nil {
			tx.Rollback()
			return err
		}

	}

	tx.Commit()

	//日志记录
	return nil
}

//订单支付成功状态更新
func orderPaymentSuccessUpdate(db *gorm.DB, q *OrderPaymentSuccessUpdateModel) error {
	var (
		err         error
		orderMaster model.OrderMaster
		orders      model.Order
	)

	err = db.Where("order_no=?", q.BusinessNo).First(&orders).Error
	if err != nil {
		return err
	}

	//查询订单
	err = db.Where("order_master_no=?", orders.OrderMasterNo).First(&orderMaster).Error
	if err != nil {
		return err
	}

	tx := db.Begin()
	//子订单状态更新
	//orderSub.PayStatus = int(orderEnum.PaySuccess)
	//err = tx.Save(orderSub).Error
	err = db.Model(&orders).Update(model.Order{PayStatus: int(orderEnum.PaySuccess), DeliveryStatus: int(orderEnum.DeliveryWait)}).Error

	if err != nil {
		tx.Rollback()
		return err
	}

	//查询订单下所有子订单状态是否ok
	//更新主订单状态
	//order.PayStatus = int(orderEnum.PaySuccess)
	//err = tx.Save(order).Error
	err = db.Model(&orderMaster).Update("pay_status", int(orderEnum.PaySuccess)).Error

	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()

	return nil
}
