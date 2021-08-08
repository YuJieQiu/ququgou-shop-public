package model

import "github.com/ququgou-shop/library/base_model"

//规格属性SKU组 TODO:无用
type AttributeValueSKU struct {
	base_model.IDAutoModel
	ProductSkuId     uint64 `json:"productSkuId" gorm:"column:product_sku_id"`
	AttributeId      uint64 `json:"attributeId" gorm:"column:attribute_id"`            //规格ID
	AttributeValueId uint64 `json:"attributeValueId" gorm:"column:attribute_value_id"` //规格属性ID
	base_model.TimeAllModel
}

// Set table name
func (AttributeValueSKU) TableName() string {
	return "attribute_value_sku"
}
