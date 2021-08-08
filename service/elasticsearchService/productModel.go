package elasticsearchService

import (
	"github.com/ququgou-shop/library/ext/ext_struct"
)

type (
	//前期如果商户关闭或者商品下架，直接删除，上架 重新添加 TODO:后期优化
	ElasticProductModel struct {
		ID            uint64              `json:"id"`
		Guid          string              `json:"guid" `
		MerId         uint64              `json:"merId"`         //商户ID
		TypeId        uint64              `json:"typeId"`        //类型编号   默认 0  [1、套餐 2、虚拟产品 3、积分产品]
		BrandId       uint64              `json:"brandId"`       //品牌Id
		BrandName     string              `json:"brandName"`     //品牌名称
		Name          string              `json:"name"`          //商品名称
		Status        int                 `json:"status"`        //0默认 未上架 1 上架 -1下架
		Description   string              `json:"description"`   //描述信息
		Keywords      string              `json:"keywords"`      //商品关键字 展示使用  例如：["",""]
		Tags          string              `json:"tags"`          //标签 例如：["",""]
		OriginalPrice float64             `json:"originalPrice"` //原始价格 展示使用 无其它作用
		MinPrice      float64             `json:"minPrice"`      //最低价格 展示使用 无其它作用
		MaxPrice      float64             `json:"maxPrice"`      //最高价格 展示使用 无其它作用
		CurrentPrice  float64             `json:"currentPrice"`  //当前销售价格 展示使用 无其它作用
		Sales         int                 `json:"sales"`         //销量  统计 该产品 所有 sku的销量
		ProductType   int                 `json:"productType"`   //产品类型   0 默认 产品暂时就这一种  0:商品(默认) 、1:服务
		Integral      int                 `json:"integral"`      //可以使用积分抵消
		Active        bool                `json:"active"`        //是否启用
		City          string              `json:"city"`          //城市(来源商户信息)
		Region        string              `json:"region"`        //区域(来源商户信息)
		Town          string              `json:"town"`          //街道(来源商户信息)
		Address       string              `json:"Address"`       //详细地址(来源商户信息)
		Location      LocationModel       `json:"location"`      //纬经度(来源商户信息)
		Priority      int                 `json:"priority"`      //优先推荐 0 、1、2、3 数字越高 推荐级别越高
		IsSingle      bool                `json:"isSingle"`      //是否单商品
		CategoryIds   []uint64            `json:"categoryIds"`   //分类ID
		CategoryInfo  string              `json:"categoryInfo"`  //分类信息
		CreatedTime   ext_struct.JsonTime `json:"createdTime"`
		UpdatedTime   ext_struct.JsonTime `json:"updatedTime"`
		//SkuInfo       []ElasticProductSKUModel `json:"skuInfo"`
		//CreatedTime   ext_struct.JsonTime        `json:"createdTime"`
		//base_model.ImageJsonModel
		//距离
		//相关性
	}

	ElasticProductSKUModel struct {
		ID                uint64  `json:"id"`
		Guid              string  `json:"guid"`
		Name              string  `json:"name"`              //名称
		OriginalPrice     float64 `json:"originalPrice"`     //原价
		Price             float64 `json:"price"`             //价格
		Width             float32 `json:"width"`             //宽
		Height            float32 `json:"height"`            //高
		Depth             float32 `json:"depth"`             //深度
		Weight            float32 `json:"weight"`            //重量
		Sales             int     `json:"sales"`             //sku的销量
		Status            int16   `json:"status"`            //默认 0
		Sort              int     `json:"sort"`              //排序
		AttributeInfo     string  `json:"attributeInfo"`     //规格属性信息 ，冗余字段 更新商品sku 规格属性需要更新该字段 "[{}]"
		IsSingleAttribute bool    `json:"isSingleAttribute"` //单规格
	}

	LocationModel struct {
		Lat float64 `json:"lat"` //维度
		Lon float64 `json:"lon"` //经度
	}
)
