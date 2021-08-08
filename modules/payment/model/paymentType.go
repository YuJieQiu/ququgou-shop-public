package model

import (
	"github.com/ququgou-shop/library/base_model"
)

//支付类型配置 其它字段需要再加
type PaymentType struct {
	base_model.IDAutoModel
	Name   string `json:"name" gorm:"column:name"`
	Code   string `json:"code" gorm:"column:code;"`
	Sort   int    `json:"sort" gorm:"column:sort"`     //排序
	Status int16  `json:"status" gorm:"column:status"` //0 启用 -1 未启用
	base_model.TimeAllModel
}

// Set table name
func (PaymentType) TableName() string {
	return "payment_types"
}
