package orderService

import (
	"github.com/ququgou-shop/library/base_model"
	"github.com/ququgou-shop/library/ext/ext_struct"
	"github.com/ququgou-shop/modules/product/shop_ext_struct"
)

type (
	//创建订单信息
	//这是提交订单的时候才创建的，所有很多东西是在提交订单之前查询出来的
	CreateOrderInfoModel struct {
		Products             []CreateOrderProductModel   `json:"products" binding:"required"`           //商品信息 多个sku
		ProductAmountTotal   float64                     `json:"productAmountTotal" binding:"required"` //商品总价 未优惠
		DiscountsAmountTotal float64                     `json:"discountsAmountTotal"`                  //优惠金额
		OrderAmountTotal     float64                     `json:"orderAmountTotal" binding:"required"`   //订单总价  优惠
		AddressId            uint64                      `json:"addressId"`                             //收货地址Id
		Remark               string                      `json:"remark"`                                //备注信息
		DeliveryTypeId       int                         `json:"deliveryTypeId" binding:"required"`     //配送方式 支持的配送方式（快递，自取，送货上门）
		DeliveryFee          float64                     `json:"deliveryFee" `                          //配送费用
		PaymentTypeId        uint64                      `json:"paymentTypeId" binding:"required"`      //支付方式
		Discounts            []CreateOrderDiscountsModel `json:"discounts"`                             //优惠卷信息(可能使用多个优惠卷)

		//DeliveryAddressId    int                         `json:"deliveryAddressId"`                     //配送点地址信息(用户如果是个自提地点，则可以选择自提点信息)
		//优惠卷应该是单独的一个对象，优惠卷编号、对应的商家编号、对应的商品类型、商品sku 、 优惠条件(满减、无门槛等) 、优惠金额
		//现在不需要使用优惠卷，但是需要留一个入口
		//商品和商户需要挂钩？ no ，后台根据购买的商品和商户id 做个验证判断
		// 需要对传如对商品进行价格判断，商品数量判断、商品是否有效、商户是否在改区域售卖、商品是否属于这个商品、
		//这里需要用到go 的协程 进行提升速度
	}

	//创建订单产品信息
	CreateOrderProductModel struct {
		MerId            uint64  ` json:"merId"`           //商户Id
		ProductNo        string  `json:"productNo"`        //产品编号
		ProductSkuId     uint64  `json:"productSkuId"`     //产品sku
		ProductNumber    int     `json:"productNumber"`    //产品数量
		ProductAmount    float64 `json:"productAmount"`    //商品总价价格
		ProductUnitPrice float64 `json:"productUnitPrice"` //商品单价
		ShopCartId       uint64  `json:"shopCartId"`       //购物车信息
	}

	//创建订单是使用的优惠卷信息
	CreateOrderDiscountsModel struct {
		DiscountsNo          string  `json:"discounts_no"`           //优惠卷编号
		DiscountsAmountTotal float64 `json:"discounts_amount_total"` //优惠金额
	}

	//创建订单地址信息 (暂留)
	CreateOrderAddressModel struct {
		AddressId uint64 `json:"address_id"` //地址Id
	}

	//创建订单返回信息
	CreateOrderInfoResultModel struct {
		OrderId         uint64 `json:"orderId"`
		OrderNo         string `json:"orderNo"`
		UserId          uint64 `json:"userId"`
		PaymentTypeId   uint64 `json:"paymentTypeId"`
		PaymentTypeCode string `json:"paymentTypeCode"`
	}

	//获取订单创建前信息
	GetBeforeOrderCreateInfoModel struct {
		Products []CreateOrderProductModel `json:"products" binding:"required"` //所购买的商品
	}

	GetUserOrderSmallInfoListModel struct {
		base_model.QueryParamsPage
		All    bool `form:"all" json:"all"`
		Status int  `form:"status" json:"status"` //OrderBusinessStatus
	}

	GetOrderSmallInfoListModel struct {
		base_model.QueryParamsPage
		All            bool     `form:"all" json:"all"`
		DisplayStatus  int      `json:"displayStatus" form:"displayStatus"` //显示订单状态
		UserId         uint64   `json:"userId"`
		MerId          uint64   `json:"merId"`
		OrderNo        string   `json:"orderNo" form:"orderNo"` //订单编号
		BusinessStatus []string //业务订单状态
	}

	GetUserOrderDetailModel struct {
		OrderNo string `form:"orderNo" json:"orderNo"`
		MerId   uint64 `json:"merId" form:"merId"`
	}

	CancelUserOrderModel struct {
		OrderNo string `form:"orderNo" json:"orderNo"`
	}

	UserOrderPayModel struct {
		OrderMasterNo string `json:"orderMasterNo"`
		OrderMasterId uint64 `json:"orderMasterId"`
		OrderNo       string `json:"orderNo"`
		OrderId       uint64 `json:"orderId"`
		UserId        uint64 `json:"userId"`
	}

	MerSuccessUserOrderModel struct {
		OrderNo       string `form:"orderNo" json:"orderNo"`
		MerId         uint64 `json:"merId" form:"merId"`
		OpenMerUserId uint64 `json:"openMerUserId" json:"openMerUserId"` //操作商户管理员ID
	}

	SmallOrderInfModel struct {
		NO               string                       `json:"no"`      //订单编号 SubOrderNo
		Id               uint64                       `json:"_"`       //订单编号 SubOrderId
		MerId            uint64                       `json:"merId"`   //商户id
		MerName          string                       `json:"merName"` //商户名称
		Status           string                       `json:"status"`  //状态 默认 0 暂留
		StatusText       string                       `json:"statusText"`
		OrderAmountTotal float64                      `json:"orderAmountTotal"` //实际付款金额(优惠后)
		Products         []SmallOrderInfoProductModel `json:"products"`
		CreatedTime      ext_struct.JsonTime          `json:"createdTime"`
		Type             int                          `json:"type"`
		OrderStatus      int                          `json:"_"`
		PayStatus        int                          `json:"_"`
		DeliveryStatus   int                          `json:"_"`
	}

	SmallOrderInfoProductModel struct {
		No          string                                       `json:"no"`
		Cover       string                                       `json:"cover"`
		Name        string                                       `json:"name"`
		Description string                                       `json:"description"`
		SkuInfo     shop_ext_struct.SkuAttributeValuesArrayModel `json:"skuInfo"`
		Count       int                                          `json:"count"`
		UnitPrice   float64                                      `json:"unitPrice"`   //产品单价
		AmountTotal float64                                      `json:"amountTotal"` //付款金额 (优惠后)
		Status      int                                          `json:"status"`      //状态 默认 0 暂留
	}

	UserOrderDetailModel struct {
		NO                   string                       `json:"no" gorm:"column:order_sub_no"`  //订单编号 SubOrderNo
		Id                   uint64                       `json:"-"`                              //订单编号 SubOrderId
		MerId                uint64                       `json:"merId" gorm:"column:mer_id"`     //商户id
		MerName              string                       `json:"merName" gorm:"column:mer_name"` //商户名称
		Status               string                       `json:"status" gorm:"column:status"`    //状态 默认 0 暂留
		StatusText           string                       `json:"statusText"`
		OriginalAmountTotal  float64                      `json:"originalAmountTotal" gorm:"column:product_amount_total"`
		DiscountsAmountTotal float64                      `json:"discountsAmountTotal" gorm:"column:discounts_amount_total"` //优惠金额
		OrderAmountTotal     float64                      `json:"orderAmountTotal" gorm:"column:order_amount_total"`         //实际付款金额(优惠后)
		CreatedTime          ext_struct.JsonTime          `json:"createdTime" gorm:"column:created_at"`
		CancelTime           ext_struct.JsonTime          `json:"cancelTime" gorm:"column:cancel_time"` //取消时间
		Remark               string                       `json:"remark" gorm:"column:remark"`          //备注
		PayTypeText          string                       `json:"payTypeText"`
		PaymentTime          ext_struct.JsonTime          `json:"paymentTime" gorm:"-"` //支付时间
		Address              UserOrderAddressModel        `json:"address" gorm:"-"`     //收货地址
		Products             []SmallOrderInfoProductModel `json:"products" gorm:"-"`
		DeliveryType         int                          `json:"deliveryType" gorm:"-"` //收货方式
	}

	UserOrderAddressModel struct {
		City    string `json:"city" gorm:"column:city"`
		Region  string `json:"region" gorm:"column:region"`
		Town    string `json:"town" gorm:"column:town"`
		Address string `json:"address" gorm:"column:address"`
		Phone   string `json:"phone" gorm:"column:phone"`
		Name    string `json:"name" gorm:"column:name"`
	}

	OrderPaymentSuccessUpdateModel struct {
		BusinessType int    `json:"businessType"` //业务类型
		BusinessNo   string `json:"businessNo"`   //业务编号
	}

	MerGetOrderUserInfoModel struct {
		MerId   uint64 `json:"merId"`
		OrderNo string `json:"orderNo" form:"orderNo"`
	}

	MerOrderUserInfoModel struct {
		UserName string `json:"userName" gorm:"column:user_name;"`
		Mobile   string `json:"mobile"  gorm:"column:mobile;"`
	}
)
