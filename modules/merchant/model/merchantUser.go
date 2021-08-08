package model

import "github.com/ququgou-shop/library/base_model"

//商户关联的用户
type MerchantUser struct {
	base_model.IDAutoModel
	MerchantId uint64 `json:"merchantId" gorm:"column:merchant_id;"` //商户Id
	UserId     uint64 `json:"userId" gorm:"column:user_id"`          //用户Id
	IsAdmin    bool   `json:"isAdmin" gorm:"column:is_admin;"`       //是否超级管理员
	Status     int    `json:"status" gorm:"column:status"`           //0 状态 ( 0 审核通过)    -1 审核中
	Active     bool   `json:"active" gorm:"column:active"`           //激活 是否使用
	base_model.TimeAllModel
}

// Set table name
func (MerchantUser) TableName() string {
	return "merchant_users"
}
