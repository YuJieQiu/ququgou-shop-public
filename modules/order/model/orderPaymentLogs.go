package model

import "github.com/ququgou-shop/library/base_model"

//TODO：订单支付操作记录( orderId 还是 orderMaster Id？
//订单支付操作记录
type OrderPaymentLogs struct {
	base_model.IDAutoModel
	UserId  uint64 `json:"userId" gorm:"column:user_id"`
	OrderId uint64 `json:"orderId" gorm:"column:order_id"`
	Remark  string `json:"remark" gorm:"column:remark"`
	base_model.TimeAllModel

	//UserName string `json:"userName" gorm:"column:user_name"`
	//OrderPaymentId     uint64 `json:"orderPaymentId" gorm:"column:order_payment_id"`
	//OrderPaymentStatus int    `json:"orderPaymentStatus" gorm:"column:order_payment_status"`
}

// Set table name
func (OrderPaymentLogs) TableName() string {
	return "order_payment_logs"
}
