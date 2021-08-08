package model

import (
	"github.com/ququgou-shop/library/base_model"
	"github.com/ququgou-shop/service/productService"
)

//app 配置
type AppConfig struct {
	base_model.IDAutoModel
	Name        string `json:"name" gorm:"column:name;"`               //名字
	Text        string `json:"text" gorm:"column:text;"`               //显示文本
	Description string `json:"description" gorm:"column:description;"` //描述
	LinkType    int    `json:"linkType" gorm:"column:link_type;"`      //连接类型 默认 0
	LinkUrl     string `json:"linkUrl" gorm:"column:link_url;"`        //链接网址
	Type        string `json:"type" gorm:"column:type;"`               //分类 类型 homeTab、home_category、new_use_exclusive、new_product_recommend
	Code        string `json:"code" gorm:"column:code;"`               //自定义 code
	CategoryId  uint64 `json:"categoryId" gorm:"column:category_id;"`  //分类Id
	ConfigType  int    `json:"configType" gorm:"column:config_type;"`  //规则  默认0 获取全部 1、获取productConfig 2、获取CategoryId
	Sort        int    `json:"sort" gorm:"column:sort;"`
	Status      int    `json:"status"  gorm:"column:status;"`         //0未启用 、默认 1启用
	ResourceId  int    `json:"resourceId" gorm:"column:resource_id;"` //图片
	base_model.ImageJsonSingleModel
	base_model.TimeAllModel

	//业务参数 不是数据库数据
	ProductSmallInfos *[]productService.ProductSmallInfoModel `gorm:"-" json:"productSmallInfos"`
}

// Set table name
func (AppConfig) TableName() string {
	return "app_configs"
}

