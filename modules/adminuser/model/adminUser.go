package model

import "github.com/ququgou-shop/library/base_model"

type AdminUser struct {
	base_model.IDAutoModel
	Guid     string `json:"guid"`
	UserName string `json:"userName" gorm:"column:user_name"`
	PassWord string `json:"passWord" gorm:"column:pass_word"` //密码
	Status   byte   `json:"status" gorm:"column:status"`      //0 未激活、1 已启用
	base_model.TimeAllModel
}

// Set table name
func (AdminUser) TableName() string {
	return "admin_users"
}
