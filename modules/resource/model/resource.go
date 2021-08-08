package model

import "github.com/ququgou-shop/library/base_model"

type Resource struct {
	base_model.IDAutoModel
	Guid        string `json:"guid" gorm:"unique_index;not null;unique;column:guid;"`
	MerId       uint64 `json:"merId" gorm:"column:mer_id"`    //商户ID
	UserId      uint64 `json:"userId" gorm:"column:user_id;"` //用户ID
	Path        string `json:"path" gorm:"column:path"`
	Url         string `json:"url" gorm:"column:url"`
	IconUrl     string `json:"iconUrl" gorm:"column:icon_url"`
	ThumbUrl    string `json:"thumbUrl" gorm:"column:thumb_url"`
	FileName    string `json:"fileName" gorm:"column:file_name"`
	Ext         string `json:"ext" gorm:"column:ext"` //扩展名
	Hosted      string `json:"hosted" gorm:"column:hosted"`
	ContentType string `json:"contentType" gorm:"column:content_type"` //类型
	Size        int64  `json:"size" gorm:"column:size"`                //大小 K
	HashCode    string `json:"hashCode" gorm:"column:hash_code"`       //文件 hash 值 之 MD5-Hash
	Type        int    `json:"type" gorm:"column:type"`                //#ResourceType
	LinkPoint   string `json:"linkPoint" gorm:"column:link_point"`
	base_model.TimeAllModel
}

// Set table name
func (Resource) TableName() string {
	return "resources"
}

//设置图片路径 TODO:待优化
func (r *Resource) SetImagesUrl(imgServiceUrl string) {
	r.Url = imgServiceUrl + r.Url
}
