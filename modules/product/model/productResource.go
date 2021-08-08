package model


import (
	"fmt"
	"github.com/ququgou-shop/library/base_model"
	modelResource "github.com/ququgou-shop/modules/resource/model"
)

//产品资源图片
type ProductResource struct {
	ResourceId   uint64 `json:"resourceId" gorm:"column:resource_id;index:resource_id"`
	ResourceGuid string `json:"resourceGuid" gorm:"column:resource_guid"`
	ProductId    uint64 `json:"productId" gorm:"column:product_id;index:product_id"`
	ProductGuid  string `json:"productGuid" gorm:"column:product_guid"`
	Cover        bool   `json:"cover" gorm:"column:cover"`       //是否封面
	Type         int16  `json:"type" gorm:"column:type"`         //类型 0 默认 图片 、1 视频
	Sort         int    `json:"sort" gorm:"column:sort"`         //默认根据sort 排序
	Position     int    `json:"position" gorm:"column:position"` //位置 默认0 暂不使用
	base_model.TimeAllModel

	//业务字段 不是数据库字段
	Resource modelResource.Resource `json:"resource" gorm:"-"`
}

// Set table name
func (ProductResource) TableName() string {
	return "product_resources"
}

func (p ProductResource) GetTableFields(args []string) string {
	tableName := p.TableName()

	fields := ""

	for i := 0; i < len(args); i++ {
		field := args[i]
		fields += fmt.Sprintf(" s%.%s,", tableName, field)
	}
	return ""
}

//func (p ProductResource) GetTableFields(arr ...interface{}) string {
//
//	tableName := p.TableName()
//
//	fields := ""
//
//	getType := reflect.TypeOf(p)
//
//	for i := 0; i < getType.NumField(); i++ {
//		field := getType.Field(i)
//		fields += fmt.Sprintf(" s%.%s,", tableName, field.Name)
//	}
//
//	fields = strings.Trim(fields, ",")
//
//	return fields
//}
