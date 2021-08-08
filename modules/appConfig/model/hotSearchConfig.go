package model

import "github.com/ququgou-shop/library/base_model"

//热门搜索配置
type HotSearchConfig struct {
	base_model.IDAutoModel
	OpenUserId uint64 `json:"openUserId" gorm:"column:open_user_id;"` //操作用户
	Text       string `json:"text" gorm:"column:text;"`               //搜索文本
	IsHome     bool   `json:"isHome" gorm:"column:is_home;"`          //是否展示首页
	Sort       int    `json:"sort" gorm:"column:sort"`                //默认根据sort 排序
	base_model.TimeAllModel
}

// Set table name
func (HotSearchConfig) TableName() string {
	return "hot_search_config"
}
