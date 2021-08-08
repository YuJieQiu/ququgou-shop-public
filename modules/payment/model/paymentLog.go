package model

import "github.com/ququgou-shop/library/base_model"

//支付日志
type PaymentLog struct {
	base_model.IDAutoModel
	Note string `json:"note" gorm:"column:note"`
	base_model.TimeAllModel
}

// Set table name
func (PaymentLog) TableName() string {
	return "payment_logs"
}
