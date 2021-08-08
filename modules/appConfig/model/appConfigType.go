package model

import (
	"github.com/ququgou-shop/library/base_model"
)

//app 配置
type AppConfigType struct {
	base_model.IDAutoModel
	Name   string `json:"name" gorm:"column:name;"`      //名字
	Text   string `json:"text" gorm:"column:text;"`      //
	Status int    `json:"status"  gorm:"column:status;"` //0未启用 、默认 1启用
	base_model.TimeAllModel
}

// Set table name
func (AppConfigType) TableName() string {
	return "app_config_types"
}
