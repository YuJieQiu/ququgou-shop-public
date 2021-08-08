package enum

//搜索排序类型
type SearchSortType int

//搜索排序类型 1、默认 3、销量 正序 5、销量 倒叙  7、价格 正序 9、价格 倒叙 11、距离 最近
const (
	SearchSortTypeDefault SearchSortType = iota + 1 //1  默认
	_
	SearchSortTypeSalesASC //3  销量 正序
	_
	SearchSortTypeSalesDESC //5  销量 倒叙
	_
	SearchSortTypePriceASC //7  价格 正序
	_
	SearchSortTypePriceDESC //9  价格 倒叙
	_
	SearchSortTypeDIST //11 距离 最近

)
