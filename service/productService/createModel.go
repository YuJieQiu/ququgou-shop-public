package productService

import (
	"github.com/ququgou-shop/library/ext/ext_struct"
	"github.com/ququgou-shop/modules/product/shop_ext_struct"
)

type (
	UpdateProductInfoModel struct {
		Id uint64 `json:"id" binding:"required"` //product id
		CreateProductInfoModel
	}

	//创建商品参数
	CreateProductInfoModel struct {
		//商户ID
		MerId uint64 `json:"merId"` //商户ID
		//品牌id
		BrandId uint64 `json:"brandId"` //品牌id
		//产品所属分类
		CategoryIds []uint64 `json:"categoryIds" binding:"required"` //产品所属分类
		//支付类型ID
		PaymentTypeIds []uint64 `json:"paymentTypeIds"` //支付类型Id
		//名称
		Name string `json:"name" binding:"required"` //名称
		//类型编号    0 默认
		TypeId uint64 `json:"typeId"` //类型编号    0 默认
		//未上架、上架、下架 0默认未上架 1 上架 3 下架
		Status int `json:"status"` //未上架、上架、下架 0默认未上架 1 上架 3 下架  TODO:定时上架功能
		//
		Content CreateContentModel `json:"content"`
		//描述
		Description string `json:"description"` //描述
		//商品关键字
		Keywords ext_struct.JsonStringArray `json:"keywords"` //商品关键字
		//标签
		Tags []CreateProductTagsModel `json:"tags"` //标签
		//原始价格 (下划线价格) 展示使用
		OriginalPrice float64 `json:"originalPrice"` //原始价格 (下划线价格) 展示使用
		//最低价 展示使用
		MinPrice float64 `json:"minPrice"` //最低价 展示使用
		//最高价 展示使用
		MaxPrice float64 `json:"maxPrice"` //最高价 展示使用
		//当前销售价格 展示使用
		CurrentPrice float64 `json:"currentPrice" binding:"required"` //当前销售价格 展示使用
		//销量  统计 该产品 所有 sku的销量
		Sales int `json:"sales"` //销量  统计 该产品 所有 sku的销量 可以手动更改销量😁
		//产品类型   0 默认 产品暂时就这一种  0:商品(默认) 1:服务
		ProductType int `json:"productType"` //产品类型   0 默认 产品暂时就这一种  0:商品(默认) 1:服务
		//宽
		Width float32 `json:"width"` //宽
		//高
		Height float32 `json:"height"` //高
		//深度 (长)
		Depth float32 `json:"depth"` //深度 (长)
		//重量
		Weight float32 `json:"weight"` //重量
		//可以使用积分抵消
		Integral int `json:"integral"` //可以使用积分抵消
		//是否启用
		Active bool `json:"active"` //是否启用
		//图片资源
		Resources []CreateProductResourceModel `json:"resources" binding:"required"` //图片资源 TODO:视频资源
		//sku 信息
		SKU []CreateProductSKUMode `json:"sku" binding:"required"` //sku 信息
		//描述属性
		Property []CreateProductPropertyModel `json:"property"`
		//是否是单品  会有一个SKU，SKU 不会有属性信息
		IsSingle bool `json:"isSingle"`

		RecommendPriority int `json:"recommendPriority"` //推荐级别
	}

	//IsPackage bool `json:"is_package"` //是否套餐
	//IsVirtual bool `json:"is_virtual"` //是否虚拟产品
	//IsIntegral bool `json:"is_integral"` //是否积分产品

	CreateProductResourceModel struct {
		ResourceId uint64 `json:"resourceId"`
		Type       int16  `json:"type"`     //类型 默认 0 图片 、可扩展出  1 视频
		Cover      bool   `json:"cover"`    //是否封面
		Position   int    `json:"position"` //位置 默认 0
	}

	CreateProductSKUMode struct {
		Id uint64 `json:"id"` //skuID
		//SkuId             uint64   `json:"skuId"`
		AttributeValueIds []uint64 `json:"attributeValueIds" binding:"required"` //多个属性集合 (不为单规格时生效)
		Name              string   `json:"name"`                                 //名称
		Code              string   `json:"code"`                                 //TODO:商品编码 应该是后台生成
		BarCode           string   `json:"barCode"`                              //条形码
		OriginalPrice     float64  `json:"originalPrice"`                        //原价
		Price             float64  `json:"price" binding:"required"`             //销售价格
		Stock             int      `json:"stock"`                                //库存
		LowStock          int      `json:"lowStock"`                             //预警库存
		Sort              int      `json:"sort" defutal:"1"`                     //排序
		ResourceId        uint64   `json:"resourceId"`                           //图片
		Width             float32  `json:"width"`                                //宽
		Height            float32  `json:"height"`                               //高
		Depth             float32  `json:"depth"`                                //深度 (长)
		Weight            float32  `json:"weight"`                               //重量
		//AttributeInfo        map[string]string                            `json:"attributeInfo"`
		//SingleAttributeValue string                                       `json:"singleAttributeValue"` //单规格属性值 IsSingleAttribute 为true
		IsSingleAttribute bool                                         `json:"isSingleAttribute"` //是否单规格SKU (如果是 则使用系统默认的Attribute 、只添加属性Value值)
		AttributeInfo     shop_ext_struct.SkuAttributeValuesArrayModel `json:"attributeInfo"`     //
	}

	//SkuAttributeValuesModel struct {
	//	Aid       uint64 `json:"aid"`
	//	AttName   string `json:"attName"`
	//	Vid       uint64 `json:"vid"`
	//	ValueName string `json:"valueName"`
	//}

	CreateProductPropertyModel struct {
		PropertyId uint64 `json:"propertyId"`
		Name       string `json:"name"`
		Title      string `json:"title"`
	}

	CreateProductTagsModel struct {
		TagId uint64 `json:"tagId"`
		Name  string `json:"name"`
	}

	CreateContentModel struct {
		//富文本内容
		Content string `json:"content"` //富文本内容
		//富文本内容 备注
		ContentRemark string `json:"contentRemark"` //富文本内容 备注
	}

	//CreateProductAttributeModel struct {
	//	AttributeId  uint64 `json:"attribute_id"` //属性ID 非必填，如果为0 则通过 attributeName 创建
	//	AttributeName string `json:"attribute_name"`
	//	AttributeValueId uint64 `json:"attribute_value_id"` //属性值ID 非必填，如果为0 则通过 attributeValueName 创建
	//	AttributeValueName string `json:"attribute_value_name"`
	//}
	//创建商品参数 E
)
