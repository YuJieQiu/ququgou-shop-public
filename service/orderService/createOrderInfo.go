package orderService

import (
	"github.com/jinzhu/gorm"
	"github.com/ququgou-shop/library/utils"
	merchantModel "github.com/ququgou-shop/modules/merchant/model"
	"github.com/ququgou-shop/modules/order/common"
	orderModel "github.com/ququgou-shop/modules/order/model"
	"github.com/ququgou-shop/modules/order/orderEnum"
	paymentModel "github.com/ququgou-shop/modules/payment/model"
	"github.com/ququgou-shop/modules/payment/paymentEnum"
	productModel "github.com/ququgou-shop/modules/product/model"
	"github.com/ququgou-shop/modules/shopcart"
	userModel "github.com/ququgou-shop/modules/user/model"
	"github.com/ququgou-shop/service/shopCartService"
)

var shopCartModule = shopcart.ModuleShopCart{}

//创建订单信息
func CreateOrderInfo(db *gorm.DB, q *CreateOrderInfoModel, u *userModel.User) (error, *CreateOrderInfoResultModel) {

	var (
		err            error
		productList    *[]productModel.Product
		productSkuList *[]productModel.ProductSKU
		merchantList   *[]merchantModel.Merchant
		order          orderModel.Order
		paymentType    *paymentModel.PaymentType
		res            CreateOrderInfoResultModel
	)

	//TODO:验证用户信息(暂时不需要，主要功能是判断用户是否有购买权限，是否多次购买 限购商品等等)

	//获取商品和sku的信息
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

	//TODO:验证优惠卷信息(暂无 优惠系统)
	err = checkProductDiscounts(q, productSkuList, u)
	if err != nil {
		return err, nil
	}

	//验证总价 未优惠
	err = checkProductAmountPrice(q, productSkuList)
	if err != nil {
		return err, nil
	}

	//验证订单总价
	err = checkOrderAmountPrice(q)
	if err != nil {
		return err, nil
	}

	//验证地址信息 这个应该是用户选择地址的时候就进行验证 ，这里就不验证了

	//验证支付类型信息 根据存入的支付类型ID 进行查询判断
	// TODO:根据商品或者商户、用户的支付方式判断（暂不需要）
	err, paymentType = checkOrderPayType(db, q)
	if err != nil {
		return err, nil
	}

	//TODO:运费、配送 类型验证 ，暂无

	//获取商家信息 包含多商户
	err, merchantList = getProdcutMerchantInfo(db, productList)
	if err != nil {
		return err, nil
	}

	//TODO：订单类型 从前端传入 支付状态根据 订单类型： 如是线上订单 支付状态为待支付 ，如果线下支付类型 支付状态默认
	orderType := getOrderType(paymentType)

	//订单装配
	order = orderModel.Order{
		OrderNo: common.CreateOrderNo(u.ID),
		UserId:  u.ID,

		OrderStatus:          int(orderEnum.OrderCreate),
		PayStatus:            int(orderEnum.PayWait),
		ProductAmountTotal:   q.ProductAmountTotal,
		DiscountsAmountTotal: q.DiscountsAmountTotal,
		OrderAmountTotal:     q.OrderAmountTotal,
		//DeliveryFee:          q.DeliveryFee,
		DeliveryType: q.DeliveryTypeId,
		AddressId:    q.AddressId,
		Remark:       q.Remark,
		//SourceType:           0,
		Type: int(orderType),
		//PaymentTypeId:        paymentType.ID,
		//PaymentTypeCode:      paymentType.Code,
	}

	//创建子订单根据是否多商户，或者物流分包等规则 暂根据是否多商户订单进行分包 TODO: 可以添加更多的分包规则
	//子订单就是用户看到的信息，相当于总的订单

	//开启事务 进行创建
	tx := db.Begin()

	err = tx.Create(&order).Error
	if err != nil {
		tx.Rollback()
		return err, nil
	}

	for _, v := range *merchantList {

		ps, pskus := getMerProductAndSku(v.ID, productList, productSkuList)

		ds := q.DiscountsAmountTotal / float64(len(*merchantList))

		err = createOrderSub(tx, q, &order, &v, ps, pskus, u, ds, orderType)
		if err != nil {
			tx.Rollback()
			return err, nil
		}
	}

	tx.Commit()

	res = CreateOrderInfoResultModel{
		OrderId:         order.ID,
		OrderNo:         order.OrderNo,
		UserId:          u.ID,
		PaymentTypeId:   paymentType.ID,
		PaymentTypeCode: paymentType.Code,
	}

	return nil, &res
}

//创建子订单
func createOrderSub(tx *gorm.DB,
	q *CreateOrderInfoModel,
	o *orderModel.Order, //当前主订单信息
	mer *merchantModel.Merchant, //当前商户
	productList *[]productModel.Product, //单商户的产品列表
	productSkuList *[]productModel.ProductSKU, //单商户的产品sku列表
	u *userModel.User, //当前购买用户
	discountsAmount float64, //优惠金额(根据子订单平摊的优惠金额)
	orderType orderEnum.OrderType,
) error {

	var (
		err              error
		orderSub         orderModel.Order
		orderDetailList  []interface{}
		orderProductList []interface{}
		amountPrice      float64 //总金额

	)

	//TODO:暂时不考虑物流拆包情况

	for _, v := range *productSkuList {
		for _, k := range q.Products {
			if v.ID == k.ProductSkuId {
				amountPrice += v.Price * float64(k.ProductNumber)
			}
		}
	}

	orderSub = orderModel.Order{
		UserId:               u.ID,
		OrderNo:              common.CreateSubOrderNo(mer.ID, u.ID),
		OrderMasterId:        o.ID,
		OrderMasterNo:        o.OrderNo,
		MerId:                mer.ID,
		OrderStatus:          int(orderEnum.OrderCreate),
		PayStatus:            int(orderEnum.PayWait),
		DeliveryStatus:       int(orderEnum.DeliveryDefaultStatus),
		ProductAmountTotal:   amountPrice,
		DiscountsAmountTotal: discountsAmount,
		OrderAmountTotal:     amountPrice - discountsAmount,
		Remark:               o.Remark,
		DeliveryType:         o.DeliveryType,
		AddressId:            o.AddressId,
		Type:                 int(orderType),
	}

	err = tx.Create(&orderSub).Error
	if err != nil {
		return err
	}

	//订单详情
	for _, v := range *productSkuList {
		p := getProduct(productList, v.ProductId)
		cp := getCreateOrderProductModel(&q.Products, v.ID)

		da := discountsAmount / float64(len(*productSkuList))

		od := orderModel.OrderDetail{
			UserId:                  u.ID,
			OrderId:                 orderSub.ID,
			ProductId:               p.ID,
			ProductSkuId:            v.ID,
			ProductSkuAttributeInfo: v.AttributeInfo,
			ProductCount:            cp.ProductNumber,
			ProductUnitPrice:        v.Price,
			OriginalAmountTotal:     v.Price * float64(cp.ProductNumber),
			DiscountsAmountTotal:    da,
			AmountTotal:             v.Price*float64(cp.ProductNumber) - da,
			Status:                  0,
			Type:                    0,
		}

		orderDetailList = append(orderDetailList, od)

		//订单产品信息添加
		oproduct := orderModel.OrderProduct{
			UserId:             u.ID,
			OrderId:            orderSub.ID,
			ProductId:          p.ID,
			ProductSkuId:       v.ID,
			MerId:              mer.ID,
			ProductName:        p.Name,
			ProductDescription: p.Description,
			OriginalPrice:      v.OriginalPrice,
			Price:              v.Price,
			LocationId:         0, //TODO: 位置ID 暂留
			Width:              v.Width,
			Height:             v.Height,
			Depth:              v.Depth,
			Weight:             v.Weight,
			ProductType:        p.ProductType,
			CoverImage:         p.ImageJson[0].Url,
		}

		orderProductList = append(orderProductList, oproduct)

		//库存更新
		//TODO:商品库存 减少 目前直接在ProductSKU 表中删除，后面添加库存系统
		err = tx.Model(v).Where("stock > ?", cp.ProductNumber).UpdateColumn("stock", gorm.Expr("stock - ?", cp.ProductNumber)).Error
		if err != nil {
			return err
		}

	}

	//TODO: TTT gormbulk
	//err = gormbulk.BulkInsert(tx, orderDetailList, 3000)
	//if err != nil {
	//	return err
	//}
	//
	//err = gormbulk.BulkInsert(tx, orderProductList, 3000)
	//if err != nil {
	//	return err
	//}

	//判断是否购物车商品(删除购物车商品)
	var cartIds []uint64
	for _, k := range q.Products {
		if k.ShopCartId > 0 {
			cartIds = append(cartIds, k.ShopCartId)
		}
	}

	if len(cartIds) > 0 {
		err = shopCartService.DeleteShopCart(tx, &shopCartService.DeleteShopCartModel{
			UserId: u.ID,
			Ids:    cartIds,
		})
		if err != nil {
			return err
		}
	}

	//订单日志
	olog := orderModel.OrderLogs{
		UserId:   u.ID,
		UserGuid: u.Guid,
		OrderId:  o.ID,
		//OrderSubId: orderSub.ID,
		Status: o.OrderStatus,
		Type:   1,
		Remark: "订单创建",
	}

	err = tx.Create(&olog).Error
	if err != nil {
		return err
	}

	return nil
}

//获取商户商品和sku
func getMerProductAndSku(merId uint64, productList *[]productModel.Product, productSkuList *[]productModel.ProductSKU) (*[]productModel.Product, *[]productModel.ProductSKU) {
	var (
		pList    []productModel.Product
		pSkuList []productModel.ProductSKU
	)

	for _, v := range *productList {
		if v.MerId == merId {
			pList = append(pList, v)
			for _, k := range *productSkuList {
				if k.ProductId == v.ID {
					pSkuList = append(pSkuList, k)
				}
			}
		}
	}

	return &pList, &pSkuList
}

//获取产品的商家信息
func getProdcutMerchantInfo(db *gorm.DB, productList *[]productModel.Product) (error, *[]merchantModel.Merchant) {

	var (
		err    error
		merIds []uint64
		mers   []merchantModel.Merchant //商家信息
	)

	for _, v := range *productList {
		if merIds == nil || len(merIds) <= 0 || !utils.Uint64ArrayExistItem(&merIds, v.MerId) {
			merIds = append(merIds, v.MerId)
			continue
		}
	}

	err = db.Where("id in (?)", merIds).Find(&mers).Error
	if err != nil {
		return err, nil
	}

	if mers == nil || len(mers) <= 0 {
		return ErrMerchantsInvalid, nil
	}
	return nil, &mers
}

//验证商品总价 未优惠前的
func checkProductAmountPrice(q *CreateOrderInfoModel, productSkuList *[]productModel.ProductSKU) error {

	var amountPrice float64

	for _, v := range q.Products {
		amountPrice += float64(v.ProductNumber) * v.ProductUnitPrice
	}

	if amountPrice != q.ProductAmountTotal {
		return ErrProductAmountPriceInvalid
	}

	return nil
}

//验证商品优惠 价格和优惠卷等信息
func checkProductDiscounts(q *CreateOrderInfoModel, productSkuList *[]productModel.ProductSKU, u *userModel.User) error {
	//TODO : 查询错所有的优惠卷是否正确，然后判断优惠卷金额相加是否等于总的优惠金额

	if q.Discounts == nil || len(q.Discounts) <= 0 {
		//优惠卷信息都没有，还能有优惠金额？？？
		if q.DiscountsAmountTotal != 0 {
			return ErrDiscountsAmountPriceInvalid
		}
	}
	return nil
}

//验证订单总价
func checkOrderAmountPrice(q *CreateOrderInfoModel) error {

	amount := q.ProductAmountTotal - q.DiscountsAmountTotal

	if q.OrderAmountTotal != amount {
		return ErrOrderAmountPriceInvalid
	}

	return nil
}

//验证订单支付类型信息 暂时就一个微信支付 默认返回true
func checkOrderPayType(db *gorm.DB, q *CreateOrderInfoModel) (error, *paymentModel.PaymentType) {
	var (
		err error
		p   paymentModel.PaymentType
	)
	err = db.Where("id= ? ", q.PaymentTypeId).First(&p).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return err, nil
	}

	if p.ID <= 0 {
		return ErrOrderPaymentTypeInvalid, nil
	}

	return nil, &p
}

//获取订单类型
//根据支付类型判断
//如果是微信等线上支付的 为线上订单 否则为线下订单类型
func getOrderType(p *paymentModel.PaymentType) orderEnum.OrderType {
	if paymentEnum.PaymentType(p.Code) == paymentEnum.PaymentTypeOfflinePay {
		return orderEnum.OrderTypeOffline
	} else {
		return orderEnum.OrderTypeOnline
	}
}

//
//func BulkInsertPids(db *gorm.DB, posts []*orderModel.OrderDetail) error {
//
//	fieldNums := 6
//	quesMarkString := "("
//
//	for i := 0; i < fieldNums; i++ {
//		quesMarkString += "?, "
//	}
//
//	quesMarkString = quesMarkString[: len(quesMarkString) - 2] + ")"
//
//	valueStrings := make([]string, 0, len(posts))
//	valueArgs := make([]interface{}, 0, len(posts) * fieldNums)
//
//	for _, post := range posts {
//		valueStrings = append(valueStrings, quesMarkString)
//		valueArgs = append(valueArgs, post.Id)
//		valueArgs = append(valueArgs, post.Name)
//		valueArgs = append(valueArgs, post.Cid)
//		valueArgs = append(valueArgs, post.CreatedTime)
//		valueArgs = append(valueArgs, post.X)
//		valueArgs = append(valueArgs, post.Y)
//	}
//	stmt := fmt.Sprintf("INSERT INTO pid (id, name, cid, created_time, x, y) VALUES %s", strings.Join(valueStrings, ","))
//	err := db.Exec(stmt, valueArgs...).Error
//	return err
//}
