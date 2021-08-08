package model

import "github.com/ququgou-shop/library/base_model"

//规格属性
type Attribute struct {
	base_model.IDAutoModel
	MerId      uint64 `json:"merId" gorm:"column:mer_id"`           //商户ID 0默认系统
	Name       string `json:"name" gorm:"column:name"`              //名称
	CategoryId uint64 `json:"categoryId" gorm:"column:category_id"` //所属分类ID
	Status     int16  `json:"status" gorm:"column:status"`          //  0 启用 -1 未启用
	Sort       int    `json:"sort" gorm:"column:sort"`              //默认根据sort 排序
	Remark     string `json:"remark" gorm:"column:remark"`          //备注
	ResourceId uint64 `json:"resourceId" gorm:"column:resource_id"` //图片资源ID 选填
	IsSystem   bool   `json:"isSystem" gorm:"column:is_system"`     //是否系统规格属性
	base_model.ImageJsonSingleModel
	base_model.TimeAllModel

	//业务字段 不是数据库字段
	Values []AttributeValue `json:"values,omitempty" gorm:"-"`
}

// Set table name
func (Attribute) TableName() string {
	return "attributes"
}
