package base_model

import (
	"github.com/ququgou-shop/library/ext/ext_struct"
	"time"
)

type (
	// IDAutoModel is
	IDAutoModel struct {
		ID uint64 `json:"id" sql:"AUTO_INCREMENT" gorm:"unique_index;not null;unique;primary_key;column:id"`
	}

	GuidModel struct {
		Guid string `json:"guid" gorm:"unique_index;not null;unique;column:guid;"`
	}

	// CreateModel is
	CreateModel struct {
		CreatedAt time.Time `json:"createdAt" gorm:"column:created_at" sql:"DEFAULT:current_timestamp"`
	}
	// UpdatedAtModel is
	UpdatedAtModel struct {
		UpdatedAt time.Time `json:"updatedAt" gorm:"column:updated_at" sql:"DEFAULT:current_timestamp"`
	}
	// DeletedAtModel is
	DeletedAtModel struct {
		DeletedAt *time.Time `json:"deletedAt,omitempty" gorm:"column:deleted_at"`
	}
	// TimeAllModel is
	TimeAllModel struct {
		CreateModel
		UpdatedAtModel
		DeletedAtModel
	}

	ImageJsonModel struct {
		ImageJson ext_struct.JsonImageArrayString `json:"images"  gorm:"type:text(1000);column:image_json;"` //图片json 数据库字段名称 image_json
	}

	//Single
	ImageJsonSingleModel struct {
		ImageJson ext_struct.JsonImageString `json:"images"  gorm:"column:image_json;type:text(1000);"` //图片json 数据库字段名称 image_json
	}

	//分页
	QueryParamsPage struct {
		Page   int `form:"page" json:"page,omitempty"`
		Limit  int `form:"limit" json:"limit,omitempty"`
		Offset int `form:"offset" json:"offset,omitempty"`
	}
)

func (page *QueryParamsPage) PageSet() {
	if page.Page <= 0 {
		page.Page = 1
	}

	if page.Limit <= 0 {
		page.Limit = 20
	}

	page.Offset = (page.Page - 1) * page.Limit
}
