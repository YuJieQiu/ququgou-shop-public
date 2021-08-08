package model

import "github.com/ququgou-shop/library/base_model"

//PropertyValue
type PropertyValue struct {
	base_model.IDAutoModel
	PropertyId uint64 `json:"propertyId" gorm:"column:property_id"`
	ProductId  uint64 `json:"productId" gorm:"column:product_id"`
	Name       string `json:"name" gorm:"column:name"`
	Title      string `json:"title" gorm:"column:title"`
	Status     int16  `json:"status" gorm:"column:status"` //默认 0
	Sort       int    `json:"sort" gorm:"column:sort"`
	base_model.TimeAllModel
}

// Set table name
func (PropertyValue) TableName() string {
	return "property_value"
}
