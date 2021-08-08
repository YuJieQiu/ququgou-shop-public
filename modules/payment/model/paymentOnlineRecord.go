package model

import "github.com/ququgou-shop/library/base_model"

//在线支付记录
type PaymentOnlineRecord struct {
	base_model.IDAutoModel
	TradeNo         string `json:"tradeNo" gorm:"column:trade_no"`                   //交易号
	PaymentTypeId   uint64 `json:"paymentTypeId" gorm:"column:payment_type_id;"`     //支付类型Id
	PaymentTypeCode string `json:"paymentTypeCode" gorm:"column:payment_type_code;"` //支付方式Code
	Success         bool   `json:"success" gorm:"column:success;"`                   //成功
	Result          string `json:"result" gorm:"column:result;type:text;"`           //结果
	base_model.TimeAllModel
}

// Set table name
func (PaymentOnlineRecord) TableName() string {
	return "payment_online_records"
}
