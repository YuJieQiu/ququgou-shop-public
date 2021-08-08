package model

import "github.com/ququgou-shop/library/base_model"

type Banner struct {
	base_model.IDAutoModel
	Name            string `json:"name" gorm:"column:name"`
	Description     string `json:"description"`                                     //描述
	LinkUrl         string `json:"linkUrl" gorm:"column:link_url"`                  //链接网址
	Type            string `json:"type" gorm:"column:type"`                         //类型默认 0
	Position        string `json:"position" gorm:"column:position"`                 //位置信息
	ResourceId      int    `json:"resourceId" gorm:"column:resource_id"`            //图片
	Sort            int    `json:"sort" gorm:"column:sort"`                         //默认根据sort 排序
	BackgroundColor string `json:"backgroundColor" gorm:"column:background_color;"` //背景颜色
	FontColor       string `json:"fontColor" gorm:"column:font_color;"`             //字体颜色
	base_model.ImageJsonSingleModel
	base_model.TimeAllModel
}

// Set table name
func (Banner) TableName() string {
	return "banners"
}
