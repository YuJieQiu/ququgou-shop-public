package model

import "github.com/ququgou-shop/library/base_model"

//Property
type Property struct {
	base_model.IDAutoModel
	MerId    uint64 `json:"merId" gorm:"column:mer_id"`
	Name     string `json:"name" gorm:"column:name"`
	Status   int16  `json:"status" gorm:"column:status"` //默认 0
	Sort     int    `json:"sort" gorm:"column:sort"`
	IsSystem bool   `json:"isSystem" gorm:"column:is_system"`
	base_model.TimeAllModel

	//业务字段 不是数据库字段
	Values []PropertyValue `json:"values" gorm:"-"`
}

// Set table name
func (Property) TableName() string {
	return "property"
}
