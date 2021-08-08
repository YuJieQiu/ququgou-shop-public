package model

import (
	"github.com/ququgou-shop/library/base_model"
	"github.com/ququgou-shop/modules/product/shop_ext_struct"
)

//产品sku
type ProductSKU struct {
	base_model.IDAutoModel
	Guid              string                                       `json:"guid" gorm:"column:guid;unique_index;not null;unique;"`
	ProductId         uint64                                       `json:"productId" gorm:"column:product_id"`
	Name              string                                       `json:"name" gorm:"column:name"`                                    //名称
	Code              string                                       `json:"code" gorm:"column:code"`                                    //商品编码 TODO: 商品编码规则
	BarCode           string                                       `json:"barCode" gorm:"column:bar_code"`                             //商品条形码
	OriginalPrice     float64                                      `json:"originalPrice" gorm:"column:original_price"`                 //原价
	Price             float64                                      `json:"price" gorm:"column:price;type:decimal(18,2)"`               //价格
	Stock             int                                          `json:"stock" gorm:"column:stock"`                                  //库存
	LockStock         int                                          `json:"lockStock" gorm:"column:lock_stock"`                         //锁定库存
	LowStock          int                                          `json:"lowStock" gorm:"column:low_stock"`                           //预警库存
	Width             float32                                      `json:"width" gorm:"column:width;type:decimal(10,2)"`               //宽
	Height            float32                                      `json:"height" gorm:"column:height;type:decimal(10,2)"`             //高
	Depth             float32                                      `json:"depth" gorm:"column:depth;type:decimal(10,2)"`               //深度
	Weight            float32                                      `json:"weight" gorm:"column:weight;type:decimal(10,2)"`             //重量
	Sales             int                                          `json:"sales" gorm:"column:sales"`                                  //sku的销量
	Status            int16                                        `json:"status" gorm:"column:status"`                                //默认 0
	Sort              int                                          `json:"sort" gorm:"column:sort"`                                    //排序
	ResourceId        uint64                                       `json:"resourceId" gorm:"column:resource_id"`                       //图片 选填
	AttributeInfo     shop_ext_struct.SkuAttributeValuesArrayModel `json:"attributeInfo" gorm:"column:attribute_info;type:text(1000)"` //规格属性信息 ，冗余字段 更新商品sku 规格属性需要更新该字段
	IsSingleAttribute bool                                         `json:"isSingleAttribute" gorm:"column:is_single_attribute"`        //单规格
	base_model.ImageJsonSingleModel
	base_model.TimeAllModel

	//业务字段 不是数据库字段
	AttributeValues []AttributeValue `json:"attributeValues,omitempty" gorm:"-"`
}

// Set table name
func (ProductSKU) TableName() string {
	return "product_sku"
}

//设置图片路径 TODO:待优化
func (s *ProductSKU) SetImagesUrl(imgServiceUrl string) {
	s.ImageJson.Url = imgServiceUrl + s.ImageJson.Url
}
