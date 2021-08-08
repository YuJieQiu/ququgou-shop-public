package cache

//the cache keys
const (
	key           = "USER_WEB_"
	UserInfo      = key + "USER_INFO_"
	HomeHotSearch = key + "HOME_HOT_SEARCH_" //首页热搜
	HotSearch     = key + "HOT_SEARCH_"      //热搜

	DistributedLock              = key + "Distributed_LOCK_"
	OrderCreateLock              = key + "Order_Create_LOCK_"               //订单创建
	CartAddLock                  = key + "Cart_Add_LOCK_"                   //购物车添加
	TimeoutNotPayOrderCancelLock = key + "Order_Timeout_NotPay_Cancel_LOCK" //超时订单取消
)
