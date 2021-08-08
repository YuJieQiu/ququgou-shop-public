package model

import "github.com/ququgou-shop/library/base_model"

//产品支持的交付方式
//DeliveryType
type ProductDeliveryType struct {
	base_model.IDAutoModel
	ProductId      uint64 `json:"productId" gorm:"column:product_id;index:product_id"`
	DeliveryTypeId uint64 `json:"deliveryTypeId" gorm:"column:delivery_type_id;"`
	Remark         string `json:"remark" gorm:"column:remark"` //备注
	base_model.TimeAllModel
}

// Set table name
func (ProductDeliveryType) TableName() string {
	return "product_delivery_type"
}
