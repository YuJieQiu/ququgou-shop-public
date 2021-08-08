package productService

import (
	"github.com/ququgou-shop/library/base_model"
	"github.com/ququgou-shop/modules/product/model"
	"github.com/ququgou-shop/modules/product/productEnum"
)

type (
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

	GetCategoryListModel struct {
		base_model.QueryParamsPage `gorm:"-"`
		Pid                        uint64 `form:"pid" gorm:"-" json:"pid,omitempty"`
		Id                         uint64 `form:"id" gorm:"-" json:"id,omitempty"`
		Name                       string `form:"name" gorm:"-" json:"name,omitempty"`
		Status                     int16  `form:"status" gorm:"-" json:"status,omitempty"`
		MerId                      uint64 `json:"merId"`
		IsSystem                   bool   `json:"isSystem"`
	}

	GetMerchantCategoryListModel struct {
		base_model.QueryParamsPage `gorm:"-"`
		MerId                      uint64 `form:"merId" json:"merId"`
		Pid                        uint64 `form:"pid" gorm:"-" json:"pid,omitempty"`
		Id                         uint64 `form:"id" gorm:"-" json:"id,omitempty"`
		Name                       string `form:"name" gorm:"-" json:"name,omitempty"`
	}

	GetAttributeListModel struct {
		base_model.QueryParamsPage `gorm:"-"`
		CategoryId                 uint64 `form:"categoryId" gorm:"-" json:"categoryId,omitempty"`
		Name                       string `form:"name" gorm:"-" json:"name,omitempty"`
		MerId                      uint64 `json:"merId"`
	}

	GetAttributeOptionListModel struct {
		base_model.QueryParamsPage `gorm:"-"`
		AttributeId                uint64 `form:"attributeId" gorm:"-" json:"attributeId,omitempty"`
		Name                       string `form:"name" gorm:"-" json:"name,omitempty"`
	}

	GetAttributeValueListModel struct {
		base_model.QueryParamsPage `gorm:"-"`
		AttributeId                uint64 `form:"attributeId" gorm:"-" json:"attributeId,omitempty"`
		Name                       string `form:"name" gorm:"-" json:"name,omitempty"`
		MerId                      uint64 `form:"merId" json:"merId"`
	}

	GetPropertyListModel struct {
		base_model.QueryParamsPage `gorm:"-"`
		PropertyId                 uint64 `form:"propertyId" gorm:"-" json:"propertyId,omitempty"`
		Name                       string `form:"name" gorm:"-" json:"name,omitempty"`
		MerId                      uint64 `json:"merId"`
	}

	GetTagsListModel struct {
		base_model.QueryParamsPage `gorm:"-"`
		MerId                      uint64 `form:"merId" json:"merId"`
	}
	GetTagModel struct {
		Name  string `form:"name" json:"name" ` //标签名称
		MerId uint64 `json:"merId"`
	}

	GetAttributeModel struct {
		Name  string `form:"name" json:"name" ` //名称
		MerId uint64 `json:"merId"`
	}

	GetAttributeValueModel struct {
		Name        string `form:"name" json:"name" ` //名称
		AttributeId uint64 `form:"attributeId" json:"attributeId"`
		MerId       uint64 `form:"merId" json:"merId"`
	}
	GetPropertyValueListModel struct {
		base_model.QueryParamsPage `gorm:"-"`
		PropertyId                 uint64 `form:"propertyId" gorm:"-" json:"propertyId,omitempty"`
		Name                       string `form:"name" gorm:"-" json:"name,omitempty"`
	}

	GetProductListModel struct {
		base_model.QueryParamsPage `gorm:"-"`
		CategoryId                 int     `form:"categoryId" gorm:"-" json:"categoryId,omitempty"`
		Name                       string  `form:"name" gorm:"-" json:"name,omitempty"`
		MerId                      uint64  `form:"merId" json:"merId"`
		Status                     int     `json:"status" form:"status"` //未上架、上架、下架 0默认未上架 1 上架 -1下架
		IsDesc                     bool    `json:"isDesc" form:"isDesc"` //排序 是否最近添加
		Lat                        float64 `form:"lat" json:"lat"`       //维度
		Lon                        float64 `form:"lon" json:"lon"`       //经度
	}

	GetProductSmallInfoListModel struct {
		base_model.QueryParamsPage `gorm:"-"`
		MerId                      uint64   `form:"merId" json:"merId"`
		CategoryId                 int      `form:"categoryId" gorm:"-" json:"categoryId,omitempty"`
		Name                       string   `form:"name" gorm:"-" json:"name,omitempty"`
		Type                       int      `form:"type" json:"type"`
		ProductIds                 []uint64 `form:"productIds" json:"productIds"`
		Lat                        float64  `form:"lat" json:"lat"`                         //维度
		Lon                        float64  `form:"lon" json:"lon"`                         //经度
		ComputeDistance            bool     `json:"computeDistance" form:"computeDistance"` //判断是否需要根据经纬度 计算距离
	}

	GetProductDetailInfoSingleModel struct {
		Guid   string `form:"guid" json:"guid"`
		UserId uint64 `json:"userId"`
	}

	GetProductInfoModel struct {
		ProductId uint64 `form:"productId" json:"productId"`
		MerId     uint64 `form:"merId" json:"merId"`
	}

	UpdateProductStatusModel struct {
		ProductId uint64                    `form:"productId" json:"productId"`
		MerId     uint64                    `form:"merId" json:"merId"`
		UserId    uint64                    `json:"userId" form:"userId"`
		Status    productEnum.ProductStatus `json:"status" form:"status"`
	}

	ProductSalesUpdateModel struct {
		ProductId    uint64 `form:"productId" json:"productId"`
		ProductSkuId uint64 `json:"productSkuId" form:"productSkuId"`
		Count        int    `json:"count" form:"count"`
		Type         int    `json:"type" form:"type"` //类型 0 增加 -1 减少
	}
	ProductStockUpdateModel struct {
		ProductId    uint64 `form:"productId" json:"productId"`
		ProductSkuId uint64 `json:"productSkuId" form:"productSkuId"`
		Count        int    `json:"count" form:"count"`
		Type         int    `json:"type" form:"type"` //类型 0 增加 -1 减少
	}

	GetUserProductCollectionListModel struct {
		base_model.QueryParamsPage `gorm:"-"`
		UserId                     uint64  `json:"userId" form:"userId"`
		Lat                        float64 `form:"lat" json:"lat"` //维度
		Lon                        float64 `form:"lon" json:"lon"` //经度

	}

	GetProductPaymentTypeListModel struct {
		UserId    uint64 `json:"userId"`
		ProductId uint64 `json:"productId" form:"productId"`
	}
	ProductPaymentTypeModel struct {
		PaymentTypeId uint64 `json:"id"`
		Code          string `json:"code"`
		Name          string `json:"name"`
	}
	UserProductCollectionAddModel struct {
		UserId uint64 `json:"userId"`
		//ProductId   uint64 `json:"productId"`
		ProductCode string `json:"productCode"`
	}

	UserProductCollectionRemoveModel struct {
		UserId uint64 `json:"userId"`
		//ProductId   uint64 `json:"productId"`
		ProductCode string `json:"productCode"`
	}
	ProductDomainModel struct {
		model.Product
		PaymentTypeIds []string                         `json:"paymentTypeIds"`
		CategoryIds    []uint64                         `json:"categoryIds"`
		CategoryInfos  []ProductDomainCategoryInfoModel `json:"categoryInfos"`
	}

	ProductDomainCategoryInfoModel struct {
		Id   uint64 `json:"categoryId"`
		Name string `json:"categoryName"`
	}
	ProductSkuDomainModel struct {
		model.ProductSKU
	}
)
