package model

import "github.com/ququgou-shop/library/base_model"

//品牌
//TODO:Brand
type Brand struct {
	base_model.IDAutoModel
	MerId      uint64 `json:"merId" gorm:"column:mer_id"`
	Name       string `json:"name" gorm:"column:name"`
	Status     int16  `json:"status" gorm:"column:status"`
	Sort       int    `json:"sort" gorm:"column:sort"`
	ResourceId uint64 `json:"resourceId" gorm:"column:resource_id"` //图片 选填
	base_model.ImageJsonModel
	base_model.TimeAllModel
}

// Set table name
func (Brand) TableName() string {
	return "brands"
}
