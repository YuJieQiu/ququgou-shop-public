package model

import "github.com/ququgou-shop/library/base_model"

//商户地址
type MerchantAddress struct {
	base_model.IDAutoModel
	MerchantId uint64  `json:"merchantId" gorm:"column:merchant_id"`
	City       string  `json:"city" gorm:"column:city"`
	Region     string  `json:"region" gorm:"column:region"`
	Town       string  `json:"town" gorm:"column:town"`
	Address    string  `json:"Address" gorm:"column:address"`
	Latitude   float64 `json:"latitude" gorm:"column:latitude"`   //纬度
	Longitude  float64 `json:"longitude" gorm:"column:longitude"` //经度
	Name       string  `json:"name" gorm:"column:name"`           //名称
	Remark     string  `json:"remark" gorm:"column:remark"`
	Active     bool    `json:"active" gorm:"column:active"` //默认 启用
	base_model.TimeAllModel
}

// Set table name
func (MerchantAddress) TableName() string {
	return "merchant_addresses"
}
