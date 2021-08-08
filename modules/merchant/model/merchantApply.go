package model

import "github.com/ququgou-shop/library/base_model"

//商户申请信息表
type MerchantApply struct {
	base_model.IDAutoModel
	UserId       uint64  `json:"userId" gorm:"column:user_id;"`
	Phone        string  `json:"phone" gorm:"column:phone;"`
	Name         string  `json:"name" gorm:"column:name"`
	Description  string  `json:"description" gorm:"column:description"` //描述
	TypeId       uint64  `json:"typeId" gorm:"column:type_id"`          //类型 默认 0  MerchantTypeId
	City         string  `json:"city" gorm:"column:city"`
	Region       string  `json:"region" gorm:"column:region"`
	Town         string  `json:"town" gorm:"column:town"`
	Address      string  `json:"Address" gorm:"column:address"`
	Latitude     float64 `json:"latitude" gorm:"column:latitude"`           //纬度
	Longitude    float64 `json:"longitude" gorm:"column:longitude"`         //经度
	ResourcesIds string  `json:"resourcesIds" gorm:"column:resources_ids;"` //图片资源ID 信息
	Status       int     `json:"status" gorm:"column:status"`               //0 审核中  1审核通过  -1 审核未通过
	Remark       string  `json:"remark" gorm:"column:remark"`
	base_model.TimeAllModel
}

// Set table name
func (MerchantApply) TableName() string {
	return "merchant_apply"
}
