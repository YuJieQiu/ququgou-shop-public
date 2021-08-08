package model

import "github.com/ququgou-shop/library/base_model"

//用户微信信息信息
type UserWeChatInfo struct {
	base_model.IDAutoModel
	UserId uint64 `json:"userId" gorm:"column:user_id;index:user_id"` //用户ID
	WeChatOpenId string `json:"weChatOpenId" gorm:"column:wechat_open_id;"`//微信OpenID
	base_model.TimeAllModel
}

// Set table name
func (UserWeChatInfo) TableName() string {
	return "user_wechat_info"
}
