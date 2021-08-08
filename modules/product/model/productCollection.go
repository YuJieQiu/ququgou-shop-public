package model

import "github.com/ququgou-shop/library/base_model"

//产品收藏
type ProductCollection struct {
	base_model.IDAutoModel
	UserId    uint64 `json:"userId" gorm:"column:user_id;index:user_id"`
	ProductId uint64 `json:"productId" gorm:"column:product_id;"`
	base_model.TimeAllModel
}

// Set table name
func (ProductCollection) TableName() string {
	return "product_collection"
}
