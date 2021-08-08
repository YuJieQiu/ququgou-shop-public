package model

import "github.com/ququgou-shop/library/base_model"

//类别 （系统类目）
type Category struct {
	base_model.IDAutoModel
	MerId      uint64 `json:"merId" gorm:"column:mer_id"`           //商户ID 0 默认系统
	Pid        uint64 `json:"pid" gorm:"column:pid"`                //父类别 id (暂两个级别)
	Name       string `json:"name" gorm:"column:name"`              //
	Status     int16  `json:"status" gorm:"column:status"`          //0 启用 -1 未启用
	Sort       int    `json:"sort" gorm:"column:sort"`              //默认根据sort 排序
	Remark     string `json:"remark" gorm:"column:remark"`          //备注
	ResourceId uint64 `json:"resourceId" gorm:"column:resource_id"` //图片资源ID 选填
	IsSystem   bool   `json:"isSystem" gorm:"column:is_system"`     //是否系统类别
	base_model.ImageJsonSingleModel
	base_model.TimeAllModel

	Child []Category `json:"child" gorm:"-"`
}

// Set table name
func (Category) TableName() string {
	return "categorys"
}
