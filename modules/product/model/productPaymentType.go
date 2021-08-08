package model

import "github.com/ququgou-shop/library/base_model"

//产品支付方式
type ProductPaymentType struct {
	base_model.IDAutoModel
	ProductId     uint64 `json:"productId" gorm:"column:product_id;index:product_id"`
	PaymentTypeId uint64 `json:"paymentTypeId" gorm:"column:payment_type_id;"`
	base_model.TimeAllModel
}

// Set table name
func (ProductPaymentType) TableName() string {
	return "product_payment_type"
}
