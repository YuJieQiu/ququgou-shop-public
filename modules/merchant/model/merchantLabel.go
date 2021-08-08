package model

import "github.com/ququgou-shop/library/base_model"

//商家标签
type MerchantLabel struct {
	base_model.IDAutoModel
	MerchantId uint64 `json:"merchantId" gorm:"column:merchant_id"`
	LabelId    uint64 `json:"labelId" gorm:"column:label_id"`
	Sort       int    `json:"sort" gorm:"column:sort"` //排序
	Text       string `json:"text" gorm:"column:text"` //冗余字段 => label.text
	base_model.TimeAllModel
}

// Set table name
func (MerchantLabel) TableName() string {
	return "merchant_labels"
}
