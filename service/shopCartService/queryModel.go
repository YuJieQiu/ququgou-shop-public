package shopCartService

import (
	"time"

	"github.com/ququgou-shop/library/ext/ext_struct"
	"github.com/ququgou-shop/modules/product/shop_ext_struct"
)

type (
	AddShopCartItemModel struct {
		UserId       uint64  `json:"userId"`       //用户Id
		MerId        uint64  `json:"merId"`        //商户ID
		ProductNo    string  `json:"productNo"`    //商品编号 guid
		ProductSkuId uint64  `json:"productSkuId"` //商品skuID
		Number       int     `json:"number"`       //数量
		Price        float64 `json:"price"`        //加入价
		TotalPrice   float64 `json:"totalPrice"`   //加入总价
	}

	GetUserShopCartProductInfoModel struct {
		UserId uint64 `json:"userId"` //用户Id
	}

	DeleteShopCartModel struct {
		Ids    []uint64 `json:"ids"`    //shop cart Id
		UserId uint64   `json:"userId"` //用户Id
		MerId  uint64   `json:"merId"`  //商户ID
	}

	UserShopCartProductModel struct {
		CartID         uint64                                       `json:"cartId"`         //shop cart id
		ProductNo      string                                       `json:"productNo"`      //商品编号
		ProductSkuId   uint64                                       `json:"productSkuId"`   //商品skuID
		Number         int                                          `json:"number"`         //数量
		JoinPrice      float64                                      `json:"joinPrice"`      //加入时的价格
		JoinTotalPrice float64                                      `json:"joinTotalPrice"` //加入时的总价
		JoinTime       time.Time                                    `json:"joinTime"`       //加入时间
		Name           string                                       `json:"name"`           //商品名称
		ProductStatus  int                                          `json:"productStatus"`  //商品状态 未上架、上架、下架 0默认未上架 1 上架 -1下架
		Description    string                                       `json:"description"`    //商品描述信息
		OriginalPrice  float64                                      `json:"originalPrice"`  //原价
		Price          float64                                      `json:"price" `         //价格
		Stock          int                                          `json:"stock"`          //库存
		SkuStatus      int16                                        `json:"skuStatus" `     //sku 状态 默认 0
		AttributeInfo  shop_ext_struct.SkuAttributeValuesArrayModel `json:"attributeInfo" ` //规格属性信息
		Img            ext_struct.JsonImage                         `json:"img"`
	}

	UserShopCartMerModel struct {
		MerId           uint64                     `json:"merId"` //商户ID
		MerCode         string                     `json:"merCode"`
		MerName         string                     `json:"merName"`
		Products        []UserShopCartProductModel `json:"products"`
		InvalidProducts []UserShopCartProductModel `json:"invalidProducts"`
	}
)
