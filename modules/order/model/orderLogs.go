package model

import "github.com/ququgou-shop/library/base_model"

//订单操作日志表 TODO:优化 后面可以放到 mongodb 中

type OrderLogs struct {
	base_model.IDAutoModel
	UserId        uint64 `json:"userId" gorm:"column:user_id"`
	UserGuid      string `json:"userGuid" gorm:"column:user_guid"`
	OrderId       uint64 `json:"orderId"gorm:"column:order_id"`
	OrderMasterId uint64 `json:"orderMasterId" gorm:"column:order_master_id;"` //订单id
	Status        int    `json:"status" gorm:"column:status"`
	Type          int    `json:"type" gorm:"column:type"` //1 订单创建 2、订单取消
	Remark        string `json:"remark" gorm:"column:remark"`
	base_model.TimeAllModel
}

// Set table name
func (OrderLogs) TableName() string {
	return "order_logs"
}
