package productService

import (
	"github.com/ququgou-shop/library/ext/ext_struct"
	"github.com/ququgou-shop/modules/product/shop_ext_struct"
)

//获取单个商品详细信息
type ProductDetailInfoSingle struct {
	Guid                string                      `json:"guid"`
	TypeId              uint64                      `json:"typeId"`             //类型ID
	TypeName            string                      `json:"typeName,omitempty"` //类型名称 暂无
	BrandId             uint64                      `json:"brandId"`            //品牌编号
	Brand               interface{}                 `json:"brand"`              //品牌信息 暂无
	MerId               uint64                      `json:"merId"`              //商家Id
	MerCode             string                      `json:"merCode"`
	MerName             string                      `json:"merName"`       //商家名称
	MerInfo             interface{}                 `json:"merInfo"`       //商家详细信息 暂无
	Name                string                      `json:"name"`          //商品名称
	Description         string                      `json:"description"`   //描述信息
	Content             string                      `json:"content"`       //富文本描述
	Keywords            ext_struct.JsonStringArray  `json:"keywords"`      //商品关键字
	Tags                *[]ProductTagsModel         `json:"tags"`          //标签
	OriginalPrice       float64                     `json:"originalPrice"` //原始价格
	CurrentPrice        float64                     `json:"currentPrice"`  //当前销售价格 展示使用 无其它作用
	MinPrice            float64                     `json:"minPrice"`      //最低价格
	MaxPrice            float64                     `json:"maxPrice"`      //最高价格
	Sales               int                         `json:"sales"`         //销量
	Image               string                      `json:"image"`         //封面 url
	Stock               int                         `json:"stock"`         //库存
	IsSingle            bool                        `json:"isSingle"`      //是否单品
	Attributes          *[]ProductAttributeModel    `json:"attributes"`    //规则选项 信息
	SkuInfo             *[]ProductSKUDetailModel    `json:"skuInfo"`       //sku 信息
	Resources           *[]ProductResourceModel     `json:"resources"`     //图片资源信息
	ProductDeliveryType *[]ProductDeliveryTypeModel `json:"deliveryTypes"` //商品交付方式 1、快递 5、线下自提 10、
	Status              int                         `json:"status"`        //状态
	StatusText          string                      `json:"statusText"`
	Collected           bool                        `json:"collected"` //是否收藏商品
}

//商品交付方式 	1、快递 5、线下自提 10、
type ProductDeliveryTypeModel struct {
	DeliveryTypeId uint64 `json:"deliveryTypeId" gorm:"column:delivery_type_id;"` //
	Name           string `json:"name"`
}

//商品标签
type ProductTagsModel struct {
	Id   uint64 `json:"id"`
	Name string `json:"name"`
}

//商品sku
type ProductSKUDetailModel struct {
	Id                uint64                                       `json:"id" gorm:"column:sku_id"`
	SkuImage          ext_struct.JsonImageString                   `json:"skuImage" gorm:"column:sku_image"`
	SkuName           string                                       `json:"skuName" gorm:"column:sku_name"`
	Price             float64                                      `json:"price" gorm:"column:price"`
	Stock             int                                          `json:"stock" gorm:"column:stock"`
	Status            int                                          `json:"status" gorm:"column:status"`
	Sort              int                                          `json:"sort" gorm:"column:sort"`
	AttributeInfo     shop_ext_struct.SkuAttributeValuesArrayModel `json:"attributeInfo" gorm:"column:attribute_info"`          //规格属性信息 ，冗余字段 更新商品sku 规格属性需要更新该字段
	IsSingleAttribute bool                                         `json:"isSingleAttribute" gorm:"column:is_single_attribute"` //单规格
	//PropPath map[uint64]uint64          `json:"propPath"`
	//AttributeId uint64 `json:"attribute_id"`
	//AttributeOptionId uint64 `json:"attribute_option_id"`
}

type ProductAttributeModel struct {
	AId     uint64                         `json:"aid"`
	Name    string                         `json:"name"`
	Options []ProductAttributesOptionModel `json:"options"`
}

type ProductAttributesOptionModel struct {
	VId   uint64 `json:"vid"`
	Name  string `json:"name"`
	Image string `json:"image"`
}

type ProductOptionModel struct {
	Id   uint64          `json:"id"`
	Name string          `json:"name"`
	SKU  ProductSKUModel `json:"sku"`
}

type ProductSKUModel struct {
	Id     uint64                          `json:"id"`
	Name   string                          `json:"name"`
	Image  ext_struct.JsonImageArrayString `json:"image"`
	Price  float64                         `json:"price"`
	Stock  int                             `json:"stock"`
	Status int                             `json:"status"`
	Sort   int                             `json:"sort"`
}

type ProductResourceModel struct {
	Guid string `json:"guid"`
	Url  string `json:"url"`
}

type productAttrOptionModel struct {
	AttributeId         uint64 `json:"attributeId"`
	AttributeName       string `json:"attributeName"`
	AttributeOptionId   uint64 `json:"attributeOptionId"`
	AttributeOptionName string `json:"attributeOptionName"`
}

//Product Info
type ProductSmallInfoModel struct {
	Guid          string                     `json:"guid" gorm:"column:guid"`
	MerId         uint64                     `json:"merId" gorm:"column:mer_id;"`
	MerName       string                     `json:"merName" gorm:"column:mer_name;"`
	TypeId        uint64                     `json:"typeId" gorm:"column:type_id"`
	BrandId       uint64                     `json:"brandId" gorm:"column:brand_id"`             //品牌编号 暂留
	Brand         interface{}                `json:"brand"`                                      //品牌信息 暂无
	Name          string                     `json:"name"`                                       //商品名称
	Description   string                     `json:"description"`                                //描述信息
	Keywords      ext_struct.JsonStringArray `json:"keywords"`                                   //商品关键字
	Tags          ext_struct.JsonStringArray `json:"tags"`                                       //标签
	CurrentPrice  float64                    `json:"currentPrice" gorm:"column:current_price"`   //当前销售价格
	OriginalPrice float64                    `json:"originalPrice" gorm:"column:original_price"` //原始价格
	MinPrice      float64                    `json:"minPrice" gorm:"column:min_price"`           //现在 最低价格
	MaxPrice      float64                    `json:"maxPrice" gorm:"column:max_price"`           //
	ProductType   int                        `json:"productType" gorm:"column:product_type"`     //产品类型     0:商品(默认) 、1:服务
	Sales         int                        `json:"sales"`                                      //销量
	Image         string                     `json:"image"`                                      //封面 url
	Stock         int                        `json:"stock"`                                      //库存
	Latitude      float64                    `json:"latitude" gorm:"column:latitude"`            //纬度
	Longitude     float64                    `json:"longitude" gorm:"column:longitude"`          //经度                          //经度
	Distance      float64                    `json:"distance"`                                   //距离 单位 KM
}
