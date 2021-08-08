package model

import "github.com/ququgou-shop/library/base_model"

type MerchantType struct {
	base_model.IDAutoModel
	Name        string `json:"name" gorm:"column:name"`
	Description string `json:"description" gorm:"column:description"` //描述
	Status      int    `json:"status" gorm:"column:status"`           //0
	Active      bool   `json:"active" gorm:"column:active"`           //启用
	base_model.TimeAllModel
}

// Set table name
func (MerchantType) TableName() string {
	return "merchant_types"
}
