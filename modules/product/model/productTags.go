package model

import "github.com/ququgou-shop/library/base_model"

//产品标签
type ProductTags struct {
	base_model.IDAutoModel
	ProductId uint64 `json:"productId" gorm:"column:product_id;index:product_id"`
	TagId     uint64 `json:"tagId" gorm:"column:tag_id;index:tag_id"` //标签ID
	Sort      int    `json:"sort" gorm:"column:sort"`                 //
	base_model.TimeAllModel
	//Name string `json:"name" gorm:"index:name"` //标签名称

	//业务字段 不是数据库字段
	Tag Tags `json:"tag" gorm:"-"`
}

// Set table name
func (ProductTags) TableName() string {
	return "product_tags"
}
