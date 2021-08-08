package model

import "github.com/ququgou-shop/library/base_model"

//规格属性SKU组
type AttributeGroupSKU struct {
	base_model.IDAutoModel
	ProductSkuId     uint64 `json:"productSkuId" gorm:"column:product_sku_id"`
	AttributeGroupId uint64 `json:"attributeGroupId" gorm:"column:attribute_group_id"`
	base_model.TimeAllModel
}

// Set table name
func (AttributeGroupSKU) TableName() string {
	return "attribute_group_sku"
}
