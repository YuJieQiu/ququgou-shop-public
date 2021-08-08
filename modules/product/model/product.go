package model

import (
	"github.com/ququgou-shop/library/base_model"
	"github.com/ququgou-shop/library/ext/ext_struct"
)

//产品
type Product struct {
	base_model.IDAutoModel
	Guid              string                     `json:"guid" gorm:"column:guid;unique_index;not null;unique;"`
	MerId             uint64                     `json:"merId" gorm:"column:mer_id"`                                    //商户ID
	TypeId            uint64                     `json:"typeId" gorm:"column:type_id"`                                  //类型编号   默认 0  [1、套餐 2、虚拟产品 3、积分产品]
	BrandId           uint64                     `json:"brandId" gorm:"column:brand_id"`                                //品牌Id
	Name              string                     `json:"name" gorm:"column:name;index:name"`                            //商品名称
	Status            int                        `json:"status" gorm:"column:status"`                                   //未上架、上架、下架  0默认 未上架 1 上架 -1下架
	Description       string                     `json:"description" gorm:"column:description;type:varchar(2000)"`      //描述信息
	Keywords          ext_struct.JsonStringArray `json:"keywords" gorm:"column:keywords"`                               //商品关键字 展示使用 暂无其它作用 TODO:后期优化
	OriginalPrice     float64                    `json:"originalPrice" gorm:"column:original_price;type:decimal(18,2)"` //原始价格 展示使用 无其它作用
	MinPrice          float64                    `json:"minPrice" gorm:"column:min_price;type:decimal(18,2)"`           //最低价格 展示使用 无其它作用
	MaxPrice          float64                    `json:"maxPrice" gorm:"column:max_price;type:decimal(18,2)"`           //最高价格 展示使用 无其它作用
	CurrentPrice      float64                    `json:"currentPrice"  gorm:"column:current_price;type:decimal(18,2)"`  //当前销售价格 展示使用 无其它作用
	Sales             int                        `json:"sales" gorm:"column:sales"`                                     //销量  统计 该产品 所有 sku的销量
	ProductType       int                        `json:"productType" gorm:"column:product_type"`                        //产品类型   0 默认 产品暂时就这一种  0:商品(默认) 、1:服务
	Width             float32                    `json:"width" gorm:"column:width;type:decimal(10,2)"`                  //宽
	Height            float32                    `json:"height" gorm:"column:height;type:decimal(10,2)"`                //高
	Depth             float32                    `json:"depth" gorm:"column:depth;type:decimal(10,2)"`                  //深度 (长)
	Weight            float32                    `json:"weight" gorm:"column:weight;type:decimal(10,2)"`                //重量
	Integral          int                        `json:"integral" gorm:"column:integral"`                               // 可以使用积分抵消
	Active            bool                       `json:"active" gorm:"column:active"`                                   //是否启用
	IsSingle          bool                       `json:"isSingle" gorm:"column:is_single"`
	RecommendPriority int                        `json:"recommendPriority" gorm:"column:recommend_priority;"` //推荐优先级别 级别越高 升序
	base_model.ImageJsonModel
	base_model.TimeAllModel

	//业务字段 不是数据库字段
	StatusText      string              `json:"statusText" gorm:"-"`
	Content         *ProductContent     `json:"content" gorm:"-"`
	SKU             []ProductSKU        `json:"sku" gorm:"-"`
	Property        []Property          `json:"property" gorm:"-"`
	Tags            []ProductTags       `json:"tags" gorm:"-"`
	CategoryProduct []CategoryProduct   `json:"categoryProduct" gorm:"-"`
	Attributes      []Attribute         `json:"attributes" gorm:"-"`
	Resources       []ProductResource   `json:"resources" gorm:"-"`
	CreatedTime     ext_struct.JsonTime `json:"createdTime" gorm:"-"`
}

//TODO:添加商品类型 套餐产品、虚拟产品、积分产品
//Tags ext_struct.JsonStringArray `json:"tags"` //标签
//Content string `json:"content" gorm:"type:longtext"` //商品富文本信息
//ResourceId int `json:"resource_id"` //封面图片
//IsSingle bool `json:"is_single"` //是否是单品 ，如果单品则不启用sku 单品 不需要查询sku的产品 不需要，单品默认就一个sku
//IsPackage bool `json:"is_package"` //是否套餐
//IsVirtual bool `json:"is_virtual"` //是否虚拟产品
//IsIntegral bool `json:"is_integral"` //是否积分产品
//Location string `json:"location"` //位置 暂留

// Set table name
func (Product) TableName() string {
	return "products"
}

//设置图片路径 TODO:待优化
func (p *Product) SetImagesUrl(imgServiceUrl string) {
	for k := 0; k < len(p.ImageJson); k++ {
		p.ImageJson[k].Url = imgServiceUrl + p.ImageJson[k].Url
	}
}
