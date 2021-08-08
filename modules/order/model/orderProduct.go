package model

import "github.com/ququgou-shop/library/base_model"

//订单产品表
type OrderProduct struct {
	base_model.IDAutoModel
	MerId              uint64  `json:"merId" gorm:"column:mer_id"` //商户id
	UserId             uint64  `json:"userId" gorm:"column:user_id"`
	OrderId            uint64  `json:"orderId" gorm:"column:order_id"`
	ProductId          uint64  `json:"productId" gorm:"column:product_id"`
	ProductSkuId       uint64  `json:"productSkuId" gorm:"column:product_sku_id"`
	CategoryId         uint64  `json:"categoryId" gorm:"column:category_id"`
	ProductName        string  `json:"productName"  gorm:"column:product_name;index:name"`                      //商品名称
	ProductDescription string  `json:"productDescription" gorm:"column:product_description;type:varchar(2000)"` //商品描述信息
	OriginalPrice      float64 `json:"originalPrice" gorm:"column:original_price;type:decimal(18,2)"`           //原始价格
	Price              float64 `json:"price"  gorm:"column:price;type:decimal(18,2)"`                           //实际价格
	LocationId         uint64  `json:"location" gorm:"column:location_id"`                                      //位置 暂留
	Width              float32 `json:"width" gorm:"column:width;type:decimal(10,2)"`                            //宽
	Height             float32 `json:"height" gorm:"column:height;type:decimal(10,2)"`                          //高
	Depth              float32 `json:"depth" gorm:"column:depth;type:decimal(10,2)"`                            //深度
	Weight             float32 `json:"weight" gorm:"column:weight;type:decimal(10,2)"`                          //重量
	ProductType        int     `json:"productType" gorm:"column:product_type"`
	CoverImage         string  `json:"coverImage"` //封面图片
	base_model.TimeAllModel
}

// Set table name
func (OrderProduct) TableName() string {
	return "order_products"
}
