package model

import "github.com/ququgou-shop/library/base_model"

//规格属性组
type AttributeGroup struct {
	base_model.IDAutoModel
	ProductId uint64 `json:"productId" gorm:"column:product_id"` //产品ID
	Sort      int    `json:"sort" gorm:"column:sort"`
	base_model.TimeAllModel

	//业务字段 不是数据库字段
	AttributeValueGroup []AttributeValueGroup `json:"attributeValueGroup"  gorm:"-"`
}

// Set table name
func (AttributeGroup) TableName() string {
	return "attribute_groups"
}
