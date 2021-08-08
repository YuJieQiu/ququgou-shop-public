package model

import "github.com/ququgou-shop/library/base_model"

type MerchantResources struct {
	base_model.IDAutoModel
	MerchantId  uint64 `json:"merchantId" gorm:"column:merchant_id"`
	ResourcesId uint64 `json:"resourcesId" gorm:"column:resources_id"`
	Cover       bool   `json:"cover" gorm:"column:cover"`
	Type        int16  `json:"type" gorm:"column:type"`      //暂留字段
	IsLogo      bool   `json:"isLogo" gorm:"column:is_logo"` //是否logo 图片
	Sort        int    `json:"sort" gorm:"column:sort"`      //排序
	base_model.TimeAllModel
}

// Set table name
func (MerchantResources) TableName() string {
	return "merchant_resources"
}
