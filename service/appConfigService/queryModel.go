package appConfigService

import (
	"github.com/ququgou-shop/library/base_model"
	appConfigModel "github.com/ququgou-shop/modules/appConfig/model"
	bannerModel "github.com/ququgou-shop/modules/banner/model"
)

type (
	GetAppConfigListModel struct {
		base_model.QueryParamsPage `gorm:"-"`
		Name                       string `form:"name" gorm:"-" json:"name,omitempty"`
		Type                       string `form:"type" gorm:"-" json:"type,omitempty"`
	}

	AppConfigSaveModel struct {
		AppConfigs []appConfigModel.AppConfig `json:"appConfigs"`
		Type       string                     `json:"type"` //分类 类型 homeTab、homeCategory
	}

	HotSearchConfigListSaveModel struct {
		UserId uint64                           `json:"userId"`
		List   []appConfigModel.HotSearchConfig `json:"list"`
	}

	GetHotSearchConfigListModel struct {
		base_model.QueryParamsPage `gorm:"-"`
		IsHome                     bool `json:"isHome" form:"isHome"`
	}

	GetHomeProductConfigListModel struct {
		base_model.QueryParamsPage `gorm:"-"`
		//ProductId       uint64 `json:"product_id"`
		AppConfigId uint64 `json:"appConfigId"` //appconfig iD
	}
	GetHomeConfigProductInfoModel struct {
		base_model.QueryParamsPage `gorm:"-"`
		//ProductId       uint64 `json:"product_id"`
		AppConfigId uint64 `form:"appConfigId" json:"appConfigId"` //appconfig iD
	}

	GetHomeCategoryListModel struct {
		base_model.QueryParamsPage `gorm:"-"`
		Name                       string `form:"name" gorm:"-" json:"name,omitempty"`
		MerId                      uint64 `json:"merId"`
	}

	GetCategoryProductListModel struct {
		base_model.QueryParamsPage `gorm:"-"`
		CategoryId                 uint64  `form:"categoryId" json:"categoryId"`
		Order                      int     `json:"order"`          //排序方式 0默认 1价格  3销量 ...
		Lat                        float64 `form:"lat" json:"lat"` //维度
		Lon                        float64 `form:"lon" json:"lon"` //经度
		Distance                   int     `json:"distance" form:"distance"`
	}

	GetSearchProductListModel struct {
		base_model.QueryParamsPage `gorm:"-"`
		CategoryId                 uint64  `form:"categoryId" json:"categoryId"` //商品分类Id
		Text                       string  `form:"text" json:"text"`             //搜索名称 泛搜索 包括 商品名称
		Lat                        float64 `form:"lat" json:"lat"`               //维度
		Lon                        float64 `form:"lon" json:"lon"`               //经度
		SearchSortType             int     `form:"sortType" json:"sortType"`     //排序类型 1、默认 3、销量 正序 5、销量 倒叙  7、价格 正序 9、价格 倒叙 11、距离 最近
		Distance                   int     `json:"distance" form:"distance"`
		MerId                      uint64  `json:"merId" form:"merId"` //商家ID
	}

	GetBannerListModel struct {
		base_model.QueryParamsPage `gorm:"-"`
		Name                       string `form:"name" gorm:"-" json:"name,omitempty"`
	}

	BannerSaveModel struct {
		Banners []bannerModel.Banner `json:"banners"`
	}
)
