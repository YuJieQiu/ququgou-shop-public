package orderService

import (
	"github.com/jinzhu/gorm"
	"github.com/ququgou-shop/library/utils"
	productModel "github.com/ququgou-shop/modules/product/model"
	"github.com/ququgou-shop/modules/product/shop_ext_struct"
	userModel "github.com/ququgou-shop/modules/user/model"
)

//创建订单前的信息
type BefOrderInfo struct {
	Merchants       *[]BefOrderMerchantInfo `json:"merchants"`       //商户信息
	FreeDeliveryFee bool                    `json:"freeDeliveryFee"` //免费包邮 TODO://后期加入，暂时都是免邮
}

//商家信息
type BefOrderMerchantInfo struct {
	MerId    uint64                `json:"merId"`
	MerName  string                `json:"merName"`
	Products []BefOrderProductInfo `json:"products"` //产品信息
	MerCover string                `json:"merCover"`
	//可以使用的优惠卷等信息
}

//商品详细信息
type BefOrderProductInfo struct {
	ProductId     uint64                                       `json:"productId"`
	ProductSkuId  uint64                                       `json:"productSkuId"`
	Price         float64                                      `json:"price"`         //现在价格
	OriginalPrice float64                                      `json:"originalPrice"` //原价
	DiscountPrice float64                                      `json:"discountPrice"` //TODO: 优惠金额 暂留 后面优惠卷可能使用到
	Number        int                                          `json:"number"`
	ProductCover  string                                       `json:"productCover"`
	AttributeInfo shop_ext_struct.SkuAttributeValuesArrayModel `json:"attributeInfo"` //规格属性信息(json 字符串)
}

//TODO:如果在生成订单的时候做验证，则不使用该方法
//订单创建前
//获取订单创建前所需要的东西
func GetBeforeOrderCreateInfo(db *gorm.DB, q *GetBeforeOrderCreateInfoModel, user *userModel.User) (error, *BefOrderInfo) {

	var (
		err            error
		productList    *[]productModel.Product
		productSkuList *[]productModel.ProductSKU
		breOrerInfo    BefOrderInfo
		breMers        *[]BefOrderMerchantInfo
	)

	//多个商品的sku
	err, productList, productSkuList = getProductSkuList(db, &q.Products)
	if err != nil {
		return err, nil
	}

	//验证价格信息是否正 发生变化
	err = checkProductPrice(&q.Products, productSkuList)
	if err != nil {
		return err, nil
	}

	//验证库存信息
	err = checkProductStock(&q.Products, productSkuList)
	if err != nil {
		return err, nil
	}

	//根据商家来进行分组

	breMers = befMerProductInfoLoad(&q.Products, productList, productSkuList)

	breOrerInfo.Merchants = breMers

	//TODO:运费计算规则 待写 暂时都是免邮产品
	breOrerInfo.FreeDeliveryFee = true //

	return nil, &breOrerInfo
	//用户地址信息

	//用户可以使用优惠卷信息

	//商品信息

	//商品配送方式

	//商品配置方式下的可选地址信息

	//用户的会员折扣信息 //暂无

	//商家信息

	//平台可选支付信息

	//.......

}

//获取商品sku信息
func getProductSkuList(db *gorm.DB,
	q *[]CreateOrderProductModel) (error, *[]productModel.Product, *[]productModel.ProductSKU) {

	//chSkus chan *[]shopModel.ProductSKU ,
	//	chErr chan error
	var (
		productsNos    = []string{}
		skusIds        = []uint64{}
		productList    []productModel.Product
		productSkuList []productModel.ProductSKU
		err            error
	)
	//query product and sku info , and change infos

	for _, v := range *q {
		if !utils.StringArrayExistItem(&productsNos, v.ProductNo) {
			productsNos = append(productsNos, v.ProductNo)
		}

		if !utils.Uint64ArrayExistItem(&skusIds, v.ProductSkuId) {
			skusIds = append(skusIds, v.ProductSkuId)
		}
	}

	err = db.Where("guid in ( ? ) and status=1 and active=1 ", productsNos).Find(&productList).Error
	if err != nil {
		return err, nil, nil
	}

	if productList == nil || len(productList) <= 0 {
		return ErrProductInvalid, nil, nil
	}

	err = db.Where("id in ( ? )", skusIds).Find(&productSkuList).Error
	if err != nil {
		return err, nil, nil
	}

	if productSkuList == nil || len(productSkuList) <= 0 {
		return ErrProductInvalid, nil, nil
	}

	return nil, &productList, &productSkuList
}

//检查商品的价格是否发生变化
func checkProductPrice(q *[]CreateOrderProductModel, productSkuList *[]productModel.ProductSKU) error {

	for _, v := range *q {
		for _, k := range *productSkuList {
			if v.ProductSkuId == k.ID {
				if v.ProductUnitPrice != k.Price {
					return ErrProductPriceChange
				}
			}
		}
	}

	return nil
}

//验证商品库存信息
func checkProductStock(q *[]CreateOrderProductModel, productSkuList *[]productModel.ProductSKU) error {
	for _, v := range *q {
		for _, k := range *productSkuList {
			if v.ProductSkuId == k.ID {
				if v.ProductNumber > k.Stock {
					return ErrProductStockInvalid
				}
			}
		}
	}

	return nil
}

//商家和商品信息填充
func befMerProductInfoLoad(q *[]CreateOrderProductModel, productList *[]productModel.Product, productSkuList *[]productModel.ProductSKU) *[]BefOrderMerchantInfo {
	var (
		breMers = []BefOrderMerchantInfo{}
	)

	for _, v := range *productSkuList {

		c := getCreateOrderProductModel(q, v.ID)

		p := getProduct(productList, v.ProductId)

		mer := getBbefOrderMerchantInfo(&breMers, p.MerId)

		pInfo := BefOrderProductInfo{
			ProductId:     p.ID,
			ProductSkuId:  v.ID,
			OriginalPrice: v.OriginalPrice,
			Price:         v.Price,
			Number:        c.ProductNumber,
			AttributeInfo: v.AttributeInfo,
			ProductCover:  "", //TODO:商品图片待添加

		}

		if mer.MerId == 0 {
			mer = BefOrderMerchantInfo{
				MerId:    p.MerId,
				MerName:  "", //TODO:商家名称和商家封面信息 暂留 后面商家信息表加入后添加
				MerCover: "",
				Products: []BefOrderProductInfo{},
			}
			mer.Products = append(mer.Products, pInfo)
			breMers = append(breMers, mer)

		} else {
			mer.Products = append(mer.Products, pInfo)
		}
	}

	return &breMers
}

func getBbefOrderMerchantInfo(breMers *[]BefOrderMerchantInfo, merId uint64) BefOrderMerchantInfo {
	for _, v := range *breMers {
		if v.MerId == merId {
			return v
		}
	}
	return BefOrderMerchantInfo{}
}

func getProduct(productList *[]productModel.Product, pid uint64) productModel.Product {

	for _, v := range *productList {
		if v.ID == pid {
			return v
		}
	}
	return productModel.Product{}
}

func getCreateOrderProductModel(q *[]CreateOrderProductModel, productSkuId uint64) CreateOrderProductModel {
	for _, v := range *q {
		if v.ProductSkuId == productSkuId {
			return v
		}
	}
	return CreateOrderProductModel{}
}
