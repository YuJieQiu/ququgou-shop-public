package model

import "github.com/ququgou-shop/library/base_model"

type ShopCart struct {
	base_model.IDAutoModel
	UserId         uint64  `json:"userId" gorm:"column:user_id;index;"`
	MerId          uint64  `json:"merId" gorm:"column:mer_id"`                    //商户ID
	ProductNo      string  `json:"productNo" gorm:"column:product_no"`            //商品编号
	ProductSkuId   uint64  `json:"productSkuId" gorm:"column:product_sku_id"`     //商品skuID
	Number         int     `json:"number" gorm:"column:number"`                   //数量
	JoinPrice      float64 `json:"joinPrice" gorm:"column:join_price"`            //加入时的价格
	JoinTotalPrice float64 `json:"joinTotalPrice" gorm:"column:join_total_price"` //加入时的总价
	base_model.TimeAllModel
}

// Set table name
func (ShopCart) TableName() string {
	return "shop_cart"
}
