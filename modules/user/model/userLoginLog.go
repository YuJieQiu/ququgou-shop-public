package model

import "github.com/ququgou-shop/library/base_model"

//用户登录日志
type UserLoginLog struct {
	base_model.IDAutoModel
	UserId    uint64  `json:"userId" gorm:"column:user_id;index:user_id"`
	Latitude  float64 `json:"latitude" gorm:"column:latitude"`   //纬度
	Longitude float64 `json:"longitude" gorm:"column:longitude"` //经度
	Ip        string  `json:"ip" gorm:"column:ip"` 				 //ID
	Remark    string  `json:"remark" gorm:"column:remark;type:text;"` //备注 其它信息
	base_model.TimeAllModel
}

// Set table name
func (UserLoginLog) TableName() string {
	return "user_login_logs"
}
