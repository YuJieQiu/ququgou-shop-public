package model

import "github.com/ququgou-shop/library/base_model"

//收货地址 TODO:暂时这样，数据信息有点少
type Address struct {
	base_model.IDAutoModel
	UserId    uint64 `json:"user_id"`
	City      string `json:"city"`
	Region    string `json:"region"`
	Town      string `json:"town"`
	Address   string `json:"address"`
	Phone     string `json:"phone"`
	Name      string `json:"name"`
	Remark    string `json:"remark"`
	IsDefault bool   `json:"is_default"`
	Active    bool   `json:"active"` //默认 启用
	base_model.TimeAllModel
}

// Set table name
func (Address) TableName() string {
	return "addresses"
}
