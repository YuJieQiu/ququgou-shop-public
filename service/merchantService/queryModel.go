package merchantService

import "github.com/ququgou-shop/library/base_model"

type (
	MerResourceModel struct {
		Id     uint64 `json:"id"`
		Type   int16  `json:"type"`   //类型 默认 0 图片 、可扩展出  1 视频
		Cover  bool   `json:"cover"`  //是否封面
		IsLogo bool   `json:"isLogo"` //是否Logo
		Url    string `json:"url"`
	}

	CreateMerApplyInfoModel struct {
		UserId      uint64             `json:"userId" binding:"required"`
		Name        string             `json:"name" binding:"required"`
		Description string             `json:"description"` //描述
		TypeId      uint64             `json:"typeId"`      //类型 默认 0  MerchantTypeId
		City        string             `json:"city"`
		Region      string             `json:"region"`
		Town        string             `json:"town"`
		Address     string             `json:"address"`
		Latitude    float64            `json:"latitude"`  //纬度
		Longitude   float64            `json:"longitude"` //经度
		Remark      string             `json:"remark"`
		Phone       string             `json:"phone"`
		Resources   []MerResourceModel `json:"resources"` //图片等资源
	}

	MerApplyInfoModel struct {
		Id           uint64  `json:"id"`
		UserId       uint64  `json:"userId"`
		Name         string  `json:"name"`
		Description  string  `json:"description"` //描述
		TypeId       uint64  `json:"typeId"`      //类型 默认 0  MerchantTypeId
		City         string  `json:"city"`
		Region       string  `json:"region"`
		Town         string  `json:"town"`
		Address      string  `json:"Address"`
		Latitude     float64 `json:"latitude"`     //纬度
		Longitude    float64 `json:"longitude"`    //经度
		ResourcesIds string  `json:"resourcesIds"` //图片资源ID 信息
		Remark       string  `json:"remark"`
		Phone        string  `json:"phone"`
		Status       int     `json:"status"`

		StatusText string             `json:"statusText" gorm:"-"`
		Resources  []MerResourceModel `json:"resources" gorm:"-"` //图片等资源
	}

	InitMerchantInfoModel struct {
		Name string `json:"name" form:"name"`
	}

	GetMerchantInfoModel struct {
		MerCode       string `json:"merCode" form:"merCode"` //商户key
		MerId         uint64 `json:"merId" form:"merId"`
		IsWhereUserId bool   `json:"isWhereUserId"` //是否需要关联查询UserId
	}

	//更新商户信息对象
	MerInfoModel struct {
		Id           string             `json:"id"`          //商户key guid
		Name         string             `json:"name"`        //名称
		Description  string             `json:"description"` //描述
		Address      MerAddressModel    `json:"address"`
		BusinessTime MerBusinessTime    `json:"businessTime"`                 //营业时间
		Label        map[uint64]string  `json:"label"`                        //标签 数组 []
		Resources    []MerResourceModel `json:"resources" binding:"required"` //图片等资源
		Phones       string             `json:"phones"`                       //手机 数组["",""]

		ActiveGoods   bool `json:"activeGoods"`   //启用商品
		ActiveService bool `json:"activeService"` //启用服务
		ActiveComment bool `json:"activeComment"` //启用评论

		Logo   MerResourceModel `json:"logo"`   //商标
		TypeId uint64           `json:"typeId"` //商家类型Id
	}

	MerBusinessTime struct {
		StartTime string `json:"startTime"` //开始时间
		EndTime   string `json:"endTime"`   //结束时间
	}

	MerAddressModel struct {
		City      string  `json:"city"`
		Region    string  `json:"region"`
		Town      string  `json:"town"`
		Address   string  `json:"address"`
		Latitude  float64 `json:"latitude"`  //纬度
		Longitude float64 `json:"longitude"` //经度
		Name      string  `json:"name"`      //名称
		Remark    string  `json:"remark"`
	}

	GetLabelListModel struct {
		base_model.QueryParamsPage `gorm:"-"`
		MerId                      uint64 `json:"merId" form:"merId"`
		Type                       int    `form:"type" json:"type"` //1 merchant(商家)
	}

	GetGetLabelModel struct {
		MerId uint64 `json:"merId" form:"merId"`
		Text  string `form:"text" json:"text"`
	}

	ProductMerchantInfoModel struct {
		MerId       uint64          `json:"merId" form:"merId"`
		Name        string          `json:"name"`
		Address     MerAddressModel `json:"address"`
		Phones      string          `json:"phones"` //手机 数组["",""]
		Cover       string          `json:"cover"`
		Description string          `json:"description"` //描述
	}

	GetMerchantForUserModel struct {
		MerId       uint64 `json:"merId"`
		MerGuid     string `json:"merGuid"`
		Name        string `json:"name"`
		Cover       string `json:"cover"`
		Description string `json:"description"`
		TypeId      uint64 `json:"typeId" ` //类型 默认 0  MerchantTypeId
		IsAdmin     bool   `json:"isAdmin"` //是否超级管理员
	}

	QueryMerchantAddressInfoListModel struct {
		MerIds []uint64 `json:"merIds"`
	}

	MerchantAddressInfoModel struct {
		MerId   uint64 `json:"merId" form:"merId" gorm:"column:merchant_id;"`
		MerName string `json:"merName" gorm:"column:mer_name;"`
		Phones  string `json:"phones"` //手机 数组["",""]
		//Cover       string `json:"cover"`
		Description string `json:"description"` //描述
		MerAddressModel
	}

	GetMerchantListModel struct {
		base_model.QueryParamsPage `gorm:"-"`
		Text                       string  `form:"text" json:"text"` //搜索名称 泛搜索 包括 商户名称
		Lat                        float64 `form:"lat" json:"lat"`   //维度
		Lon                        float64 `form:"lon" json:"lon"`   //经度
		ComputeDistance            bool    `form:"computeDistance" json:"computeDistance"`
		SearchSortType             int     `form:"sortType" json:"sortType"` //排序类型 1、默认 3、销量 正序 5、销量 倒叙  7、价格 正序 9、价格 倒叙 11、距离 最近
		Distance                   int     `json:"distance" form:"distance"`
	}

	MerchantListModel struct {
		MerId       uint64  `json:"merId" gorm:"column:mer_id;"`
		Name        string  `json:"name" gorm:"column:name;"`
		Description string  `json:"description" gorm:"column:description;"`
		Image       string  `json:"image" gorm:"column:image;"`         //封面 url
		Latitude    float64 `json:"latitude" gorm:"column:latitude;"`   //纬度
		Longitude   float64 `json:"longitude" gorm:"column:longitude;"` //经度                          //经度
		Distance    float64 `json:"distance" gorm:"column:distance;"`   //距离
	}

	GetLabelModel struct {
		Type int    `form:"type" json:"type"`
		Text string `form:"text" json:"text"`
	}
)
