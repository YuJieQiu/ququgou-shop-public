package orderService

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/ququgou-shop/modules/order/model"
	"github.com/ququgou-shop/modules/order/orderEnum"
	userModel "github.com/ququgou-shop/modules/user/model"
	"github.com/ququgou-shop/service/productService"

	"time"
)

//订单取消
//待付款的订单
func OrderCancel(db *gorm.DB, q *CancelUserOrderModel, u *userModel.User) error {
	var (
		err   error
		order model.Order
		//orderDetails []model.OrderDetail
	)

	//查询订单状态
	err = db.Where("order_no =? and user_id=? ", q.OrderNo, u.ID).First(&order).Error
	if err != nil {
		return err
	}

	if order.OrderStatus != int(orderEnum.OrderCreate) &&
		order.PayStatus != int(orderEnum.PayWait) {
		return ErrOrderStatusInvalid
	}

	err = orderCancel(db, &order)
	if err != nil {
		return err
	}

	return nil
}

func orderCancel(db *gorm.DB, order *model.Order) (err error) {
	var orderDetails []model.OrderDetail
	tx := db.Begin()

	timeNow := time.Now()

	err = tx.Model(order).
		Where("order_status=? and pay_status=?", int(orderEnum.OrderCreate), int(orderEnum.PayWait)).
		Update(model.Order{OrderStatus: int(orderEnum.OrderCancel), CancelTime: &timeNow}).Error

	if err != nil {
		tx.Rollback()
		return err
	}

	//订单取消 库存修改
	//获取订单详情
	err = db.Where("order_id=?", order.ID).Find(&orderDetails).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	//更改商品库存
	for _, i := range orderDetails {
		err = productService.ProductStockUpdate(tx, &productService.ProductStockUpdateModel{
			ProductId:    i.ProductId,
			ProductSkuId: i.ProductSkuId,
			Count:        i.ProductCount,
			Type:         0, //增加
		})
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	tx.Commit()

	//日志记录
	oLog := model.OrderLogs{
		UserId:        order.UserId,
		UserGuid:      "",
		OrderId:       order.ID,
		OrderMasterId: order.OrderMasterId,
		Status:        int(orderEnum.OrderCancel),
		Type:          2,
		Remark:        "订单取消",
	}

	db.Save(&oLog)

	return nil
}

//超时未支付订单取消
////获取待支付订单  一般超过20 分钟没有支付的
func TimeoutNotPayOrderCancel(db *gorm.DB, min int) error {
	var (
		err   error
		order []model.Order
	)

	sqlSelect := `select * `

	sql := fmt.Sprintf(` from orders  
			where deleted_at is null 
			and created_at < DATE_ADD(UTC_TIMESTAMP(), INTERVAL %v MINUTE)`, min)

	sql += loadQueryStatusSqlString(orderEnum.OrderBusinessStatusWaitPay)

	err = db.Raw(sqlSelect + sql).Scan(&order).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return err
	}

	for _, o := range order {
		err = orderCancel(db, &o)
		if err != nil {
			return err
		}
	}

	return err
}
