package model

import "github.com/ququgou-shop/library/base_model"

//规格属性
type AttributeValue struct {
	base_model.IDAutoModel
	MerId       uint64 `json:"merId" gorm:"column:mer_id"` //商户ID 0默认系统
	Name        string `json:"name" gorm:"column:name"`
	AttributeId uint64 `json:"attributeId" gorm:"column:attribute_id"` //规格ID
	Status      int16  `json:"status" gorm:"column:status"`            // 0 启用 -1 未启用
	Sort        int    `json:"sort" gorm:"column:sort"`                //默认根据sort 排序
	Remark      string `json:"remark" gorm:"column:remark"`            //备注
	ResourceId  uint64 `json:"resourceId" gorm:"column:resource_id"`   //图片资源ID 选填
	IsSystem    bool   `json:"isSystem" gorm:"column:is_system"`
	base_model.ImageJsonSingleModel
	base_model.TimeAllModel

	//业务字段 不是数据库字段
	Attribute *Attribute `json:"attribute,omitempty" gorm:"-"`
}

// Set table name
func (AttributeValue) TableName() string {
	return "attribute_values"
}
