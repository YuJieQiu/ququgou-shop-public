package paymentService

import (
	"github.com/ququgou-shop/modules/payment/paymentEnum"
	"time"
)

type (
	CreatePaymentModel struct {
		UserId        uint64          `json:"userId" `       //交易用户Id
		PaymentTypeId uint64          `json:"paymentTypeId"` //交易列类型Id
		OrderId       uint64          `json:"orderId"`       //交易订单Id
		Source        int             `json:"source" `       //来源 默认0
		Note          string          `json:"note"`          //记录内容
		ClientInfo    ClientInfoModel `json:"clientInfo"`
		BusinessType  int             `json:"businessType"` //业务类型
		BusinessNo    string          `json:"businessNo"`   //业务编号
		Amount        float64         `json:"amount"`       //交易金额
	}

	CreatePaymentInfoModel struct {
		UserId uint64 `json:"userId" ` //交易用户Id
		//UserGuid   string  `json:"userGuid" gorm:"column:user_guid"`     //交易用户guid 暂留
		BusinessNo string          `json:"businessNo"` //业务编号 (可以是订单系统中的订单编号等....)
		Amount     float64         `json:"amount" `    //交易金额
		PayType    int             `json:"payType" `   //支付类型 0:余额 1:微信 2:支付宝 ...', 暂时默认 1 微信支付 ##PaymentType
		Source     int             `json:"source" `    //来源 默认0
		Note       string          `json:"note" `      //记录内容
		ClientInfo ClientInfoModel `json:"clientInfo"`
	}

	ClientInfoModel struct {
		Ip         string `json:"ip"`
		DeviceInfo string `json:"deviceInfo"  ` // 设备号，支付可传 WEB
	}

	GetPaymentTypeListModel struct {
		MerId uint64 `json:"merId" form:"merId"`
	}

	PaymentNotifyResultModel struct {
		PaymentType   paymentEnum.PaymentType `json:"paymentType"`
		TransactionNo string                  `json:"transactionNo"`
		OutTradeNo    string                  `json:"outTradeNo"`
		BankType      string                  `json:"bankType"`
		UserId        uint64                  `json:"userId"`
		TotalFee      float64                 `json:"totalFee"`
		CashFee       float64                 `json:"cashFee"`
		MchId         string                  `json:"mchId"`
		AppId         string                  `json:"appId"`
		TimeEnd       time.Time               `json:"timeEnd"`

		//CashFeeType string  `json:"cashFeeType"`
		//FeeType       string                  `json:"feeType"`
	}
)
