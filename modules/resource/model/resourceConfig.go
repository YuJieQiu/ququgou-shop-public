package model

import "github.com/ququgou-shop/library/base_model"

type ResourceConfig struct {
	base_model.IDAutoModel
	ResourceType int    `json:"resourceType" gorm:"column:resource_type;"` // #ResourceType
	StoreType    int    `json:"storeType" gorm:"column:store_type;"`       //#UploadStoreType
	ServiceUrl   string `json:"serviceUrl" gorm:"column:service_url;"`     //访问URL
	LimitSize    int    `json:"limitSize" gorm:"column:limit_size;"`       //限制大小
	LimitTypes   string `json:"limitTypes" gorm:"column:limit_types;"`     //类型限制
	base_model.TimeAllModel
}

// Set table name
func (ResourceConfig) TableName() string {
	return "resource_config"
}
