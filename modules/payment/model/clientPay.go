package model

//客户端支付信息
type ClientPay struct {
	// 设备号，支付可传 WEB
	DeviceInfo string `json:"deviceInfo" gorm:"column:device_info"`
	// 商品描述，可用商品标题
	Body string `json:"body" gorm:"column:body"`
	// 订单标识.
	Reference string `json:"reference" gorm:"column:reference"`
	// 订单总价.
	Total string `json:"total" gorm:"column:total"`
	// 支付回调.
	NotifyURL string `json:"notifyURL" gorm:"column:notify_url"`
}

// Set table name
func (ClientPay) TableName() string {
	return "client_pays"
}
