package model

import "github.com/ququgou-shop/library/base_model"

//产品类别
type CategoryProduct struct {
	base_model.IDAutoModel
	ProductId  uint64 `json:"productId" gorm:"column:product_id;index:product_id"`
	CategoryId uint64 `json:"categoryId" gorm:"column:category_id;index:category_id"`
	base_model.TimeAllModel

	//业务字段 不是数据库字段
	Category Category `json:"category" gorm:"-"`
}

// Set table name
func (CategoryProduct) TableName() string {
	return "category_products"
}
