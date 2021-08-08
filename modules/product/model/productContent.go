package model

import "github.com/ququgou-shop/library/base_model"

//产品富文本描述信息
type ProductContent struct {
	base_model.IDAutoModel
	ProductId uint64 `json:"productId" gorm:"column:product_id;index:product_id"`
	Content   string `json:"content" gorm:"type:longtext"` //商品富文本信息
	Remark    string `json:"remark" gorm:"column:remark"`  //备注
	base_model.TimeAllModel
}

// Set table name
func (ProductContent) TableName() string {
	return "product_content"
}
