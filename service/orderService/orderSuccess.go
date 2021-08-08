package orderService

import (
	"github.com/jinzhu/gorm"
	"github.com/ququgou-shop/modules/order/model"
	"github.com/ququgou-shop/modules/order/orderEnum"
	"github.com/ququgou-shop/service/productService"

	userModel "github.com/ququgou-shop/modules/user/model"
	"time"
)

//商户完成用户订单
func MerSuccessUserOrder(db *gorm.DB, q *MerSuccessUserOrderModel, u *userModel.User) error {

	var (
		err          error
		orders       model.Order
		orderDetails []model.OrderDetail
	)

	//查询订单状态
	err = db.Where("order_sub_no =? and mer_id=? ", q.OrderNo, q.MerId).First(&orders).Error
	if err != nil {
		return err
	}

	orderStatus := GetOrderBusinessStatus(orders.OrderStatus, orders.PayStatus, orders.DeliveryStatus, orders.Type)

	if orderStatus != orderEnum.OrderBusinessStatusWaitProcess {
		return ErrOrderStatusInvalid
	}

	tx := db.Begin()

	timeNow := time.Now()

	err = tx.Model(&orders).
		Where("mer_id=? ",
			q.MerId).
		Update(model.Order{OrderStatus: int(orderEnum.OrderSuccess), CancelTime: &timeNow}).Error

	if err != nil {
		tx.Rollback()
		return err
	}

	//获取订单详情
	err = db.Where("order_id=?", orders.ID).Find(&orderDetails).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	//更改商品销量
	for _, i := range orderDetails {

		err = productService.ProductSalesUpdate(tx, &productService.ProductSalesUpdateModel{
			ProductId:    i.ProductId,
			ProductSkuId: i.ProductSkuId,
			Count:        i.ProductCount,
			Type:         0,
		})
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	//订单详情 状态不用改变
	//err = tx.Model(&model2.OrderDetail{}).Where("order_sub_id =? and user_id=? and status=?  and deleted_at is null",
	//	ordersub.ID, u.ID, int(orderEnum.OrderStatusWaitPay)).
	//	Update("status", int(orderEnum.OrderStatusOrderCancel)).Error

	//if err != nil {
	//	tx.Rollback()
	//	return err
	//}

	tx.Commit()

	//日志记录
	oLog := model.OrderLogs{
		UserId:        orders.UserId,
		UserGuid:      u.Guid,
		OrderId:       orders.ID,
		OrderMasterId: orders.OrderMasterId,
		Status:        int(orderEnum.OrderSuccess),
		Type:          2,
		Remark:        "订单完成，操作用户:" + u.UserName,
	}

	db.Save(&oLog)

	return nil
}
