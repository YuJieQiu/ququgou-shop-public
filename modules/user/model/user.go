package model

import (
	"github.com/ququgou-shop/library/base_model"
)

//用户信息
type User struct {
	base_model.IDAutoModel
	AppKey string `json:"appKey" gorm:"column:app_key;"` //APP_Key
	Guid         string `json:"guid" gorm:"column:guid"`
	UserName     string `json:"userName" gorm:"column:user_name"`
	Mobile       string `json:"mobile" gorm:"column:mobile"` //手机号
	Gender       byte   `json:"gender" gorm:"column:gender"` //性别
	AvatarUrl    string `json:"avatarUrl" gorm:"column:avatar_url"` //头像地址
	WechatOpenId string `json:"wechatOpenId" gorm:"column:wechat_open_id"`
	Status       byte   `json:"status" gorm:"column:status"` //0 未激活、1 已启用
	Type         byte   `json:"type" gorm:"column:type"`     //1、商城小程序(User_web) 3、商户小程序(Merchant_web)
	base_model.TimeAllModel
}

// Set table name
func (User) TableName() string {
	return "users"
}

//TODO:用户地址 、用户积分 、用户访问历史
