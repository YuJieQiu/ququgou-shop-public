package model

import "github.com/ququgou-shop/library/base_model"

//规格属性 选项 组
type AttributeValueGroup struct {
	AttributeGroupId uint64 `json:"attributeGroupId" gorm:"column:attribute_group_id"`
	AttributeValueId uint64 `json:"attributeValueId" gorm:"column:attribute_value_id"`
	Sort             int    `json:"sort" gorm:"column:sort"` //默认根据sort 排序
	base_model.TimeAllModel
}

// Set table name
func (AttributeValueGroup) TableName() string {
	return "attribute_value_groups"
}
